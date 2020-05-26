package main

import (
	"MysqlRealTimeSync/util"
	"context"
	"github.com/siddontang/go-mysql/mysql"
	"github.com/siddontang/go-mysql/replication"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

var (
	cfg = replication.BinlogSyncerConfig{
		ServerID: 9527,
		Flavor:   "mysql",
		Host:     util.GetConfigString("mysqlSourceHostIp"),
		Port:     uint16(util.GetConfigInt("mysqlSourcePort")),
		User:     util.GetConfigString("mysqlSourceUsername"),
		Password: util.GetConfigString("mysqlSourcePassword"),
	}

	database = "test"
	sqlStr   = "SELECT 1"
)

func main() {

	log.WithFields(log.Fields{
		"Task":         "执行MySQL数据同步",
		"MySQL source": util.GetConfigString("mysqlSourceHostIp") + ":" + util.GetConfigString("mysqlSourcePort"),
		"MySQL target": util.GetConfigString("mysqlTargetHostIp") + ":" + util.GetConfigString("mysqlTargetPort"),
	}).Info("开始同步")

	streamer, _ := replication.NewBinlogSyncer(cfg).StartSync(
		mysql.Position{
			Name: util.GetConfigString("mysqlSourceFile"),
			Pos:  uint32(util.GetConfigInt("mysqlSourcePosition")),
		})

	for {
		e, _ := streamer.GetEvent(context.Background())

		sql, database := ParseBinlogToSqlStr(e)
		log.WithFields(log.Fields{
			"Task":           "MySQL数据同步日志记录",
			"MySQL database": database,
			"MySQL SQL":      sql,
		}).Info("同步的SQL语句")

		_ = util.ExecuteDatabaseTargetSql(sql, database)

	}

	// or we can use a timeout context
	for {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		ev, err := streamer.GetEvent(ctx)
		cancel()
		if err == context.DeadlineExceeded {
			// meet timeout
			continue
		}
		ev.Dump(os.Stdout)
	}

}

func ParseBinlogToSqlStr(e *replication.BinlogEvent) (string, string) {

	switch e.Header.EventType {

	case replication.TABLE_MAP_EVENT:
		tableMapEvent := e.Event.(*replication.TableMapEvent)
		database = string(tableMapEvent.Schema)
		break

	case replication.ROWS_QUERY_EVENT:
		rowsQueryEvent := e.Event.(*replication.RowsQueryEvent)
		sql := string(rowsQueryEvent.Query)
		if sql != "SELECT 1" {
			sqlStr = sql
		}
		break

	default:
		break
	}

	return sqlStr, database
}

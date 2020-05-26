package util

import (
	resetlogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	InfoLevelFileName  = "logs/info.log"
	ErrorLevelFileName = "logs/error.log"
	RotateFileName     = "logs/go.log"

	LogMaxAge       = 3600 * 24 * 15
	LogRotationTime = 3600 * 24 * 30
)

var Log *logrus.Logger

func NewLogger() *logrus.Logger {
	if Log != nil {
		return Log
	}

	pathMap := lfshook.PathMap{
		logrus.InfoLevel:  InfoLevelFileName,
		logrus.ErrorLevel: ErrorLevelFileName,
	}

	Log = logrus.New()

	Log.Hooks.Add(lfshook.NewHook(
		pathMap,
		&logrus.JSONFormatter{},
	))

	return Log
}

func NewRotateLogger() *logrus.Logger {
	if Log != nil {
		return Log
	}
	writer, _ := resetlogs.New(
		RotateFileName+"_%Y%m%d%H%M",
		resetlogs.WithLinkName(RotateFileName),
		resetlogs.WithMaxAge(time.Duration(LogMaxAge)*time.Second),
		resetlogs.WithRotationTime(time.Duration(LogRotationTime)*time.Second),
	)

	Log.Hooks.Add(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.InfoLevel:  writer,
			logrus.ErrorLevel: writer,
		},
		&logrus.JSONFormatter{},
	))

	pathMap := lfshook.PathMap{
		logrus.InfoLevel:  InfoLevelFileName,
		logrus.ErrorLevel: ErrorLevelFileName,
	}

	Log = logrus.New()
	Log.Hooks.Add(lfshook.NewHook(
		pathMap,
		&logrus.JSONFormatter{},
	))

	return Log
}

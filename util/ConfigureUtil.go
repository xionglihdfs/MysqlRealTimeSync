package util

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	configMap map[string]string
}

func GetConfigString(key string) string {
	config := new(Config)
	config.InitConfig("config/config.properties")
	return config.Read(key)
}

func GetConfigInt(key string) int {
	config := new(Config)
	config.InitConfig("config/config.properties")
	value, _ := strconv.Atoi(GetConfigString(key))
	return value
}

func (c *Config) InitConfig(path string) {
	c.configMap = make(map[string]string)

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer func() {
		if f != nil {
			_ = f.Close()
		}
	}()

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		s := strings.TrimSpace(string(b))

		if strings.Index(s, "#") == 0 {
			continue
		}

		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}

		key := strings.TrimSpace(s[:index])
		if len(key) == 0 {
			continue
		}
		value := strings.TrimSpace(s[index+1:])

		pos := strings.Index(value, "\t#")
		if pos > -1 {
			value = value[0:pos]
		}

		pos = strings.Index(value, " #")
		if pos > -1 {
			value = value[0:pos]
		}

		pos = strings.Index(value, "\t//")
		if pos > -1 {
			value = value[0:pos]
		}

		pos = strings.Index(value, " //")
		if pos > -1 {
			value = value[0:pos]
		}

		if len(value) == 0 {
			continue
		}

		c.configMap[key] = strings.TrimSpace(value)
	}
}

func (c Config) Read(key string) string {
	v, found := c.configMap[key]
	if !found {
		return ""
	}
	return v
}

package conf

import (
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
)

const (
	ConfigLocal = "config.yaml"
	ConfigProd  = "config.prod.yaml"
)

type Config struct {
	Server *Server `json:"server" mapstructure:"server" yaml:"server"`
	Data   *Data   `json:"data" mapstructure:"data" yaml:"data"`
}

func LoadConfig() (*Config, error) {
	file := ConfigProd
	if gin.Mode() == gin.DebugMode {
		file = ConfigLocal
	}
	file = "config/" + file
	v := viper.New()
	v.SetConfigFile(file)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	c := &Config{}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Println("配置文件发送变化:", e.Name)
		if err = v.Unmarshal(c); err != nil {
			panic(err)
		}
	})
	if err = v.Unmarshal(c); err != nil {
		return nil, err
	}
	return c, nil
}

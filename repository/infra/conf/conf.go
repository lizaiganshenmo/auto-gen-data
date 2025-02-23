package conf

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	defaultConfPath = "./conf"
)

var (
	Viper = viper.New()
)

func initConf(holder *viper.Viper, confPath, name string) {
	if confPath == "" {
		holder.SetConfigName(name)
		holder.AddConfigPath(defaultConfPath)
	} else {
		holder.SetConfigFile(confPath)
	}
	err := holder.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}

func InitConf(confPath string) {
	initConf(Viper, confPath, "app")
}

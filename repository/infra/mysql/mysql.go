package mysql

import (
	"github.com/lizaiganshenmo/auto-gen-data/repository/infra/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// var dsn = "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local"

var Client *gorm.DB

func Init() {
	var err error
	Client, err = gorm.Open(mysql.Open(conf.Viper.GetString("mysql.dsn")), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
}

package logger

import (
	"log"
	"os"
	"path"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzlogrus "github.com/hertz-contrib/logger/logrus"
	"github.com/lizaiganshenmo/auto-gen-data/repository/infra/conf"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	defaultLogDir     = "./logs/"
	defaultFilePrefix = "auto-gen-data."
)

var (
	Logger *hertzlogrus.Logger
)

func Init() {
	logDir := conf.Viper.GetString("log.dir")
	if "" == logDir {
		logDir = defaultLogDir
	}

	filePrefix := conf.Viper.GetString("log.filePrefix")
	if "" == filePrefix {
		filePrefix = defaultFilePrefix
	}
	fileName := path.Join(logDir, filePrefix+time.Now().Format("2006-01-02T15")+".log")
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			log.Println(err.Error())
			return
		}
	}

	lumberjackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    20,   // A file can be up to 20M.
		MaxBackups: 5,    // Save up to 5 files at the same time.
		MaxAge:     10,   // A file can exist for a maximum of 10 days.
		Compress:   true, // Compress with gzip.
		LocalTime:  true,
	}

	Logger = hertzlogrus.NewLogger()
	Logger.SetOutput(lumberjackLogger)
	Logger.SetLevel(hlog.LevelDebug)

	// 定时切割
	go rotateLog(Logger, logDir+filePrefix)
}

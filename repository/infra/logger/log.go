package logger

import (
	"time"

	hertzlogrus "github.com/hertz-contrib/logger/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

// 切割日志
func rotateLog(logger *hertzlogrus.Logger, baseFilename string) {
	var (
		lj = &lumberjack.Logger{
			MaxSize:    20,
			MaxBackups: 5,
			MaxAge:     10,
			Compress:   true,
			LocalTime:  true,
		}
	)
	now := time.Now()
	// 计算距离下一个整点的时间差
	nextHour := time.Date(now.Year(), now.Month(), now.Day(), now.Hour()+1, 0, 0, 0, now.Location())
	durationToNextHour := nextHour.Sub(now)

	timer := time.NewTimer(durationToNextHour)
	defer timer.Stop()

	<-timer.C

	// 首次切割日志文件
	currentTime := time.Now()
	filename := baseFilename + "." + currentTime.Format("2006-01-02T15")
	lj.Filename = filename
	logger.SetOutput(lj)

	// 定时切割
	ticker := time.NewTicker(time.Hour)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			currentTime := time.Now()
			filename := baseFilename + currentTime.Format("2006-01-02T15") + ".log"
			lj.Filename = filename
			logger.SetOutput(lj)
		}
	}
}

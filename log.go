package glog

import (
	"fmt"
	"path"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

/*
   log config
*/
func ConfigLocalFilesystemLogger(l *logrus.Logger, logPath string, logFileName string, maxAge time.Duration, rotationTime time.Duration) {

	//info log config
	infoLogFileName := fmt.Sprintf("%s.%s.", "info", logFileName)
	infoLogPath := path.Join(logPath, infoLogFileName)
	infoLogWriter, err := rotatelogs.New(
		infoLogPath+"%Y-%m-%d",
		rotatelogs.WithLinkName(infoLogPath),      // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)
	if err != nil {
		l.Errorf("config local file system info logger error. %+v", err)
	}

	//error log config
	errLogFileName := fmt.Sprintf("%s.%s.", "error", logFileName)
	errLogPath := path.Join(logPath, errLogFileName)
	errLogWriter, err := rotatelogs.New(
		errLogPath+"%Y-%m-%d",
		rotatelogs.WithLinkName(errLogPath),       // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(maxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
	)
	if err != nil {
		l.Errorf("config local file system error logger error. %+v", err)
	}

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: infoLogWriter, // 为不同级别设置不同的输出目的
		logrus.InfoLevel:  infoLogWriter,
		logrus.WarnLevel:  infoLogWriter,
		logrus.ErrorLevel: errLogWriter,
		logrus.FatalLevel: errLogWriter,
		logrus.PanicLevel: errLogWriter,
	}, &logrus.TextFormatter{DisableColors: true, TimestampFormat: "2006-01-02 15:04:05.000"})
	l.AddHook(lfHook)
}

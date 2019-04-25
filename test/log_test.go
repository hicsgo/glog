package test

import (
	"glog"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()
	glog.ConfigLocalFilesystemLogger(Log, "your log save  path", "your log name", time.Hour*24*30, time.Hour*24)
}

/*
  测试日志输出
*/
func TestOutPutLog(t *testing.T) {
	Log.Info("test info log")
	Log.Error("test error log")
}

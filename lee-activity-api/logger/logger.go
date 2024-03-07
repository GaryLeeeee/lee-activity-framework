package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

var _logger *logrus.Logger

func init() {
	logger := logrus.New()
	// 设置日志级别
	logger.SetLevel(logrus.DebugLevel)
	// 输出到控制台
	logger.Out = os.Stdout

	_logger = logger
}

func Info(args ...interface{}) {
	_logger.Info(args...)
}

func Infof(format string, args ...interface{}) {
	_logger.Infof(format, args...)
}

func Error(args ...interface{}) {
	_logger.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	_logger.Errorf(format, args...)
}

func Debug(args ...interface{}) {
	_logger.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	_logger.Debugf(format, args...)
}

package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"time"
)

type LeeFormatter struct {
}

func (f LeeFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	levelText := ""
	switch entry.Level {
	case logrus.DebugLevel:
		levelText = "\x1b[34m" + "[DEBUG]" + "\x1b[0m" // 蓝色
	case logrus.InfoLevel:
		levelText = "\x1b[32m" + "[INFO]" + "\x1b[0m" // 绿色
	case logrus.WarnLevel:
		levelText = "\x1b[33m" + "[WARN]" + "\x1b[0m" // 黄色
	case logrus.ErrorLevel:
		levelText = "\x1b[31m" + "[ERROR]" + "\x1b[0m" // 红色
	default:
		panic("unhandled default case")
	}

	// 定制时间格式
	logTime := time.Now().Format("2006-01-02 15:04:05")

	// 只保留当前文件名，而不输出路径
	fileName := filepath.Base(entry.Caller.File)
	// 输出如下：[2024-03-08 19:08:59] [logger.go:20] [INFO] info test
	logStr := fmt.Sprintf("[%s] %s [%s:%d %s] %s\n", logTime, levelText, fileName, entry.Caller.Line, entry.Caller.Function, entry.Message)
	return []byte(logStr), nil
}

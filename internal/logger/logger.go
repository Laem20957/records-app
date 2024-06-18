package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

type ILogger interface {
	Log() *logrus.Logger
	Error(args ...interface{})
	Fatal(args ...interface{})
	Info(args ...interface{})
}

type logger struct {
	*logrus.Logger
}

func (log *logger) Log() *logrus.Logger {
	return log.Logger
}

func (log *logger) Error(args ...interface{}) {
	log.Logger.Error(args...)
}

func (log *logger) Fatal(args ...interface{}) {
	log.Logger.Fatal(args...)
}

func (log *logger) Info(args ...interface{}) {
	log.Logger.Info(args...)
}

func CreateLogs() ILogger {
	logs := new(logger)
	logs.Logger = logrus.New()
	logs.Out = os.Stderr

	formatter := new(logrus.TextFormatter)
	formatter.TimestampFormat = "2006-01-02 15:04:05"
	formatter.FullTimestamp = true
	logs.SetFormatter(formatter)
	return logs
}

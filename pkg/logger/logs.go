package logger

import "github.com/sirupsen/logrus"

func CreateLogs() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	return logger
}

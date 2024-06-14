package database

import (
	"fmt"

	logger "records-app/internal/logger"
	settings "records-app/settings"

	gorm "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var sets = settings.GetSettings()
var logs = logger.CreateLogs()

func ConnectDB() (*gorm.DB, error) {
	strConn := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		sets.DBHost, sets.DBPort, sets.DBName, sets.DBUsername, sets.DBPassword)

	db, err := gorm.Open("postgres", strConn)
	if err != nil {
		logs.Log().Fatal(err)
	}
	if err = db.DB().Ping(); err != nil {
		logs.Log().Fatal(err)
	}
	return db, nil
}

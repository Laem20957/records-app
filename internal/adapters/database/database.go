package database

import (
	"fmt"
	"time"

	"records-app/internal/logger"
	"records-app/settings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var logs = logger.CreateLogs()

type IConnectionDB interface {
	ConnectDB() *gorm.DB
	CloseDB()
}

type ConnectionDB struct {
	conn *gorm.DB
}

func (db *ConnectionDB) ConnectDB() *gorm.DB {
	return db.conn
}

func (db *ConnectionDB) CloseDB() {
	if db.conn != nil {
		defer db.conn.Close()
	} else {
		logs.Fatal("runtime error:" +
			"invalid memory address or nil pointer dereference")
	}
}

func NewConnectionDB(settings *settings.Settings) (IConnectionDB, error) {
	strConn := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
		settings.DBHost, settings.DBPort, settings.DBName,
		settings.DBUsername, settings.DBPassword, settings.DBSSLMode)

	db, err := gorm.Open("postgres", strConn)
	if err != nil {
		logs.Log().Fatal(err)
	}

	db.DB().SetConnMaxLifetime(5 * time.Minute)
	db.DB().SetMaxOpenConns(10)
	db.DB().SetMaxIdleConns(10)

	if err = db.DB().Ping(); err != nil {
		logs.Log().Fatal(err)
	}
	return &ConnectionDB{conn: db}, nil
}

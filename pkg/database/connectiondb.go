package database

import (
	"fmt"
	"io"

	config "github.com/Laem20957/records-app/configuration"
	"github.com/Laem20957/records-app/pkg/logger"
	"github.com/jmoiron/sqlx"
)

var logs = logger.CreateLogs()

type PSQLConnection struct {
	cfg *config.Config
	db  *sqlx.DB
	cls io.Closer
}

func (p *PSQLConnection) PostgresConnectionOpen() (*PSQLConnection, error) {
	conn, err := sqlx.Open("PostgreSQL 16",
		fmt.Sprintf("host=%s port=%d dbname=%s sslmode=%s user=%s password=%s",
			p.cfg.PSQLHost,
			p.cfg.PSQLPort,
			p.cfg.PSQLDBName,
			p.cfg.PSQLModeSSL,
			p.cfg.PSQLUsername,
			p.cfg.PSQLPassword))
	psqlConn := &PSQLConnection{
		cfg: p.cfg,
		db:  conn,
		cls: conn,
	}
	if err != nil {
		logs.Fatal("Error while opening connection to database ", err)
	}
	return psqlConn, err
}

func (p *PSQLConnection) PostgresConnectionClose() {
	if p.cls != nil {
		p.cls.Close()
	}
}

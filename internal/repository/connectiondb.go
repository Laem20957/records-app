package repository

import (
	"fmt"

	cfg "github.com/Laem20957/records-app/configurations"
	"github.com/Laem20957/records-app/internal/logger"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

var logs = logger.CreateLogs()

type PSQLConnection struct {
	cfg cfg.Configurations
	db  *sqlx.DB
}

func (p *PSQLConnection) PostgreConnection() *PSQLConnection {
	conn, err := sqlx.Open("pgx",
		fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
			p.cfg.PSQLHost,
			p.cfg.PSQLPort,
			p.cfg.PSQLDBName,
			p.cfg.PSQLUsername,
			p.cfg.PSQLPassword,
			p.cfg.PSQLModeSSL))

	psqlConn := &PSQLConnection{
		cfg: p.cfg,
		db:  conn,
	}
	if err != nil {
		logs.Log().Fatal("Error while opening connection to database ", err)
	}

	defer psqlConn.db.Close()
	return psqlConn
}

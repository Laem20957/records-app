package main

import (
	config "github.com/Laem20957/records-app/configs"
	server "github.com/Laem20957/records-app/internal/servers"

	_ "github.com/lib/pq"
)

// @Title Records-App
// @Version 1.0
// @Description Server API

// @Contact.name Daniel Kotelnikov
// @Contact.email danielkotelnikov20@gmail.com

// @Host localhost:8080
// @BasePath /

// @SecurityDefinitions.apikey ApiKeyAuth
// @In header
// @Name Authorization

func main() {

	// db, err := database.NewPostgresConnection(database.ConnectionInfo{
	// 	Host:     cfg.DB.Host,
	// 	Port:     cfg.DB.Port,
	// 	Username: cfg.DB.Username,
	// 	DBName:   cfg.DB.Name,
	// 	SSLMode:  cfg.DB.SSLMode,
	// 	Password: cfg.DB.Password,
	// })
	// if err != nil {
	// 	logrus.Fatal(err)
	// }

	// defer func(db *sqlx.DB) {
	// 	if err := db.Close(); err != nil {
	// 		logrus.Fatal(err)
	// 	}
	// }(db)

	// noteCache := gcache.New(20).LRU().Build()
	// noteRepo := repository.NewRepository(db)
	// noteService := service.NewService(cfg, noteCache, noteRepo)

	config.InitConfigs()

	httpServer := server.HttpServer{}
	httpServer.HttpServerSettings().HttpServerStart()
	httpServer.HttpServerSettings().HttpServerStop()

}

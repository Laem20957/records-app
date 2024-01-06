package main

import (
	"github.com/Laem20957/records-app/internal/server"

	_ "github.com/lib/pq"
)

// @title Note app API
// @version 1.0
// @description API server for Note app

// @contact.name Dmitry Mikhaylov
// @contact.email ru.system.ru@gmail.com

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

const (
	CONFIG_DIR  = "configs"
	CONFIG_FILE = "main"
)

func main() {

	// cfg, err := config.New(CONFIG_DIR, CONFIG_FILE)
	// if err != nil {
	// 	logrus.Fatal(err)
	// }

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
	// noteHandler := rest.NewHandler(noteService)
	// noteRouter := noteHandler.InitRoutes()

	httpServer := server.HttpServerSettings()
	httpServer.HttpServerStart()
	httpServer.HttpServerStop()

}

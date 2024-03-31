package main

import (
	config "github.com/Laem20957/records-app/configuration"
	"github.com/Laem20957/records-app/internal/server"
	psql "github.com/Laem20957/records-app/pkg/database"
	_ "github.com/lib/pq"
)

// @Title Records-App
// @Version 1.0
// @Description Server API
// @Contact.name Daniel Kotelnikov
// @Host localhost:8080
// @BasePath /
// @SecurityDefinitions.apikey ApiKeyAuth
// @In header
// @Name Authorization

func main() {

	config.InitConfigs()

	httpServer := server.HttpServer{}
	httpServer.HttpServerSettings().HttpServerStart()
	httpServer.HttpServerSettings().HttpServerStop()

	database := psql.PSQLConnection{}
	database.PostgreConnection()
}

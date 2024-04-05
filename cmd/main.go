package main

import (
	config "github.com/Laem20957/records-app/configuration"
	"github.com/Laem20957/records-app/internal/api/rest"
	psql "github.com/Laem20957/records-app/pkg/database"
	_ "github.com/lib/pq"
)

// @Title Records-App
// @Version 1.0
// @Description Server API
// @Host localhost:8080
// @BasePath /
// SecurityDefinitions.apikey ApiKeyAuth
// In Header
// Name Authorization
func main() {

	config.InitConfigs()

	httpServer := rest.HttpServer{}
	httpServer.HttpServerSettings().HttpServerStart()
	httpServer.HttpServerSettings().HttpServerStop()

	database := psql.PSQLConnection{}
	database.PostgreConnection()
}

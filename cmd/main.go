package main

import (
	rest "github.com/Laem20957/records-app/api/rest"
	cfg "github.com/Laem20957/records-app/configurations"
	psql "github.com/Laem20957/records-app/internal/repository"
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

	cfg.InitConfigs()

	httpServer := rest.HttpServer{}
	httpServer.HttpServerSettings().HttpServerStart()
	httpServer.HttpServerSettings().HttpServerStop()

	database := psql.PSQLConnection{}
	database.PostgreConnection()
}

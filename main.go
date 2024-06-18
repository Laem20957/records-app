package main

import (
	"records-app/api/rest"
	"records-app/internal/adapters/database"
	"records-app/settings"

	_ "github.com/lib/pq"
)

// @Title Records-App
// @Version 1.0
// @Description Server API
// @Host localhost:8080
// @license.name MIT License
// @license.url  https://opensource.org/license/MIT
// @SecurityDefinitions.apikey ApiKeyAuth
// @In Header
// @Name Authorization
func main() {
	settings.GetSettings()

	httpServer := rest.HttpServer{}
	httpServer.NewHttpServer().HttpServerStart()
	httpServer.NewHttpServer().HttpServerStop()

	database := database.ConnectionDB{}
	database.ConnectDB()
	database.CloseDB()

	// var record domain.Records
	// if err := database.Table("records_app.records").Where("id = ?", 1).First(&record).Error; err != nil {
	// 	panic("Record not found")
	// }
	// fmt.Println(record)

	// rows, err := database.DB.Query("SELECT * FROM records_app.records")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer rows.Close()
	// var results []string
	// for rows.Next() {
	// 	var column1, column2, column3, column4 string
	// 	err := rows.Scan(&column1, &column2, &column3, &column4)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	results = append(results, fmt.Sprintf("%s, %s, %s, %s",
	// 		column1, column2, column3, column4))
	// }
	// if err := rows.Err(); err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(results)
}

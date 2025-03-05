package main

import (
	_ "crud-mongo/docs"
	"crud-mongo/internal/database"
	"crud-mongo/internal/routes"
)

// @title           API Go Gin CRUD
// @version         1.0
// @description     Exemplo de API CRUD com Swagger no Gin
// @host            localhost:8080
// @BasePath        /
func main() {
	database.InitDatabase()
	database.InitRedis()
	routes.SetupRouter()
}

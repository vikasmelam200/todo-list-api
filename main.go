package main

import (
	"todo-list-api/db"
	"todo-list-api/routes"
)

func main() {
	db.Init()
	router := routes.SetupRoutes()
	router.Run(":8080")
}

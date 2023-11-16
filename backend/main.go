package main

import (
	"knowledge_checkup/backend/router"
	"knowledge_checkup/backend/view"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	view.GetTpl()        // Парсинг всіх шаблонів на початку роботи серверу
	router.SetupRoutes() // Налаштування маршрутів
}

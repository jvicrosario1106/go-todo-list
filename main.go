package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jvicrosario1106/todo-list/database"
	"github.com/jvicrosario1106/todo-list/routes"
)

func main() {

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())
	database.DbConnection()
	routes.SetUpRoutes(app)

	app.Listen(":8000")
}

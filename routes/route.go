package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jvicrosario1106/todo-list/handlers"
)

func SetUpRoutes(app *fiber.App) {

	api := app.Group("/todos")

	api.Get("/", handlers.GetAll)
	api.Get("/:id", handlers.GetOne)
	api.Post("/", handlers.CreateTodo)
	api.Delete("/:id", handlers.DeleteTodo)
	api.Post("/:id", handlers.UpdateTodo)

}

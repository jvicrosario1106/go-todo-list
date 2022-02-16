package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jvicrosario1106/todo-list/database"
	"github.com/jvicrosario1106/todo-list/models"
)

func GetAll(c *fiber.Ctx) error {

	var todo []models.Todo

	if err := database.DB.Find(&todo).Error; err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Failed to get all todos",
		})
	}

	return c.Status(fiber.StatusOK).JSON(todo)

}

func GetOne(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo models.Todo

	if err := database.DB.Find(&todo, id).Error; err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "No data with this ID",
		})
	}

	return c.Status(fiber.StatusOK).JSON(todo)

}

func CreateTodo(c *fiber.Ctx) error {

	todo := new(models.Todo)

	if err := c.BodyParser(todo); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Unable to create new todolist",
		})
	}

	if err := database.DB.Create(&todo).Error; err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Failed in creating todos",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(todo)

}

func DeleteTodo(c *fiber.Ctx) error {

	id := c.Params("id")
	var todo models.Todo

	if err := database.DB.Where("ID = ?", id).First(&todo).Error; err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Unable to delete with this ID",
		})
	}

	database.DB.Delete(&todo, id)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully Deleted",
	})

}

func UpdateTodo(c *fiber.Ctx) error {

	type updateTodoData struct {
		Title string `json:"title"`
	}

	todo := new(models.Todo)
	id := c.Params("id")

	//store struct type request body to variable
	var updateData updateTodoData

	if err := c.BodyParser(&updateData); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Failed to Update",
		})
	}

	if err := database.DB.Where("ID = ?", id).First(&todo).Error; err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "No data with this ID",
		})
	}

	todo.Title = updateData.Title

	database.DB.Save(&todo)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Successfully Updated",
	})

}

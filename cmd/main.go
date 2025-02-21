package main

import (
	"github.com/Asker231/todo-list-fiber.git/config"
	"github.com/Asker231/todo-list-fiber.git/internal/todos"
	"github.com/gofiber/fiber/v2"
)


func main() {
	//init server
	app := fiber.New()
	//init config
	_ = config.NewConfig()
	//init db
	
	//init handler
	todos.NewTodoHandler(app)


	app.Listen(":3002")
}
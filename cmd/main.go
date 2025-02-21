package main

import (
	"github.com/Asker231/todo-list-fiber.git/config"
	"github.com/Asker231/todo-list-fiber.git/internal/todos"
	"github.com/Asker231/todo-list-fiber.git/internal/todos/repo"
	"github.com/Asker231/todo-list-fiber.git/pkg/db"
	"github.com/gofiber/fiber/v2"
)


func main() {
	//init server
	app := fiber.New()
	//init config
	cnf := config.NewConfig()

	//init db
	db ,err := db.InitDataBAse(cnf)	
	if err != nil{
		return 
	}
	defer db.Close()
	//init repo 
	todoRepo  := repo.NewTodoRepo(db)
	//init service
	serv := todos.NewService(*todoRepo)
	//init handler

	todos.NewTodoHandler(app,serv)


	app.Listen(":3002")
}
package todos

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type TodoHandler struct{
		Service ServiceTodo
}

func NewTodoHandler(router fiber.Router,service *ServiceTodo){
	todoHandler := &TodoHandler{
		Service: *service,
	}
	api := router.Group("/task")
	api.Post("/create",todoHandler.AddTodo)
	api.Get("/tasks/:id",todoHandler.GetAll)
	api.Put("/task/:id",todoHandler.UpdateTodo)
	api.Put("/task/:id",todoHandler.DeleteTodo)
}
//add todo method
func(t *TodoHandler)AddTodo(ctx *fiber.Ctx)error{

	todo := new(TodoRequest)

	if err := ctx.BodyParser(todo);err != nil{	
		fmt.Println(err.Error()) 
	}

	t.Service.CretaeService(todo.Title,todo.Description)
	
	return ctx.SendString("Add todo")
}

//getid
func(t *TodoHandler)GetAll(ctx *fiber.Ctx)error{
	idStr := ctx.Params("id")

	id,err:= strconv.Atoi(idStr)
	if err != nil{
		fmt.Println(err.Error())
	}

	todo := t.Service.FindTaskByid(id)

	fmt.Println(todo)
	return ctx.SendString("All todo")
}

//update todo
func(t *TodoHandler)UpdateTodo(ctx *fiber.Ctx)error{
	return ctx.SendString("Update")
}

func(t *TodoHandler)DeleteTodo(ctx *fiber.Ctx)error{
	return ctx.SendString("Delete")
}
package todos

import (
	"fmt"
	"strconv"

	"github.com/Asker231/todo-list-fiber.git/pkg/validate"
	"github.com/gofiber/fiber/v2"
)

type TodoHandler struct{
		Service ServiceTodo
}

func NewTodoHandler(router fiber.Router,service *ServiceTodo){
	todoHandler := &TodoHandler{
		Service: *service,
	}
	api := router.Group("/api")
	api.Post("/tasks",todoHandler.AddTodo)
	api.Get("/tasks",todoHandler.GetAll)
	api.Put("/tasks/:id",todoHandler.UpdateTodo)
	api.Delete("/tasks/:id",todoHandler.DeleteTodo)
}
//add todo method
func(t *TodoHandler)AddTodo(ctx *fiber.Ctx)error{

	todo := new(TodoRequest)

	if err := ctx.BodyParser(todo);err != nil{	
		fmt.Println(err.Error()) 
	}

	err := validate.ValidateBody(todo)
	if err != nil{
		ctx.JSON(err.Error())
		return err
	}

	t.Service.CretaeService(todo.Title,todo.Description,todo.Status)
	
	return ctx.SendString("Add todo")
}

//getall
func(t *TodoHandler)GetAll(ctx *fiber.Ctx)error{
	allTodo := t.Service.GetAllTodoService()

	return ctx.JSON(allTodo)
}

//update todo
func(t *TodoHandler)UpdateTodo(ctx *fiber.Ctx)error{
	idStr := ctx.Params("id")
	_,_ = strconv.Atoi(idStr)
	return ctx.SendString("Update")
}

//delete
func(t *TodoHandler)DeleteTodo(ctx *fiber.Ctx)error{
	idStr := ctx.Params("id")

	id,err:= strconv.Atoi(idStr)

	if err != nil{
		fmt.Println(err.Error())
	}

	_ = t.Service.DeleteById(id)

	return ctx.SendString("Delete")
}
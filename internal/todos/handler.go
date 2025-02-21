package todos

import "github.com/gofiber/fiber/v2"

type TodoHandler struct{
	
}

func NewTodoHandler(router fiber.Router){
	todoHandler := &TodoHandler{
	}
	api := router.Group("/todo")
	api.Post("/add",todoHandler.AddTodo)
	api.Get("/allTodo",todoHandler.GetAll)
	api.Put("/update/:id",todoHandler.UpdateTodo)
	api.Put("/delete/:id",todoHandler.DeleteTodo)
}
//add todo method
func(t *TodoHandler)AddTodo(ctx *fiber.Ctx)error{
	return ctx.SendString("Add todo")
}

//getAll
func(t *TodoHandler)GetAll(ctx *fiber.Ctx)error{
	return ctx.SendString("All todo")
}

//update todo
func(t *TodoHandler)UpdateTodo(ctx *fiber.Ctx)error{
	return ctx.SendString("Update")
}

func(t *TodoHandler)DeleteTodo(ctx *fiber.Ctx)error{
	return ctx.SendString("Delete")
}
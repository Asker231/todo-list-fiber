package todos

import (
	"fmt"
	"time"

	"github.com/Asker231/todo-list-fiber.git/internal/todos/repo"
)


type ServiceTodo struct{
	Repo repo.RepositoryTodo
}

func NewService(repo repo.RepositoryTodo)*ServiceTodo{
	return &ServiceTodo{
		Repo: repo,
	}
}

//create 
func(s *ServiceTodo)CretaeService(title,description string){

	todo := repo.TodoModel{
		Title: title,
		Description: description,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	s.Repo.Cretate(todo)

}

//find by id

func(s *ServiceTodo)FindTaskByid(id int)([]repo.TodoModel){

	t ,err := s.Repo.FindByid(id)
	if err != nil{
		fmt.Println(err.Error())
	}
	return t
}
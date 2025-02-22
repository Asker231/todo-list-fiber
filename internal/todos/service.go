package todos

import (
	"fmt"
	"time"

	"github.com/Asker231/todo-list-fiber.git/internal/todos/repo"
)

type ServiceTodo struct {
	Repo repo.RepositoryTodo
}

func NewService(repo repo.RepositoryTodo) *ServiceTodo {
	return &ServiceTodo{
		Repo: repo,
	}
}

// create
func (s *ServiceTodo) CretaeService(title, description string) {

	todo := repo.TodoModel{
		Title:       title,
		Description: description,
		Status:      "new",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	s.Repo.Cretate(todo)

}

//find by id
func (s *ServiceTodo) DeleteById(id int) repo.TodoModel {

	todo, err := s.Repo.DeleteByid(id)
	if err != nil {
		fmt.Println(err.Error())
	}
	return todo
}

// get all
func (s *ServiceTodo) GetAllTodoService() []repo.TodoModel {

	data, err := s.Repo.GetAllRepo()
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return data
}

// update by id
func (s *ServiceTodo) UpdateTodoService(id int, status string) error {
	err := s.Repo.UpdateRepo(id, status)
	if err != nil {
		return err
	}
	return err
}

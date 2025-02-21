package todos

import "github.com/Asker231/todo-list-fiber.git/internal/todos/repo"


type ServiceTodo struct{
	Repo repo.RepositoryTodo
}

func NewService(repo repo.RepositoryTodo)*ServiceTodo{
	return &ServiceTodo{
		Repo: repo,
	}
}
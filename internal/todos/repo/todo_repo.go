package repo

import "github.com/jackc/pgx/v5/pgxpool"


type RepositoryTodo struct{
	DataBase *pgxpool.Pool
}


func NewTodoRepo(dbref *pgxpool.Pool)*RepositoryTodo{
	return &RepositoryTodo{
		DataBase: dbref,
	}
}

//cretate todo


//find todo
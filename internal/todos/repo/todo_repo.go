package repo

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)


type RepositoryTodo struct{
	DataBase *pgxpool.Pool
}


func NewTodoRepo(dbref *pgxpool.Pool)*RepositoryTodo{
	return &RepositoryTodo{
		DataBase: dbref,
	}
}

//cretate todo
func(r *RepositoryTodo)Cretate(todo TodoModel){
	query := `INSERT INTO todo(title,description,created_at,updated_at)VALUES(@title,@description,@created_at,@updated_at);`
	args := pgx.NamedArgs{
		"title":todo.Title,
		"description":todo.Description,
		"created_at":todo.CreatedAt,
		"updated_at":todo.UpdatedAt,
	}
	res,err := r.DataBase.Exec(context.Background(),query,args)
	if err != nil{
		fmt.Println(err.Error())
	}
	fmt.Println(res)
}

//find todo
func(r *RepositoryTodo)FindByid(id int)([]TodoModel,error){
	query := `SELECT * FROM todo WHERE id = $1`

	args := pgx.NamedArgs{
		"id":id,
	}

	rows ,_ := r.DataBase.Query(context.Background(),query,args)

	todo,err := pgx.CollectRows(rows,pgx.RowToStructByNameLax[TodoModel])
	if err != nil{
		fmt.Println(err.Error())
	}
	return todo,nil
}
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
	fmt.Println(res.String())
}

//delete todo
func(r *RepositoryTodo)DeleteByid(id int)(TodoModel,error){
 	var todo  TodoModel
	err := r.DataBase.QueryRow(context.Background(),"DELETE  FROM todo WHERE id = $1",id).Scan(&todo.Id,&todo.Title,&todo.Description,&todo.CreatedAt,&todo.UpdatedAt)
	if err != nil{
		fmt.Println(err.Error())
	}
	return todo,nil
}

//getall todo
func(r *RepositoryTodo)GetAllRepo()([]TodoModel,error){
	var todos []TodoModel
	rows,err := r.DataBase.Query(context.Background(),"SELECT * FROM todo")
	if err != nil{
		return nil,err
	}
	defer rows.Close()
	for rows.Next(){
		var todo TodoModel
		if err := rows.Scan(&todo.Id,&todo.Title,&todo.Description,&todo.CreatedAt,&todo.UpdatedAt);err != nil{
		   fmt.Println(err.Error())
		}
		todos = append(todos, todo)
	}
	return todos,nil
}

//update todo
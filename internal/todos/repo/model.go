package repo

import "time"


type TodoModel struct{
   Id int  
   Title string 
   Description string 
   CreatedAt time.Time
   UpdatedAt time.Time
}


func NewTodomodel(title,description string)*TodoModel{
	return &TodoModel{
		Title: title,
		Description: description,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

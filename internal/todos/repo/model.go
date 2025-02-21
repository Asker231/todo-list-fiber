package repo

import "time"


type TodoModel struct{
   Id int  `json:"id"`
   Title string `json:"title"`	
   Description string `json:"description"`
   CreatedAt time.Time `json:"created_at"`
   UpdatedAt time.Time `json:"updated_at"`
}


func NewTodomodel(title,description string)*TodoModel{
	return &TodoModel{
		Title: title,
		Description: description,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

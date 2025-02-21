package db

import (
	"context"

	"github.com/Asker231/todo-list-fiber.git/config"
	"github.com/jackc/pgx/v5/pgxpool"
)



func InitDataBAse(conf *config.AppConfig)(*pgxpool.Pool,error){
	conn,err := pgxpool.New(context.Background(),conf.DB.DNS)
	if err != nil{
		return nil,err
	}
	return conn,nil
} 
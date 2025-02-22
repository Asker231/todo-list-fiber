package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type (
	AppConfig struct {
		DB DataBaseConfig
	}

	DataBaseConfig struct {
		DNS string
	}
)

func NewConfig() *AppConfig {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println(err.Error())
	}
	return &AppConfig{
		DB: DataBaseConfig{
			DNS: os.Getenv("DNS"),
		},
	}
}

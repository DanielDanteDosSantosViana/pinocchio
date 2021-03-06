package config

import (
	"fmt"
	"os"
)

type service struct {
	Port string `json:"port"`
}

type db struct {
	Mongo string `json:"mongo"`
	Name  string `json:"name"`
}

type config struct {
	Service  service  `json:"service"`
	Db       db       `json:"db"`
}

var Conf config

func Load() {
	PORT_ENV := os.Getenv("PORT_ENV")
	MONGO_HOST := os.Getenv("MONGO_HOST")
	DB_NAME := os.Getenv("DB_NAME")

	if PORT_ENV == "" || MONGO_HOST == "" || DB_NAME == "" {
		panic(fmt.Errorf("Não foram encontradas as variáveis de ambiente para inicialização do sistema. Verifique a váriaveis 'PORT_ENV', 'MONGO_HOST', 'DB_NAME' ."))
	}

	Conf = config{service{PORT_ENV}, db{MONGO_HOST, DB_NAME}}
}
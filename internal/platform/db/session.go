package db

import (
	"github.com/DanielDanteDosSantosViana/pinocchio/internal/platform/config"
	"errors"
	"fmt"
	"gopkg.in/mgo.v2"
)

func NewSession() (mgo.Session, error) {
	db, err := mgo.Dial(config.Conf.Db.Mongo)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Ocorreu um erro ao tentar abrir conexão com o dbWrite . %v", err))
	}

	err = db.Ping()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Ocorreu um erro ao tentar verificar conexão dbWrite . %v", err))
	}

	return db, nil
}
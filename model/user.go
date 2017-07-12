package model

import (
	"github.com/fberrez/forum/datastore"
)

type User struct {
	Id       int    `json:"id"`
	Pseudo   string `json:"pseudo"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GetUserByPseudo(pseudo string) (User, error) {
	result := User{}
	var err error

	switch datastore.ReadConfig().Type {
	case datastore.TypeMySQL:
		err = datastore.SQL.Get(&result, "SELECT * FROM user WHERE pseudo = ?", pseudo)
	}

	return result, err
}

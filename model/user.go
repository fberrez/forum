package model

import (
	"database/sql"
	"github.com/fberrez/forum/datastore"
	"time"
)

type User struct {
	Id       int       `json:"id" db:"id"`
	Pseudo   string    `json:"pseudo" db:"pseudo"`
	Password string    `json:"password" db:"password"`
	Email    string    `json:"email" db:"email"`
	Date     time.Time `json:"date" db:"date"`
	Karma    float64   `json:"karma" db:"karma"`
	Ip       string    `json:"ip" db:"ip"`
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

func GetUserByEmail(email string) (User, error) {
	result := User{}
	var err error

	switch datastore.ReadConfig().Type {
	case datastore.TypeMySQL:
		err = datastore.SQL.Get(&result, "SELECT * FROM user WHERE email = ?", email)
	}

	return result, err
}

func CreateUser(newUser User) error {
	var err error

	switch datastore.ReadConfig().Type {
	case datastore.TypeMySQL:
		_, err = datastore.SQL.Exec("INSERT INTO user (pseudo, password, email, karma, ip) VALUES(?, ?, ?, 5, ?)", newUser.Pseudo, newUser.Password, newUser.Email, newUser.Ip)
	}

	return err
}

func EditUser(newUser User) (sql.Result, error) {
	var result sql.Result
	var err error

	switch datastore.ReadConfig().Type {
	case datastore.TypeMySQL:
		result, err = datastore.SQL.NamedExec("UPDATE user SET pseudo=:pseudo, password=:password, email=:email, karma=:karma, ip=:ip WHERE id=:id", newUser)

	}

	return result, err
}

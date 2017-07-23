package model

import (
	"database/sql"
	"github.com/fberrez/forum/datastore"
	"time"
)

type User struct {
	Id              int       `json:"id" db:"user_id"`
	Pseudo          string    `json:"pseudo" db:"user_pseudo"`
	Password        string    `json:"password" db:"user_password"`
	Email           string    `json:"email" db:"user_email"`
	Date            time.Time `json:"date" db:"user_date"`
	Date_lastConnec time.Time `json:"date_lastConnec" db:"user_date_lastConnection"`
	Group           int       `json:"group" db:"user_groupId"`
	Karma           float64   `json:"karma" db:"user_karma"`
	Ip              string    `json:"ip" db:"user_ip"`
}

func GetUserByPseudo(pseudo string) (User, error) {
	result := User{}
	var err error

	switch datastore.ReadConfig().Type {
	case datastore.TypeMySQL:
		err = datastore.SQL.Get(&result, "SELECT * FROM forum_user WHERE user_pseudo = ? LIMIT 1", pseudo)
	}

	return result, err
}

func GetUserByEmail(email string) (User, error) {
	result := User{}
	var err error

	switch datastore.ReadConfig().Type {
	case datastore.TypeMySQL:
		err = datastore.SQL.Get(&result, "SELECT * FROM forum_user WHERE user_email = ? LIMIT 1", email)
	}

	return result, err
}

func CreateUser(newUser User) error {
	var err error

	switch datastore.ReadConfig().Type {
	case datastore.TypeMySQL:
		_, err = datastore.SQL.Exec("INSERT INTO forum_user (user_pseudo, user_password, user_email, user_karma, user_ip) VALUES(?, ?, ?, 5, ?)", newUser.Pseudo, newUser.Password, newUser.Email, newUser.Ip)
	}

	return err
}

func EditUser(newUser User) (sql.Result, error) {
	var result sql.Result
	var err error

	switch datastore.ReadConfig().Type {
	case datastore.TypeMySQL:
		result, err = datastore.SQL.NamedExec("UPDATE forum_user SET user_pseudo=:pseudo, user_password=:password, user_email=:email, user_karma=:karma, user_ip=:ip, user_groupId:=:group, user_date_lastConnection=:date_lastConnec WHERE id=:id", newUser)

	}

	return result, err
}

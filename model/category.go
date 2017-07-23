package model

import (
	"github.com/fberrez/forum/datastore"
	"time"
)

type Category struct {
	Id          int       `json:"id" db:"category_id"`
	Title       string    `json:"title" db:"category_title"`
	Description string    `json:"description" db:"category_description"`
	Date        time.Time `json:"date" db:"category_date"`
}

func GetCategory() ([]Category, error) {
	result := []Category{}
	var err error

	switch datastore.ReadConfig().Type {
	case datastore.TypeMySQL:
		err = datastore.SQL.Select(&result, "SELECT * FROM forum_category")
	}

	return result, err
}

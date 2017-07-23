package model

import (
	"github.com/fberrez/forum/datastore"
	"time"
)

type SubCategory struct {
	Id          int       `json:"id" db:"subCategory_id"`
	IdCategory  int       `json:"idCategory" db:"subCategory_idCategory"`
	Title       string    `json:"title" db:"subCategory_title"`
	Description string    `json:"description" db:"subCategory_description"`
	Date        time.Time `json:"date" db:"subCategory_date"`
}

func GetSubCategory() ([]SubCategory, error) {
	result := []SubCategory{}
	var err error

	switch datastore.ReadConfig().Type {
	case datastore.TypeMySQL:
		err = datastore.SQL.Select(&result, "SELECT * FROM forum_subCategory")
	}

	return result, err
}

func GetSubCategoryByIdCategory(idCategory int) ([]SubCategory, error) {
	result := []SubCategory{}
	var err error

	switch datastore.ReadConfig().Type {
	case datastore.TypeMySQL:
		err = datastore.SQL.Select(&result, "SELECT * FROM forum_subCategory WHERE subCategory_idCategory = ?", idCategory)
	}

	return result, err
}

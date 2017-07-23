package model

import (
	"github.com/fberrez/forum/datastore"
	"time"
)

type Post struct {
	Id            int       `json:"id" db:"post_id"`
	IdSubcategory int       `json:"idSubcategory" db:"post_idSubCategory"`
	IdUser        int       `json:"idUser" db:"post_idUser"`
	IdParentPost  int       `json:"idParentPost" db:"post_idParentPost"`
	Title         string    `json:"title" db:"post_title"`
	Content       string    `json:"content" db:"post_content"`
	IsPoll        bool      `json:"isPoll" db:"post_isPoll"`
	PollTitle     string    `json:"pollTitle" db:"post_pollTitle"`
	IsEdited      bool      `json:"isEdited" db:"post_isEdited"`
	CreateDate    time.Time `json:"createDate" db:"post_createDate"`
	EditDate      time.Time `json:"editDate" db:"post_editDate"`
	Ip            string    `json:"ip" db:"post_ip"`
}

func GetPost() ([]Post, error) {
	result := []Post{}
	var err error

	switch datastore.ReadConfig().Type {
	case datastore.TypeMySQL:
		err = datastore.SQL.Select(&result, "SELECT * FROM forum_post")
	}

	return result, err
}

func GetPostByIdSubCat(id int) ([]Post, error) {
	result := []Post{}
	var err error

	switch datastore.ReadConfig().Type {
	case datastore.TypeMySQL:
		err = datastore.SQL.Select(&result, "SELECT * FROM forum_post WHERE post_idSubCategory = ?", id)
	}

	return result, err
}

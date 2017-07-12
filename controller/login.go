package controller

import (
	"github.com/fberrez/forum/model"
	"github.com/gin-gonic/gin"
	"log"
)

func ConnectUser(c *gin.Context) {
	user, err := model.GetUserByPseudo("toast")
	if err != nil {
		log.Printf("db.Query error :\n\t-query: GetUserByPseudo\n\t-error: %v\n\n", err)
	} else {
		log.Println(user)
	}
}

/*
func getUserInfo(c *gin.Context) {
	dbObject := datastore.DbObject{Address: "", Dbname: "", Username: "", Password: ""}

	dbObject.Flow = datastore.Connect(dbObject)

	rows, err := store.GetAllUsers(dbObject.Flow)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error - Bad Request"})
		return
	}

	var users []model.User
	for rows.Next() {
		var id int
		var pseudo, email, password string
		err := rows.Scan(&id, &pseudo, &email, &password)

		if err != nil {
			log.Printf("rows.Scan for user error : %v", err)
		}

		user := model.User{Id: id, Pseudo: pseudo, Email: email, Password: password}
		users = append(users, user)
	}

	err = datastore.Close(dbObject.Flow)
	if err != nil {
		log.Fatalf("db.Close failed : %v", err)
	}

	c.JSON(http.StatusOK, users)
	return
}
*/

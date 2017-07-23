package controller

import (
	"fmt"
	"github.com/fberrez/forum/model"
	"github.com/fberrez/forum/shared/passhash"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var user model.User

	if missingParam("pseudo", c) || missingParam("password", c) || missingParam("email", c) {
		return
	}

	hashedPassword, err := passhash.HashString(c.PostForm("password"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"executed": false, "message": "Password cannot be hashed"})
		return
	}

	user.Pseudo = c.PostForm("pseudo")
	user.Password = hashedPassword
	user.Email = c.PostForm("email")
	user.Ip = c.ClientIP()

	fmt.Println("Password : " + c.PostForm("password") + "\nHashed Password : " + user.Password + "\nValidity : " + strconv.FormatBool(passhash.MatchString(user.Password, c.PostForm("password"))))

	err = model.CreateUser(user)
	if err != nil {
		log.Printf("db.Query error :\n\t-query: CreateUser\n\t-desc: %v\n\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"executed": false, "message": "Cannot add user into database"})
		return
	}

	user, err = model.GetUserByPseudo(user.Pseudo)
	if err != nil {
		log.Printf("db.Query error :\n\t-query: GetUserByPseudo\n\t-desc: %v\n\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"executed": false, "message": "SQL Query Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"executed": true, "message": "User created !"})
}

func EditUser(c *gin.Context) {
	var user model.User

	if missingParam("id", c) || missingParam("pseudo", c) || missingParam("password", c) || missingParam("email", c) || missingParam("karma", c) {
		return
	}

	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"executed": false, "message": "ID must be an int"})
		return
	}

	karma, err := strconv.ParseFloat(c.PostForm("karma"), 32)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"executed": false, "message": "Karma must be an number"})
		return
	}

	user = model.User{Id: id, Pseudo: c.PostForm("pseudo"), Password: c.PostForm("password"), Email: c.PostForm("email"), Karma: karma, Ip: c.ClientIP()}
	_, err = model.EditUser(user)
	if err != nil {
		log.Printf("db.Query error :\n\t-query: EditUser\n\t-desc: %v\n\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"executed": false, "message": "SQL Query Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"executed": true, "message": "User edited !"})
}

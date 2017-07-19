package controller

import (
	"database/sql"
	"github.com/fberrez/forum/model"
	"github.com/fberrez/forum/shared/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const (
	sessLoginAttempt = "login_attempt"
)

func loginAttempt(sess *sessions.Session) {
	if sess.Get(sessLoginAttempt) == nil {
		sess.Get(sessLoginAttempt) = 1
	} else {
		sess.Get(sessLoginAttempt) = sess.Get(sessLoginAttempt).(int) + 1
	}
}

func LoginPOST(c *gin.Context) {
	session := sessions.Default(c)

	if session.Get(sessLoginAttempt) != nil && session.Get[sessLoginAttempt].(int) >= 5 {
		log.Println("Brute force prevented")
		c.JSON(http.StatusOK, gin.H{"executed": true, "message": "No brute force please :)"})
		session.Save()
		return
	}

	if missingParam("login", c) || missingParam("password", c) {
		return
	}

	login := c.PostForm("login")
	password := c.PostForm("password")

	var user model.User
	var err error
	if utils.IsAnEmail(login) {
		user, err := model.GetUserByEmail(login)
	} else {
		user, err := model.GetUserByPseudo(login)
	}

	if err == sql.ErrNoRows {
		loginAttempt(session)
		c.JSON(http.StatusOK, gin.H{"executed": true, "message": "Incorrect IDs"})
	}

}

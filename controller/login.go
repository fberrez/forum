package controller

import (
	"database/sql"
	"fmt"
	"github.com/fberrez/forum/model"
	"github.com/fberrez/forum/shared/passhash"
	"github.com/fberrez/forum/shared/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const (
	sessLoginAttempt = "login_attempt"
)

func loginAttempt(sess sessions.Session) {
	if sess.Get(sessLoginAttempt) == nil {
		sess.Set(sessLoginAttempt, 1)
	} else {
		sess.Set(sessLoginAttempt, sess.Get(sessLoginAttempt).(int)+1)
	}
}

func LoginUser(c *gin.Context) {
	session := sessions.Default(c)

	if session.Get(sessLoginAttempt) != nil && session.Get(sessLoginAttempt).(int) >= 5 {
		log.Println("Brute force prevented")
		c.JSON(http.StatusOK, gin.H{"executed": true, "message": "No brute force please :)"})
		session.Save()
		return
	}

	if missingParam("pseudo", c) || missingParam("password", c) {
		return
	}

	login := c.PostForm("pseudo")
	password := c.PostForm("password")

	user := model.User{}
	var err error
	if utils.IsAnEmail(login) {
		user, err = model.GetUserByEmail(login)
	} else {
		user, err = model.GetUserByPseudo(login)
	}

	if err == sql.ErrNoRows {
		loginAttempt(session)
		c.JSON(http.StatusOK, gin.H{"executed": true, "message": "Incorrect IDs"})
		session.Save()
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"executed": false, "message": "Connection service encountered an error. Please try again later"})
		session.Save()
		return
	}

	if passhash.MatchString(user.Password, password) {
		if user.Group == -1 {
			c.JSON(http.StatusOK, gin.H{"executed": true, "message": "Account is disabled."})
			session.Save()
			return
		} else {
			session.Set("id", user.Id)
			session.Set("pseudo", user.Pseudo)
			session.Set("email", user.Email)
			session.Set("group", user.Group)
			session.Set("karma", user.Karma)
			session.Save()
			log.Println("Successful connection !")
			c.JSON(http.StatusOK, gin.H{"executed": true, "message": "Successful connection !"})
			return
		}
	} else {
		loginAttempt(session)
		c.JSON(http.StatusOK, gin.H{"executed": true, "message": "Password is incorrect - Attemp: " + fmt.Sprintf("%v", session.Get(sessLoginAttempt))})
		session.Save()
		return
	}

}

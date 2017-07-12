package main

import (
	"github.com/fberrez/forum/controller"
	"github.com/fberrez/forum/middleware"
	"github.com/gin-gonic/gin"
)

func getRouter() *gin.Engine {
	router := gin.New()
	router.Use(middleware.CORSMiddleware())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	userRouter := router.Group("/u")
	{
		userRouter.GET("/connect", controller.ConnectUser)
		//userRouter.GET("/disconnect")
	}

	/*profileRouter := userRouter.Group("/profile")
	{
		userRouter.GET("/edit")
		userRouter.GET("/view")
		userRouter.GET("/delete")
	}*/

	return router
}

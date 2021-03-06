package main

import (
	"github.com/fberrez/forum/controller"
	"github.com/fberrez/forum/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func getRouter() *gin.Engine {
	router := gin.New()
	router.Use(middleware.CORSMiddleware())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("forum", store))

	userRouter := router.Group("/u")
	{
		userRouter.POST("/login", controller.LoginUser)
		userRouter.GET("/logout", controller.LogoutUser)
		userRouter.POST("/create", controller.CreateUser)
		userRouter.POST("/edit", controller.EditUser)
	}

	// postRouter := router.Group("/p")
	// {
	// 	postRouter.POST("/create", controller.CreatePost)
	// 	postRouter.POST("/answer/:id", controller.CreateAnswer)
	// }

	categoryRouter := router.Group("/c")
	{
		categoryRouter.GET("/list/", controller.GetCategory)
		categoryRouter.GET("/list/:idCat/", controller.GetSubCategoryByIdCategory)
		categoryRouter.GET("/list/:idCat/:idSubCat", controller.GetPostsByIdSubCategory)
		// categoryRouter.GET("/:idCat/:idSubCat/:idPost/list", controller.GetPostMessages)
	}

	/*profileRouter := userRouter.Group("/profile")
	{
		userRouter.GET("/edit")
		userRouter.GET("/view")
		userRouter.GET("/delete")
	}*/

	return router
}

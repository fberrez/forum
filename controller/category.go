package controller

import (
	"github.com/fberrez/forum/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetCategory(c *gin.Context) {
	categories, err := model.GetCategory()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"executed": false, "message": err})
		log.Fatalf("getCategory error : %v", err)
	}

	c.JSON(http.StatusOK, gin.H{"executed": true, "content": categories})
}

// func GetContent(c *gin.Context) {
// 	if c.Param("idSub") == "" {
// 		getSubCategoryByIdCategory(c.Param("idCat"), c)
// 	} else {
// 		getPostsByIdSubCategory(c.Param("idSub"), c)
// 	}
// }

func GetSubCategoryByIdCategory(c *gin.Context) {
	idCategory, err := strconv.Atoi(c.Param("idCat"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"executed": false, "message": err})
		log.Fatalf("Parse idCategory to int error : %v", err)
	}

	subCategories, err := model.GetSubCategoryByIdCategory(idCategory)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"executed": false, "message": err})
		log.Fatalf("getSubCategory error : %v", err)
	}

	c.JSON(http.StatusOK, gin.H{"executed": true, "content": subCategories})
}

func GetPostsByIdSubCategory(c *gin.Context) {
	idSC, err := strconv.Atoi(c.Param("idSubCat"))
	log.Println(idSC)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"executed": false, "message": err})
		log.Fatalf("Parse idSubCat to int error : %v", err)
	}

	posts, err2 := model.GetPostByIdSubCat(idSC)

	if err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"executed": false, "message": err2})
		log.Fatalf("getPostByIdSubCat error : %v", err2)
	}

	c.JSON(http.StatusOK, gin.H{"executed": true, "content": posts})
}

package route

import (
	"app/main/controller"

	"github.com/gin-gonic/gin"
)

func Init() {
	router := router()
	router.Run(":8000")
}

func router() *gin.Engine {
	router := gin.Default()

	book := router.Group("/book")
	{
		router.LoadHTMLGlob("templates/book/*.html")
		book.GET("/", controller.BookIndex)
		book.GET("/create", controller.BookCreate)
		book.POST("/store", controller.BookStore)
	}

	return router
}
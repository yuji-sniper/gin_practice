package route

import (
	"app/main/controller"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func Init() {
	router = gin.Default()
	setCors()
	route()
	router.Run(":8000")
}

func route() {
	book := router.Group("/book")
	{
		router.LoadHTMLGlob("templates/book/*.html")
		book.GET("/", controller.BookIndex)
		book.GET("/create", controller.BookCreate)
		book.POST("/store", controller.BookStore)
		book.GET("/edit/:id", controller.BookEdit)
		book.POST("/update/:id", controller.BookUpdate)
		book.POST("/delete/:id", controller.BookDelete)
	}
}

func setCors() {
	router.Use(cors.New(cors.Config{
		AllowMethods: []string{
            "POST",
            "GET",
            "OPTIONS",
            "PUT",
            "DELETE",
        },
		AllowHeaders: []string{
            "Access-Control-Allow-Headers",
            "Content-Type",
            "Content-Length",
            "Accept-Encoding",
            "X-CSRF-Token",
            "Authorization",
        },
		AllowOrigins: []string{
            "http://localhost:8000/",
        },
		MaxAge: 24 * time.Hour,
	}))
}
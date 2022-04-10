package route

import (
	"app/main/controller"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func Init() {
	router = gin.Default()
	setCors()
	setSession()
	route()
	router.Run(":8000")
}

// Routing
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

// CORS
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

// Session
func setSession() {
	store, err := redis.NewStore(10, "tcp", "redis:6379", "", []byte("secret"))
	if err != nil {
		log.Fatalln(err.Error(), "セッションの接続に失敗しました")
	}
	router.Use(sessions.Sessions("session", store))
}
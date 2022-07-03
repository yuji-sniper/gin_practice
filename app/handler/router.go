package handler

import (
	"app/main/handler/controller"
	"log"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func Init() {
	router = gin.Default()
	setSession()
	route()
	router.Run(":8000")
}

// ルーティング
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

// セッション接続
func setSession() {
	store, err := redis.NewStore(10, "tcp", "redis:6379", "", []byte("secret"))
	if err != nil {
		log.Fatalln(err.Error(), "セッションの接続に失敗しました")
		return
	}
	router.Use(sessions.Sessions("session", store))
}

package controller

import (
	"app/main/model"
	"app/main/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

var bookRepository repository.BookRepository

func init() {
	bookRepository = repository.BookRepository{}
}

// Book一覧
func BookIndex(ctx *gin.Context) {
	books := bookRepository.FetchBooks()
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"books": books,
	})
}

// Book作成画面
func BookCreate(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "create.html", gin.H{})
} 

// Book作成処理
func BookStore(ctx *gin.Context) {
	book := model.Book{}
	if err := ctx.Bind(&book); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	bookRepository.CreateBook(&book)

	ctx.Redirect(http.StatusFound, "/book/")
}

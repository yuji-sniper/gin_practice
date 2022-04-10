package controller

import (
	"app/main/model"
	"app/main/repository"
	"net/http"
	"strconv"

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

	ctx.Redirect(http.StatusFound, "/book")
}

// Book編集画面
func BookEdit(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.String(http.StatusBadRequest, "idが数字ではありません!")
		return
	}
	book := bookRepository.FindBook(id)
	ctx.HTML(http.StatusOK, "edit.html", gin.H{
		"book": book,
	})
}

// Book更新処理
func BookUpdate(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.String(http.StatusBadRequest, "idが数字ではありません!")
		return
	}
	book := model.Book{}
	if err := ctx.Bind(&book); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}
	bookRepository.UpdateBook(id, &book)

	ctx.Redirect(http.StatusFound, "/book")
}

// Book削除
func BookDelete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.String(http.StatusBadRequest, "idが数字ではありません!")
		return
	}
	bookRepository.DeleteBook(id)

	ctx.Redirect(http.StatusFound, "/book")
}

package controller

import (
	"app/main/model"
	"app/main/module"
	"app/main/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var bookRepository repository.BookRepository

var sessionModule module.SessionModule

func init() {
	bookRepository = repository.BookRepository{}

	sessionModule = module.SessionModule{}
}

// Book一覧
func BookIndex(ctx *gin.Context) {
	books := bookRepository.FetchBooks()

	flashMessage := sessionModule.GetFlashMessage(ctx)

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"books": books,
		"flashMessage": flashMessage,
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

	sessionModule.SetFlashMessage(ctx, "作成しました!")

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

	sessionModule.SetFlashMessage(ctx, "更新しました!")

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

	sessionModule.SetFlashMessage(ctx, "削除しました!")

	ctx.Redirect(http.StatusFound, "/book")
}

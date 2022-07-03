package controller

import (
	"app/main/domain/model"
	"app/main/domain/repository"
	"app/main/form"
	"app/main/infra"
	"app/main/module"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var bookRepository repository.BookContract
var sessionModule module.SessionModule

func init() {
	bookRepository = infra.BookRepository{}
	sessionModule = module.SessionModule{}
}

// Book一覧
func BookIndex(ctx *gin.Context) {
	books := bookRepository.FetchBooks()
	flashMessage := sessionModule.Pull(ctx, "flash_message")
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"books": books,
		"flashMessage": flashMessage,
	})
}

// Book作成画面
func BookCreate(ctx *gin.Context) {
	errorMessages := sessionModule.Pull(ctx, "_error_messages")
	old := sessionModule.Pull(ctx, "_old")
	ctx.HTML(http.StatusOK, "create.html", gin.H{
		"errorMessages": errorMessages,
		"old": old,
	})
}

// Book作成
func BookStore(ctx *gin.Context) {
	if err := form.Validate(form.BookForm{}, ctx); err != nil {
		ctx.Redirect(http.StatusFound, "/book/create")
		return
	}
	book := model.Book{}
	ctx.Bind(&book)
	fmt.Println(book)
	bookRepository.CreateBook(&book)
	sessionModule.Set(ctx, "flash_message", "作成しました！")
	ctx.Redirect(http.StatusFound, "/book")
}

// Book編集画面
func BookEdit(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.String(http.StatusBadRequest, "IDが数字ではありません!")
		return
	}
	book := bookRepository.FindBook(id)
	errorMessages := sessionModule.Pull(ctx, "_error_messages")
	old := sessionModule.Pull(ctx, "_old")
	ctx.HTML(http.StatusOK, "edit.html", gin.H{
		"book": book,
		"errormessages": errorMessages,
		"old": old,
	})
}

// Book更新
func BookUpdate(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.String(http.StatusBadRequest, "IDが数字ではありません!")
		return
	}
	if err := form.Validate(form.BookForm{}, ctx); err != nil {
		ctx.Redirect(http.StatusFound, fmt.Sprintf("/book/edit/%d", id))
		return
	}
	book := model.Book{}
	ctx.Bind(&book)
	bookRepository.UpdateBook(id, &book)
	sessionModule.Set(ctx, "flash_message", "更新しました！")
	ctx.Redirect(http.StatusFound, "/book")
}

// Book削除
func BookDelete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.String(http.StatusBadRequest, "IDが数字ではありません!")
		return
	}
	bookRepository.DeleteBook(id)
	sessionModule.Set(ctx, "flash_message", "削除しました!")
	ctx.Redirect(http.StatusFound, "/book")
}

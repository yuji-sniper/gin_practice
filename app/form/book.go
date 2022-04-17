package form

import "github.com/gin-gonic/gin"

type BookForm struct {
	Title string `form:"title" binding:"required,min=3"`
	Content string `form:"content" binding:"required,min=5"`
}

func (BookForm) Bind(ctx *gin.Context) error {
	return ctx.ShouldBind(&BookForm{})
}

func (BookForm) Messages() map[string]map[string]string {
	return map[string]map[string]string{
		"Title": {
			"required": "タイトルを入力してください",
			"min": "タイトルは3文字以上で入力してください",
		},
		"Content": {
			"required": "内容を入力してください",
			"min": "内容は5文字以上で入力してください",
		},
	}
}

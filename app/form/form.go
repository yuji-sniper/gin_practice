package form

import (
	"app/main/module"
	"encoding/gob"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var sessionModule module.SessionModule

type Form interface {
	Bind(ctx *gin.Context) error 
	Messages() map[string]map[string]string
}

func init() {
	sessionModule = module.SessionModule{}
	gob.Register(gin.H{})
}

// バリデート
func Validate(form Form, ctx *gin.Context) error {
	err := form.Bind(ctx)
	if err != nil {
		allMessages := form.Messages()
		errorMessages := gin.H{}
		if validationErrors, exists := err.(validator.ValidationErrors); exists {
			for _, e := range validationErrors {
				if msg, exists := allMessages[e.Field()][e.Tag()]; exists {
					errorMessages[e.Field()] = msg
				}
			}
		}
		setSessionMessages(ctx, errorMessages)
		setSessionOld(ctx)
	}
	return err
}

// バリデーションメッセージをセッションに保存してリダイレクト
func setSessionMessages(ctx *gin.Context, messages gin.H) {
	sessionModule.Set(ctx, "_error_messages", messages)
}

// old値をセッションに保存
func setSessionOld(ctx *gin.Context) {
	olds := gin.H{}
	for key, _ := range ctx.Request.Form {
		olds[key] = ctx.PostForm(key)
	}
	sessionModule.Set(ctx, "_old", olds)
}

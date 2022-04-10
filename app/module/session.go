package module

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type SessionModule struct{}

// フラッシュメッセージをセッションに保存
func (SessionModule) SetFlashMessage(ctx *gin.Context, message string) {
	session := sessions.Default(ctx)
	session.Set("flash_message", message)
	if err := session.Save(); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error(), "セッションのセーブに失敗しました!")
		return
	}
}

// フラッシュメッセージをセッションから取得
func (SessionModule) GetFlashMessage(ctx *gin.Context) interface{} {
	session := sessions.Default(ctx)
	flashMessage := session.Get("flash_message")
	session.Delete("flash_message")
	if err := session.Save(); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error(), "セッションのセーブに失敗しました!")
	}
	if flashMessage != nil {
		return flashMessage.(string)
	}
	return flashMessage
}

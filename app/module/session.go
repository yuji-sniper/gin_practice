package module

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type SessionModule struct{}

// セッションに保存
func (SessionModule) Set(ctx *gin.Context, key string, value interface{}) {
	session := sessions.Default(ctx)
	session.Set(key, value)
	if err := session.Save(); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error(), "セッションのセーブに失敗しました!")
	}
}

// セッションから値を取得して削除
func (SessionModule) Pull(ctx *gin.Context, key string) interface{} {
	session := sessions.Default(ctx)
	value := session.Get(key)
	session.Delete(key)
	if err := session.Save(); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error(), "セッションエラーが発生しました!")
	}
	return value
}

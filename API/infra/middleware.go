package infra

import (
	"log"

	"github.com/gin-gonic/gin"
)

const (
	userIDKey string = "userID"
)

// 認証ミドルウェア
func authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// cookieからuseridを取得
		cookie, err := ctx.Request.Cookie("userID")
		if err != nil {
			log.Println(err)
			return
		}
		// コンテキストにuserIDを放り込む
		SetUserID(ctx, cookie.Value)
		ctx.Next()
	}
}

// コンテキストにuseridを格納
func SetUserID(ctx *gin.Context, userID string) {
	ctx.Set(userIDKey, userID)
}

// コンテキストからuseridを取得
func GetUserIDFromContext(ctx *gin.Context) string {
	var userID string
	if ctx.Value(userIDKey) != nil {
		userID = ctx.Value(userIDKey).(string)
	}
	return userID
}

package controllers

type Context interface {
	Param(string) string
	Bind(interface{}) error
	Status(int)
	JSON(int, interface{})
	Value(interface{}) interface{}
	Set(string, interface{})
	//GET
	DefaultQuery(key, defaultValue string) string
	Query(string) string
}

const (
	userIDKey string = "userID"
)

// コンテキストからuseridを取得
func GetUserIDFromContext(ctx Context) string {
	var userID string
	if ctx.Value(userIDKey) != nil {
		userID = ctx.Value(userIDKey).(string)
	}
	return userID
}

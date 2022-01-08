package controllers

type Context interface {
	Param(string) string
	Bind(interface{}) error
	Status(int)
	JSON(int, interface{})
	//GET
	DefaultQuery(key, defaultValue string) string
	Query(string) string
}

package interfaces

import "github.com/gin-gonic/gin"

type Route struct {
	Method  HttpMethod
	Path    string
	Handler func(c *gin.Context)
}

var Routes = []Route{}

func init() {
	accountController := NewAccountController()

	Routes = append(Routes, Route{
		Method:  "GET",
		Path:    "/ping",
		Handler: accountController.Ping,
	})
}

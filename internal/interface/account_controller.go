package interfaces

import "github.com/gin-gonic/gin"

type AccountController struct {
}

func NewAccountController() *AccountController {
	return &AccountController{}
}

func (controller *AccountController) Ping(c *gin.Context) {
	c.String(200, "pong")
}

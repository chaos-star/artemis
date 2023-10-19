package Controller

import (
	"Artemis/Common/Response"
	"github.com/gin-gonic/gin"
)

type User struct {
}

var UserController = new(User)

func (u *User) Example(c *gin.Context) {
	Response.Success(c)
}

package Group

import (
	"Artemis/App/Controller"
	"github.com/gin-gonic/gin"
)

type User struct {
}

func (u *User) Router(group *gin.RouterGroup) {
	x := group.Group("user")
	x.POST("ping", Controller.UserController.Example)
}

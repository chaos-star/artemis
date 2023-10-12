package Group

import (
	"Artemis/Common/Response"
	"github.com/gin-gonic/gin"
)

type User struct {
}

func (u *User) Router(group *gin.RouterGroup) {
	x := group.Group("user")
	x.POST("ping", func(c *gin.Context) {
		Response.Success(c)
	})
}

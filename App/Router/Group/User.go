package Group

import (
	"Artemis/App/Middleware"
	"Artemis/Common/Response"
	"github.com/gin-gonic/gin"
)

type User struct {
}

func (u *User) Router(group *gin.RouterGroup) {
	x := group.Group("user").Use(Middleware.ResponseRecord())
	x.GET("ping", func(c *gin.Context) {
		Response.Success(c)
	})
}

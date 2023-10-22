package Group

import (
	"Artemis/App/Controller/UserController"
	"github.com/gin-gonic/gin"
)

type User struct {
}

func (u *User) Router(group *gin.RouterGroup) {
	router := group.Group("user")
	controller := UserController.Instance
	router.POST("example", controller.Example)
}

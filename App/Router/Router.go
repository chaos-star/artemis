package Router

import (
	"Artemis/App/Router/Group"
	"github.com/chaos-star/marvel"
	"github.com/gin-gonic/gin"
)

type IGroupRouter interface {
	Router(group *gin.RouterGroup)
}

var (
	DenyAuthentication = []IGroupRouter{}

	Authentication = []IGroupRouter{
		&Group.User{},
	}
)

func init() {
	//不授权路由
	denyRouter := marvel.Web.Group("")
	{
		for _, item := range DenyAuthentication {
			item.Router(denyRouter)
		}
	}

	//授权路由
	authRouter := marvel.Web.Group("")
	{
		for _, item := range Authentication {
			item.Router(authRouter)
		}
	}

}

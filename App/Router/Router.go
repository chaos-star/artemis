package Router

import (
	"Artemis/App/Middleware"
	"Artemis/App/Router/Group"
	"github.com/chaos-star/marvel"
	"github.com/chaos-star/marvel/Web"
)

var (
	Default = []Web.IGroupRouter{
		new(Group.User),
	}
)

func init() {
	marvel.Web.LoadRouterGroup("", Default, Middleware.TraceLog())
}

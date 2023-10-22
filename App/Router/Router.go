package Router

import (
	"Artemis/App/Middleware"
	"github.com/chaos-star/marvel"
)

func init() {
	marvel.Web.LoadRouterGroup("", Default, Middleware.TraceLog(), Middleware.Translate())
}

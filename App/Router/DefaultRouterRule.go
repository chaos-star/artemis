package Router

import (
	"Artemis/App/Router/Group"
	"github.com/chaos-star/marvel/Web"
)

var Default = []Web.IGroupRouter{
	new(Group.User),
}

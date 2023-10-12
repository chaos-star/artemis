package Rpc

import (
	"github.com/chaos-star/marvel"
	"github.com/chaos-star/marvel/Server"
)

// {
// Sd:   &proto.Mailer_ServiceDesc,
// Ss:   service.NewEmailServer(),
// Name: "EmailServer",
// },
var Services = []Server.ServiceOption{}

func init() {
	marvel.Server.Load(Services)
}

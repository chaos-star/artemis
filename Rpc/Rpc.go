package Rpc

import (
	"github.com/chaos-star/marvel"
	"github.com/chaos-star/marvel/Server"
)

var Services = []Server.ServiceOption{
	// {
	// Sd:   &proto.Mailer_ServiceDesc,
	// Ss:   service.NewEmailServer(),
	// Name: "EmailServer",
	// },
}

func init() {
	marvel.Server.Load(Services)
}

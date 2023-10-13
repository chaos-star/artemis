//go:build web

package main

import (
	_ "Artemis/App/Router"
	"github.com/chaos-star/marvel"
)

func main() {
	marvel.Web.RunServer()
}

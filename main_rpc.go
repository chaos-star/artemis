//go:build rpc

package main

import _ "Artemis/Rpc"

func main() {
	marvel.Logger.Info("Rpc Running")
	err := marvel.Server.Run()
	if err != nil {
		panic(err)
	}
}

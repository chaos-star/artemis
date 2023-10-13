//go:build console

package main

import (
	"Artemis/Console"
	"flag"
	"github.com/chaos-star/marvel"
)

func main() {
	flag.Parse()

	CommandName := Console.CommandName
	if CommandName == "start" {
		marvel.Cron.Run()
	} else {
		if CommandName == "" {
			tips := []string{
				"\033[1m-name command\033[0m execute single job.",
				"tips: -name start 开启定时任务",
			}
			marvel.Cron.Usage(tips...)
		}
		Command := marvel.Cron.GetJob(CommandName)
		if Command != nil {
			Command.Cmd.Run()
		} else {

		}
	}

}

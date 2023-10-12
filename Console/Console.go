package Console

import (
	"Artemis/Console/Command"
	"flag"
	"github.com/chaos-star/marvel"
	"github.com/chaos-star/marvel/CronJob"
)

var SpecCron = map[string]CronJob.SpecJob{
	"example": {"@every 1s", new(Command.Example), "test example"},
}
var CommandName string

func init() {
	flag.StringVar(&CommandName, "name", "", "command name")
	marvel.Cron.JoinJobs(SpecCron)
}

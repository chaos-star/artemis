package Console

import (
	"Artemis/Console/Command"
	"github.com/chaos-star/marvel"
	"github.com/chaos-star/marvel/CronJob"
)

var SpecCron = map[string]CronJob.SpecJob{
	"example": {"@every 1s", new(Command.Example)},
}

func init() {
	marvel.Cron.JoinJobs(SpecCron)
}

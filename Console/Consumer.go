package Console

import (
	"Artemis/Console/Consumer"
	mate "github.com/chaos-star/queue-mate"
)

var Consumers = []mate.MQConsume{
	new(Consumer.Example),
}

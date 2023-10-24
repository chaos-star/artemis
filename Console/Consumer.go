package Console

import (
	"Artemis/Console/Consumer/ExampleConsumer"
	mate "github.com/chaos-star/queue-mate"
)

var Consumers = []mate.MQConsume{
	new(ExampleConsumer.ExampleConsumer),
}

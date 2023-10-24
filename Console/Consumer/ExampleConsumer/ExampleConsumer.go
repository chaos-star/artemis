package ExampleConsumer

import (
	"fmt"
	"github.com/chaos-star/marvel"
	mate "github.com/chaos-star/queue-mate"
)

type ExampleConsumer struct {
}

func (e *ExampleConsumer) GetOptions() []mate.Option {
	var options []mate.Option
	return options
}
func (e *ExampleConsumer) RunConsume(option mate.Option) (err error) {

	client := marvel.MQ.Instance("mq_alias").NewClient()
	//延时消费
	//err = client.Use(e,option).Retry(3).DelayReceive(client.Fanout, "ex_artemis_test", nil, "queue_artemis_test")
	err = client.Use(e, option).Retry(3).Receive(client.Fanout, "ex_artemis_test", nil, "queue_artemis_test")
	return
}

func (e *ExampleConsumer) Process(body []byte) (err error) {
	fmt.Println(fmt.Sprintf("message:%s", string(body)))
	return
}

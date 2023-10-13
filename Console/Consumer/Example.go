package Consumer

import (
	"fmt"
	"github.com/chaos-star/marvel"
)

type Example struct {
}

func (e *Example) RunConsume() (err error) {
	client := marvel.MQ.Instance("super").NewClient()
	//延时消费
	//err = client.Use(e).Retry(3).DelayReceive(client.Fanout, "ex_artemis_test", nil, "queue_artemis_test")
	err = client.Use(e).Retry(3).Receive(client.Fanout, "ex_artemis_test", nil, "queue_artemis_test")
	return
}

func (e *Example) Process(body []byte) (err error) {
	fmt.Println(fmt.Sprintf("message:%s", string(body)))
	return
}

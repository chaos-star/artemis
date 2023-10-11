package Command

import "fmt"

type Example struct {
}

func (t *Example) Run() {
	fmt.Println("@every 1s example")
}

package abstract_factory

import "fmt"

type IMouse interface {
	GetColor()
}

func GetMouseDetails(mouse abstract_factory.IMouse) {
	fmt.Println(mouse.GetColor())
}

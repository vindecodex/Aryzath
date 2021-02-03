package abstract_factory

import "fmt"

type IMouse interface {
	GetColor() string
}

func GetMouseDetails(mouse IMouse) {
	fmt.Println(mouse.GetColor())
}

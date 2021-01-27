package device

import "fmt"

type Light struct{}

func (l Light) On() {
	fmt.Println("ON")
}

func (l Light) Off() {
	fmt.Println("OFF")
}

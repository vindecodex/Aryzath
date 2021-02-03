package main

import (
	"Aryzath/abstract_factory/abstract_factory"
)

func main() {
	razer, _ := abstract_factory.GetBrand("razer")
	rmouse := razer.CreateMouse()
	rkeyboard := razer.CreateKeyboard()

	abstract_factory.GetMouseDetails(rmouse)
	abstract_factory.GetKeyboardDetails(rkeyboard)
}

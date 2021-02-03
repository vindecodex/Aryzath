package main

import (
	"Aryzath/abstract_factory/abstract_factory"
)

func main() {
	razer, _ := abstract_factory.GetBrand("razer")
	steelseries, _ := abstract_factory.GetBrand("steelseries")

	rmouse := razer.CreateMouse()
	rkeyboard := razer.CreateKeyboard()
	smouse := steelseries.CreateMouse()
	skeyboard := steelseries.CreateKeyboard()

	abstract_factory.GetMouseDetails(rmouse)
	abstract_factory.GetKeyboardDetails(rkeyboard)

	abstract_factory.GetMouseDetails(smouse)
	abstract_factory.GetKeyboardDetails(skeyboard)
}

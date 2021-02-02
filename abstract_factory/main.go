package main

import (
	"Aryzath/abstract_factory/abstract_factory"
	"fmt"
)

func main() {
	razer, _ := abstract_factory.GetBrand("razer")
	razerMouse := razer.CreateMouse()
	razerKeyboard := razer.CreateKeyboard()

	steelseries, _ := abstract_factory.GetBrand("steelseries")
	steelseriesMouse := steelseries.CreateMouse()
	steelseriesKeyboard := steelseries.CreateKeyboard()

	GetMouseDetails(razerMouse)
	GetKeyboardDetails(razerKeyboard)

	GetMouseDetails(steelseriesMouse)
	GetKeyboardDetails(steelseriesKeyboard)

}

func GetMouseDetails(mouse abstract_factory.IMouse) {
	fmt.Println("Color: ", mouse.GetColor)
}

func GetKeyboardDetails(keyboard abstract_factory.IKeyboard) {
	fmt.Println("Color: ", keyboard.GetColor)
}

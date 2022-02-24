package main

import "github.com/vindecodex/Aryzath/abstract_factory/abstract_factory"

func main() {
	// Talks to the factory which returns an interface
	razer, _ := abstract_factory.GetBrand("razer")
	steelseries, _ := abstract_factory.GetBrand("steelseries")

	// Created the Object by interacting with the interface
	rmouse := razer.CreateMouse()
	rkeyboard := razer.CreateKeyboard()
	smouse := steelseries.CreateMouse()
	skeyboard := steelseries.CreateKeyboard()

	// Get Information through factory method that accepts interface as argument
	abstract_factory.GetMouseDetails(rmouse)
	abstract_factory.GetKeyboardDetails(rkeyboard)

	// Get Information through factory method that accepts interface as argument
	abstract_factory.GetMouseDetails(smouse)
	abstract_factory.GetKeyboardDetails(skeyboard)
}

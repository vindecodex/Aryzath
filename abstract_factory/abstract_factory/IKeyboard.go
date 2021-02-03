package abstract_factory

import "fmt"

type IKeyboard interface {
	GetColor() string
}

func GetKeyboardDetails(keyboard IKeyboard) {
	fmt.Println(keyboard.GetColor())
}

package abstract_factory

import "fmt"

type IKeyboard interface {
	GetColor()
}

func GetKeyboardDetails(keyboard abstract_factory.IKeyboard) {
	fmt.Println(keyboard.GetColor())
}

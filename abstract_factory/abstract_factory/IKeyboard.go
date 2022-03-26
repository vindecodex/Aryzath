package abstract_factory

type IKeyboard interface {
	GetColor() string
}

func GetKeyboardDetails(keyboard IKeyboard) string {
	return keyboard.GetColor()
}

package template

type Template interface {
	InsertCard()
	InputPassword(password int)
	Withdraw(amount int)
	RemoveCard()
}

func Withdraw(t Template, password int, amount int) {
	t.InsertCard()
	t.InputPassword(password)
	t.Withdraw(amount)
	t.RemoveCard()
}

package template

type Person struct {
	Password       int
	WithdrawAmount int
}

func NewPerson(password, withdrawAmount int) *Person {
	return &Person{Password: password, WithdrawAmount: withdrawAmount}
}

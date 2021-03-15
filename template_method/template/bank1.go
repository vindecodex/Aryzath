package template

import "fmt"

type Bank1 struct {
	Person *Person
}

func NewBank1(person *Person) *Bank1 {
	return &Bank1{Person: person}
}

func (b *Bank1) InsertCard() {
	fmt.Println("Welcome from Bank 1")
}

func (b *Bank1) InputPassword(password int) {
	if password == b.Person.Password {
		fmt.Println("Authenticated by Bank1 Secure Algorithm")
	}
}

func (b *Bank1) Withdraw(amount int) {
	fmt.Printf("Withdrawing with amount of %d \n", amount)
}

func (b *Bank1) RemoveCard() {
	fmt.Println("Removing Card \nThank you!")
}

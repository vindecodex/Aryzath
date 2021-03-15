package template

import "fmt"

type Bank2 struct {
	Person *Person
}

func NewBank2(person *Person) *Bank2 {
	return &Bank2{Person: person}
}

func (b *Bank2) InsertCard() {
	fmt.Println("Welcome from Bank 2")
}

func (b *Bank2) InputPassword(password int) {
	if password == b.Person.Password {
		fmt.Println("Authenticated by Bank2 Secure Algorithm")
	}
}

func (b *Bank2) Withdraw(amount int) {
	fmt.Printf("Withdrawing with amount of %d \n", amount)
}

func (b *Bank2) RemoveCard() {
	fmt.Println("Removing Card \nThank you!")
}

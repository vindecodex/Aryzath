package main

import (
	"Aryzath/template_method/template"
	"fmt"
)

func main() {
	person := template.NewPerson(1234, 5000)

	bank1 := template.NewBank1(person)
	bank2 := template.NewBank2(person)

	template.Withdraw(bank1, person.Password, person.WithdrawAmount)
	fmt.Println("----------------")
	template.Withdraw(bank2, person.Password, person.WithdrawAmount)
}

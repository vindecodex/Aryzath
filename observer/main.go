package main

import "github.com/vindecodex/Aryzath/observer/observer"

func main() {
	customerA := observer.NewCustomer(1, "John")
	customerB := observer.NewCustomer(2, "Mark")

	store := observer.NewStore("PS5")

	store.Register(customerA)
	store.Register(customerB)
	store.Notify()

	customerC := observer.NewCustomer(3, "Andrie")
	store.Register(customerC)
	store.Unregister(customerB.GetId())
	store.Notify()
}

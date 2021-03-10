package observer

import "fmt"

type Customer struct {
	ID   int
	Name string
}

func NewCustomer(id int, name string) *Customer {
	return &Customer{id, name}
}

func (c *Customer) GetNotified(product string) {
	fmt.Printf("%s got notified for new product: %s \n", c.Name, product)
}

func (c *Customer) GetId() int {
	return c.ID
}

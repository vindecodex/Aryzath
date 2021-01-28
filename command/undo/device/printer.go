package device

import "fmt"

type Printer struct {
	Content string
}

func (p Printer) Print() {
	fmt.Println(p.Content)
}

package composite

import "fmt"

type Product struct {
	Name string
}

func (p *Product) Scan() {
	fmt.Println(p.Name)
}

package composite

import "fmt"

type Box struct {
	component []Component
	Name      string
}

func (b *Box) Scan() {
	fmt.Println("Scanning from " + b.Name)
	for _, compos := range b.component {
		fmt.Print("-")
		compos.Scan()
	}
	fmt.Println("-------------")
}

func (b *Box) Add(c Component) {
	b.component = append(b.component, c)
}

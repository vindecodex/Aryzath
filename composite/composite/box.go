package composite

import "fmt"

type Box struct {
	Component []Component
	Name      string
}

func (b *Box) Scan() {
	fmt.Println("Scanning from " + b.Name)
	for _, compos := range b.Component {
		fmt.Print("-")
		compos.Scan()
	}
	fmt.Println("-------------")
}

func (b *Box) Add(c Component) {
	b.Component = append(b.Component, c)
}

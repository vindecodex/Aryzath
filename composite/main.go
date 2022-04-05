package main

import (
	"github.com/vindecodex/Aryzath/composite/composite"
)

func main() {
	phone := &composite.Product{Name: "IPhone"}

	keyboard := &composite.Product{Name: "Keyboard"}
	mouse := &composite.Product{Name: "Mouse"}

	oil := &composite.Product{Name: "Oil"}

	box1 := &composite.Box{Name: "box1"}
	box1.Add(phone)

	box2 := &composite.Box{Name: "box2"}
	box2.Add(keyboard)
	box2.Add(mouse)
	box2.Add(box1)

	box3 := &composite.Box{Name: "box3"}
	box3.Add(oil)

	box4 := &composite.Box{Name: "parent box"}
	box4.Add(box1)
	box4.Add(box2)
	box4.Add(box3)

	box4.Scan()
}

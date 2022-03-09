package main

import (
	"fmt"

	"github.com/vindecodex/Aryzath/iterator/iterator"
)

func main() {
	data_a := &iterator.Data{}
	data_b := &iterator.Data{}

	data_a.SetData("ABC")
	data_b.SetData("XYZ")

	collections := iterator.DataCollection{Data: []*iterator.Data{data_a, data_b}}

	iterator := collections.CreateIterator()

	for iterator.HasNext() {
		data := iterator.GetNext()
		fmt.Printf("Data: $%s \n", data.GetData())
	}
}

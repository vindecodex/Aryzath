package main

import (
	"fmt"

	"github.com/vindecodex/Aryzath/bridge/abstraction_heirarchy"
	"github.com/vindecodex/Aryzath/bridge/implementation_heirarchy"
)

func main() {
	eclipse := implementation_heirarchy.Eclipse{}
	netbeans := implementation_heirarchy.Netbeans{}

	java_file := &abstraction_heirarchy.Java{}
	cpp_file := &abstraction_heirarchy.Cpp{}

	java_file.SetIde(eclipse)
	fmt.Println(java_file.Compile())

	java_file.SetIde(netbeans)
	fmt.Println(java_file.Compile())

	cpp_file.SetIde(eclipse)
	fmt.Println(cpp_file.Compile())

	cpp_file.SetIde(netbeans)
	fmt.Println(cpp_file.Compile())

}

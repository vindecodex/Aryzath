package main

import (
	"Aryzath/bridge/abstraction_heirarchy"
	"Aryzath/bridge/implementation_heirarchy"
)

func main() {
	eclipse := implementation_heirarchy.Eclipse{}
	netbeans := implementation_heirarchy.Netbeans{}

	java_file := &abstraction_heirarchy.Java{}
	cpp_file := &abstraction_heirarchy.Cpp{}

	java_file.SetIde(eclipse)
	java_file.Compile()

	java_file.SetIde(netbeans)
	java_file.Compile()

	cpp_file.SetIde(eclipse)
	cpp_file.Compile()

	cpp_file.SetIde(netbeans)
	cpp_file.Compile()

}

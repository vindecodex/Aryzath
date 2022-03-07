package main

import (
	"fmt"

	"github.com/vindecodex/Aryzath/factory_method/tool"
)

func main() {
	// now our main file is not tightly coupled with our tool
	tool := tool.SetTool()          // if we set it to switch then we can pass what tool we are going to use
	fmt.Println(tool.GetToolName()) // Prints what tool we are using
}

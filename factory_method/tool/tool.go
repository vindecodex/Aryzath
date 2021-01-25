package tool

type ITool interface {
	GetToolName() string
}

/*

If we plan to change from Pencil to Pen then we can just change the return type to PEN

Or We can just make a switch statement inside GetToolName(toolName) that accepts toolName
and will return a tool

Any type that will be used as return in SetTool Method should implement the ITool methods.

*/
func SetTool() ITool { // Also known as the factory method
	return Pencil{}
}

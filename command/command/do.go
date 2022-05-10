package command

type Do struct {
	c Command
}

func (d *Do) SetCommand(c Command) { // Accepts Class that implements the Command interface
	d.c = c
}

func (d *Do) It() string {
	return d.c.Execute() // Execute the method whatever the class was passed
}

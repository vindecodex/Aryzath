package command

type Do struct {
	c []Command
}

func (d *Do) SetCommand(c Command) {
	d.c = append(d.c, c)
}

func (d *Do) It() {
	d.c[len(d.c)-1].Execute()
}

func (d *Do) Undo() {
	newCommand := d.c[:len(d.c)-1]
	d.c = newCommand
	d.It()
}

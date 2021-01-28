package device

type PrinterCommand struct {
	P Printer
}

func (p PrinterCommand) Execute() {
	p.P.Print()
}

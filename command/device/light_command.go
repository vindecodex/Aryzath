package device

type LightOn struct { // focus on On method of Light Class
	l Light
}

type LightOff struct { // focus on Off method of Light Class
	l Light
}

func (lon *LightOn) Execute() { // Implements the command interface
	lon.l.On() // Execute the On method of Light Class
}

func (lof *LightOff) Execute() { // Implements the command interface
	lof.l.Off() // Execute the Off method of Light Class
}

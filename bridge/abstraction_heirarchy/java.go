package abstraction_heirarchy

import (
	"Aryzath/bridge/implementation_heirarchy"
	"fmt"
)

type Java struct {
	ide implementation_heirarchy.Ide
}

// Injecting the implementation
func (j *Java) SetIde(ide implementation_heirarchy.Ide) {
	j.ide = ide
}

// The implementation of IDE was executed here
func (j *Java) Compile() {
	fmt.Println("Java Compiling using " + j.ide.CompileLanguage())
}

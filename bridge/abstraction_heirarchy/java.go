package abstraction_heirarchy

import (
	"github.com/vindecodex/Aryzath/bridge/implementation_heirarchy"
)

type Java struct {
	ide implementation_heirarchy.Ide
}

// Injecting the implementation
func (j *Java) SetIde(ide implementation_heirarchy.Ide) {
	j.ide = ide
}

// The implementation of IDE was executed here
func (j *Java) Compile() string {
	return "Java Compiling using " + j.ide.CompileLanguage()
}

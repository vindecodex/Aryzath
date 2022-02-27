package abstraction_heirarchy

import (
	"fmt"

	"github.com/vindecodex/Aryzath/bridge/implementation_heirarchy"
)

type Cpp struct {
	ide implementation_heirarchy.Ide
}

// Injecting the implementation
func (c *Cpp) SetIde(ide implementation_heirarchy.Ide) {
	c.ide = ide
}

// The implementation of IDE was executed here
func (c *Cpp) Compile() {
	fmt.Println("Cpp Compiling using " + c.ide.CompileLanguage())
}

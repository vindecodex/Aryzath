package abstraction_heirarchy

import "github.com/vindecodex/Aryzath/bridge/implementation_heirarchy"

type ProgrammingLanguage interface {
	Compile()
	SetIde(implementation_heirarchy.Ide)
}

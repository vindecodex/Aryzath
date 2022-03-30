package abstraction_heirarchy

import "github.com/vindecodex/Aryzath/bridge/implementation_heirarchy"

type ProgrammingLanguage interface {
	Compile() string
	SetIde(implementation_heirarchy.Ide)
}

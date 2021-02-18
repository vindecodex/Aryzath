package abstraction_heirarchy

import "Aryzath/bridge/implementation_heirarchy"

type ProgrammingLanguage interface {
	Compile()
	SetIde(implementation_heirarchy.Ide)
}

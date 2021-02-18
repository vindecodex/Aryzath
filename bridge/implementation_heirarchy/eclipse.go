package implementation_heirarchy

type Eclipse struct{}

// The concrete implementation
func (e Eclipse) CompileLanguage() string {
	return "Eclipse"
}

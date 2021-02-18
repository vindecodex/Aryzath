package implementation_heirarchy

type Netbeans struct{}

// The concrete implementation
func (e Netbeans) CompileLanguage() string {
	return "Netbeans"
}

package model

// Galinha representa uma ave do tipo galinha
type Galinha interface {
	Cacareja() string
}

// Pata representa uma ave do tipo pato
type Pata interface {
	Quack() string
}

//Ave Representa um animal
type Ave struct {
	Nome string
}

// Cacareja retorna um som emitido por galinha
func (a Ave) Cacareja() string {
	return "Cocoricó..."
}

// Quack retorna um som emitido por uma pata
func (a Ave) Quack() string {
	return "Quack, quack..."
}

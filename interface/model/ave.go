package model

// Galinha representa uma ave do tipo galinha
type Galinha interface {
	Carcareja() string
}

// Pato representa uma ave do tipo Pato
type Pato interface {
	Quack() string
}

//Ave struct
type Ave struct {
	Nome string
}

//Carcareja retorna o som feito por uma galinha
func (a Ave) Carcareja() string {
	return "Cocoricó!"
}

//Quack retorna o som feito por uma galinha
func (a Ave) Quack() string {
	return "Quack!!!"
}

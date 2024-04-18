package interfaces

type Humano interface {
	Respirar()
	Pensar()
	Comer()
	Sexo() string
}

type EsParlante interface {
	Hablar()
}

type Perro struct {
}

func (p Perro) Respirar() {
	println("Soy un perro y respiro")
}

func (p Perro) Pensar() {
	println("Soy un perro y pienso")
}

func (p Perro) Comer() {
	println("Soy un perro y como")
}

func (p Perro) Hablar() {
	println("Soy un perro y ladro")
}

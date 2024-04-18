package funciones

func Tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia++
		return numero * secuencia
	}
}

func LlamarClosure() {
	multiplicar := Tabla(2)
	for i := 0; i < 10; i++ {
		println(multiplicar())
	}
}
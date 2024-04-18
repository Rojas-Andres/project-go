package variables

import (
	"strconv"
	"time"
)
var Nombre string
var Estado bool
var Sueldo float32
var Sueldo2 float64
var Fecha time.Time

func RestoVariables() {
	// ShowIntegers() se puede llamar desde aquí
	// Variables
	Nombre = "Pedro"
	Estado = true
	Sueldo = 500.50
	Fecha = time.Now()
	var a int = 10
	var b int = 3

	// Operaciones
	var suma int = a + b
	var resta int = a - b
	var multiplicacion int = a * b
	var division int = a / b
	var resto int = a % b
	// Mostrar resultados
	println("Suma: ", suma)
	println("Resta: ", resta)
	println("Multiplicacion: ", multiplicacion)
	println("Division: ", division)
	println("Resto: ", resto)
}


func ConviertoaTexto(numero int)(bool, string){
	texto := strconv.Itoa(numero)
	return true, texto
}
package ejer_interfaces

import (
	"fmt"

	"github.com/Rojas-Andres/project-go/interfaces"
)

func HumanoRespirando(humano interfaces.Humano) {
	humano.Respirar()
	fmt.Printf("Soy un humano %s y respiro \n", humano.Sexo())
}
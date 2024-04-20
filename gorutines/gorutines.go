package gorutines

import (
	"fmt"
	"strings"
	"time"
)

// func ImprimirNombre(nombre string){
// 	letras := strings.Split(nombre, "")
// 	for _, letra := range letras {
// 		time.Sleep(1000 * time.Millisecond)
// 		fmt.Println(letra)
// 	}
// }

func MiNombreLentooo(nombre string){
	letras := strings.Split(nombre, "")
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
}
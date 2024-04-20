package defer_panic

import (
	"fmt"
	"log"
)

func VemosDefer(){
	fmt.Println("Este es mi primer mensaje")
	defer fmt.Println("Este es el Final mensaje")
	fmt.Println("Este es el segundo mensaje")
}

func EjemploPanic(){
	defer func(){
		reco := recover()
		if reco != nil {
			log.Fatalf("Ocurrio un error que genero Panic \n %v", reco) // El log guarda la fecha y hora del error , %v es para mostrar el valor de la variable
		}
	}()
	a:=1
	if a == 1 {
		panic("Se encontro el valor de 1")
	}

}
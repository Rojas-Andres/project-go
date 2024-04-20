package main

import (
	// "github.com/Rojas-Andres/project-go/ejercicios"
	// "github.com/Rojas-Andres/project-go/iteraciones"
	// "github.com/Rojas-Andres/project-go/variables"
	// "github.com/Rojas-Andres/project-go/files"
	// "github.com/Rojas-Andres/project-go/funciones"
	// "github.com/Rojas-Andres/project-go/arreglos_slices"
	// "github.com/Rojas-Andres/project-go/mapas"
	// "github.com/Rojas-Andres/project-go/users"
	// "github.com/Rojas-Andres/project-go/ejer_interfaces"
	// "github.com/Rojas-Andres/project-go/modelos"
	// "github.com/Rojas-Andres/project-go/defer_panic"
	"fmt"

	"github.com/Rojas-Andres/project-go/gorutines"
)


func main() {
	// variables.ShowIntegers()
	// variables.RestoVariables()
	// booleano , texto := variables.ConviertoaTexto(12312312)
	// println(booleano)
	// println(texto)

	// SWHITCH V1
	// os := runtime.GOOS
	// if os == "windows" {
	// 	println("El sistema operativo es Windows")
	// } else {
	// 	println("El sistema operativo es otro")
	// }
	// println(os)

	// switch os {
	// case "windows":
	// 	println("El sistema operativo es Windows")
	// case "darwin":
	// 	println("El sistema operativo es MacOS")
	// default:
	// 	println("El sistema operativo es otro")
	// }


	// SWHITCH V2
	// switch sytem_os := runtime.GOOS; sytem_os {
	// case "darwin":
	// 	println("El sistema operativo es darwin")
	// case "linux":
	// 	println("El sistema operativo es MacOS")
	// default:
	// 	println("El sistema operativo es ", sytem_os)
	// }
	// iteraciones.Iterar()
	// value := ejercicios.TablaMultiplicar()

	// files.GrabaTabla()
	// files.SumaTabla()
	// files.LeerArchivo()
	// files.LeoArchivo()
	// funciones.Calculos()
	// funciones.LlamarClosure()
	// funciones.Exponencia(2)


	// arreglos_slices.MostrarTabla()
	// arreglos_slices.MuestroSlice()
	// arreglos_slices.Capacidad()
	// mapas.MostrarMapas()


	// users.AltaUsuario()

	// Pedro:= new(modelos.Hombre)
	// ejer_interfaces.HumanoRespirando(Pedro)

	// Maria:= new(modelos.Mujer)
	// ejer_interfaces.HumanoRespirando(Maria)

	// defer_panic.VemosDefer()
	// defer_panic.EjemploPanic()
	go gorutines.MiNombreLentooo("Andres")
	fmt.Println("Estoy aqui")
	var x string
	fmt.Scanln(&x)

}

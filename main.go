package main

import (
	"runtime"
	// "github.com/Rojas-Andres/project-go/ejercicios"
	// "github.com/Rojas-Andres/project-go/iteraciones"
	// "github.com/Rojas-Andres/project-go/variables"
	// "github.com/Rojas-Andres/project-go/files"
	"github.com/Rojas-Andres/project-go/funciones"
)


func main() {
	// variables.ShowIntegers()
	// variables.RestoVariables()
	// booleano , texto := variables.ConviertoaTexto(12312312)
	// println(booleano)
	// println(texto)
	os := runtime.GOOS
	if os == "windows" {
		println("El sistema operativo es Windows")
	} else {
		println("El sistema operativo es otro")
	}
	println(os)

	// switch os {
	// case "windows":
	// 	println("El sistema operativo es Windows")
	// case "darwin":
	// 	println("El sistema operativo es MacOS")
	// default:
	// 	println("El sistema operativo es otro")
	// }

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
	funciones.LlamarClosure()
}

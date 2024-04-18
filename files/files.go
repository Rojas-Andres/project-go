package files

import (
	"bufio"
	"fmt"
	"io/ioutil"

	"os"

	"github.com/Rojas-Andres/project-go/ejercicios"
)

var fileName string = "./files/txt/tabla.txt"
func GrabaTabla()  {
	var texto string = ejercicios.TablaMultiplicar()
	archivo, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error al crear el archivo", err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)

	archivo.Close()
}


func SumaTabla(){
	var texto string = ejercicios.TablaMultiplicar()
	if !Append(texto) {
		fmt.Println("Error al concatenar el contenido")
	} else {
		fmt.Println("El contenido se ha concatenado correctamente")
	}

}

func Append(texto string ) bool {
	archivo, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error al abrir el archivo", err.Error())
		return false
	}
	_, err = archivo.WriteString(texto)
	if err != nil {
		fmt.Println("Error al escribir el archivo", err.Error())
		return false
	}
	archivo.Close()
	return true
}

func LeerArchivo(){
	archivo, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error al abrir el archivo", err.Error())
		return
	}
	fmt.Println(string(archivo))
}

func LeoArchivo(){
	archivo, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error al abrir el archivo", err.Error())
		return
	}
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
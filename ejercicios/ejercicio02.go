package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TablaMultiplicar(){
	var err error 
	var numero int 

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Ingrese un número para mostrar su tabla de multiplicar")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Error al leer el número")
			} else{
				break
			}
		}
	}
	fmt.Println("Tabla de multiplicar del número ", numero)

	for i:=1; i <= 10; i++{
		fmt.Printf("%d x %d = %d\n", numero, i, numero*i)
	}
}
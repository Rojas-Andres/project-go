package arreglos_slices

import "fmt"

var tabla []int 
var matriz [20][20] int
func MostrarTabla() {
	tabla = []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}
	tabla2:= [10] string {"Pablo", "Juan"}
	for i := 0; i < len(tabla2); i++ {
		fmt.Println(tabla2[i])
	}

	matriz[7][12] = 15

	fmt.Println(matriz)
}
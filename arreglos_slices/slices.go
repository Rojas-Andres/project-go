package arreglos_slices

import "fmt"

var tablas []int = []int{2,41,21,2}
var arreglo [10] int = [10] int{1,2,3,4,5,6,7,8,9,10}
func MuestroSlice() {
	fmt.Println("Longitud inicial del slice", len(tablas))
	fmt.Println("Arreglo original")
	fmt.Println(arreglo)
	fmt.Println("capacidad", cap(arreglo))
	porcion:= arreglo[3:]
	porcion2:= arreglo[:5]
	fmt.Println(porcion)
	fmt.Println(porcion2)
}

func Capacidad(){
	elementos := make([]int, 5, 20) // 5 elementos y capacidad de 20 , es mas para reservar en memoria 

	fmt.Printf("Longitud %d, Capacidad %d\n", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)
	for i := 0; i < 130; i++ {
		nums = append(nums, i)
		fmt.Printf("Longitud %d, Capacidad %d\n", len(nums), cap(nums))
	}
}
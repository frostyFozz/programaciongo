package main

import "fmt"

func main() {

	// Slice no poseen datos, son apuntadores a una Array
	set := [7]string{"Leon", "Gorilla", "Caballo", "Baca", "Mariposa", "Abeja", "Pajarito"}
	animales := set[0:4]
	fly := set[4:7]
	fly[0] = "Aguila"
	fmt.Println("Array: ", set)
	fmt.Println("Animales: ", animales)
	fmt.Println("Voladores: ", fly)

}

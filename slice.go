package main

import (
	"fmt"
)

func main() {
	// len(): # es el tamano del elemento en el slice
	// cap(): # es la capacidad que tiene el slice dentro del codigo

	/*comida := [5]string{"Arroz", "Habichuela", "Guandules", "Maiz", "Res"}
	cemillas := comida[1:4]
	cemillas = append(cemillas, "Almejas", "verduras", "Girita")

	fmt.Println("Comidas: ", comida)
	fmt.Println("Cemillas: ", cemillas)
	fmt.Println("Tamano: ", len(cemillas))
	fmt.Println("Capacidad: ", cap(cemillas))
	*/

	comida := []string{"Moro", "Locrio", "Carnemolida"}
	fmt.Println("Comidas: ", comida)
	fmt.Println("Tamano: ", len(comida))
	fmt.Println("Capacidad: ", cap(comida))

	frutas := make([]string, 0, 4)
	frutas = append(frutas, "mango", "Peraz", "manzanas", "bananas")

	fmt.Println("Tipos de fruta: ", frutas)
	fmt.Println("Tamano: ", len(frutas))
	fmt.Println("Capacidad del slice: ", cap(frutas))

	var cemillas []string
	fmt.Println("Cemillas: ", cemillas)
	fmt.Println("Tamano: ", len(cemillas))
	fmt.Println("Capacidad: ", cap(cemillas))

}

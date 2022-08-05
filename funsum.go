package main

import "fmt"

func main() {
	fmt.Println(suma(4, 5, 7, 9))

}

func suma(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v

	}
	return total
}

// esta es una funcion que permite almacenar un slice con mas de un numero para sumarlos
// esto se usa en caso de que tengamos que hacer una operacion agregando y quitando valores

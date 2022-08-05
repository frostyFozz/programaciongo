package main

import "fmt"

func main() {
	func(texto string) {
		fmt.Println("hello", texto)
	}("Paco")

}

// Esto es una funcion anonima

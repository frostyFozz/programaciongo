package main

import "fmt"

func main() {
	// Bool, String, Numerico
	var user = "Edwin"
	var ver = true
	var num1 uint8 = 10
	var num2 float32 = 2.6
	var num3 uint16 = 300
	// Este formato que se le dio a num1 se llama casting
	// sirve para comvertir una variable temporalmente de un tipo a otro
	c := uint16(num1) + num3
	// aqui vamos a imprimir los tipos de variables
	fmt.Printf("\nTipo: %T, Valor: %v", ver, ver)
	fmt.Printf("\nTipo: %T, Valor: %v", user, user)
	fmt.Printf("\nTipo: %T, Valor: %v", num1, num1)
	fmt.Printf("\nTipo: %T, Valor: %v", num2, num2)
	fmt.Printf("\nTipo: %T, valor: %v", c, c)

}

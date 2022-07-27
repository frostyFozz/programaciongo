// estas son variables en short format

package main

import "fmt"

func main() {

	var dog, cat = "perrito", "\ngatito"
	var num1 = 2
	var num2 = 5.3
	var bol = true
	fmt.Println(dog, cat)
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(bol)
	fmt.Printf("\nTipo: %T, Valor: %v", dog, dog)
	fmt.Printf("\nTipo: %T, Valor: %v", cat, cat)
	fmt.Printf("\nTipo: %T, Valor: %v", num1, num2)
	fmt.Printf("\nTipo: %T, Valor: %v", bol, bol)
}

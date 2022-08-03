package main

import "fmt"

func main() {
	/*	Hello("Edwin", "Mercedes")*/
	// esta funcion esta recibiendo un valor y lo cambia por otro !
	carita := "Fachero"
	Cambio(&carita)
	fmt.Println(carita)
}

func Cambio(value *string) {
	*value = "Facherito"
}

/*func Hello(primerNom string, Apellido string) {
fmt.Printf("Hello %s %s", primerNom, Apellido)
*/

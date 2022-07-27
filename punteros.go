package main

import "fmt"

func main() {
	fruit := "Banana"
	var p *string
	p = &fruit
	*p = "Pina"
	fmt.Printf("Tipo: %T, Valor: %s, Direccion: %v\n", fruit, fruit, &fruit)
	fmt.Printf("Tipo: %T, Valor: %v, Diferencia: %s", p, p, *p)
}

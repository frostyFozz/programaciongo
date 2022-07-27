package main

import (
	"fmt"
)

func main() {
	// Operadores Aritmeticos (), *, /, +, -
	var a = 4 + 2*5
	fmt.Println(a)

	// Operadores de Agisnacion =, *=, /=, +=, -=

	var b = 10
	b /= 2
	println(b)

	// Declaraciones post-incremento y post-decremento ++, --
	// (no son una expresion si no una declaracion)
	// ++ le agrega uno mas a la variable y -- le resta
	// ejemplo
	var c = 3
	c++
	fmt.Println(c)
}

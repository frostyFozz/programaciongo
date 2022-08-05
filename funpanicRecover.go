package main

import "fmt"

func main() {
	dividir(9, 4)
	dividir(100, 7)
	dividir(200, 0)
	dividir(40, 8)

}

// Esto es un ejemplo de como trabajar con un panic controlado y un recover para que
// nuestro programa siga su ejecucion aunque tenga un error
func dividir(dividendo, divisor int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recuperandome Del panic", r)
		}
	}()
	validador(divisor)
	fmt.Println(dividendo / divisor)
}
func validador(divisor int) {
	if divisor == 0 {
		panic("Corran!!")
	}
}

package main

import (
	"errors"
	"fmt"
)

func main() {
	resultado, err := division(100, 0)
	if err != nil {
		fmt.Printf("Ha ocurrido un error %v", err)
		return
	}
	fmt.Println(resultado)
}

func division(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("No se puede dividir por 0!")
	}

	return dividendo / divisor, nil
}

package main

import (
	"fmt"
)

func main() {
	numero := []int{1, 15, 4, 30, 7, 80, 9, 100}
	resultado := filter(numero, func(num int) bool {
		if num <= 10 {
			return true
		}
		return false
	})

	fmt.Println(resultado) //[ 1, 4, 7, 9]
}

func filter(nums []int, callback func(int) bool) []int {
	resultado := []int{}
	for _, v := range nums {
		if callback(v) {
			resultado = append(resultado, v)
		}
	}
	return resultado
}

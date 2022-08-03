package main

import "fmt"

func main() {
	total := suma(7, 9)
	fmt.Println(total)

}

func suma(num1, num2 int) int {
	return num1 + num2
}

package main

import "fmt"

func main() {
	var bebida [3]string
	bebida[0] = "Brugal\n"
	bebida[1] = "Cerveza\n"
	bebida[2] = "mas Cerveza"

	alchool := [3]string{"Vodka\n", "Mucha Cerveza\n", "Mucha mas Cerveza"}

	fmt.Println(bebida, alchool)

}

package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	content, err := ioutil.ReadFile("/Users/edwinmercedes/Desktop/Edwin.rtf")
	if err != nil {
		fmt.Printf("Ha ocurrido un error %v", err)
		return

	}
	fmt.Println("El contenido del texto es: ", string(content))
}

package main

import (
	"fmt"
	"os"
)

func main() {
	archivo, err := os.Create("Hola.txt")
	if err != nil {
		fmt.Printf("Se a generado un error al crear el archivo %v", err)
		return
	}
	defer archivo.Close()

	_, err = archivo.Write([]byte("hola Golang"))
	if err != nil {
		fmt.Printf("Se a generado un error al escribir en  el archivo %v", err)
		return
	}
}

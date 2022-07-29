package main

import "fmt"

func main() {
	animales := make(map[string]string)
	animales["Gato"] = "Cat"
	animales["perro"] = "Dog"

	fmt.Println(animales)

	frutas := map[string]string{
		"Apple": "Manzanas",
		"Uvas":  "Grape",
		"peras": "pe",
	}
	fmt.Println(frutas)
	delete(frutas, "Apple")
	fmt.Println(frutas)

	if _, ok := animales["Gorilla"]; !ok {
		animales["Gorilla"] = "Mono"
	}
	fmt.Println(animales)
}

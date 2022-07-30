package main

import "fmt"

func main() {
	type class struct {
		Name    string
		Maestro string
		Ciudad  string
	}
	db := class{
		Name:    "Almacen De Datos",
		Maestro: "Edwin Mercedes",
		Ciudad:  "Santo Domingo",
	}
	fmt.Printf("%+v\n", db)
	html := class{
		Name:    "Datos Almacenados",
		Maestro: "Dannesa",
		Ciudad:  "La Vega",
	}
	fmt.Printf("%+v\n", html)

	b := &db
	b.Name = "Data Base"
	fmt.Printf("%+v\n", b)

}

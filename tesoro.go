package main

import "fmt"

func main() {
	reservas := make(map[string]string)
	reservas["Petroleo: "] = "Gasolina"
	reservas["Gaslicuado: "] = "Gas"
	reservas["Gasnatural: "] = "GasVerde"

	fmt.Println(reservas)

	if _, ok := reservas["Electricidad"]; !ok {
		reservas["Electricidad: "] = "Luz"

	}

	fmt.Println(reservas)
	frutas := [8]string{"Mango", "Pera", "Manzanas", "Chinas", "mandarina", "limon"}
	citricos := frutas[3:6]
	citricos[0] = "Pina"
	fmt.Println(frutas)
	fmt.Println(citricos)
	fmt.Println("Cantidad:", len(citricos))
	fmt.Println("Capacidad: ", cap(citricos))

	ciudades := map[string]string{
		"Santo Domingo: ": "Capital\n",
		"Santiago: ":      "Ciudad Corazon\n",
		"La vega: ":       "Bella\n",
		"Barahona: ":      "La perla del Sur\n",
	}
	fmt.Println(ciudades)
}

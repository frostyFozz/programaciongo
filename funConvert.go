package main

import (
	"fmt"
	"strings"
)

// esto es una funcion que convierte el texto total en min y el mismo total en mayusc
func main() {
	texto := "La Casa De la Canna "
	minusc, mayusc := convert(texto)
	fmt.Println(minusc, mayusc)
}

func convert(text string) (string, string) {
	min := strings.ToLower(text)
	may := strings.ToUpper(text)

	return min, may
}

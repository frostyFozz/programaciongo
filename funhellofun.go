package main

import "fmt"

func main() {
	x := hola("Edwin")("Mercedes")
	fmt.Println(x) // Hello Edwin Mercedes

}

func hola(name string) func(string) string {
	return func(text string) string {
		return "hello " + name + " " + text
	}
}

package main

import "fmt"

func main() {
	if retorno := test(); retorno == "teste" {
		fmt.Print("Sim!")
	} else {
		fmt.Print("Não!")
	}
}

func test() string {
	return "teste"
}

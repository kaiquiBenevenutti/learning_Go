package main

import "fmt"

func main() {
	test := "test"

	switch test {

	case "teste", "teste3":
		fmt.Print("Primeira condição!")

	case "test":
		fmt.Print("Segunda condição!")
	}
}
package main

import "fmt"

func main() {
	test := "teste7"

	switch test {

	case "teste", "teste3":
		fmt.Print("Primeira condição!")
		fallthrough //faz ele cair para o case de baixo tambem
	case "test":
		fmt.Print("Segunda condição!")
		fallthrough //continua caindo
	
	case "test5":
		fmt.Print("Terceira condição!")
		break // faz ele parar de cair

	case "teste6":
		fmt.Print("Teste")

	default:
		fmt.Print("Default")
	}
}
package main

import "fmt"

func main() {

	teste := []string{"teste1", "teste2", "teste3"}

	//Substituir uma variavel por "_" para ignora-la, não pode apenas deixar vazio
	for _, value := range teste {
		fmt.Println(value)
	}
}

package main

import "fmt"

func main() {

	teste := []string{"teste1", "teste2", "teste3"}

	for i, value := range teste {
		fmt.Println(value, i)
	}
}

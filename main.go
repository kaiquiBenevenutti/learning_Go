package main

import "fmt"

func main() {

	teste := false

	for ok := true; ok; ok = teste {
		fmt.Println("FOI")
	}
}

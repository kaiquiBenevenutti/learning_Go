package main

import (
	"fmt"
)

func main() {

	value, err := test()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(value)
}

func test() (string, error) {
	return "certo", nil
}

package main

import (
	"fmt"
)

func main() {

	funcaoTest := func(test string, testInt int) {
		fmt.Println(test, testInt)
	}

	test(funcaoTest)
}

func test(value func(string, int)) {
	value("Teste", 89)
}

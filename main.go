package main

import "fmt"

func main() {

	test() //funciona
	test("kaiqui") //funciona
	test("kaiqui", "rafael") //funciona ...

}


//Pode mandar 0 ou N parametros
func test(valoresString ...string) {

	for _, x := range valoresString {
		fmt.Println(x)
	}
}

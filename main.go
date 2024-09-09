package main

import "fmt"

func main() {

	var user user = user{
		name: "Rafinha",
		age:  18,
	}

	fmt.Print(user)
}

//Pode ser instanciado fora do arquivo
type User struct {
	Email    string //pode ser visto quando instanciar
	password string // não pode ser visto ao instanciar
}

//Não pode ser instanciado fora do arquivo
type user struct {
	name string
	age  int
}

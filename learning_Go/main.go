package main

import "fmt"

func main() {

	var test []string = []string{"test", "test", "test", "test"}

	fmt.Println(len(test))
	fmt.Println(cap(test))

	test = append(test, "test3")

	fmt.Println(len(test))
	fmt.Println(cap(test))

	test = append(test, "test3")

	fmt.Println(len(test))
	fmt.Println(cap(test))

	test = append(test, "test3")
	test = append(test, "test3")
	test = append(test, "test3")

	fmt.Println(len(test))
	fmt.Println(cap(test))
}

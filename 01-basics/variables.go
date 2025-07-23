package main

import "fmt"

func variables() {
	var value1 = "variable"
	value2 := "short variable"
	const value3 = "constant variable"
	fmt.Println(value1, value2, value3)

	fmt.Println(greet("pabz"))
}

func greet(name string) string {
	return fmt.Sprintf("hello %s!", name)
}

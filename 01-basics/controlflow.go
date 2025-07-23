package main

import "fmt"

func controlFlow() {
	even := 2
	odd := 3

	fmt.Printf("is %d even? %t\n", even, isEven(even))
	fmt.Printf("is %d even? %t\n", odd, isEven(odd))
}

func isEven(n int) bool {
	return n%2 == 0
}

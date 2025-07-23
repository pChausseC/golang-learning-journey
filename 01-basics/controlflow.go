package main

import (
	"fmt"
	"strings"
	"time"
)

func controlFlow() {
	defer fmt.Println("I guess we're all done here?")
	firstValue := 2
	secondValue := 3

	fmt.Printf("is %d even? %t\n", firstValue, isEven(firstValue))
	fmt.Printf("is %d even? %t\n", secondValue, isEven(secondValue))

	// If/Else
	if isEven(firstValue + secondValue) {
		fmt.Println("both even or both odd")
	} else {
		fmt.Println("one even and one odd")
	}

	// Switch
	day := time.Now().Weekday().String()
	switch day {
	case "Monday":
		fmt.Println("I hate mondays")
	case "Sunday":
		fmt.Println("I love lasagna")
	default:
		fmt.Printf("I cant believe its %s already!\n", day)
	}

	//For loop
	for i := 1; i <= 5; i++ {
		fmt.Println(strings.Repeat("i", i))
	}
	var age = 0
	for {
		if age > 2 {
			fmt.Println("Hello mother")
			break
		}
		fmt.Println("Googoo gaga")
		age++
	}
}

func isEven(n int) bool {
	return n%2 == 0
}

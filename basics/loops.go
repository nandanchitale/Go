package main

import "fmt"

func main() {
	// Go has only 1 type of loop : For
	// while can be used with modifying syntax of for loop

	// For loop
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// While loop
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// Nested loops
	for i := 1; i <= 5; i++ {
		for j := 1; j < i; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}

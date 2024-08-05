package main

import "fmt"

func main() {
	// Defer
	defer FirstRun()
	SecondRun()

	// Recover and Panic
	fmt.Println(div(3, 0))
	fmt.Println(3, 5)

	demPanic()
}

func FirstRun() {
	fmt.Println("I Executed First")
}

func SecondRun() {
	fmt.Println("I Executed Second")
}

func div(num1, num2 int) int {
	defer func() {
		fmt.Println(recover())
	}()

	solution := num1 / num2

	return solution
}

func demPanic() {
	defer func() {
		fmt.Println(recover())
	}()

	panic("Panic")
}

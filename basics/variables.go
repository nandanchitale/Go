package main

import "fmt"

func main() {
	// you can also declare multiple variables like this
	var (
		varA = 2
		varB = 3
	)

	fmt.Println(varA, varB)

	// variables can also be declared as
	x, y := 14, 15

	fmt.Println(x)
	fmt.Println(y)

	var a int = 5
	var b float32 = 4.32
	const pi float64 = 3.1415139475

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(pi)
}

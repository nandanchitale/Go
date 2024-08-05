package main

import "fmt"

func main() {
	x, y := 5, 6

	fmt.Println("x + y = ", x+y)
	fmt.Println("x - y = ", x-y)
	fmt.Println("x * y = ", x*y)
	fmt.Println("x / y = ", x/y)
	fmt.Println("x / y = ", x/y)
	fmt.Println("x mod y = ", x%y)

	isBool := true
	// Long way : var isBool bool = true

	hate := false

	fmt.Println(isBool && hate)
	fmt.Println((isBool || hate))
	fmt.Println(isBool)
}

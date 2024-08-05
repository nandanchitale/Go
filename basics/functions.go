package main

import "fmt"

func main() {
	x, y := 5, 6
	fmt.Println(add(x, y))

	num := 5
	fmt.Println(factorial(num))

	fmt.Println(addNums(10, 20, 30, 40, 50))
}

func add(num1, num2 int) int {
	return num1 + num2
}

// recursive function
func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

func addNums(args ...int) int {
	sum := 0
	for _, value := range args {
		sum += value
	}

	return sum
}

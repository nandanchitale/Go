package main

import "fmt"

func main() {
	// arrays

	// way 1
	var evenNum [5]int
	evenNum[0] = 0
	evenNum[1] = 2
	evenNum[2] = 4
	evenNum[3] = 6
	evenNum[4] = 8

	fmt.Println(evenNum[2])
	fmt.Println("-------------------------------")

	// way 2
	var evenNums = [6]int{0, 2, 4, 6, 8, 10}
	fmt.Println(evenNums[5])
	fmt.Println("-------------------------------")

	// iterator on array
	for _, value := range evenNum {
		fmt.Println(value)
	}

	fmt.Println("-------------------------------")
	for i, value := range evenNum {
		fmt.Println(i, value)
	}

	fmt.Println("-------------------------------")

	// slicing
	slicedArray := evenNums[3:5]
	fmt.Println(slicedArray)

	fmt.Println("-------------------------------")

	// empty sliced array
	slice2 := make([]int, 5, 10)
	fmt.Println(slice2)

	fmt.Println("-------------------------------")

	// copy array
	numSlice := []int{5, 4, 3, 2, 1}
	copy(slice2, numSlice)
	fmt.Println(slice2)
	fmt.Println(numSlice)

	fmt.Println("-------------------------------")

	// append values to array
	slice3 := append(numSlice, 3, 0, -1)
	fmt.Println(slice3)

	fmt.Println("-------------------------------")
}

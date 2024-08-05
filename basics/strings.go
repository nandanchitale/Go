package main

import "fmt"

func main() {
	// String
	var Name string = "Nandan Chitale"

	const pi float64 = 3.14132145
	win := true
	x := 5

	// Get string length
	fmt.Println(len(Name))

	fmt.Println(Name + " is a chill dude")

	fmt.Printf("%.3f \n", pi)
	fmt.Printf("%T \n", Name)

	fmt.Printf("%t \n", win) // bool
	fmt.Printf("%d \n", x)   // decimal
	fmt.Printf("%b \n", 25)  // binary
	fmt.Printf("%c \n", 35)  // ascii
	fmt.Printf("%x \n", 15)  // hexcode
	fmt.Printf("%e \n", pi)  // scientific notation
}

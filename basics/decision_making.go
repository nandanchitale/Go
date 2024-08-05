package main

import "fmt"

func main() {

	// If-else
	var age int = 21
	if age > 18 {
		fmt.Println("You can vote")
	} else { // else has to be on this line. If it's on new line, it'll give syntax error
		fmt.Println("you cannot vote")
	}

	// switch case
	switch age {
	case 16:
		fmt.Println("Prepare for college")
	case 20:
		fmt.Println("Get a job")
	default:
		fmt.Println("Are you alive?")
	}
}

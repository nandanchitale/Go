package main

import "fmt"

func main() {
	StudentAge := make(map[string]int)
	StudentAge["Nandan"] = 26
	StudentAge["Aarya"] = 23
	StudentAge["Sahil"] = 25
	StudentAge["Kirti"] = 22

	fmt.Println(StudentAge)
	fmt.Println(StudentAge["Kirti"])
	fmt.Println(len(StudentAge))

	// Nested maps
	superHero :=
		map[string]map[string]string{
			"Superman": map[string]string{
				"RealName": "Clark Kent",
				"City":     "Metropolis",
			},
			"Batman": map[string]string{
				"RealName": "Bruce Wayne",
				"City":     "Gothom",
			},
		}

	if temp, hero := superHero["Superman"]; hero {
		fmt.Println(temp["RealName"], temp["City"])
	}
}

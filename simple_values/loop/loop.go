package main

import "fmt"

func greet() {
	age := 19

	if age >= 18 {
		fmt.Println("Peraon is an adult")

	}
	if age < 18 {
		fmt.Println("Person is not an adult")
	}
}

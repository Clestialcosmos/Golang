package main

import "fmt"

func main() {
	var names [4]int
	//Array length
	// fmt.Println(len(names))
	fmt.Print("Enter the values of array :")
	for i := 0; i < 4; i++ {
		fmt.Scanln(&names[i])
	}
	fmt.Println(names)
	//If we have fixed size or predictable than er majorily usr array but if we want dynamic array then in that case we use slices but also we can use dynamic array but their storge get affected that's not worthg it.
	

}

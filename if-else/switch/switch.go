package main

import (
	"fmt"
)

func main() {
	//simple switch

	// i := 5

	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// case 4:
	// 	fmt.Println("four")
	// case 5:
	// 	fmt.Println("five")
	// default:
	// 	fmt.Println("other")
	// }

	//multiple condition switch
	// time is library in go

	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("Holiday")
	// default:
	// 	fmt.Println(time.Now().Weekday())

	// }

	//type switch
	WhoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("It's an integer")
		case string:
			fmt.Println("it's string")
		case bool:
			fmt.Println("it's a booean")
		default:
			fmt.Println("other", t)
		}
	}
	WhoAmI("GOLANG")

}

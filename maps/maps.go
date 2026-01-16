package main

import (
	"fmt"
	"maps"
)

// maps -> hash(java),objects(java),dictionary (python) are basically maps
func main() {
	//Creating map

	// m := make(map[string]string)
	//here map is used as name := make(map[key type]value type)

	//setting an element

	// m["name"] = "golang"
	// m["area"] = "Backend"
	// fmt.Println(m["name"], m["area"])
	// fmt.Println(m["class"])
	// fmt.Println(len(m))
	// fmt.Println(m)
	// delete(m, "area")
	// fmt.Println(m)

	// m := map[string]int{"date": 10, "price": 100000000, "phone": 98957}
	// fmt.Println(m)

	//_, ok := m["price"]
	// v, ok := m["phone"]
	// fmt.Println(v)
	// if ok {
	// 	fmt.Println("all ok")
	// } else {
	// 	fmt.Println("not ok")
	// }

	m1 := map[string]int{"ID": 12345, "Price": 500}
	m2 := map[string]int{"ID": 23456, "Price": 1000}

	fmt.Println(maps.Equal(m1, m2))
}

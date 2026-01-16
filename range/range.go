package main

import "fmt"

//main -> iterating for data structure
func main() {
	// nums := []int{1, 2, 34, 4}

	// for _, num := range nums {
	// 	fmt.Println(num)
	// }

	// sum := 0
	// for _, num := range nums {
	// 	sum = sum + num
	// }
	// fmt.Println(sum)
	// here "_" is used for index and we can keep as 'i'

	// for i, num := range nums {
	// 	fmt.Println(num, i)
	// }

	m1 := map[string]string{"name": "Tanishq Tomar", "chinese name": "谭明坚"}

	// for k, v := range m1 {
	// 	fmt.Println(k, v)
	// }

	for k := range m1 {
		fmt.Println(k)
	}

	// if only one variable is there then only key value will appear

	for i, c := range "golang" {
		// fmt.Println(i, c)
		fmt.Println(i, string(c))
	}
	//here c => unique code point rune and rune is data structure (discuss later on)
	//starting byte of rune is 255
	//if it's less than stores under 1 byte
	//as g => 103
	//o => 111
	//l => 108 like this and so on.....

}

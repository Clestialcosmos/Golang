package main

import "fmt"

// func add(a string, b string) string {
// 	return a + b

// }
func getlanguages() (string, string, string, string, string) {
	return "C", "Python", "javascript", "CSS", "HTML"
}

func main() {
	a, b, c, d, e := getlanguages()
	fmt.Println(a, b, c, d, e)
}

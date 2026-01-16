package main

import (
	"fmt"
	"slices"
)

// slice -> dynamic
// most used construct in go
// +useful method
func main() {
	//we declare but not initialize to be precise slice is nil

	// var nums []int

	// fmt.Println(nums == nil)
	// fmt.Print(len(nums))

	var digit = make([]int, 2, 5)
	//for two D arrays use vare name = [][]int{{1,2,3},{4,5,6}}
	//cpacity should always greater than length
	fmt.Println(digit)
	digit = append(digit, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	digit[0] = 12
	digit[1] = 13
	//copy function
	var nums = make([]int, 3, 5)
	copy(nums, digit)
	fmt.Println(digit, nums)
	fmt.Println(cap(digit))
	fmt.Println(len(digit))
	fmt.Println(digit[3:7])
	fmt.Println(digit[:5])
	fmt.Println(digit[4:])
	fmt.Println(slices.Equal(nums, digit))
	fmt.Println(slices.Equal(digit, digit))

}

//In Go, slices are dynamic, flexible views into arrays that let you work with sequences of data without worrying about fixed sizes.
// They are widely used for collections because they support resizing, efficient iteration, and built-in functions like append, copy,
//  and slicing operations.

//In Go, the make function is used to create and initialize slices, maps, and channels. Unlike new, which only allocates memory, make sets up the internal data structures so the resulting object is ready to use.
//make(type, length, capacity)
//capacity = max numbers of elements can fit
//for input in GO -> scanln,scan

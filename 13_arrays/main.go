package main

import "fmt"
func main() {
	// Declare an array of integers with a length of 5
	//aray is fixed size, 
	// 	we cannot change the size of an array after it is declared
	var marks [5]int
	marks[0] = 85
	marks[1] = 90
	marks[2] = 78
	marks[3] = 92
	marks[4] = 88
	fmt.Println(marks)

	// Declare an array with initial values
	marks2 := [5]int{85, 90, 78, 92, 88}
	fmt.Println(marks2)
}
package main

import "fmt"

func main() {
	// In Go, a pointer is a variable that holds the memory address of another variable.
	// To declare a pointer, you use the * operator followed by the type of the variable it points to.
	// 	| Symbol | Meaning               |
	// | ------ | --------------------- |
	// | `&x`   | address of x          |
	// | `*p`   | value at that address |

	score := 10
	fmt.Println("before:", score)

	addScore(&score)
	fmt.Println("after:", score)

}
func addScore(score *int) {
	*score += 5
}

//normal variables -> value store karta hai
//pointer variables -> address store karta hai
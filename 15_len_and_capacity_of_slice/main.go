package main

import "fmt"

func main() {
	//length and capacity of slice
	//length is the number of elements in the slice
	//capacity is the number of elements in the underlying array
	//capacity is always greater than or equal to length
	//when we append to a slice, if the length exceeds the capacity,
	// a new underlying array is created with double the capacity and
	// the elements are copied to the new array

    s := make([]int, 2, 2) // len=2, cap=2
    fmt.Println(len(s), cap(s)) // 2 2

    s = append(s, 10) // capacity exceeded

    fmt.Println(len(s), cap(s)) // 3 4 (capacity doubled)


	todos := []string{"do youtube", "workout every day"}
	more := []string{"learn go"}
	
	//...
	todos = append(todos, more...)
	fmt.Println(todos) // [do youtube workout every day learn go]
}
package main

import "fmt"

func add(a int, b int)int{
	return a + b
}
func sumAndProduct(a int, b int)(int, int){
	return a + b, a * b
}
func main() {
	// function
	// func functionName(parameters) returnType {
	// 	// function body
	// }	
	sum := add(10, 20)
	fmt.Println("sum:", sum)
	sum, product := sumAndProduct(10, 20)
	fmt.Println("sumAndProduct:", sum, product)

	// _ is used to ignore a return value
	_, product = sumAndProduct(10, 20)
	fmt.Println("product:", product)
}
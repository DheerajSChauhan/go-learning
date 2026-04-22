package main

import "fmt"

func divide(a int, b int)(rohan int,  sanagam int){
	rohan	= a / b
	sanagam = a +b
	return
}
func main(){
	q, r := divide(10, 2)
	fmt.Println("quotient:", q)
	fmt.Println("sum:", r)
}
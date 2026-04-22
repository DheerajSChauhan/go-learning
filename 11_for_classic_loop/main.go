package main

import "fmt"

func main() {
	// for i := 1; i <= 5; i++ {
	// 	fmt.Println(i)
	// }

	// println("Sum count:")
	// n := 10
	// sum := 0
	// for  i := 0; i<=n; i++{
	// 	sum += i;
	// }
	// fmt.Println(sum)

	//printing pattern:
	// *
	// **
	// ***
	// ****
	for i:=1; i<=4; i++{
		for j:= 1; j<= i; j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}
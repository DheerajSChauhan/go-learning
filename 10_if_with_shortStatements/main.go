package main

import "fmt"

func main() {
	items := 3
	pricePerItem := 49
	if total := items * pricePerItem; total >= 100 {
		fmt.Println("Eligible for ship.")
	}else{
		fmt.Println("Not-Eligible for shiping***")
	}
}
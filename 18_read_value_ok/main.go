package main

import "fmt"

func main() {
	points := map[string]int{
		"a": 10,
		"b": 0, //valid value
	}
	fmt.Println("a", points["a"])
	fmt.Println("b", points["b"])
	fmt.Println("c", points["c"])

	vala, ok := points["a"]
	fmt.Println("a", vala, ok)
	valb, ok := points["b"]
	fmt.Println("b", valb, ok)
	valc, ok := points["c"]
	fmt.Println("c", valc, ok)

	if val, ok := points["c"]; ok {
		fmt.Println("c", val)
	}else{
		fmt.Println("c", "not found")
	}

	price := map[string]int{
		"item1" : 500,
		"item2" : 1000,
	}
	total := 0
	for items, price := range price{
		fmt.Println(items, price)
		total += price
	}
	fmt.Println("Total:", total)
}

package main

import (
	"fmt"
	"strings"
)

func main(){
	firstName := "Dheeraj"
	lastName := "Chauhan"
	fullName := firstName + " " + lastName

	fmt.Println(fullName)
	fmt.Println(strings.ToUpper(fullName))
}
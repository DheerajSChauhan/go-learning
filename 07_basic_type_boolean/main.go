package main

import "fmt"

func main() {
	isLogged := true
	isAdmin := false
	hasSubscription := true

	//and
	canOpenDashboard := isLogged && hasSubscription
	canDeletePost := isAdmin || (isLogged && hasSubscription)

	fmt.Println(canOpenDashboard, canDeletePost)

	age := 5
	fmt.Println(age > 18)
}
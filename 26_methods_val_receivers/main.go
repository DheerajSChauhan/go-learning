package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	u := User{Name: "Dheeraj", Age: 20}
	fmt.Println(u.Intro())
}

// value receiver means this method will get a copy of the value and any changes made to the value inside the method will not affect the original value outside the method. This is useful when you want to ensure that the original value remains unchanged when the method is called.
func (u User) Intro() string {
	return fmt.Sprintf("Hi, I am %s", u.Name)
}
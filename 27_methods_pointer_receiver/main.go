package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	u := User{"Dheeraj", 20}
	fmt.Println("Before birthday:", u.Age)

	u.Birthday()
	fmt.Println("After birthday:", u.Age)
}

func (u *User) Birthday(){
	u.Age++
}
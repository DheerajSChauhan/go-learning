package main

import "fmt"

//structs are user defined data types that allow us to group together related data of different types. They are similar to classes in other programming languages, but they do not have methods or inheritance. A struct is defined using the `type` keyword followed by the name of the struct and the fields it contains.

type User struct{
	ID int
	Name string
	Email string
	Age int
}
func main(){
	u1 := User{
		ID: 1,
		Name: "Dheeraj Singh Chauhan",
		Email: "dheeraj@gmail.com",
		Age: 20,
	}
	fmt.Println(u1)

	//struc is mutable data type, we can change the value of the fields of a struct after it has been created.
	fmt.Println("before:", u1.Name)
	u1.Name = "Dheeraj Singh"
	fmt.Println("after:", u1.Name)

}
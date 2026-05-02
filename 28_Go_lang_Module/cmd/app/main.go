package main

import "go-modules/internal/greet"

func main() {
	msg1 := greet.Hello("Dheeraj")
	println(msg1)
}
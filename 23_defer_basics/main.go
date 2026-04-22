package main

import "fmt"

func main() {
	fmt.Println("case1: success")
	if err := doWork(true); err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("case2: failure")
	if err := doWork(false); err != nil {
		fmt.Println("error:", err)
	}
}

func doWork(success bool) error {
	fmt.Println("start: acquiring resource")

	// defer ensures cleanup always runs
	defer fmt.Println("cleanup: releasing resource")

	if !success {
		return fmt.Errorf("work failed")
	}

	fmt.Println("work: doing something important")
	fmt.Println("work: this work is done")
	return nil
}
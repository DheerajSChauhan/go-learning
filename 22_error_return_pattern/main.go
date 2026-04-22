// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	//go do not use exeptions for normal failures
// 	//function ->returns error as noremal return value

// 	//val, err := something()
// 	//if err != nil{handle error}

// }

// func run() error {

// }

// func parseLevel(s string) (int, error) {
// 		//{vslue, error}
// 		//nil error means no error
// 		//non nil error means error

// 		//pattern: return value, error
// 	n, err := strconv.Atoi(s)
// 	if err != nil {
// 		return 0, fmt.Errorf("Level must be number")
// 	}

// 	if n < 1 || n > 5 {
// 		return 0, fmt.Errorf("Level must be between 1 and 5")
// 	}
// 	return n, nil
// }

package main

import "fmt"

func divide(num1 int , num2 int)(int, error){
	if num2 == 0{
		return -1, fmt.Errorf("Cannot divide by zero")
	}
	return num1/num2, nil
}

func main(){
	result, err := divide(20, 0)
	if err != nil{
		fmt.Println("Error:", err)
	} else {
		fmt.Println("The division of two number is: ", result)
	}
}
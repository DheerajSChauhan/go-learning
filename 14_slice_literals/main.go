package main
import "fmt"
func main(){
	//common collection type
	//dynamic size
	result := []string{ "Dheeraj", "Rohit", "Satyarth", "Shivam" }
	fmt.Println(result, "first-element:", result[0], "Length:", len(result), "lastElement:" ,result[len(result)-1])	

	result[1] = "priyanka"
	fmt.Println(result)
}
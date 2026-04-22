package main
import "fmt"

func main(){
	views := []int{10, 20, 45, 50, 60}
	total := 0
	for i, v := range views {
		fmt.Println("Day: ", i,"Views: ", v)
		total += v
	}
	fmt.Println("Total views: ", total)
}
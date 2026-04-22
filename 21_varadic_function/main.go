package main
import "fmt"

func sum(nums...int)int{
	total := 0
	for _, num := range nums{
		total += num
	}
	return total
}

func main(){
	fmt.Println(sum(1, 2, 3, 4, 5))

	values := []int{10, 20, 30}
	fmt.Println(sum(values...))

	res := func (n int) int {
		return n * n
	}
	fmt.Println(res(2))

	res1 := func ( a int, b int)int{
		return a + b
	}(10, 20)
	fmt.Println(res1)

}
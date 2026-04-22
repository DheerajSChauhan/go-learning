package main
import "fmt"

func main(){
	//map[keyType]valueType
	//map[string]int
	//map[string]string
	//map[int]string
	//map[int]int

	ages := map[string]int{
		"dheeraj": 25,
		"rohit": 30,
		//"rohit": 35, //duplicate key, last value will be used
	}
	fmt.Println(ages["dheeraj"]) //25
	fmt.Println(ages) //map[dheeraj:25 rohit:30]

	//make(map[keyType]valueType)
	var scores map[string]int //nil map, cannot add key-value pairs
	scores = make(map[string]int) //initialize the map
	fmt.Println(scores, scores["math"]) //map[] 0	
	scores["math"] = 90
	scores["english"] = 85
	fmt.Println(scores) //map[english:85 math:90]

	user := map[string]string{
		"u1": "dheeraj",
		"u2": "rahul",
		"u3": "jhon",
	}
	fmt.Println("before deletion:", user) //map[u1:dheeraj u2:rahul u3:jhon]
	delete(user, "u3")
	fmt.Println("after deletion:", user) //map[u1:dheeraj u2:rahul]
}
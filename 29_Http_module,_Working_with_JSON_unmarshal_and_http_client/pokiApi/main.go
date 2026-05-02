package main

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
)

type Pokemon struct {
	Name   string `json:"name"`
	Height int    `json:"height"`
	Weight int    `json:"weight"`
}	

func main() {
	url := "https://pokeapi.co/api/v2/pokemon/squirtle"
	//all pokemon names are in https://pokeapi.co/api/v2/pokemon?limit=1000
	// eg : pikachu , squirtle , charmander , bulbasaur , etc

	resp, err := http.Get(url)
	if err != nil{
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	if( resp.StatusCode != http.StatusOK){
		fmt.Println("Error:", resp.Status)
		return
	}
	
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var data Pokemon

	if err := json.Unmarshal(bodyBytes, &data); err != nil{
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}	
	fmt.Printf("Name: %s, Height: %d, Weight: %d\n", data.Name, data.Height, data.Weight)


}
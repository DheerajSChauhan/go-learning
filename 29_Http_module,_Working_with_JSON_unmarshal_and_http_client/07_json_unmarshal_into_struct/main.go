package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type catFactResponseStructure struct{
	Fact string `json:"fact"`
	Length int `json:"length"`
}

func main() {
	url := "https://catfact.ninja/fact"

	resp, err := http.Get(url)
	if err != nil{
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	if(resp.StatusCode != http.StatusOK){
		fmt.Println("Error:", resp.Status)
		return
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read body failed ",err)
		return
	}

	var data catFactResponseStructure

	if err := json.Unmarshal(bodyBytes, &data); err != nil{
		fmt.Println("json unmarshal is failed")
		return
	}
	fmt.Println(data.Fact, data.Length)
}
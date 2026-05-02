package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	apiURL := "https://jsonplaceholder.typicode.com/todos"

	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Status:", resp.Status)
		return
	}

	// ✅ Read response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}

	bodyText := string(bodyBytes)

	// ✅ Print only first 250 chars
	max := 250
	if len(bodyText) < max {
		max = len(bodyText)
	}

	fmt.Println(bodyText[:max])
}
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)
type catResponse struct{
	Fact string `json:"fact"`
	Length int `json:"length"`
}

func writeJson( w http.ResponseWriter, status int, data any){
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(data)
}

func externalHandeler(w http.ResponseController, r *http.Request){

}

func main() {
	http.HandleFunc("/external", externalHandler)
	err := http.ListenAndServe(":5000", nil)
	
	fmt.Println(err)
}
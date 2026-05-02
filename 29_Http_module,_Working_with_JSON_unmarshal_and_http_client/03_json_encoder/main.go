package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func successHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	res := map[string]any{
		"ok" : true,
		"message" : "JSON encode successfully",
		"date/time" : time.Now().UTC(),
	}
	_ = json.NewEncoder(w).Encode(res)
}
func main() {

	http.HandleFunc("/ok", successHandler)
	err := http.ListenAndServe(":5000", nil)
	fmt.Println(err)

}
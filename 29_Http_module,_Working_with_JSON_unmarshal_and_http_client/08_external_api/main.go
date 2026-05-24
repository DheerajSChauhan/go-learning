package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type catFactResponse struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

func writeJson(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(data)
}

func fetchCatFact() (catFactResponse, error) {
	url := "https://catfact.ninja/fact"

	res, err := http.Get(url)
	if err != nil {
		return catFactResponse{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return catFactResponse{}, fmt.Errorf("external api error: %s", res.Status)
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return catFactResponse{}, err
	}

	var data catFactResponse
	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		return catFactResponse{}, err
	}

	return data, nil
}

func externalHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJson(w, http.StatusMethodNotAllowed, map[string]any{
			"ok":    false,
			"error": "only GET method is allowed",
		})
		return
	}

	data, err := fetchCatFact()
	if err != nil {
		writeJson(w, http.StatusInternalServerError, map[string]any{
			"ok":    false,
			"error": "failed to fetch data",
		})
		return
	}

	writeJson(w, http.StatusOK, map[string]any{
		"ok":        true,
		"timestamp": time.Now().UTC(),
		"external": map[string]any{
			"name":   "catfact.ninja",
			"fact":   data.Fact,
			"length": data.Length,
		},
		"data": data,
	})
}

func main() {
	http.HandleFunc("/external", externalHandler)

	fmt.Println("Server running on port 5000")

	err := http.ListenAndServe(":5000", nil)
	fmt.Println(err)
}
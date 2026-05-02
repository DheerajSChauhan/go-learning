package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func writeJson(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

type testRequest struct {
	Name string `json:"name"`
}

func testHandler(w http.ResponseWriter, r *http.Request) {

	// Step 1: Method check
	if r.Method != http.MethodPost {
		writeJson(w, http.StatusMethodNotAllowed, map[string]any{
			"ok":    false,
			"error": "method not allowed",
		})
		return
	}

	// Step 2: Read body
	defer r.Body.Close()

	var req testRequest
	dec := json.NewDecoder(r.Body)

	if err := dec.Decode(&req); err != nil {
		writeJson(w, http.StatusBadRequest, map[string]any{
			"ok":    false,
			"error": "invalid json",
		})
		return
	}

	// Step 3: Validation
	req.Name = strings.TrimSpace(req.Name)

	if req.Name == "" {
		writeJson(w, http.StatusBadRequest, map[string]any{
			"ok":    false,
			"error": "name must not be empty",
		})
		return
	}

	// Step 4: Success response
	writeJson(w, http.StatusOK, map[string]any{
		"ok":        true,
		"data":      req,
		"timeStamp": time.Now().UTC(),
	})
}

func main() {
	http.HandleFunc("/test", testHandler)

	fmt.Println("Server running on :5000")
	err := http.ListenAndServe(":5000", nil)
	fmt.Println(err)
}
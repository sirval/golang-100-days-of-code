package main

import (
	"encoding/json"
	"net/http"

	"github.com/sirval/todo/structs/structs.go"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("todo", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			resp := structs.Response{
				StatusCode: http.StatusOK,
				Data:       "Retrieved",
			}
			json.NewEncoder(w).Encode(resp)
		}
	})
	http.ListenAndServe(":8000", mux)
}

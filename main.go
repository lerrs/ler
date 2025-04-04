package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var task string

type requestBody struct {
	Task string `json:"task"`
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var rb requestBody

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&rb)
		if err != nil {
			http.Error(w, "Неправильный запрос", http.StatusBadRequest)
			return
		}
		task = rb.Task
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Task received: %s", task)
	} else {
		http.Error(w, "Поддерживается только метод POST", http.StatusMethodNotAllowed)
	}
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintf(w, "Hello, %s", task)
	} else {
		http.Error(w, "Поддерживается только метод GET", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/task", PostHandler)
	http.HandleFunc("/task/get", GetHandler)

	http.ListenAndServe("localhost:8080", nil)
}

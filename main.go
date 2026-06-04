package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Simple struct {
	Name        string
	Description string
	Url         string
}

func handler(w http.ResponseWriter, r *http.Request) {
	payload := Simple{
		Name:        "Hello",
		Description: "World",
		Url:         r.Host,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Printf("render response: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = fmt.Fprintln(w, string(body))
}

func main() {
	fmt.Println("Server listening on port 4444")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":4444", nil))
}

package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/cotacao", nil)
	err := http.ListenAndServe(":8087", nil)

	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"log"
	"net/http"

	"github.com/Matheus-ALima/web-form.git/handlers"
)

func main() {
	http.HandleFunc("/", handlers.SubscriptionHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

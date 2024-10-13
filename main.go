package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/health", checkHealth)
	http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)

	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server Started listening on port: 8000")
}

func handleCreatePaymentIntent(w http.ResponseWriter, r *http.Request) {

}

func checkHealth(w http.ResponseWriter, r *http.Request) {

}

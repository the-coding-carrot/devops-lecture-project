package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/the-coding-carrot/devops-lecture-project/checkout-service/internal"
)

func main() {
	mux := http.NewServeMux()
	// Checkout
	mux.HandleFunc("/checkout/placeorder", internal.CheckoutPlaceOrderHandler)
	port := 8080
	log.Printf("Server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}

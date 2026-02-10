package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/the-coding-carrot/devops-lecture-project/internal"
)

func main() {
	mux := http.NewServeMux()
	// Auth
	mux.HandleFunc("/auth/login", internal.AuthLoginHandler)
	mux.HandleFunc("/auth/logout", internal.AuthLogoutHandler)
	// Product
	mux.HandleFunc("/products", internal.ProductListHandler)
	mux.HandleFunc("/products/{id}", internal.ProductDetailHandler)
	// Checkout
	mux.HandleFunc("/checkout/placeorder", internal.CheckoutPlaceOrderHandler)
	port := 8080
	log.Printf("Server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/the-coding-carrot/devops-lecture-project/product-service/internal"
)

func main() {
	mux := http.NewServeMux()
	// Product
	mux.HandleFunc("/products", internal.ProductListHandler)
	mux.HandleFunc("/products/{id}", internal.ProductDetailHandler)
	port := 8080
	log.Printf("Server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}

package internal

import (
	"net/http"
	"strings"

	"github.com/the-coding-carrot/devops-lecture-project/checkout-service/pkg"
)

// Static data for three products
var products = []pkg.Product{
	{ID: 1, Name: "Minecraft", Price: 24.99},
	{ID: 2, Name: "Grand Theft Auto V", Price: 59.99},
	{ID: 3, Name: "Assassin's Creed II", Price: 39.99},
}

func CheckoutPlaceOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error":"Missing Authorization header"}`))
		return
	}

	const bearerPrefix = "Bearer "
	if !strings.HasPrefix(authHeader, bearerPrefix) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error":"Authorization header must use Bearer scheme"}`))
		return
	}

	tokenString := strings.TrimPrefix(authHeader, bearerPrefix)

	if !pkg.VerifyToken(tokenString) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error":"Invalid token"}`))
		return
	}

	w.Write([]byte(`{"message":"Order placed successfully"}`))
}

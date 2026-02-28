package internal

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/the-coding-carrot/devops-lecture-project/product-service/pkg"
)

// Static data for three products
var products = []pkg.Product{
	{ID: 1, Name: "Minecraft", Price: 24.99},
	{ID: 2, Name: "Grand Theft Auto V", Price: 59.99},
	{ID: 3, Name: "Assassin's Creed II", Price: 39.99},
	{ID: 4, Name: "Star Wars: Jedi Survivor", Price: 89.99},
}

func ProductListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response, err := json.Marshal(products)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Internal Server Error"}`))
		return
	}
	w.Write(response)
}

func ProductDetailHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"error":"Product ID has wrong format"}`, http.StatusBadRequest)
		return
	}

	product := pkg.FindProductByID(products, id)
	if product == nil {
		http.Error(w, `{"error":"Product not found"}`, http.StatusNotFound)
		return
	}

	resp, err := json.Marshal(product)
	if err != nil {
		http.Error(w, `{"error":"Internal Server Error"}`, http.StatusInternalServerError)
		return
	}

	w.Write(resp)
}

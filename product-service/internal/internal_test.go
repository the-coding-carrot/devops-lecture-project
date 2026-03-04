package internal

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/the-coding-carrot/devops-lecture-project/product-service/pkg"
)

func TestProductListHandler(t *testing.T) {

	t.Run("False Method", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/products", nil)
		rec := httptest.NewRecorder()
		ProductListHandler(rec, req)

		if rec.Code != http.StatusMethodNotAllowed {
			t.Errorf("Expected status code %d, got %d", http.StatusMethodNotAllowed, rec.Code)
		}
	})

	t.Run("JSON Marshalling Error", func(t *testing.T) {
		originalProducts := products
		defer func() { products = originalProducts }()

		products = []pkg.Product{
			{ID: -1, Name: "", Price: math.NaN()},
		}

		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		rec := httptest.NewRecorder()
		ProductListHandler(rec, req)

		if rec.Code != http.StatusInternalServerError {
			t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, rec.Code)
		}
	})

	t.Run("Happy Path", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		rec := httptest.NewRecorder()

		ProductListHandler(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Code)
		}

		expectedHeader := "application/json"
		if rec.Header().Get("Content-Type") != expectedHeader {
			t.Errorf("Expected Content-Type header %s, got %s", expectedHeader, rec.Header().Get("Content-Type"))
		}
		var resProducts []pkg.Product
		err := json.Unmarshal(rec.Body.Bytes(), &resProducts)
		if err != nil {
			t.Errorf("Failed to unmarshal response body: %v", err)
		}

		if len(resProducts) != len(products) {
			t.Errorf("Expected %d products, got %d", len(products), len(resProducts))
		}
	})

}

func TestProductDetailHandler(t *testing.T) {

	t.Run("False Method", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/product/1", nil)
		rec := httptest.NewRecorder()
		ProductDetailHandler(rec, req)

		if rec.Code != http.StatusMethodNotAllowed {
			t.Errorf("Expected status code %d, got %d", http.StatusMethodNotAllowed, rec.Code)
		}
	})

	t.Run("JSON Marshalling Error", func(t *testing.T) {
		originalProducts := products
		defer func() { products = originalProducts }()

		products = []pkg.Product{
			{ID: 1, Name: "", Price: math.NaN()},
		}

		req := httptest.NewRequest(http.MethodGet, "/products/1", nil)
		req.SetPathValue("id", "1")
		rec := httptest.NewRecorder()
		ProductDetailHandler(rec, req)

		if rec.Code != http.StatusInternalServerError {
			t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, rec.Code)
		}
	})

	t.Run("Happy Path", func(t *testing.T) {
		if len(products) > 0 {
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/products/%v", products[0].ID), nil)
			req.SetPathValue("id", fmt.Sprintf("%v", products[0].ID))
			rec := httptest.NewRecorder()
			ProductDetailHandler(rec, req)

			if rec.Code != http.StatusOK {
				t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Code)
			}
		}
	})

	t.Run("Non-existent Product ID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/products/-1", nil)
		req.SetPathValue("id", "-1")
		rec := httptest.NewRecorder()
		ProductDetailHandler(rec, req)

		if rec.Code != http.StatusNotFound {
			t.Errorf("Expected status code %d, got %d", http.StatusNotFound, rec.Code)
		}
	})

	t.Run("Invalid Product ID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/products/a", nil)
		req.SetPathValue("id", "a")
		rec := httptest.NewRecorder()
		ProductDetailHandler(rec, req)

		if rec.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, rec.Code)
		}
	})

}

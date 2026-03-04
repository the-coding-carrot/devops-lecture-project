package internal

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang-jwt/jwt/v5"
)

func TestCheckoutPlaceOrderHandler(t *testing.T) {

	// False Method
	t.Run("False Method", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/checkout/place-order", nil)
		rec := httptest.NewRecorder()
		CheckoutPlaceOrderHandler(rec, req)
		if rec.Code != http.StatusMethodNotAllowed {
			t.Errorf("Expected status code %d, got %d", http.StatusMethodNotAllowed, rec.Code)
		}
	})

	// Missing Authorization header
	t.Run("Missing Authorization header", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/checkout/place-order", nil)
		rec := httptest.NewRecorder()
		CheckoutPlaceOrderHandler(rec, req)
		if rec.Code != http.StatusUnauthorized {
			t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, rec.Code)
		}
	})

	// Invalid Bearer token
	t.Run("Invalid Bearer token", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/checkout/place-order", nil)
		req.Header.Set("Authorization", "a")
		rec := httptest.NewRecorder()
		CheckoutPlaceOrderHandler(rec, req)
		if rec.Code != http.StatusUnauthorized {
			t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, rec.Code)
		}
	})

	// False Bearer token
	t.Run("False Bearer token", func(t *testing.T) {
		falseKey := []byte("false")
		falseToken := jwt.New(jwt.SigningMethodHS256)
		falseTokenString, _ := falseToken.SignedString(falseKey)

		req := httptest.NewRequest(http.MethodPost, "/checkout/place-order", nil)
		req.Header.Set("Authorization", "Bearer "+falseTokenString)
		rec := httptest.NewRecorder()
		CheckoutPlaceOrderHandler(rec, req)
		if rec.Code != http.StatusUnauthorized {
			t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, rec.Code)
		}
	})

	// Happy Path
	t.Run("Happy Path", func(t *testing.T) {
		secretKey := []byte("secret-key")
		token := jwt.New(jwt.SigningMethodHS256)
		tokenString, _ := token.SignedString(secretKey)

		req := httptest.NewRequest(http.MethodPost, "/checkout/place-order", nil)
		req.Header.Set("Authorization", "Bearer "+tokenString)
		rec := httptest.NewRecorder()
		CheckoutPlaceOrderHandler(rec, req)

		if rec.Body.String() != `{"message":"Order placed successfully"}` {
			t.Errorf("Expected response body %s, got %s", `{"message":"Order placed successfully"}`, rec.Body.String())
		}
	})
}

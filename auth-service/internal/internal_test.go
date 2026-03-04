package internal

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"testing"
)

func TestAuthLoginHandler(t *testing.T) {

	t.Run("False Method", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/auth/login", nil)
		rec := httptest.NewRecorder()
		AuthLoginHandler(rec, req)
		if rec.Code != http.StatusMethodNotAllowed {
			t.Errorf("Expected status code %d, got %d", http.StatusMethodNotAllowed, rec.Code)
		}
	})

	t.Run("Invalid Credentials", func(t *testing.T) {
		formData := "username=invalid&password=invalid"
		req := httptest.NewRequest(http.MethodPost, "/auth/login", strings.NewReader(formData))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		rec := httptest.NewRecorder()
		AuthLoginHandler(rec, req)

		if rec.Code != http.StatusUnauthorized {
			t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, rec.Code)
		}
	})

	t.Run("Happy Path", func(t *testing.T) {

		formData := "username=user&password=pass"
		req := httptest.NewRequest(http.MethodPost, "/auth/login", strings.NewReader(formData))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		rec := httptest.NewRecorder()
		AuthLoginHandler(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Code)
		}
	})
}

func TestAuthLogoutHandler(t *testing.T) {
	t.Run("False Method", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/auth/logout", nil)
		rec := httptest.NewRecorder()
		AuthLogoutHandler(rec, req)
		if rec.Code != http.StatusMethodNotAllowed {
			t.Errorf("Expected status code %d, got %d", http.StatusMethodNotAllowed, rec.Code)
		}
	})

	t.Run("Happy Path", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/auth/logout", nil)
		rec := httptest.NewRecorder()
		AuthLogoutHandler(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Code)
		}
	})
}

package pkg

import (
	"crypto/rand"
	"crypto/rsa"
	"testing"

	"github.com/golang-jwt/jwt/v5"
)

func TestVerifyToken(t *testing.T) {

	validToken := jwt.New(jwt.SigningMethodHS256)
	validTokenString, _ := validToken.SignedString(secretKey)

	tempKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	invalidSignedToken := jwt.New(jwt.SigningMethodRS256)
	invalidSignedTokenString, _ := invalidSignedToken.SignedString(tempKey)

	falseToken := jwt.New(jwt.SigningMethodHS256)
	falseTokenString, _ := falseToken.SignedString([]byte("a"))

	testCases := []struct {
		name        string
		tokenString string
		expected    bool
	}{
		{"Invalid token", "a", false},
		{"Valid token", validTokenString, true},
		{"Invalid signed token", invalidSignedTokenString, false},
		{"False token", falseTokenString, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if err := VerifyToken(tc.tokenString); err != tc.expected {
				t.Errorf("Expected token %s to be %v, got %v", tc.tokenString, tc.expected, err)
			}
		})
	}
}

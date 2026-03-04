package pkg

import (
	"crypto/rand"
	"crypto/rsa"
	"testing"

	"github.com/golang-jwt/jwt/v5"
)

func TestCreateToken(t *testing.T) {

	username := "test"

	tokenString, err := CreateToken(username)

	if err != nil {
		t.Errorf("Failed to create token: %v", err)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		t.Errorf("Generated token is invalid: %v", err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if claims["username"] != username {
			t.Errorf("Expected username claim to be %s, got %s", username, claims["username"])
		}

		if _, exists := claims["exp"]; !exists {
			t.Error("Expected exp claim to be present")
		}
	} else {
		t.Error("Failed to parse token claims")
	}
}

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

package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Custom claims struct
type CustomClaims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

var (
	// HMAC Secret key (in production, use environment variable)
	hmacSecret = []byte("your-256-bit-secret")

	// RSA keys for RSA256 signing
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

func init() {
	// Generate RSA key pair for demonstration
	var err error
	privateKey, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal("Failed to generate RSA key:", err)
	}
	publicKey = &privateKey.PublicKey
}

func main() {
	fmt.Println("🔐 JWT (JSON Web Token) Demo")
	fmt.Println("============================\n")

	// Demo 1: Basic HMAC Token
	fmt.Println("1. Basic HMAC Token Creation and Validation")
	fmt.Println("-------------------------------------------")
	basicHMACDemo()

	// Demo 2: Custom Claims
	fmt.Println("\n2. Custom Claims Example")
	fmt.Println("-------------------------")
	customClaimsDemo()

	// Demo 3: RSA Signing
	fmt.Println("\n3. RSA256 Signing Example")
	fmt.Println("--------------------------")
	rsaSigningDemo()

	// Demo 4: Token Expiration
	fmt.Println("\n4. Token Expiration Demo")
	fmt.Println("------------------------")
	expirationDemo()

	// Demo 5: Invalid Token Handling
	fmt.Println("\n5. Invalid Token Handling")
	fmt.Println("--------------------------")
	invalidTokenDemo()

	// Demo 6: Refresh Token Pattern
	fmt.Println("\n6. Refresh Token Pattern")
	fmt.Println("------------------------")
	refreshTokenDemo()
}

// Demo 1: Basic HMAC token creation and validation
func basicHMACDemo() {
	// Create a new token with HMAC SHA256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  "1234567890",
		"name": "John Doe",
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign the token with the secret
	tokenString, err := token.SignedString(hmacSecret)
	if err != nil {
		log.Fatal("Error signing token:", err)
	}

	fmt.Printf("Generated Token: %s\n", tokenString)

	// Parse and validate the token
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verify signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSecret, nil
	})

	if err != nil {
		log.Printf("Error parsing token: %v", err)
		return
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		fmt.Printf("✅ Token is valid!\n")
		fmt.Printf("Subject: %s\n", claims["sub"])
		fmt.Printf("Name: %s\n", claims["name"])
		fmt.Printf("Issued At: %v\n", time.Unix(int64(claims["iat"].(float64)), 0))
		fmt.Printf("Expires At: %v\n", time.Unix(int64(claims["exp"].(float64)), 0))
	}
}

// Demo 2: Custom claims with structured data
func customClaimsDemo() {
	// Create custom claims
	claims := CustomClaims{
		UserID:   123,
		Username: "john_doe",
		Role:     "admin",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "jwt-demo-app",
			Subject:   "user-auth",
			Audience:  []string{"web-app", "mobile-app"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Create token with custom claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(hmacSecret)
	if err != nil {
		log.Fatal("Error signing token:", err)
	}

	fmt.Printf("Generated Token with Custom Claims: %s\n", tokenString)

	// Parse with custom claims
	parsedToken, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})

	if err != nil {
		log.Printf("Error parsing token: %v", err)
		return
	}

	if claims, ok := parsedToken.Claims.(*CustomClaims); ok && parsedToken.Valid {
		fmt.Printf("✅ Custom claims token is valid!\n")
		fmt.Printf("User ID: %d\n", claims.UserID)
		fmt.Printf("Username: %s\n", claims.Username)
		fmt.Printf("Role: %s\n", claims.Role)
		fmt.Printf("Issuer: %s\n", claims.Issuer)
		fmt.Printf("Expires: %v\n", claims.ExpiresAt.Time)
	}
}

// Demo 3: RSA256 signing
func rsaSigningDemo() {
	// Create token with RSA256 signing
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub":  "1234567890",
		"name": "Jane Doe",
		"role": "user",
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Hour).Unix(),
	})

	// Sign with RSA private key
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		log.Fatal("Error signing RSA token:", err)
	}

	fmt.Printf("RSA256 Signed Token: %s\n", tokenString)

	// Validate with RSA public key
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verify signing method is RSA
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})

	if err != nil {
		log.Printf("Error parsing RSA token: %v", err)
		return
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		fmt.Printf("✅ RSA256 token is valid!\n")
		fmt.Printf("Subject: %s\n", claims["sub"])
		fmt.Printf("Name: %s\n", claims["name"])
		fmt.Printf("Role: %s\n", claims["role"])
	}

	// Display public key for verification (in production, this would be shared)
	publicKeyPEM := exportRSAPublicKeyAsPEMStr(publicKey)
	fmt.Printf("Public Key (PEM):\n%s\n", publicKeyPEM)
}

// Demo 4: Token expiration handling
func expirationDemo() {
	// Create a token that expires in 1 second
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": "1234567890",
		"exp": time.Now().Add(time.Second * 1).Unix(),
		"iat": time.Now().Unix(),
	})

	tokenString, err := token.SignedString(hmacSecret)
	if err != nil {
		log.Fatal("Error signing token:", err)
	}

	fmt.Printf("Token expires in 1 second...\n")

	// Validate immediately (should be valid)
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})

	if err != nil {
		fmt.Printf("❌ Immediate validation failed: %v\n", err)
	} else if parsedToken.Valid {
		fmt.Printf("✅ Token is currently valid\n")
	}

	// Wait for expiration
	fmt.Println("Waiting 2 seconds for token to expire...")
	time.Sleep(2 * time.Second)

	// Validate after expiration (should fail)
	parsedToken, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})

	if err != nil {
		fmt.Printf("❌ Expected expiration error: %v\n", err)
	} else {
		fmt.Printf("Unexpected: Token should have expired\n")
	}
}

// Demo 5: Invalid token handling
func invalidTokenDemo() {
	validToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
	invalidToken := "invalid.token.here"
	tamperedToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.TAMPERED_SIGNATURE"

	testCases := []struct {
		name  string
		token string
	}{
		{"Valid Token Format (but wrong secret)", validToken},
		{"Invalid Token Format", invalidToken},
		{"Tampered Signature", tamperedToken},
	}

	for _, tc := range testCases {
		fmt.Printf("Testing: %s\n", tc.name)
		_, err := jwt.Parse(tc.token, func(token *jwt.Token) (interface{}, error) {
			return hmacSecret, nil
		})

		if err != nil {
			fmt.Printf("❌ %s: %v\n", tc.name, err)
		} else {
			fmt.Printf("✅ %s: Token is valid (unexpected)\n", tc.name)
		}
	}
}

// Demo 6: Refresh token pattern
func refreshTokenDemo() {
	// Create access token (short lived)
	accessClaims := jwt.MapClaims{
		"sub":  "1234567890",
		"name": "John Doe",
		"type": "access",
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Minute * 15).Unix(), // 15 minutes
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString(hmacSecret)
	if err != nil {
		log.Fatal("Error creating access token:", err)
	}

	// Create refresh token (long lived)
	refreshClaims := jwt.MapClaims{
		"sub":  "1234567890",
		"type": "refresh",
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Hour * 24 * 7).Unix(), // 7 days
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString(hmacSecret)
	if err != nil {
		log.Fatal("Error creating refresh token:", err)
	}

	fmt.Printf("Access Token (15 min): %s\n", accessTokenString)
	fmt.Printf("Refresh Token (7 days): %s\n", refreshTokenString)

	// Simulate token refresh
	fmt.Println("\nSimulating token refresh...")
	newAccessToken := refreshAccessToken(refreshTokenString)
	if newAccessToken != "" {
		fmt.Printf("✅ New Access Token: %s\n", newAccessToken)
	}
}

// Helper function to refresh access token
func refreshAccessToken(refreshTokenString string) string {
	// Parse refresh token
	token, err := jwt.Parse(refreshTokenString, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})

	if err != nil {
		fmt.Printf("❌ Invalid refresh token: %v\n", err)
		return ""
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Verify it's a refresh token
		if tokenType, exists := claims["type"]; !exists || tokenType != "refresh" {
			fmt.Printf("❌ Not a refresh token\n")
			return ""
		}

		// Create new access token
		newClaims := jwt.MapClaims{
			"sub":  claims["sub"],
			"name": "John Doe", // In practice, fetch from database
			"type": "access",
			"iat":  time.Now().Unix(),
			"exp":  time.Now().Add(time.Minute * 15).Unix(),
		}

		newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
		tokenString, err := newToken.SignedString(hmacSecret)
		if err != nil {
			fmt.Printf("❌ Error creating new access token: %v\n", err)
			return ""
		}

		return tokenString
	}

	return ""
}

// Helper function to export RSA public key as PEM string
func exportRSAPublicKeyAsPEMStr(pubkey *rsa.PublicKey) string {
	pubkeyBytes, err := x509.MarshalPKIXPublicKey(pubkey)
	if err != nil {
		return ""
	}
	pubkeyPem := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubkeyBytes,
		},
	)
	return string(pubkeyPem)
}

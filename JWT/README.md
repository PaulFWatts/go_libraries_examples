# JWT (JSON Web Token) Demo

A comprehensive demonstration of the `github.com/golang-jwt/jwt/v5` library, showcasing JWT creation, validation, and security best practices in Go.

## 🚀 Features

This demo covers all major JWT operations:
- **Basic HMAC token creation and validation**
- **Custom claims with structured data**
- **RSA256 signing and verification**
- **Token expiration handling**
- **Invalid token detection and error handling**
- **Refresh token pattern implementation**
- **Security best practices**

## 📦 Dependencies

```bash
go get github.com/golang-jwt/jwt/v5
```

## 🔧 Setup

1. **Install dependencies:**
   ```bash
   go mod tidy
   ```

2. **Run the complete demo:**
   ```bash
   go run main.go
   ```

3. **Build executable:**
   ```bash
   go build -o jwt-demo main.go
   ```

## 📋 What It Demonstrates

### 1. Basic HMAC Token Operations
```go
// Create token with HMAC SHA256
token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "sub":  "1234567890",
    "name": "John Doe",
    "exp":  time.Now().Add(time.Hour * 24).Unix(),
})

// Sign the token
tokenString, err := token.SignedString(hmacSecret)

// Parse and validate
parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    return hmacSecret, nil
})
```

### 2. Custom Claims Structure
```go
type CustomClaims struct {
    UserID   int    `json:"user_id"`
    Username string `json:"username"`
    Role     string `json:"role"`
    jwt.RegisteredClaims
}
```

### 3. RSA256 Signing (Asymmetric)
- Public/private key pair generation
- Signing with private key
- Verification with public key
- PEM key format handling

### 4. Token Expiration Management
- Setting expiration times
- Handling expired tokens
- Real-time expiration testing

### 5. Security Features
- Invalid token detection
- Signature tampering detection
- Algorithm verification
- Proper error handling

### 6. Refresh Token Pattern
- Short-lived access tokens (15 minutes)
- Long-lived refresh tokens (7 days)
- Token refresh workflow

## 🔒 Security Best Practices Demonstrated

### 1. **Secret Management**
```go
// ❌ Don't do this in production
var hmacSecret = []byte("your-256-bit-secret")

// ✅ Use environment variables
hmacSecret := []byte(os.Getenv("JWT_SECRET"))
```

### 2. **Algorithm Verification**
```go
parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    // Always verify the signing method
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
        return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
    }
    return hmacSecret, nil
})
```

### 3. **Proper Claims Validation**
```go
// Use RegisteredClaims for standard JWT claims
RegisteredClaims: jwt.RegisteredClaims{
    Issuer:    "your-app",
    Subject:   "user-auth",
    Audience:  []string{"web-app"},
    ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
    NotBefore: jwt.NewNumericDate(time.Now()),
    IssuedAt:  jwt.NewNumericDate(time.Now()),
}
```

## 🎯 Demo Scenarios

### Scenario 1: Basic Authentication
- User logs in
- Server generates JWT with user claims
- Client includes JWT in subsequent requests
- Server validates JWT on each request

### Scenario 2: Role-Based Access Control
- JWT contains user role information
- Different endpoints check for different roles
- Centralized authorization logic

### Scenario 3: Token Refresh
- Access token expires frequently (15 minutes)
- Refresh token has longer lifespan (7 days)
- Client uses refresh token to get new access token

### Scenario 4: Microservices Architecture
- Central auth service issues JWTs
- Microservices validate JWTs independently
- Public key distribution for RSA verification

## 🔍 Common JWT Use Cases

### 1. **Stateless Authentication**
```go
// No server-side session storage needed
// All user info contained in token
claims := jwt.MapClaims{
    "user_id": 123,
    "username": "john_doe",
    "exp": time.Now().Add(time.Hour).Unix(),
}
```

### 2. **API Authorization**
```go
// Include permissions/scopes in JWT
claims := jwt.MapClaims{
    "sub": "user123",
    "permissions": []string{"read:posts", "write:posts"},
    "exp": time.Now().Add(time.Hour).Unix(),
}
```

### 3. **Single Sign-On (SSO)**
```go
// JWT can be shared across multiple applications
claims := jwt.MapClaims{
    "sub": "user123",
    "aud": []string{"app1.com", "app2.com", "app3.com"},
    "exp": time.Now().Add(time.Hour).Unix(),
}
```

## ⚠️ Security Considerations

### 1. **Token Storage**
- **Secure storage**: Use httpOnly cookies or secure storage
- **Avoid localStorage**: Vulnerable to XSS attacks
- **Consider token rotation**: Regular token refresh

### 2. **Secret Management**
- **Strong secrets**: Use cryptographically secure random bytes
- **Secret rotation**: Regularly rotate signing keys
- **Environment variables**: Never hardcode secrets

### 3. **Token Validation**
- **Always verify signature**: Don't trust client-side validation
- **Check expiration**: Implement proper token lifecycle
- **Validate audience**: Ensure token is for your application

### 4. **Algorithm Security**
- **Specify algorithms**: Don't allow "none" algorithm
- **Use strong algorithms**: HS256, RS256, ES256
- **Algorithm confusion**: Always verify expected algorithm

## 🛠️ Production Deployment Tips

### 1. **Environment Configuration**
```bash
# Environment variables for production
JWT_SECRET=your-very-long-and-secure-secret-key-here
JWT_EXPIRATION=3600
JWT_ISSUER=your-app-name
```

### 2. **Error Handling**
```go
// Proper error handling in production
if err != nil {
    switch err := err.(type) {
    case *jwt.ValidationError:
        if err.Errors&jwt.ValidationErrorExpired != 0 {
            // Token expired
            return errors.New("token expired")
        }
        if err.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
            // Signature invalid
            return errors.New("invalid signature")
        }
    }
    return err
}
```

### 3. **Middleware Integration**
```go
// Example middleware for Gin framework
func JWTMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        // Remove "Bearer " prefix
        tokenString = strings.TrimPrefix(tokenString, "Bearer ")
        
        token, err := jwt.Parse(tokenString, keyFunc)
        if err != nil || !token.Valid {
            c.JSON(401, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }
        
        // Add claims to context
        c.Set("claims", token.Claims)
        c.Next()
    }
}
```

## 📚 Additional Resources

- **JWT Specification**: [RFC 7519](https://tools.ietf.org/html/rfc7519)
- **JWT.io Debugger**: [https://jwt.io](https://jwt.io)
- **Library Documentation**: [github.com/golang-jwt/jwt](https://github.com/golang-jwt/jwt)
- **Security Best Practices**: [OWASP JWT Guide](https://cheatsheetseries.owasp.org/cheatsheets/JSON_Web_Token_for_Java_Cheat_Sheet.html)

## 🔄 Migration from v4 to v5

If migrating from jwt-go v4:

1. **Update import path**:
   ```go
   // Old
   "github.com/dgrijalva/jwt-go"
   
   // New
   "github.com/golang-jwt/jwt/v5"
   ```

2. **Update Claims interface**:
   ```go
   // Old
   jwt.StandardClaims
   
   // New
   jwt.RegisteredClaims
   ```

3. **Time handling changes**:
   ```go
   // Old
   ExpiresAt: time.Now().Add(time.Hour).Unix()
   
   // New
   ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour))
   ```

---

This demo provides a solid foundation for implementing JWT authentication in Go applications with proper security practices and real-world usage patterns.

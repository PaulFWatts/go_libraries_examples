# HTTPRouter Demo

A comprehensive demonstration of the `github.com/julienschmidt/httprouter` library, showcasing high-performance HTTP routing capabilities in Go.

## 🚀 Features

This demo covers all major httprouter features:
- **High-performance routing** with path parameters
- **RESTful API patterns** for users and products
- **Wildcard routing** for flexible path matching
- **Multiple path parameters** in single routes
- **Custom error handling** (404, 405, panics)
- **Middleware integration** with logging example
- **Method-specific routing** (GET, POST, PUT, DELETE)
- **Search functionality** with dynamic parameters
- **JSON API responses** with proper HTTP status codes

## 📦 Dependencies

```bash
go get github.com/julienschmidt/httprouter
```

## 🔧 Setup

1. **Install dependencies:**
   ```bash
   go mod tidy
   ```

2. **Run the server:**
   ```bash
   go run main.go
   ```

3. **Build executable:**
   ```bash
   go build -o httprouter-demo main.go
   ```

The server will start on `http://localhost:8080`

## 📋 API Endpoints

### 🏠 General Endpoints
- `GET /` - Home page with API information
- `GET /api` - Detailed API information
- `GET /health` - Health check endpoint

### 👥 User Management
- `GET /api/users` - Get all users
- `GET /api/users/:id` - Get user by ID
- `POST /api/users` - Create new user
- `PUT /api/users/:id` - Update existing user
- `DELETE /api/users/:id` - Delete user

### 📦 Product Management
- `GET /api/products` - Get all products
- `GET /api/products/by-id/:id` - Get product by ID
- `GET /api/products/by-category/:category` - Get products by category
- `POST /api/products` - Create new product
- `PUT /api/products/by-id/:id` - Update existing product
- `DELETE /api/products/by-id/:id` - Delete product

### 🔍 Search Functionality
- `GET /api/search/users/:query` - Search users by name, email, or username
- `GET /api/search/products/:query` - Search products by name, description, or category

### 🎯 Special Features
- `GET /api/wildcard/*filepath` - Wildcard route demonstration
- `GET /api/params/:category/:subcategory/:id` - Multiple parameters
- `GET /api/protected` - Protected endpoint with logging middleware
- `GET /api/panic` - Panic handler demonstration

## 🧪 Testing the API

### Using curl

#### Get all users:
```bash
curl http://localhost:8080/api/users
```

#### Get specific user:
```bash
curl http://localhost:8080/api/users/1
```

#### Get specific product:
```bash
curl http://localhost:8080/api/products/by-id/1
```

#### Create new user:
```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice Johnson","email":"alice@example.com","username":"alice_j"}'
```

#### Update user:
```bash
curl -X PUT http://localhost:8080/api/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"John Updated","email":"john.updated@example.com","username":"john_updated"}'
```

#### Update product:
```bash
curl -X PUT http://localhost:8080/api/products/by-id/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Updated Laptop","description":"High-performance updated laptop","price":1099.99,"category":"Electronics"}'
```

#### Search users:
```bash
curl http://localhost:8080/api/search/users/john
```

#### Get products by category:
```bash
curl http://localhost:8080/api/products/by-category/Electronics
```

#### Test wildcard route:
```bash
curl http://localhost:8080/api/wildcard/path/to/some/file.txt
```

#### Test multiple parameters:
```bash
curl http://localhost:8080/api/params/electronics/laptops/123
```

### Using PowerShell (Windows)

#### Get all users:
```powershell
Invoke-RestMethod -Uri "http://localhost:8080/api/users" -Method GET
```

#### Create new user:
```powershell
$body = @{
    name = "Alice Johnson"
    email = "alice@example.com"
    username = "alice_j"
} | ConvertTo-Json

Invoke-RestMethod -Uri "http://localhost:8080/api/users" -Method POST -Body $body -ContentType "application/json"
```

#### Search products:
```powershell
Invoke-RestMethod -Uri "http://localhost:8080/api/search/products/laptop" -Method GET
```

## 🎯 HTTPRouter Features Demonstrated

### 1. **Path Parameters**
```go
// Single parameter
router.GET("/api/users/:id", getUserByID)

// Multiple parameters
router.GET("/api/params/:category/:subcategory/:id", multiParamHandler)

// Access parameters in handler
func getUserByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    id := ps.ByName("id")
    // Handle the request
}
```

### 2. **Wildcard Routes**
```go
// Wildcard captures everything after the prefix
router.GET("/api/wildcard/*filepath", wildcardHandler)

func wildcardHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    filepath := ps.ByName("filepath")  // Gets the entire remaining path
}
```

### 3. **Method-Specific Routing**
```go
router.GET("/api/users", getUsers)
router.POST("/api/users", createUser)
router.PUT("/api/users/:id", updateUser)
router.DELETE("/api/users/:id", deleteUser)
```

### 4. **Custom Error Handling**
```go
// Handle 404 Not Found
router.NotFound = http.HandlerFunc(notFoundHandler)

// Handle 405 Method Not Allowed
router.MethodNotAllowed = http.HandlerFunc(methodNotAllowedHandler)

// Handle panics
router.PanicHandler = panicHandler
```

### 5. **Middleware Integration**
```go
// Custom middleware wrapper
func withLogging(next httprouter.Handle) httprouter.Handle {
    return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
        // Pre-processing
        start := time.Now()
        
        next(w, r, ps)  // Call the actual handler
        
        // Post-processing
        duration := time.Since(start)
        log.Printf("Request completed in %v", duration)
    }
}

// Use middleware
router.GET("/api/protected", withLogging(protectedEndpoint))
```

## 🏁 Performance Benefits

HTTPRouter provides several performance advantages:

### 1. **Zero Garbage Collection**
- No heap allocations during routing
- Minimal memory footprint per request

### 2. **Efficient Path Matching**
- Radix tree-based routing algorithm
- O(log n) lookup complexity
- No regular expressions (faster than regexp-based routers)

### 3. **Memory Efficient**
- Small memory footprint
- Optimized data structures

### 4. **Benchmarks**
HTTPRouter is one of the fastest HTTP routers for Go:
```
BenchmarkHttpRouter_Param        10000000    140 ns/op       0 B/op    0 allocs/op
BenchmarkHttpRouter_Param5       5000000     299 ns/op       0 B/op    0 allocs/op
BenchmarkHttpRouter_Param20      2000000     825 ns/op       0 B/op    0 allocs/op
```

## 🔧 Advanced Configuration

### Router Settings
```go
router := httprouter.New()

// Enable automatic redirection for trailing slashes
router.RedirectTrailingSlash = true

// Enable automatic redirection for fixed paths
router.RedirectFixedPath = true

// Handle OPTIONS requests automatically
router.HandleOPTIONS = true

// Handle method not allowed
router.HandleMethodNotAllowed = true
```

### Serving Static Files
```go
// Serve static files from a directory
router.ServeFiles("/static/*filepath", http.Dir("./static/"))

// This serves files from ./static/ directory at /static/ URL path
```

## 🛡️ Security Considerations

### 1. **Input Validation**
```go
func getUserByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    idStr := ps.ByName("id")
    
    // Always validate input
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID format", http.StatusBadRequest)
        return
    }
    
    // Continue with validated input
}
```

### 2. **Panic Recovery**
```go
router.PanicHandler = func(w http.ResponseWriter, r *http.Request, p interface{}) {
    // Log the panic
    log.Printf("Panic: %v", p)
    
    // Return safe error response
    http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}
```

### 3. **Content Type Validation**
```go
func createUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    if r.Header.Get("Content-Type") != "application/json" {
        http.Error(w, "Content-Type must be application/json", http.StatusBadRequest)
        return
    }
    
    // Process JSON
}
```

## 🆚 Comparison with Other Routers

| Feature | HTTPRouter | Gin | Gorilla/Mux | Standard Library |
|---------|------------|-----|-------------|------------------|
| Performance | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐⭐ | ⭐⭐ |
| Memory Usage | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐⭐ | ⭐⭐⭐⭐ |
| Path Parameters | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ❌ |
| Wildcards | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ❌ |
| Middleware | ⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐ |
| Zero Allocations | ✅ | ❌ | ❌ | ❌ |

## 📚 Real-World Use Cases

### 1. **Microservices**
- Fast routing for high-throughput services
- Minimal overhead for container deployments
- Excellent for API gateways

### 2. **REST APIs**
- Clean parameter extraction
- Efficient resource routing
- Standard HTTP method handling

### 3. **File Servers**
- Static file serving with wildcards
- Asset delivery systems
- CDN edge servers

### 4. **Proxy Services**
- Fast path matching for request forwarding
- Minimal latency overhead
- High concurrent connection handling

## 🔄 Migration Guide

### From Standard Library
```go
// Before (stdlib)
http.HandleFunc("/users/", userHandler)

// After (httprouter)
router.GET("/users/:id", getUserHandler)
```

### From Gorilla/Mux
```go
// Before (gorilla/mux)
r := mux.NewRouter()
r.HandleFunc("/users/{id}", getUserHandler).Methods("GET")

// After (httprouter)
router := httprouter.New()
router.GET("/users/:id", getUserHandler)
```

## 🏗️ Production Deployment

### Docker Example
```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o httprouter-demo main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/httprouter-demo .
EXPOSE 8080
CMD ["./httprouter-demo"]
```

### Systemd Service
```ini
[Unit]
Description=HTTPRouter Demo Service
After=network.target

[Service]
Type=simple
User=www-data
WorkingDirectory=/opt/httprouter-demo
ExecStart=/opt/httprouter-demo/httprouter-demo
Restart=always

[Install]
WantedBy=multi-user.target
```

---

This demo showcases httprouter's high-performance routing capabilities with practical examples that can be adapted for real-world applications. The zero-allocation design makes it ideal for high-throughput services and microservices architectures.

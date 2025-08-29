# Echo Web Framework Demo

A comprehensive demonstration of the `github.com/labstack/echo/v4` library, a high-performance, minimalist Go web framework known for its speed, simplicity, and developer-friendly API.

## 🚀 Features

This demo showcases all major Echo framework capabilities:
- **High-Performance Routing** - Zero memory allocation HTTP router
- **Middleware Support** - Built-in and custom middleware
- **RESTful API Design** - Complete CRUD operations
- **JSON Binding & Responses** - Automatic request/response handling
- **Path & Query Parameters** - Flexible parameter extraction
- **Error Handling** - Centralized error management
- **Static File Serving** - Built-in static content support
- **Template Rendering** - HTML template support
- **Cookie & Header Management** - Full HTTP feature support

## 📦 Dependencies

```bash
go get github.com/labstack/echo/v4
go get github.com/labstack/gommon
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

3. **Access the demo:**
   - Open your browser to `http://localhost:8080`
   - The home page provides a comprehensive overview of all available endpoints

## 📋 What It Demonstrates

### 1. High-Performance HTTP Server

#### Server Configuration
```go
e := echo.New()

// Built-in middleware
e.Use(middleware.Logger())
e.Use(middleware.Recover())
e.Use(middleware.CORS())

// Custom middleware for response timing
e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        start := time.Now()
        err := next(c)
        duration := time.Since(start)
        c.Response().Header().Set("X-Response-Time", duration.String())
        return err
    }
})
```

#### Route Grouping
```go
// API group with common prefix
api := e.Group("/api")

// User routes
users := api.Group("/users")
users.GET("", getAllUsers)
users.GET("/:id", getUserByID)
users.POST("", createUser)
users.PUT("/:id", updateUser)
users.DELETE("/:id", deleteUser)
```

### 2. RESTful API Implementation

#### Complete CRUD Operations
- **Users Management** - Create, read, update, delete users
- **Products Management** - Full product lifecycle management
- **Search Functionality** - Text-based search across entities
- **Category Filtering** - Product filtering by category

#### Request/Response Handling
```go
func createUser(c echo.Context) error {
    var newUser User
    if err := c.Bind(&newUser); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "error": "Invalid request body",
        })
    }
    
    // Process and return JSON response
    return c.JSON(http.StatusCreated, newUser)
}
```

### 3. Middleware Ecosystem

#### Built-in Middleware
- **Logger** - Request/response logging
- **Recover** - Panic recovery
- **CORS** - Cross-origin resource sharing
- **Static** - Static file serving

#### Custom Middleware
- **Response Time Tracking** - Performance monitoring
- **Request ID Generation** - Request tracing
- **Authentication** - Security middleware

### 4. Data Binding & Validation

#### JSON Binding
```go
type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

var user User
c.Bind(&user) // Automatic JSON to struct binding
```

#### Query & Path Parameters
```go
// Path parameters: /users/:id
userID := c.Param("id")

// Query parameters: /search?q=john
query := c.QueryParam("q")

// Multiple query params
name := c.QueryParam("name")
age := c.QueryParam("age")
```

### 5. HTTP Features

#### Cookie Management
```go
// Set cookie
cookie := &http.Cookie{
    Name:     "session_id",
    Value:    "abc123",
    Path:     "/",
    HttpOnly: true,
    MaxAge:   3600,
}
c.SetCookie(cookie)

// Read cookies
cookies := c.Cookies()
```

#### Header Handling
```go
// Set response headers
c.Response().Header().Set("X-API-Version", "1.0.0")

// Read request headers
userAgent := c.Request().Header.Get("User-Agent")
```

#### File Uploads
```go
func uploadFile(c echo.Context) error {
    file, err := c.FormFile("file")
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "error": "No file uploaded",
        })
    }
    // Process file...
    return c.JSON(http.StatusOK, map[string]interface{}{
        "filename": file.Filename,
        "size":     file.Size,
    })
}
```

## 🌐 Available Endpoints

### Core Application
- **GET** `/` - Interactive home page with endpoint documentation
- **GET** `/health` - Health check endpoint for monitoring

### User Management API
- **GET** `/api/users` - List all users with pagination info
- **GET** `/api/users/:id` - Get specific user by ID
- **POST** `/api/users` - Create new user (JSON body required)
- **PUT** `/api/users/:id` - Update existing user
- **DELETE** `/api/users/:id` - Delete user by ID

### Product Management API
- **GET** `/api/products` - List all products with metadata
- **GET** `/api/products/:id` - Get specific product by ID
- **GET** `/api/products/category/:category` - Filter products by category
- **POST** `/api/products` - Create new product (JSON body required)
- **PUT** `/api/products/:id` - Update existing product
- **DELETE** `/api/products/:id` - Delete product by ID

### Search & Discovery
- **GET** `/api/search/users?q={query}` - Search users by name/email
- **GET** `/api/search/products?q={query}` - Search products by name/description

### Feature Examples
- **GET** `/api/examples/json` - JSON response demonstration
- **GET** `/api/examples/params/:name/:age` - Path parameter extraction
- **GET** `/api/examples/query?name=X&age=Y` - Query parameter handling
- **GET** `/api/examples/cookie` - Cookie management demonstration
- **GET** `/api/examples/headers` - HTTP header handling
- **GET** `/api/examples/status?status=404` - Custom status code responses

### Utility Endpoints
- **POST** `/api/upload` - File upload demonstration
- **GET** `/api/error` - Error handling example
- **GET** `/template` - HTML template rendering

## 🧪 Testing the API

### Using cURL

#### Basic Operations
```bash
# Get all users
curl http://localhost:8080/api/users

# Get specific user
curl http://localhost:8080/api/users/1

# Health check
curl http://localhost:8080/health
```

#### Create New User
```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Alice Johnson",
    "email": "alice@example.com"
  }'
```

#### Update User
```bash
curl -X PUT http://localhost:8080/api/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Updated",
    "email": "john.updated@example.com"
  }'
```

#### Search Users
```bash
curl "http://localhost:8080/api/search/users?q=john"
```

#### Create Product
```bash
curl -X POST http://localhost:8080/api/products \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Gaming Mouse",
    "price": 59.99,
    "category": "Electronics",
    "description": "High-precision gaming mouse"
  }'
```

#### File Upload
```bash
curl -X POST http://localhost:8080/api/upload \
  -F "file=@example.txt"
```

### Using PowerShell (Windows)

#### GET Requests
```powershell
# Get users
Invoke-RestMethod -Uri "http://localhost:8080/api/users" -Method GET

# Search products
Invoke-RestMethod -Uri "http://localhost:8080/api/search/products?q=laptop" -Method GET
```

#### POST Requests
```powershell
# Create user
$body = @{
    name = "Bob Smith"
    email = "bob@example.com"
} | ConvertTo-Json

Invoke-RestMethod -Uri "http://localhost:8080/api/users" -Method POST -Body $body -ContentType "application/json"
```

### Using Browser
- Navigate to `http://localhost:8080` for the interactive documentation
- Click on any GET endpoint link to test directly in browser
- Use browser developer tools to inspect responses

## 🏗️ Architecture Patterns

### 1. Handler Pattern
```go
func getUserByID(c echo.Context) error {
    // 1. Extract parameters
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "error": "Invalid user ID",
        })
    }
    
    // 2. Business logic
    for _, user := range users {
        if user.ID == id {
            return c.JSON(http.StatusOK, user)
        }
    }
    
    // 3. Error response
    return c.JSON(http.StatusNotFound, map[string]string{
        "error": "User not found",
    })
}
```

### 2. Middleware Chain
```go
// Request flow: Logger -> CORS -> Custom -> Handler
e.Use(middleware.Logger())
e.Use(middleware.CORS())
e.Use(customTimingMiddleware)
```

### 3. Route Grouping
```go
// Organized by resource type
api := e.Group("/api")
users := api.Group("/users")    // /api/users/*
products := api.Group("/products") // /api/products/*
```

### 4. Error Handling
```go
// Consistent error responses
return echo.NewHTTPError(http.StatusInternalServerError, "Custom error message")

// JSON error responses
return c.JSON(http.StatusBadRequest, map[string]string{
    "error": "Validation failed",
    "field": "email",
})
```

## 🚀 Performance Features

### 1. Zero Memory Allocation Router
- Echo's router is designed for minimal memory allocation
- Radix tree implementation for fast route matching
- Path parameter extraction without memory allocation

### 2. Middleware Efficiency
- Middleware chain optimized for performance
- Built-in middleware written for speed
- Custom middleware examples showing best practices

### 3. JSON Processing
- Fast JSON binding and rendering
- Automatic content type detection
- Streaming support for large responses

### 4. Memory Management
- Minimal garbage collection pressure
- Efficient context reuse
- Connection pooling support

## 🛠️ Advanced Features

### 1. Custom Context
```go
type CustomContext struct {
    echo.Context
    UserID int
}

func (c *CustomContext) GetUserID() int {
    return c.UserID
}
```

### 2. Template Rendering
```go
// HTML template support
return c.HTML(http.StatusOK, `
<h1>Welcome {{.Name}}</h1>
<p>Email: {{.Email}}</p>
`)
```

### 3. WebSocket Support
```go
e.GET("/ws", websocketHandler)

func websocketHandler(c echo.Context) error {
    websocket.Handler(func(ws *websocket.Conn) {
        // WebSocket logic
    }).ServeHTTP(c.Response(), c.Request())
    return nil
}
```

### 4. Server-Sent Events
```go
e.GET("/events", func(c echo.Context) error {
    c.Response().Header().Set("Content-Type", "text/event-stream")
    c.Response().Header().Set("Cache-Control", "no-cache")
    c.Response().Header().Set("Connection", "keep-alive")
    
    // Send events...
    return nil
})
```

## 🆚 Echo vs Other Frameworks

| Feature | Echo | Gin | Fiber | net/http |
|---------|------|-----|-------|----------|
| Performance | Excellent | Excellent | Excellent | Good |
| Memory Usage | Low | Low | Very Low | Medium |
| Learning Curve | Easy | Easy | Easy | Moderate |
| Middleware | Rich | Rich | Rich | Basic |
| Documentation | Excellent | Good | Good | Excellent |
| Community | Large | Very Large | Growing | Huge |
| Flexibility | High | High | Medium | Very High |

## 🔧 Production Considerations

### 1. Security
```go
// HTTPS redirect
e.Use(middleware.HTTPSRedirect())

// Secure headers
e.Use(middleware.Secure())

// Rate limiting
e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
```

### 2. Logging
```go
// Custom logger
e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
    Format: "${time_rfc3339} ${status} ${method} ${uri} ${latency_human}\n",
}))
```

### 3. Error Handling
```go
// Global error handler
e.HTTPErrorHandler = func(err error, c echo.Context) {
    // Custom error handling logic
    code := http.StatusInternalServerError
    message := "Internal Server Error"
    
    if he, ok := err.(*echo.HTTPError); ok {
        code = he.Code
        message = he.Message.(string)
    }
    
    c.JSON(code, map[string]interface{}{
        "error": message,
        "timestamp": time.Now(),
    })
}
```

### 4. Health Monitoring
```go
e.GET("/metrics", metricsHandler)
e.GET("/health", healthCheckHandler)
```

## 📚 Real-World Applications

### 1. RESTful APIs
- Microservices architecture
- API gateways
- Backend for mobile applications
- Third-party integrations

### 2. Web Applications
- Single-page application backends
- Progressive web app APIs
- E-commerce platforms
- Content management systems

### 3. Real-time Systems
- WebSocket applications
- Server-sent events
- Chat systems
- Live data streaming

### 4. Proxy & Gateway
- Reverse proxy
- Load balancer
- API aggregation
- Protocol translation

## 🎯 Key Benefits

### Developer Experience
- **Simple API** - Intuitive and easy to learn
- **Great Documentation** - Comprehensive guides and examples
- **Active Community** - Large ecosystem and community support
- **Flexible Architecture** - Adaptable to various use cases

### Performance
- **High Throughput** - Handles thousands of requests per second
- **Low Latency** - Optimized for speed
- **Memory Efficient** - Minimal memory footprint
- **Scalable** - Handles concurrent connections efficiently

### Features
- **Middleware Rich** - Extensive middleware ecosystem
- **HTTP/2 Support** - Modern protocol support
- **WebSocket Ready** - Real-time communication support
- **Template Engine** - Multiple template engine support

---

This comprehensive Echo demo provides a solid foundation for building high-performance web applications and APIs in Go, showcasing the framework's speed, simplicity, and rich feature set.

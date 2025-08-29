package main

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

// User represents a user in our system
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Product represents a product in our system
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
}

// In-memory storage (in production, use a database)
var users = []User{
	{ID: 1, Name: "John Doe", Email: "john@example.com"},
	{ID: 2, Name: "Jane Smith", Email: "jane@example.com"},
	{ID: 3, Name: "Bob Johnson", Email: "bob@example.com"},
}

var products = []Product{
	{ID: 1, Name: "Laptop", Price: 999.99, Category: "Electronics", Description: "High-performance laptop"},
	{ID: 2, Name: "Coffee Mug", Price: 15.50, Category: "Kitchen", Description: "Ceramic coffee mug"},
	{ID: 3, Name: "Desk Chair", Price: 199.99, Category: "Furniture", Description: "Ergonomic office chair"},
}

func main() {
	// Create Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Custom middleware for request timing
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			err := next(c)
			duration := time.Since(start)
			c.Response().Header().Set("X-Response-Time", duration.String())
			return err
		}
	})

	// Set logger level
	e.Logger.SetLevel(log.INFO)

	// Routes
	setupRoutes(e)

	// Start server
	e.Logger.Info("Starting Echo server on :8080")
	e.Logger.Fatal(e.Start(":8080"))
}

func setupRoutes(e *echo.Echo) {
	// Basic routes
	e.GET("/", homeHandler)
	e.GET("/health", healthCheckHandler)

	// API group
	api := e.Group("/api")

	// User routes
	users := api.Group("/users")
	users.GET("", getAllUsers)
	users.GET("/:id", getUserByID)
	users.POST("", createUser)
	users.PUT("/:id", updateUser)
	users.DELETE("/:id", deleteUser)

	// Product routes
	products := api.Group("/products")
	products.GET("", getAllProducts)
	products.GET("/:id", getProductByID)
	products.GET("/category/:category", getProductsByCategory)
	products.POST("", createProduct)
	products.PUT("/:id", updateProduct)
	products.DELETE("/:id", deleteProduct)

	// Search routes
	e.GET("/api/search/users", searchUsers)
	e.GET("/api/search/products", searchProducts)

	// File upload example
	e.POST("/api/upload", uploadFile)

	// Custom error handling example
	e.GET("/api/error", errorHandler)

	// Template rendering example (using built-in HTML renderer)
	e.GET("/template", templateHandler)

	// JSON response examples
	e.GET("/api/examples/json", jsonExampleHandler)
	e.GET("/api/examples/status", statusExampleHandler)

	// Parameter and query examples
	e.GET("/api/examples/params/:name/:age", paramExampleHandler)
	e.GET("/api/examples/query", queryExampleHandler)

	// Cookie and header examples
	e.GET("/api/examples/cookie", cookieExampleHandler)
	e.GET("/api/examples/headers", headerExampleHandler)
}

// Handlers
func homeHandler(c echo.Context) error {
	return c.HTML(http.StatusOK, `
	<html>
		<head>
			<title>Echo Demo Server</title>
			<style>
				body { font-family: Arial, sans-serif; margin: 40px; }
				.container { max-width: 800px; margin: 0 auto; }
				.endpoint { background: #f5f5f5; padding: 15px; margin: 10px 0; border-radius: 5px; }
				.method { color: #2c5aa0; font-weight: bold; }
				.url { color: #d73027; }
				h1 { color: #333; }
				h2 { color: #666; border-bottom: 2px solid #eee; padding-bottom: 10px; }
				a { color: #2c5aa0; text-decoration: none; }
				a:hover { text-decoration: underline; }
			</style>
		</head>
		<body>
			<div class="container">
				<h1>🚀 Echo Web Framework Demo</h1>
				<p>A high-performance, minimalist Go web framework demonstration.</p>

				<h2>📋 Available Endpoints</h2>

				<div class="endpoint">
					<span class="method">GET</span> <span class="url">/</span> - This home page
				</div>

				<div class="endpoint">
					<span class="method">GET</span> <span class="url">/health</span> - Health check endpoint
				</div>

				<h3>👥 User Management</h3>
				<div class="endpoint">
					<span class="method">GET</span> <span class="url"><a href="/api/users">/api/users</a></span> - Get all users
				</div>
				<div class="endpoint">
					<span class="method">GET</span> <span class="url"><a href="/api/users/1">/api/users/1</a></span> - Get user by ID
				</div>
				<div class="endpoint">
					<span class="method">POST</span> <span class="url">/api/users</span> - Create new user
				</div>
				<div class="endpoint">
					<span class="method">PUT</span> <span class="url">/api/users/:id</span> - Update user
				</div>
				<div class="endpoint">
					<span class="method">DELETE</span> <span class="url">/api/users/:id</span> - Delete user
				</div>

				<h3>📦 Product Management</h3>
				<div class="endpoint">
					<span class="method">GET</span> <span class="url"><a href="/api/products">/api/products</a></span> - Get all products
				</div>
				<div class="endpoint">
					<span class="method">GET</span> <span class="url"><a href="/api/products/1">/api/products/1</a></span> - Get product by ID
				</div>
				<div class="endpoint">
					<span class="method">GET</span> <span class="url"><a href="/api/products/category/Electronics">/api/products/category/Electronics</a></span> - Get products by category
				</div>

				<h3>🔍 Search & Examples</h3>
				<div class="endpoint">
					<span class="method">GET</span> <span class="url"><a href="/api/search/users?q=john">/api/search/users?q=john</a></span> - Search users
				</div>
				<div class="endpoint">
					<span class="method">GET</span> <span class="url"><a href="/api/search/products?q=laptop">/api/search/products?q=laptop</a></span> - Search products
				</div>
				<div class="endpoint">
					<span class="method">GET</span> <span class="url"><a href="/api/examples/json">/api/examples/json</a></span> - JSON response examples
				</div>
				<div class="endpoint">
					<span class="method">GET</span> <span class="url"><a href="/api/examples/params/John/25">/api/examples/params/John/25</a></span> - Path parameters example
				</div>
				<div class="endpoint">
					<span class="method">GET</span> <span class="url"><a href="/api/examples/query?name=John&age=25">/api/examples/query?name=John&age=25</a></span> - Query parameters example
				</div>

				<h2>🧪 Testing the API</h2>
				<p>Use tools like curl, Postman, or your browser to test the endpoints:</p>
				<pre>
# Get all users
curl http://localhost:8080/api/users

# Create a new user
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Alice","email":"alice@example.com"}'

# Search users
curl "http://localhost:8080/api/search/users?q=john"
				</pre>

				<h2>💡 Echo Framework Features</h2>
				<ul>
					<li><strong>High Performance:</strong> Optimized HTTP router with zero memory allocation</li>
					<li><strong>Middleware:</strong> Built-in and custom middleware support</li>
					<li><strong>Data Binding:</strong> JSON, XML, form data binding</li>
					<li><strong>Template Rendering:</strong> Support for various template engines</li>
					<li><strong>Error Handling:</strong> Centralized HTTP error handling</li>
					<li><strong>Validation:</strong> Request validation with custom validators</li>
				</ul>
			</div>
		</body>
	</html>`)
}

func healthCheckHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":    "healthy",
		"timestamp": time.Now().Format(time.RFC3339),
		"service":   "echo-demo",
		"version":   "1.0.0",
	})
}

// User handlers
func getAllUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"users": users,
		"total": len(users),
	})
}

func getUserByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid user ID",
		})
	}

	for _, user := range users {
		if user.ID == id {
			return c.JSON(http.StatusOK, user)
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{
		"error": "User not found",
	})
}

func createUser(c echo.Context) error {
	var newUser User
	if err := c.Bind(&newUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	// Simple validation
	if newUser.Name == "" || newUser.Email == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Name and email are required",
		})
	}

	// Assign new ID (in production, use proper ID generation)
	newUser.ID = len(users) + 1
	users = append(users, newUser)

	return c.JSON(http.StatusCreated, newUser)
}

func updateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid user ID",
		})
	}

	var updatedUser User
	if err := c.Bind(&updatedUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	for i, user := range users {
		if user.ID == id {
			updatedUser.ID = id
			users[i] = updatedUser
			return c.JSON(http.StatusOK, updatedUser)
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{
		"error": "User not found",
	})
}

func deleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid user ID",
		})
	}

	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			return c.JSON(http.StatusOK, map[string]string{
				"message": "User deleted successfully",
			})
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{
		"error": "User not found",
	})
}

// Product handlers
func getAllProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"products": products,
		"total":    len(products),
	})
}

func getProductByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid product ID",
		})
	}

	for _, product := range products {
		if product.ID == id {
			return c.JSON(http.StatusOK, product)
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{
		"error": "Product not found",
	})
}

func getProductsByCategory(c echo.Context) error {
	category := c.Param("category")
	var categoryProducts []Product

	for _, product := range products {
		if product.Category == category {
			categoryProducts = append(categoryProducts, product)
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"products": categoryProducts,
		"category": category,
		"total":    len(categoryProducts),
	})
}

func createProduct(c echo.Context) error {
	var newProduct Product
	if err := c.Bind(&newProduct); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	// Simple validation
	if newProduct.Name == "" || newProduct.Price <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Name and valid price are required",
		})
	}

	newProduct.ID = len(products) + 1
	products = append(products, newProduct)

	return c.JSON(http.StatusCreated, newProduct)
}

func updateProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid product ID",
		})
	}

	var updatedProduct Product
	if err := c.Bind(&updatedProduct); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	for i, product := range products {
		if product.ID == id {
			updatedProduct.ID = id
			products[i] = updatedProduct
			return c.JSON(http.StatusOK, updatedProduct)
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{
		"error": "Product not found",
	})
}

func deleteProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid product ID",
		})
	}

	for i, product := range products {
		if product.ID == id {
			products = append(products[:i], products[i+1:]...)
			return c.JSON(http.StatusOK, map[string]string{
				"message": "Product deleted successfully",
			})
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{
		"error": "Product not found",
	})
}

// Search handlers
func searchUsers(c echo.Context) error {
	query := c.QueryParam("q")
	if query == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Query parameter 'q' is required",
		})
	}

	var results []User
	for _, user := range users {
		if containsIgnoreCase(user.Name, query) || containsIgnoreCase(user.Email, query) {
			results = append(results, user)
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"query":   query,
		"results": results,
		"total":   len(results),
	})
}

func searchProducts(c echo.Context) error {
	query := c.QueryParam("q")
	if query == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Query parameter 'q' is required",
		})
	}

	var results []Product
	for _, product := range products {
		if containsIgnoreCase(product.Name, query) || 
		   containsIgnoreCase(product.Category, query) || 
		   containsIgnoreCase(product.Description, query) {
			results = append(results, product)
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"query":   query,
		"results": results,
		"total":   len(results),
	})
}

// Example handlers
func uploadFile(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "No file uploaded",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"filename": file.Filename,
		"size":     file.Size,
		"message":  "File upload simulation (not actually saved)",
	})
}

func errorHandler(c echo.Context) error {
	// Demonstrate custom error handling
	return echo.NewHTTPError(http.StatusInternalServerError, "This is a demo error")
}

func templateHandler(c echo.Context) error {
	return c.HTML(http.StatusOK, `
	<h1>Template Example</h1>
	<p>This is rendered HTML content.</p>
	<p>Current time: `+time.Now().Format(time.RFC3339)+`</p>
	`)
}

func jsonExampleHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "This is a JSON response",
		"timestamp": time.Now().Format(time.RFC3339),
		"data": map[string]interface{}{
			"string":  "Hello, Echo!",
			"number":  42,
			"boolean": true,
			"array":   []string{"item1", "item2", "item3"},
			"object": map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
		},
	})
}

func statusExampleHandler(c echo.Context) error {
	statusCode := c.QueryParam("status")
	if statusCode == "" {
		statusCode = "200"
	}

	code, err := strconv.Atoi(statusCode)
	if err != nil {
		code = 200
	}

	return c.JSON(code, map[string]interface{}{
		"requested_status": code,
		"message":          http.StatusText(code),
	})
}

func paramExampleHandler(c echo.Context) error {
	name := c.Param("name")
	age := c.Param("age")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"path_parameters": map[string]string{
			"name": name,
			"age":  age,
		},
		"message": "These values came from the URL path",
	})
}

func queryExampleHandler(c echo.Context) error {
	name := c.QueryParam("name")
	age := c.QueryParam("age")
	hobby := c.QueryParam("hobby")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"query_parameters": map[string]string{
			"name":  name,
			"age":   age,
			"hobby": hobby,
		},
		"message": "These values came from query parameters",
	})
}

func cookieExampleHandler(c echo.Context) error {
	// Set a cookie
	cookie := &http.Cookie{
		Name:     "demo_cookie",
		Value:    "echo_framework_demo",
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // Set to true in production with HTTPS
		MaxAge:   3600,  // 1 hour
	}
	c.SetCookie(cookie)

	// Read existing cookies
	cookies := c.Cookies()
	cookieMap := make(map[string]string)
	for _, cookie := range cookies {
		cookieMap[cookie.Name] = cookie.Value
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":       "Cookie set and read",
		"cookie_set":    "demo_cookie=echo_framework_demo",
		"all_cookies":   cookieMap,
		"cookie_count":  len(cookies),
	})
}

func headerExampleHandler(c echo.Context) error {
	// Set custom headers
	c.Response().Header().Set("X-Custom-Header", "Echo-Demo-Value")
	c.Response().Header().Set("X-API-Version", "1.0.0")

	// Read request headers
	userAgent := c.Request().Header.Get("User-Agent")
	contentType := c.Request().Header.Get("Content-Type")
	acceptHeader := c.Request().Header.Get("Accept")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Headers demonstration",
		"request_headers": map[string]string{
			"User-Agent":   userAgent,
			"Content-Type": contentType,
			"Accept":       acceptHeader,
		},
		"response_headers_set": []string{
			"X-Custom-Header: Echo-Demo-Value",
			"X-API-Version: 1.0.0",
		},
	})
}

// Utility functions
func containsIgnoreCase(str, substr string) bool {
	return strings.Contains(strings.ToLower(str), strings.ToLower(substr))
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

// User represents a user in our system
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

// Product represents a product in our system
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
}

// In-memory storage for demo purposes
var users = []User{
	{ID: 1, Name: "John Doe", Email: "john@example.com", Username: "john_doe"},
	{ID: 2, Name: "Jane Smith", Email: "jane@example.com", Username: "jane_smith"},
	{ID: 3, Name: "Bob Johnson", Email: "bob@example.com", Username: "bob_johnson"},
}

var products = []Product{
	{ID: 1, Name: "Laptop", Description: "High-performance laptop", Price: 999.99, Category: "Electronics"},
	{ID: 2, Name: "Mouse", Description: "Wireless mouse", Price: 29.99, Category: "Electronics"},
	{ID: 3, Name: "Book", Description: "Programming guide", Price: 39.99, Category: "Books"},
	{ID: 4, Name: "Coffee", Description: "Premium coffee beans", Price: 19.99, Category: "Food"},
}

func main() {
	fmt.Println("🚀 HTTPRouter Demo Server")
	fmt.Println("=========================\n")

	// Create a new router instance
	router := httprouter.New()

	// Configure router settings
	configureRouter(router)

	// Register routes
	registerRoutes(router)

	// Display available endpoints
	displayEndpoints()

	// Start the server
	port := ":8080"
	fmt.Printf("🌐 Server starting on http://localhost%s\n", port)
	fmt.Println("📋 Try the endpoints listed above!")
	fmt.Println("🛑 Press Ctrl+C to stop the server\n")

	log.Fatal(http.ListenAndServe(port, router))
}

// Configure router settings
func configureRouter(router *httprouter.Router) {
	// Handle method not allowed
	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Method not allowed",
			"method":  r.Method,
			"path":    r.URL.Path,
			"message": "This endpoint does not support the " + r.Method + " method",
		})
	})

	// Handle not found
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error":   "Not found",
			"path":    r.URL.Path,
			"message": "The requested endpoint does not exist",
		})
	})

	// Panic handler
	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, p interface{}) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error":   "Internal server error",
			"message": "An unexpected error occurred",
			"panic":   fmt.Sprintf("%v", p),
		})
	}
}

// Register all routes
func registerRoutes(router *httprouter.Router) {
	// Root endpoint
	router.GET("/", home)

	// API info endpoint
	router.GET("/api", apiInfo)

	// User routes
	router.GET("/api/users", getUsers)
	router.GET("/api/users/:id", getUserByID)
	router.POST("/api/users", createUser)
	router.PUT("/api/users/:id", updateUser)
	router.DELETE("/api/users/:id", deleteUser)

	// Product routes
	router.GET("/api/products", getProducts)
	router.GET("/api/products/by-id/:id", getProductByID)
	router.GET("/api/products/by-category/:category", getProductsByCategory)
	router.POST("/api/products", createProduct)
	router.PUT("/api/products/by-id/:id", updateProduct)
	router.DELETE("/api/products/by-id/:id", deleteProduct)

	// Search routes
	router.GET("/api/search/users/:query", searchUsers)
	router.GET("/api/search/products/:query", searchProducts)

	// Special routes demonstrating httprouter features
	router.GET("/api/wildcard/*filepath", wildcardHandler)
	router.GET("/api/params/:category/:subcategory/:id", multiParamHandler)
	
	// Health check
	router.GET("/health", healthCheck)

	// Demo panic endpoint (for testing panic handler)
	router.GET("/api/panic", panicHandler)

	// Middleware demonstration
	router.GET("/api/protected", withLogging(protectedEndpoint))

	// Static file serving (if you had static files)
	// router.ServeFiles("/static/*filepath", http.Dir("static/"))
}

// Display available endpoints
func displayEndpoints() {
	fmt.Println("📡 Available Endpoints:")
	fmt.Println("=====================")
	
	endpoints := []struct {
		method string
		path   string
		desc   string
	}{
		{"GET", "/", "Home page"},
		{"GET", "/api", "API information"},
		{"GET", "/health", "Health check"},
		{"", "", ""},
		{"GET", "/api/users", "Get all users"},
		{"GET", "/api/users/:id", "Get user by ID"},
		{"POST", "/api/users", "Create new user"},
		{"PUT", "/api/users/:id", "Update user"},
		{"DELETE", "/api/users/:id", "Delete user"},
		{"", "", ""},
		{"GET", "/api/products", "Get all products"},
		{"GET", "/api/products/by-id/:id", "Get product by ID"},
		{"GET", "/api/products/by-category/:category", "Get products by category"},
		{"POST", "/api/products", "Create new product"},
		{"PUT", "/api/products/by-id/:id", "Update product"},
		{"DELETE", "/api/products/by-id/:id", "Delete product"},
		{"", "", ""},
		{"GET", "/api/search/users/:query", "Search users"},
		{"GET", "/api/search/products/:query", "Search products"},
		{"", "", ""},
		{"GET", "/api/wildcard/*filepath", "Wildcard demonstration"},
		{"GET", "/api/params/:cat/:subcat/:id", "Multiple parameters"},
		{"GET", "/api/protected", "Protected endpoint (with logging)"},
		{"GET", "/api/panic", "Panic handler demonstration"},
	}

	for _, ep := range endpoints {
		if ep.method == "" {
			fmt.Println()
			continue
		}
		fmt.Printf("  %-6s %-35s %s\n", ep.method, ep.path, ep.desc)
	}
	fmt.Println()
}

// Route handlers

// Home endpoint
func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"message":     "Welcome to HTTPRouter Demo API",
		"version":     "1.0.0",
		"server_time": time.Now().Format(time.RFC3339),
		"endpoints": map[string]string{
			"api_info":  "/api",
			"users":     "/api/users",
			"products":  "/api/products",
			"health":    "/health",
		},
	}
	json.NewEncoder(w).Encode(response)
}

// API info endpoint
func apiInfo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"name":        "HTTPRouter Demo API",
		"description": "Comprehensive demonstration of httprouter features",
		"version":     "1.0.0",
		"features": []string{
			"RESTful routing",
			"Path parameters",
			"Wildcard routing",
			"Method-specific handlers",
			"Custom error handling",
			"Panic recovery",
			"Middleware support",
		},
		"resources": []string{
			"users",
			"products",
			"search",
		},
	}
	json.NewEncoder(w).Encode(response)
}

// Health check endpoint
func healthCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"status":    "healthy",
		"timestamp": time.Now().Format(time.RFC3339),
		"uptime":    "running",
		"checks": map[string]string{
			"database": "ok",
			"memory":   "ok",
			"disk":     "ok",
		},
	}
	json.NewEncoder(w).Encode(response)
}

// User handlers

func getUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"users": users,
		"count": len(users),
	}
	json.NewEncoder(w).Encode(response)
}

func getUserByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	
	idStr := ps.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid user ID format",
		})
		return
	}

	for _, user := range users {
		if user.ID == id {
			json.NewEncoder(w).Encode(user)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"error": "User not found",
	})
}

func createUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	
	var newUser User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid JSON format",
		})
		return
	}

	// Generate new ID
	newUser.ID = len(users) + 1
	users = append(users, newUser)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func updateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	
	idStr := ps.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid user ID format",
		})
		return
	}

	var updatedUser User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid JSON format",
		})
		return
	}

	for i, user := range users {
		if user.ID == id {
			updatedUser.ID = id
			users[i] = updatedUser
			json.NewEncoder(w).Encode(updatedUser)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"error": "User not found",
	})
}

func deleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	
	idStr := ps.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid user ID format",
		})
		return
	}

	for i, user := range users {
		if user.ID == id {
			// Remove user from slice
			users = append(users[:i], users[i+1:]...)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "User deleted successfully",
			})
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"error": "User not found",
	})
}

// Product handlers

func getProducts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"products": products,
		"count":    len(products),
	}
	json.NewEncoder(w).Encode(response)
}

func getProductByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	
	idStr := ps.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid product ID format",
		})
		return
	}

	for _, product := range products {
		if product.ID == id {
			json.NewEncoder(w).Encode(product)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"error": "Product not found",
	})
}

func getProductsByCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	
	category := ps.ByName("category")
	var filteredProducts []Product
	
	for _, product := range products {
		if product.Category == category {
			filteredProducts = append(filteredProducts, product)
		}
	}

	response := map[string]interface{}{
		"category": category,
		"products": filteredProducts,
		"count":    len(filteredProducts),
	}
	json.NewEncoder(w).Encode(response)
}

func createProduct(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	
	var newProduct Product
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid JSON format",
		})
		return
	}

	// Generate new ID
	newProduct.ID = len(products) + 1
	products = append(products, newProduct)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newProduct)
}

func updateProduct(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	
	idStr := ps.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid product ID format",
		})
		return
	}

	var updatedProduct Product
	if err := json.NewDecoder(r.Body).Decode(&updatedProduct); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid JSON format",
		})
		return
	}

	for i, product := range products {
		if product.ID == id {
			updatedProduct.ID = id
			products[i] = updatedProduct
			json.NewEncoder(w).Encode(updatedProduct)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"error": "Product not found",
	})
}

func deleteProduct(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	
	idStr := ps.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid product ID format",
		})
		return
	}

	for i, product := range products {
		if product.ID == id {
			// Remove product from slice
			products = append(products[:i], products[i+1:]...)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "Product deleted successfully",
			})
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"error": "Product not found",
	})
}

// Search handlers

func searchUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	
	query := ps.ByName("query")
	var matchingUsers []User
	
	for _, user := range users {
		if containsIgnoreCase(user.Name, query) || 
		   containsIgnoreCase(user.Email, query) || 
		   containsIgnoreCase(user.Username, query) {
			matchingUsers = append(matchingUsers, user)
		}
	}

	response := map[string]interface{}{
		"query":   query,
		"users":   matchingUsers,
		"count":   len(matchingUsers),
	}
	json.NewEncoder(w).Encode(response)
}

func searchProducts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	
	query := ps.ByName("query")
	var matchingProducts []Product
	
	for _, product := range products {
		if containsIgnoreCase(product.Name, query) || 
		   containsIgnoreCase(product.Description, query) || 
		   containsIgnoreCase(product.Category, query) {
			matchingProducts = append(matchingProducts, product)
		}
	}

	response := map[string]interface{}{
		"query":    query,
		"products": matchingProducts,
		"count":    len(matchingProducts),
	}
	json.NewEncoder(w).Encode(response)
}

// Special feature handlers

func wildcardHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	
	filepath := ps.ByName("filepath")
	response := map[string]interface{}{
		"message":  "Wildcard route demonstration",
		"filepath": filepath,
		"note":     "The * captures everything after /api/wildcard/",
		"example":  "Try: /api/wildcard/path/to/some/file.txt",
	}
	json.NewEncoder(w).Encode(response)
}

func multiParamHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	
	category := ps.ByName("category")
	subcategory := ps.ByName("subcategory")
	id := ps.ByName("id")
	
	response := map[string]interface{}{
		"message":     "Multiple parameters demonstration",
		"category":    category,
		"subcategory": subcategory,
		"id":          id,
		"note":        "This route captures three different path parameters",
		"example":     "Try: /api/params/electronics/laptops/123",
	}
	json.NewEncoder(w).Encode(response)
}

func protectedEndpoint(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"message":   "This is a protected endpoint",
		"note":      "Check the server logs to see the logging middleware in action",
		"timestamp": time.Now().Format(time.RFC3339),
	}
	json.NewEncoder(w).Encode(response)
}

func panicHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	panic("This is a demonstration panic!")
}

// Middleware examples

func withLogging(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		start := time.Now()
		
		// Log the request
		fmt.Printf("🔍 [%s] %s %s - Started\n", 
			time.Now().Format("15:04:05"), 
			r.Method, 
			r.URL.Path)
		
		// Call the next handler
		next(w, r, ps)
		
		// Log the completion
		duration := time.Since(start)
		fmt.Printf("✅ [%s] %s %s - Completed in %v\n", 
			time.Now().Format("15:04:05"), 
			r.Method, 
			r.URL.Path, 
			duration)
	}
}

// Helper functions

func containsIgnoreCase(str, substr string) bool {
	str = fmt.Sprintf("%s", str)
	substr = fmt.Sprintf("%s", substr)
	
	// Simple case-insensitive contains check
	for i := 0; i <= len(str)-len(substr); i++ {
		match := true
		for j := 0; j < len(substr); j++ {
			if str[i+j] != substr[j] && str[i+j] != substr[j]+32 && str[i+j] != substr[j]-32 {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}

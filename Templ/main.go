package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"templ-demo/templates"

	"github.com/gorilla/mux"
)

// Use the same types as defined in templates
type Todo = templates.Todo
type User = templates.User

// Global data stores (in a real app, you'd use a database)
var (
	todos  []Todo
	users  []User
	nextID = 1
	userID = 1
)

func init() {
	// Initialize with sample data
	todos = []Todo{
		{ID: nextID, Text: "Learn Templ basics", Completed: true},
		{ID: nextID + 1, Text: "Build a demo application", Completed: false},
		{ID: nextID + 2, Text: "Create reusable components", Completed: false},
		{ID: nextID + 3, Text: "Add interactive features", Completed: false},
	}
	nextID = 5

	users = []User{
		{ID: userID, Name: "Alice Johnson", Email: "alice@example.com", Role: "Admin"},
		{ID: userID + 1, Name: "Bob Smith", Email: "bob@example.com", Role: "User"},
		{ID: userID + 2, Name: "Carol Davis", Email: "carol@example.com", Role: "Editor"},
		{ID: userID + 3, Name: "David Wilson", Email: "david@example.com", Role: "User"},
	}
	userID = 5
}

func main() {
	r := mux.NewRouter()

	// Middleware for logging
	r.Use(loggingMiddleware)

	// Static routes
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/components", componentsHandler).Methods("GET")
	r.HandleFunc("/todo", todoHandler).Methods("GET")
	r.HandleFunc("/users", usersHandler).Methods("GET")

	// API routes for Todo
	r.HandleFunc("/todo", createTodoHandler).Methods("POST")
	r.HandleFunc("/todo/{id}", updateTodoHandler).Methods("PUT")
	r.HandleFunc("/todo/{id}", deleteTodoHandler).Methods("DELETE")
	r.HandleFunc("/todo/{id}/toggle", toggleTodoHandler).Methods("POST")

	// API routes for Users
	r.HandleFunc("/users", createUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", updateUserHandler).Methods("PUT")
	r.HandleFunc("/users/{id}", deleteUserHandler).Methods("DELETE")

	// Contact form handler
	r.HandleFunc("/contact", contactHandler).Methods("POST")

	// Health check
	r.HandleFunc("/health", healthHandler).Methods("GET")

	fmt.Println("🚀 Templ Demo Server Starting...")
	fmt.Println("📍 Server running on http://localhost:8080")
	fmt.Println("🌐 Open your browser and navigate to:")
	fmt.Println("   • Home: http://localhost:8080")
	fmt.Println("   • Components: http://localhost:8080/components")
	fmt.Println("   • Todo App: http://localhost:8080/todo")
	fmt.Println("   • Users: http://localhost:8080/users")
	fmt.Println("")

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
		// Add timeouts for production readiness
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}

// Handlers
func homeHandler(w http.ResponseWriter, r *http.Request) {
	component := templates.HomePage()
	component.Render(r.Context(), w)
}

func componentsHandler(w http.ResponseWriter, r *http.Request) {
	// For now, redirect to home (we can add a components page later)
	http.Redirect(w, r, "/", http.StatusFound)
}

func todoHandler(w http.ResponseWriter, r *http.Request) {
	component := templates.TodoPage(todos)
	component.Render(r.Context(), w)
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	// For now, return a simple JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Users page coming soon!",
		"users":   users,
	})
}

// Todo API handlers
func createTodoHandler(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	if text == "" {
		http.Error(w, "Text is required", http.StatusBadRequest)
		return
	}

	todo := Todo{
		ID:        nextID,
		Text:      text,
		Completed: false,
	}
	todos = append(todos, todo)
	nextID++

	http.Redirect(w, r, "/todo", http.StatusSeeOther)
}

func updateTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid todo ID", http.StatusBadRequest)
		return
	}

	var requestData struct {
		Text string `json:"text"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Text = requestData.Text
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	http.Error(w, "Todo not found", http.StatusNotFound)
}

func deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid todo ID", http.StatusBadRequest)
		return
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	http.Error(w, "Todo not found", http.StatusNotFound)
}

func toggleTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid todo ID", http.StatusBadRequest)
		return
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Completed = !todos[i].Completed
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	http.Error(w, "Todo not found", http.StatusNotFound)
}

// User API handlers
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	role := r.FormValue("role")

	if name == "" || email == "" || role == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	user := User{
		ID:    userID,
		Name:  name,
		Email: email,
		Role:  role,
	}
	users = append(users, user)
	userID++

	http.Redirect(w, r, "/users", http.StatusSeeOther)
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var requestData struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	for i, user := range users {
		if user.ID == id {
			users[i].Name = requestData.Name
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	http.Error(w, "User not found", http.StatusNotFound)
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	http.Error(w, "User not found", http.StatusNotFound)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	message := r.FormValue("message")

	// In a real app, you'd save this to a database or send an email
	fmt.Printf("📧 Contact Form Submission:\n")
	fmt.Printf("   Name: %s\n", name)
	fmt.Printf("   Email: %s\n", email)
	fmt.Printf("   Message: %s\n", message)

	// Return a simple success response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Thank you for your message! We'll get back to you soon.",
	})
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"status":    "healthy",
		"timestamp": time.Now().Format(time.RFC3339),
		"version":   "1.0.0",
		"data": map[string]int{
			"todos": len(todos),
			"users": len(users),
		},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Middleware
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)

		fmt.Printf("📝 %s %s - %v\n",
			r.Method,
			r.URL.Path,
			duration,
		)
	})
}

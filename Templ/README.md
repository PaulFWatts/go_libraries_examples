# Templ Demo - Type-safe HTML Templates for Go

A comprehensive demonstration of [Templ](https://templ.guide/) - a language for writing HTML user interfaces in Go that compiles to type-safe Go code.

## Overview

Templ is a revolutionary templating system that brings type safety and compile-time validation to HTML templates in Go. Unlike traditional text-based templates that fail at runtime, Templ templates are validated at compile time, ensuring your web applications are robust and error-free.

## Features Demonstrated

### 🔒 **Type Safety**
- [x] Compile-time HTML validation
- [x] Type-safe component props
- [x] No runtime template parsing errors
- [x] IDE support with syntax highlighting

### 🧩 **Component Architecture**
- [x] Reusable UI components with props
- [x] Component composition and nesting
- [x] Conditional rendering
- [x] Loop rendering with data iteration

### ⚡ **Performance**
- [x] Zero-overhead HTML generation
- [x] Templates compile to efficient Go code
- [x] No runtime parsing or interpretation
- [x] Minimal memory allocations

### 🛠️ **Modern Web Development**
- [x] Bootstrap CSS integration
- [x] Responsive design patterns
- [x] Form handling and validation
- [x] REST API integration

## Project Structure

```
Templ/
├── main.go                     # Web server with HTTP handlers
├── go.mod                     # Module dependencies
├── go.sum                     # Dependency checksums
├── README.md                  # This documentation
└── templates/                 # Templ template files
    ├── base.templ            # Base layout template
    ├── base_templ.go         # Generated Go code from base.templ
    ├── simple_home.templ     # Home page template
    ├── simple_home_templ.go  # Generated Go code
    ├── simple_todo.templ     # Todo app template
    └── simple_todo_templ.go  # Generated Go code
```

## Installation & Setup

### Prerequisites
- Go 1.21 or higher
- Templ CLI tool

### Quick Setup

1. **Install Templ CLI:**
   ```bash
   go install github.com/a-h/templ/cmd/templ@latest
   ```

2. **Clone and navigate:**
   ```bash
   cd Templ
   ```

3. **Install dependencies:**
   ```bash
   go mod tidy
   ```

4. **Generate templates (if needed):**
   ```bash
   templ generate
   ```

5. **Run the server:**
   ```bash
   go run main.go
   ```

6. **Open in browser:**
   - Home: http://localhost:8080
   - Todo App: http://localhost:8080/todo
   - Health Check: http://localhost:8080/health

## Template Examples

### 1. Basic Template Structure

```go
// simple_home.templ
package templates

templ HomePage() {
    <!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8"/>
        <title>Templ Demo</title>
    </head>
    <body>
        <h1>Welcome to Templ</h1>
        <p>Type-safe HTML templating for Go!</p>
    </body>
    </html>
}
```

### 2. Component with Props

```go
// Component definition
templ TodoItem(todo Todo) {
    <div class={ "card mb-2", getCompletedClass(todo.Completed) }>
        <div class="card-body">
            if todo.Completed {
                <s class="text-muted">{ todo.Text }</s>
            } else {
                <span>{ todo.Text }</span>
            }
        </div>
    </div>
}

// Helper function
func getCompletedClass(completed bool) string {
    if completed {
        return "border-success"
    }
    return "border-primary"
}
```

### 3. Loop Rendering

```go
templ TodoList(todos []Todo) {
    <div class="todo-container">
        if len(todos) == 0 {
            <p>No todos yet!</p>
        } else {
            for _, todo := range todos {
                @TodoItem(todo)
            }
        }
    </div>
}
```

### 4. Conditional Rendering

```go
templ Button(text, variant string, disabled bool) {
    <button 
        class={ "btn", "btn-" + variant }
        if disabled {
            disabled
        }
    >
        { text }
    </button>
}
```

## Application Features

### 🏠 **Home Page** (`http://localhost:8080`)
- Hero section with gradient background
- Feature showcase cards
- Responsive Bootstrap layout
- Clean, professional design

### ✅ **Todo Application** (`http://localhost:8080/todo`)
- **Create Todos**: Add new todo items via form submission
- **View Todos**: Display all todos with status indicators
- **Statistics**: Real-time counters for total, completed, and pending todos
- **Interactive UI**: Bootstrap styling with hover effects
- **Type Safety**: All todo operations are type-safe

### 🔧 **API Endpoints**
- `POST /todo` - Create new todo
- `PUT /todo/{id}` - Update todo text
- `DELETE /todo/{id}` - Delete todo
- `POST /todo/{id}/toggle` - Toggle completion status

## Key Templ Concepts Demonstrated

### 1. **Template Generation**
Templates are compiled to Go functions:
```bash
# Input: simple_home.templ
templ generate

# Output: simple_home_templ.go
func HomePage() templ.Component {
    return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
        // Generated HTML writing code
    })
}
```

### 2. **Type-Safe Props**
```go
// Template signature enforces types
templ UserCard(name string, email string, isActive bool) {
    <div class="user-card">
        <h3>{ name }</h3>
        <p>{ email }</p>
        if isActive {
            <span class="badge bg-success">Active</span>
        }
    </div>
}
```

### 3. **Component Composition**
```go
// Base layout
templ Base(title string) {
    <!DOCTYPE html>
    <html>
    <head><title>{ title }</title></head>
    <body>{ children... }</body>
    </html>
}

// Page using base
templ AboutPage() {
    @Base("About Us") {
        <h1>About Our Company</h1>
        <p>Content goes here...</p>
    }
}
```

### 4. **HTTP Integration**
```go
func homeHandler(w http.ResponseWriter, r *http.Request) {
    component := templates.HomePage()
    component.Render(r.Context(), w)
}
```

## Data Flow

1. **HTTP Request** → Router (Gorilla Mux)
2. **Handler Function** → Prepares data
3. **Templ Template** → Type-safe rendering
4. **Generated Go Code** → Efficient HTML output
5. **HTTP Response** → Browser

## Development Workflow

### 1. **Create Template**
```bash
# Create new .templ file
touch templates/new_page.templ
```

### 2. **Write Template Code**
```go
package templates

templ NewPage(title string) {
    <h1>{ title }</h1>
}
```

### 3. **Generate Go Code**
```bash
templ generate
```

### 4. **Use in Handler**
```go
func newPageHandler(w http.ResponseWriter, r *http.Request) {
    templates.NewPage("Hello World").Render(r.Context(), w)
}
```

### 5. **Test and Iterate**
```bash
go run main.go
# Open browser and test
```

## Testing the Demo

### 1. **Start the Server**
```bash
go run main.go
```

### 2. **Navigate to Pages**
- **Home**: http://localhost:8080
- **Todo App**: http://localhost:8080/todo
- **Health Check**: http://localhost:8080/health

### 3. **Test Todo Functionality**
1. Add a new todo item
2. View the updated statistics
3. Check the responsive design on different screen sizes

### 4. **API Testing with curl**
```bash
# Create a todo
curl -X POST -d "text=Test todo" http://localhost:8080/todo

# Toggle todo completion
curl -X POST http://localhost:8080/todo/1/toggle

# Delete a todo
curl -X DELETE http://localhost:8080/todo/1
```

## Advanced Features

### 1. **CSS Classes with Logic**
```go
templ Alert(message string, isError bool) {
    <div class={ "alert", templ.KV("alert-danger", isError), templ.KV("alert-success", !isError) }>
        { message }
    </div>
}
```

### 2. **Attributes with Conditions**
```go
templ Input(name string, required bool) {
    <input 
        type="text" 
        name={ name }
        if required {
            required
        }
    />
}
```

### 3. **Safe HTML Rendering**
```go
templ RawHTML(content string) {
    @templ.Raw(content)  // Caution: Only use with trusted content
}
```

## Performance Benefits

### **Compile-Time Advantages**
- ✅ HTML validation at build time
- ✅ Type checking for all variables
- ✅ Missing template detection
- ✅ Unused template cleanup

### **Runtime Advantages**
- ⚡ Zero template parsing overhead
- ⚡ Direct Go function calls
- ⚡ Efficient memory usage
- ⚡ Fast HTML generation

### **Developer Experience**
- 🛠️ IDE autocomplete and syntax highlighting
- 🛠️ Compile-time error detection
- 🛠️ Refactoring support
- 🛠️ Hot reload during development

## Comparison with Traditional Templates

| Feature | Traditional Go Templates | Templ |
|---------|-------------------------|-------|
| **Type Safety** | ❌ Runtime errors | ✅ Compile-time validation |
| **Performance** | 🐌 Runtime parsing | ⚡ Compiled functions |
| **IDE Support** | ❌ Limited | ✅ Full Go support |
| **Refactoring** | ❌ Difficult | ✅ IDE-assisted |
| **Component Reuse** | ❌ Text inclusion | ✅ Function composition |

## Common Use Cases

### 1. **Web Applications**
- Server-side rendered pages
- Progressive web apps
- Admin dashboards

### 2. **API Documentation**
- Interactive API explorers
- Swagger UI alternatives
- Developer portals

### 3. **Email Templates**
- Type-safe HTML emails
- Newsletter templates
- Notification emails

### 4. **Component Libraries**
- Reusable UI components
- Design system implementation
- Themed templates

## Best Practices Demonstrated

### 1. **File Organization**
```
templates/
├── components/     # Reusable components
├── pages/         # Full page templates
├── layouts/       # Base layouts
└── partials/      # Small template pieces
```

### 2. **Naming Conventions**
```go
// Page templates: noun + "Page"
templ HomePage() { }

// Components: descriptive name
templ UserCard(user User) { }

// Layouts: noun + "Layout"  
templ BaseLayout(title string) { }
```

### 3. **Type Definitions**
```go
// Define types in template files
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

// Use throughout templates
templ UserList(users []User) { }
```

## Troubleshooting

### **Common Issues**

1. **Template Generation Fails**
   ```bash
   templ generate
   # Check for syntax errors in .templ files
   ```

2. **Import Path Errors**
   ```bash
   go mod tidy
   # Ensure module paths are correct
   ```

3. **Type Mismatches**
   ```go
   // Ensure struct definitions match between packages
   type User = templates.User
   ```

### **Debug Tips**

1. **Use `templ generate -v` for verbose output**
2. **Check generated `.go` files for issues**
3. **Validate HTML structure in browser dev tools**

## Resources

- **Templ Official**: https://templ.guide/
- **GitHub Repository**: https://github.com/a-h/templ
- **Examples**: https://github.com/a-h/templ/tree/main/examples
- **VSCode Extension**: Search for "templ" in extensions

---

This demo showcases Templ as a modern, type-safe alternative to traditional Go templating solutions, bringing the reliability of Go's type system to HTML generation while maintaining excellent performance and developer experience.

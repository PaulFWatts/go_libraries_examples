# MapStructure Library Demo

A comprehensive demonstration of the `mitchellh/mapstructure` library, which provides functionality for decoding generic map values to native Go structures and vice versa.

## 🚀 Features

This demo covers all major features of the mapstructure library:
- **Basic map to struct conversion**
- **JSON to struct via map intermediary**
- **Nested structure handling**
- **Custom field mapping with tags**
- **Type conversion and custom hooks**
- **Slice and array processing**
- **Error handling and validation**
- **Advanced configuration options**
- **Real-world usage examples**

## 📦 Installation

```bash
go get github.com/mitchellh/mapstructure
```

## 🔧 Setup

1. **Run the demo:**
   ```bash
   go run main.go
   ```

2. **Build executable:**
   ```bash
   go build -o mapstructure-demo main.go
   ```

## 📋 What It Demonstrates

### 1. Basic Map to Struct Conversion
```go
type Person struct {
    Name string
    Age  int
    City string
}

input := map[string]interface{}{
    "name": "John Doe",
    "age":  30,
    "city": "New York",
}

var result Person
err := mapstructure.Decode(input, &result)
```

### 2. JSON to Struct via Map
```go
// Parse JSON to map first
var inputMap map[string]interface{}
json.Unmarshal(jsonData, &inputMap)

// Then convert map to struct
var product Product
mapstructure.Decode(inputMap, &product)
```

### 3. Custom Field Mapping
```go
type Config struct {
    DatabaseURL string `mapstructure:"db_url"`
    Port        int    `mapstructure:"server_port"`
    Debug       bool   `mapstructure:"debug_mode"`
}
```

### 4. Type Conversion Hooks
```go
timeHook := func(from reflect.Type, to reflect.Type, data interface{}) (interface{}, error) {
    if to == reflect.TypeOf(time.Time{}) && from == reflect.TypeOf("") {
        return time.Parse("2006-01-02 15:04:05", data.(string))
    }
    return data, nil
}

config := &mapstructure.DecoderConfig{
    DecodeHook: timeHook,
    Result:     &result,
}
```

### 5. Slice and Array Handling
```go
type Team struct {
    Name    string   `mapstructure:"name"`
    Members []string `mapstructure:"members"`
    Tasks   []Task   `mapstructure:"tasks"` // Nested struct slices
}
```

### 6. Advanced Configuration
```go
// Capture unknown fields
type FlexibleStruct struct {
    KnownField    string                 `mapstructure:"known"`
    UnknownFields map[string]interface{} `mapstructure:",remain"`
}

// Embedded struct flattening
type ContainerStruct struct {
    EmbeddedStruct `mapstructure:",squash"`
    Extra          string `mapstructure:"extra"`
}
```

## 🎯 Sample Output

```
🗺️ MapStructure Library Demo
=============================

1. 📦 Basic Map to Struct
   📊 Input map: map[age:30 city:New York name:John Doe]
   ✅ Result struct: {Name:John Doe Age:30 City:New York}

2. 🔄 JSON to Struct via Map
   📄 JSON: {"id": 123, "name": "Laptop", "price": 999.99}
   🗺️ Map: map[id:123 name:Laptop price:999.99]
   📦 Struct: {ID:123 Name:Laptop Price:999.99}

3. 🏗️ Nested Structures
   👤 User: {Name:Alice Smith Email:alice@example.com Address:{Street:123 Main St City:Boston ZipCode:02101}}
   🏠 Address: {Street:123 Main St City:Boston ZipCode:02101}

4. 🏷️ Custom Field Mapping
   🔧 Config: {DatabaseURL:postgresql://localhost:5432/mydb Port:8080 Debug:true Timeout:30}

5. 🔧 Type Conversion & Hooks
   📅 Event: {Name:Meeting Timestamp:2024-12-25 14:30:00 +0000 UTC Duration:2h30m0s}
   ⏰ Timestamp: 2024-12-25 14:30:00
   ⏱️ Duration: 2h30m0s
```

## 🛠️ Common Use Cases

### Configuration File Parsing
Perfect for parsing YAML, JSON, or TOML configuration files into Go structs:
```go
type AppConfig struct {
    Server   ServerConfig   `mapstructure:"server"`
    Database DatabaseConfig `mapstructure:"database"`
    Redis    RedisConfig    `mapstructure:"redis"`
}

// Parse config file to map, then to struct
var configMap map[string]interface{}
// ... load from file ...
var config AppConfig
mapstructure.Decode(configMap, &config)
```

### API Response Processing
Convert API responses from `interface{}` to typed structs:
```go
type APIResponse struct {
    Status   string      `mapstructure:"status"`
    Data     []Item      `mapstructure:"data"`
    Metadata Pagination  `mapstructure:"metadata"`
}

// Parse JSON response
var response map[string]interface{}
json.Unmarshal(apiData, &response)

// Convert to struct
var apiResp APIResponse
mapstructure.Decode(response, &apiResp)
```

### Environment Variable Processing
Combined with libraries like Viper for environment configuration:
```go
type EnvConfig struct {
    Port     int    `mapstructure:"PORT"`
    Host     string `mapstructure:"HOST"`
    LogLevel string `mapstructure:"LOG_LEVEL"`
}

envMap := viper.AllSettings()
var config EnvConfig
mapstructure.Decode(envMap, &config)
```

### Database Result Mapping
Convert database query results to structs:
```go
type User struct {
    ID       int       `mapstructure:"id"`
    Name     string    `mapstructure:"name"`
    Email    string    `mapstructure:"email"`
    Created  time.Time `mapstructure:"created_at"`
}

// From database query result
rows, _ := db.Query("SELECT id, name, email, created_at FROM users")
// ... process rows to map ...
var users []User
mapstructure.Decode(resultMaps, &users)
```

## 🔧 Advanced Features

### Custom Decode Hooks
```go
// String to time.Time conversion
func stringToTimeHookFunc() mapstructure.DecodeHookFunc {
    return func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
        if f.Kind() != reflect.String {
            return data, nil
        }
        if t != reflect.TypeOf(time.Time{}) {
            return data, nil
        }
        
        return time.Parse("2006-01-02", data.(string))
    }
}

// String to duration conversion
mapstructure.StringToTimeDurationHookFunc()

// String to slice conversion  
mapstructure.StringToSliceHookFunc(",")
```

### Decoder Configuration Options
```go
config := &mapstructure.DecoderConfig{
    // Custom decode hooks
    DecodeHook: mapstructure.ComposeDecodeHookFunc(
        mapstructure.StringToTimeDurationHookFunc(),
        mapstructure.StringToSliceHookFunc(","),
    ),
    
    // Allow weak type conversion (string "123" -> int 123)
    WeaklyTypedInput: true,
    
    // Squash embedded structs
    Squash: true,
    
    // Zero out destination before decoding
    ZeroFields: true,
    
    // Error on unused keys
    ErrorUnused: true,
    
    // Custom tag name (default: "mapstructure")
    TagName: "json",
    
    // Result destination
    Result: &result,
}

decoder, err := mapstructure.NewDecoder(config)
```

### Error Handling
```go
// Type conversion errors
err := mapstructure.Decode(input, &result)
if err != nil {
    if strings.Contains(err.Error(), "cannot parse") {
        // Handle parse error
    }
    if strings.Contains(err.Error(), "has unused keys") {
        // Handle unused keys
    }
}

// Validation after decoding
if result.Port <= 0 {
    return fmt.Errorf("invalid port: %d", result.Port)
}
```

## 🎯 Best Practices

1. **Use struct tags** for clear field mapping
2. **Implement custom hooks** for complex type conversions
3. **Handle errors gracefully** - don't ignore decode errors
4. **Validate after decoding** - mapstructure doesn't validate business logic
5. **Use WeaklyTypedInput** sparingly - prefer strong typing
6. **Consider performance** for high-frequency operations
7. **Document your struct tags** for team clarity

## 📚 Struct Tags Reference

| Tag | Description | Example |
|-----|-------------|---------|
| `mapstructure:"name"` | Map field name | `Field string \`mapstructure:"field_name"\`` |
| `mapstructure:",remain"` | Capture unused fields | `Extra map[string]interface{} \`mapstructure:",remain"\`` |
| `mapstructure:",squash"` | Flatten embedded struct | `Embedded \`mapstructure:",squash"\`` |
| `mapstructure:",omitempty"` | Skip empty fields | `Optional \`mapstructure:",omitempty"\`` |
| `mapstructure:"-"` | Ignore field | `Ignored \`mapstructure:"-"\`` |

## 🔍 Integration Examples

### With Viper (Configuration Management)
```go
import "github.com/spf13/viper"

viper.SetConfigFile("config.yaml")
viper.ReadInConfig()

var config AppConfig
err := viper.Unmarshal(&config) // Uses mapstructure internally
```

### With Cobra (CLI Applications)
```go
import "github.com/spf13/cobra"

var configFile string
var config AppConfig

rootCmd := &cobra.Command{
    Run: func(cmd *cobra.Command, args []string) {
        // Parse config file
        configData := loadConfigFile(configFile)
        mapstructure.Decode(configData, &config)
    },
}
```

### With Gin (Web Framework)
```go
func CreateUser(c *gin.Context) {
    var requestData map[string]interface{}
    c.BindJSON(&requestData)
    
    var user User
    err := mapstructure.Decode(requestData, &user)
    if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    
    // Process user...
}
```

## ⚡ Performance Considerations

- **Caching decoders** for frequently used conversions
- **Pre-allocating slices** when possible
- **Avoiding reflection-heavy operations** in hot paths
- **Using type assertions** for simple conversions when appropriate

## 🌟 Why Use MapStructure?

- **Flexibility** - Handle dynamic data structures
- **Type Safety** - Convert untyped data to typed structs
- **Configurability** - Extensive customization options
- **Ecosystem Integration** - Works well with popular libraries
- **Performance** - Efficient reflection-based conversion
- **Validation Ready** - Easy to add validation after decoding

---

*This demo showcases the power and flexibility of mapstructure for handling dynamic data in Go applications.*

## 🔗 Useful Resources

- [MapStructure GitHub](https://github.com/mitchellh/mapstructure)
- [Go Reflection Documentation](https://golang.org/pkg/reflect/)
- [Viper Configuration Library](https://github.com/spf13/viper)
- [Cobra CLI Library](https://github.com/spf13/cobra)

*Happy mapping! 🗺️*

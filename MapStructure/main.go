package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"time"

	"github.com/mitchellh/mapstructure"
)

func main() {
	fmt.Println("🗺️ MapStructure Library Demo")
	fmt.Println("=============================")

	// Basic map to struct conversion
	fmt.Println("\n1. 📦 Basic Map to Struct")
	basicMapToStruct()

	// JSON to struct via map
	fmt.Println("\n2. 🔄 JSON to Struct via Map")
	jsonToStruct()

	// Nested structures
	fmt.Println("\n3. 🏗️ Nested Structures")
	nestedStructures()

	// Custom field mapping with tags
	fmt.Println("\n4. 🏷️ Custom Field Mapping")
	customFieldMapping()

	// Type conversion and hooks
	fmt.Println("\n5. 🔧 Type Conversion & Hooks")
	typeConversionHooks()

	// Slice and array handling
	fmt.Println("\n6. 📋 Slice and Array Handling")
	sliceArrayHandling()

	// Error handling and validation
	fmt.Println("\n7. ❌ Error Handling")
	errorHandling()

	// Advanced configuration
	fmt.Println("\n8. ⚙️ Advanced Configuration")
	advancedConfiguration()

	// Real-world examples
	fmt.Println("\n9. 🌍 Real-World Examples")
	realWorldExamples()

	// Prevent terminal window from closing on Windows
	if runtime.GOOS == "windows" {
		fmt.Println("\nPress Enter to exit...")
		bufio.NewScanner(os.Stdin).Scan()
	}
}

// 1. Basic Map to Struct
func basicMapToStruct() {
	// Define a simple struct
	type Person struct {
		Name string
		Age  int
		City string
	}

	// Source map
	input := map[string]interface{}{
		"name": "John Doe",
		"age":  30,
		"city": "New York",
	}

	var result Person
	err := mapstructure.Decode(input, &result)
	if err != nil {
		fmt.Printf("   ❌ Error: %v\n", err)
		return
	}

	fmt.Printf("   📊 Input map: %+v\n", input)
	fmt.Printf("   ✅ Result struct: %+v\n", result)
}

// 2. JSON to Struct via Map
func jsonToStruct() {
	type Product struct {
		ID    int     `mapstructure:"id"`
		Name  string  `mapstructure:"name"`
		Price float64 `mapstructure:"price"`
	}

	// JSON data
	jsonData := `{
		"id": 123,
		"name": "Laptop",
		"price": 999.99
	}`

	// Parse JSON to map
	var inputMap map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &inputMap)
	if err != nil {
		fmt.Printf("   ❌ JSON parse error: %v\n", err)
		return
	}

	// Convert map to struct
	var product Product
	err = mapstructure.Decode(inputMap, &product)
	if err != nil {
		fmt.Printf("   ❌ MapStructure error: %v\n", err)
		return
	}

	fmt.Printf("   📄 JSON: %s\n", jsonData)
	fmt.Printf("   🗺️ Map: %+v\n", inputMap)
	fmt.Printf("   📦 Struct: %+v\n", product)
}

// 3. Nested Structures
func nestedStructures() {
	type Address struct {
		Street  string `mapstructure:"street"`
		City    string `mapstructure:"city"`
		ZipCode string `mapstructure:"zip_code"`
	}

	type User struct {
		Name    string  `mapstructure:"name"`
		Email   string  `mapstructure:"email"`
		Address Address `mapstructure:"address"`
	}

	input := map[string]interface{}{
		"name":  "Alice Smith",
		"email": "alice@example.com",
		"address": map[string]interface{}{
			"street":   "123 Main St",
			"city":     "Boston",
			"zip_code": "02101",
		},
	}

	var user User
	err := mapstructure.Decode(input, &user)
	if err != nil {
		fmt.Printf("   ❌ Error: %v\n", err)
		return
	}

	fmt.Printf("   🗺️ Input: %+v\n", input)
	fmt.Printf("   👤 User: %+v\n", user)
	fmt.Printf("   🏠 Address: %+v\n", user.Address)
}

// 4. Custom Field Mapping
func customFieldMapping() {
	type Config struct {
		DatabaseURL string `mapstructure:"db_url"`
		Port        int    `mapstructure:"server_port"`
		Debug       bool   `mapstructure:"debug_mode"`
		Timeout     int    `mapstructure:"request_timeout"`
	}

	input := map[string]interface{}{
		"db_url":          "postgresql://localhost:5432/mydb",
		"server_port":     8080,
		"debug_mode":      true,
		"request_timeout": 30,
	}

	var config Config
	err := mapstructure.Decode(input, &config)
	if err != nil {
		fmt.Printf("   ❌ Error: %v\n", err)
		return
	}

	fmt.Printf("   🔧 Config: %+v\n", config)
	fmt.Printf("   🗄️ Database URL: %s\n", config.DatabaseURL)
	fmt.Printf("   🌐 Port: %d\n", config.Port)
	fmt.Printf("   🐛 Debug: %t\n", config.Debug)
}

// 5. Type Conversion & Hooks
func typeConversionHooks() {
	type Event struct {
		Name      string    `mapstructure:"name"`
		Timestamp time.Time `mapstructure:"timestamp"`
		Duration  time.Duration `mapstructure:"duration"`
	}

	// Custom decode hook for time parsing
	timeHook := func(from reflect.Type, to reflect.Type, data interface{}) (interface{}, error) {
		if to == reflect.TypeOf(time.Time{}) && from == reflect.TypeOf("") {
			return time.Parse("2006-01-02 15:04:05", data.(string))
		}
		if to == reflect.TypeOf(time.Duration(0)) && from == reflect.TypeOf("") {
			return time.ParseDuration(data.(string))
		}
		return data, nil
	}

	input := map[string]interface{}{
		"name":      "Meeting",
		"timestamp": "2024-12-25 14:30:00",
		"duration":  "2h30m",
	}

	var event Event
	config := &mapstructure.DecoderConfig{
		DecodeHook: timeHook,
		Result:     &event,
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		fmt.Printf("   ❌ Decoder creation error: %v\n", err)
		return
	}

	err = decoder.Decode(input)
	if err != nil {
		fmt.Printf("   ❌ Decode error: %v\n", err)
		return
	}

	fmt.Printf("   📅 Event: %+v\n", event)
	fmt.Printf("   ⏰ Timestamp: %s\n", event.Timestamp.Format("2006-01-02 15:04:05"))
	fmt.Printf("   ⏱️ Duration: %s\n", event.Duration)
}

// 6. Slice and Array Handling
func sliceArrayHandling() {
	type Team struct {
		Name    string   `mapstructure:"name"`
		Members []string `mapstructure:"members"`
		Scores  []int    `mapstructure:"scores"`
	}

	input := map[string]interface{}{
		"name":    "Development Team",
		"members": []interface{}{"Alice", "Bob", "Charlie"},
		"scores":  []interface{}{85, 92, 78},
	}

	var team Team
	err := mapstructure.Decode(input, &team)
	if err != nil {
		fmt.Printf("   ❌ Error: %v\n", err)
		return
	}

	fmt.Printf("   👥 Team: %s\n", team.Name)
	fmt.Printf("   🧑‍💻 Members: %v\n", team.Members)
	fmt.Printf("   📊 Scores: %v\n", team.Scores)

	// Nested struct slices
	type Task struct {
		ID          int    `mapstructure:"id"`
		Title       string `mapstructure:"title"`
		Completed   bool   `mapstructure:"completed"`
	}

	type Project struct {
		Name  string `mapstructure:"name"`
		Tasks []Task `mapstructure:"tasks"`
	}

	projectInput := map[string]interface{}{
		"name": "Website Redesign",
		"tasks": []interface{}{
			map[string]interface{}{
				"id":        1,
				"title":     "Design mockups",
				"completed": true,
			},
			map[string]interface{}{
				"id":        2,
				"title":     "Implement frontend",
				"completed": false,
			},
		},
	}

	var project Project
	err = mapstructure.Decode(projectInput, &project)
	if err != nil {
		fmt.Printf("   ❌ Project decode error: %v\n", err)
		return
	}

	fmt.Printf("   📋 Project: %s\n", project.Name)
	for i, task := range project.Tasks {
		status := "❌"
		if task.Completed {
			status = "✅"
		}
		fmt.Printf("   %s Task %d: %s (ID: %d)\n", status, i+1, task.Title, task.ID)
	}
}

// 7. Error Handling
func errorHandling() {
	type StrictConfig struct {
		RequiredField string `mapstructure:"required"`
		NumberField   int    `mapstructure:"number"`
	}

	// Test with missing required field
	fmt.Println("   🧪 Testing missing field:")
	input1 := map[string]interface{}{
		"number": 42,
		// "required" field is missing
	}

	var config1 StrictConfig
	err := mapstructure.Decode(input1, &config1)
	if err != nil {
		fmt.Printf("   ❌ Expected error: %v\n", err)
	} else {
		fmt.Printf("   ⚠️ No error, result: %+v\n", config1)
	}

	// Test with type mismatch
	fmt.Println("   🧪 Testing type mismatch:")
	input2 := map[string]interface{}{
		"required": "hello",
		"number":   "not_a_number", // Should be int
	}

	var config2 StrictConfig
	err = mapstructure.Decode(input2, &config2)
	if err != nil {
		fmt.Printf("   ❌ Type conversion error: %v\n", err)
	} else {
		fmt.Printf("   ✅ Successful conversion: %+v\n", config2)
	}

	// Using WeaklyTypedInput for flexible conversion
	fmt.Println("   🧪 Using WeaklyTypedInput:")
	var config3 StrictConfig
	decoder, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           &config3,
	})

	err = decoder.Decode(input2)
	if err != nil {
		fmt.Printf("   ❌ Still failed: %v\n", err)
	} else {
		fmt.Printf("   ✅ Weakly typed conversion: %+v\n", config3)
	}
}

// 8. Advanced Configuration
func advancedConfiguration() {
	type FlexibleStruct struct {
		KnownField   string                 `mapstructure:"known"`
		UnknownFields map[string]interface{} `mapstructure:",remain"`
	}

	input := map[string]interface{}{
		"known":    "This is known",
		"extra1":   "This is extra",
		"extra2":   42,
		"extra3":   true,
	}

	var result FlexibleStruct
	config := &mapstructure.DecoderConfig{
		Result:   &result,
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		fmt.Printf("   ❌ Decoder error: %v\n", err)
		return
	}

	err = decoder.Decode(input)
	if err != nil {
		fmt.Printf("   ❌ Decode error: %v\n", err)
		return
	}

	fmt.Printf("   📦 Known field: %s\n", result.KnownField)
	fmt.Printf("   🗃️ Unknown fields: %+v\n", result.UnknownFields)

	// Using squash for embedding
	type EmbeddedStruct struct {
		ID   int    `mapstructure:"id"`
		Name string `mapstructure:"name"`
	}

	type ContainerStruct struct {
		EmbeddedStruct `mapstructure:",squash"`
		Extra          string `mapstructure:"extra"`
	}

	squashInput := map[string]interface{}{
		"id":    123,
		"name":  "Embedded Name",
		"extra": "Extra Field",
	}

	var container ContainerStruct
	err = mapstructure.Decode(squashInput, &container)
	if err != nil {
		fmt.Printf("   ❌ Squash error: %v\n", err)
		return
	}

	fmt.Printf("   🔗 Container: %+v\n", container)
	fmt.Printf("   📌 ID: %d, Name: %s, Extra: %s\n", 
		container.ID, container.Name, container.Extra)
}

// 9. Real-World Examples
func realWorldExamples() {
	// Database configuration example
	fmt.Println("   🗄️ Database Configuration:")
	
	type DatabaseConfig struct {
		Host            string        `mapstructure:"host"`
		Port            int           `mapstructure:"port"`
		Username        string        `mapstructure:"username"`
		Password        string        `mapstructure:"password"`
		Database        string        `mapstructure:"database"`
		MaxConnections  int           `mapstructure:"max_connections"`
		ConnectTimeout  time.Duration `mapstructure:"connect_timeout"`
		SSL             bool          `mapstructure:"ssl"`
	}

	dbConfigMap := map[string]interface{}{
		"host":            "localhost",
		"port":            5432,
		"username":        "admin",
		"password":        "secret",
		"database":        "myapp",
		"max_connections": 25,
		"connect_timeout": "30s",
		"ssl":             true,
	}

	var dbConfig DatabaseConfig
	config := &mapstructure.DecoderConfig{
		DecodeHook: mapstructure.StringToTimeDurationHookFunc(),
		Result:     &dbConfig,
	}

	decoder, _ := mapstructure.NewDecoder(config)
	err := decoder.Decode(dbConfigMap)
	if err != nil {
		fmt.Printf("   ❌ DB config error: %v\n", err)
		return
	}

	fmt.Printf("   🔗 Connection: %s:%d\n", dbConfig.Host, dbConfig.Port)
	fmt.Printf("   👤 User: %s\n", dbConfig.Username)
	fmt.Printf("   🗄️ Database: %s\n", dbConfig.Database)
	fmt.Printf("   🔒 SSL: %t\n", dbConfig.SSL)

	// API Response parsing
	fmt.Println("\n   📡 API Response Parsing:")
	
	type APIResponse struct {
		Status   string      `mapstructure:"status"`
		Message  string      `mapstructure:"message"`
		Data     interface{} `mapstructure:"data"`
		Metadata struct {
			Page       int `mapstructure:"page"`
			PerPage    int `mapstructure:"per_page"`
			Total      int `mapstructure:"total"`
			TotalPages int `mapstructure:"total_pages"`
		} `mapstructure:"metadata"`
	}

	apiJSON := `{
		"status": "success",
		"message": "Data retrieved successfully",
		"data": [
			{"id": 1, "name": "Item 1"},
			{"id": 2, "name": "Item 2"}
		],
		"metadata": {
			"page": 1,
			"per_page": 10,
			"total": 25,
			"total_pages": 3
		}
	}`

	var apiMap map[string]interface{}
	json.Unmarshal([]byte(apiJSON), &apiMap)

	var apiResponse APIResponse
	err = mapstructure.Decode(apiMap, &apiResponse)
	if err != nil {
		fmt.Printf("   ❌ API parse error: %v\n", err)
		return
	}

	fmt.Printf("   📊 Status: %s\n", apiResponse.Status)
	fmt.Printf("   💬 Message: %s\n", apiResponse.Message)
	fmt.Printf("   📄 Page: %d/%d (Total: %d)\n", 
		apiResponse.Metadata.Page, 
		apiResponse.Metadata.TotalPages, 
		apiResponse.Metadata.Total)

	// Configuration file parsing
	fmt.Println("\n   ⚙️ Application Configuration:")
	
	type ServerConfig struct {
		Port         int      `mapstructure:"port"`
		Host         string   `mapstructure:"host"`
		AllowedHosts []string `mapstructure:"allowed_hosts"`
		TLS          struct {
			Enabled  bool   `mapstructure:"enabled"`
			CertFile string `mapstructure:"cert_file"`
			KeyFile  string `mapstructure:"key_file"`
		} `mapstructure:"tls"`
	}

	type AppConfig struct {
		Debug    bool           `mapstructure:"debug"`
		LogLevel string         `mapstructure:"log_level"`
		Server   ServerConfig   `mapstructure:"server"`
		Database DatabaseConfig `mapstructure:"database"`
	}

	configMap := map[string]interface{}{
		"debug":     true,
		"log_level": "info",
		"server": map[string]interface{}{
			"port":          8080,
			"host":          "0.0.0.0",
			"allowed_hosts": []interface{}{"localhost", "example.com"},
			"tls": map[string]interface{}{
				"enabled":   true,
				"cert_file": "/path/to/cert.pem",
				"key_file":  "/path/to/key.pem",
			},
		},
		"database": dbConfigMap,
	}

	var appConfig AppConfig
	config = &mapstructure.DecoderConfig{
		DecodeHook: mapstructure.StringToTimeDurationHookFunc(),
		Result:     &appConfig,
	}

	decoder, _ = mapstructure.NewDecoder(config)
	err = decoder.Decode(configMap)
	if err != nil {
		fmt.Printf("   ❌ App config error: %v\n", err)
		return
	}

	fmt.Printf("   🐛 Debug: %t\n", appConfig.Debug)
	fmt.Printf("   📝 Log Level: %s\n", appConfig.LogLevel)
	fmt.Printf("   🌐 Server: %s:%d\n", appConfig.Server.Host, appConfig.Server.Port)
	fmt.Printf("   🔒 TLS: %t\n", appConfig.Server.TLS.Enabled)
	fmt.Printf("   🏠 Allowed Hosts: %v\n", appConfig.Server.AllowedHosts)
}

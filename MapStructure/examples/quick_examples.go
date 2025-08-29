// Quick MapStructure Examples
// Run with: go run quick_examples.go

package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/mitchellh/mapstructure"
)

func main() {
	fmt.Println("🗺️ Quick MapStructure Examples")
	fmt.Println("===============================")

	// 1. Basic conversion
	basicExample()

	// 2. JSON to struct
	jsonExample()

	// 3. Custom field names
	customFieldsExample()

	// 4. Time conversion
	timeConversionExample()
}

func basicExample() {
	fmt.Println("\n1. Basic Map to Struct:")

	type Person struct {
		Name string
		Age  int
		City string
	}

	input := map[string]interface{}{
		"name": "Alice",
		"age":  25,
		"city": "Boston",
	}

	var person Person
	mapstructure.Decode(input, &person)

	fmt.Printf("   Input: %+v\n", input)
	fmt.Printf("   Output: %+v\n", person)
}

func jsonExample() {
	fmt.Println("\n2. JSON to Struct:")

	type Product struct {
		ID    int     `mapstructure:"id"`
		Name  string  `mapstructure:"name"`
		Price float64 `mapstructure:"price"`
	}

	jsonData := `{"id": 1, "name": "Laptop", "price": 1299.99}`

	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(jsonData), &jsonMap)

	var product Product
	mapstructure.Decode(jsonMap, &product)

	fmt.Printf("   JSON: %s\n", jsonData)
	fmt.Printf("   Product: %+v\n", product)
}

func customFieldsExample() {
	fmt.Println("\n3. Custom Field Names:")

	type Config struct {
		DatabaseURL string `mapstructure:"db_url"`
		ServerPort  int    `mapstructure:"server_port"`
		DebugMode   bool   `mapstructure:"debug_mode"`
	}

	input := map[string]interface{}{
		"db_url":      "postgres://localhost/mydb",
		"server_port": 8080,
		"debug_mode":  true,
	}

	var config Config
	mapstructure.Decode(input, &config)

	fmt.Printf("   Config: %+v\n", config)
}

func timeConversionExample() {
	fmt.Println("\n4. Time Conversion with Hooks:")

	type Event struct {
		Name     string        `mapstructure:"name"`
		Duration time.Duration `mapstructure:"duration"`
	}

	input := map[string]interface{}{
		"name":     "Meeting",
		"duration": "2h30m",
	}

	var event Event
	config := &mapstructure.DecoderConfig{
		DecodeHook: mapstructure.StringToTimeDurationHookFunc(),
		Result:     &event,
	}

	decoder, _ := mapstructure.NewDecoder(config)
	decoder.Decode(input)

	fmt.Printf("   Event: %+v\n", event)
	fmt.Printf("   Duration: %v\n", event.Duration)
}

// Run with: go run quick_examples.go
// This file demonstrates the most common time operations in Go

package main

import (
	"fmt"
	"time"
)

// QuickTimeExamples shows the most commonly used time operations
func main() {
	fmt.Println("⏰ Quick Time Examples")
	fmt.Println("======================")

	// 1. Current time
	now := time.Now()
	fmt.Printf("Current time: %v\n", now)

	// 2. Format time
	fmt.Printf("Formatted: %s\n", now.Format("2006-01-02 15:04:05"))
	fmt.Printf("RFC3339: %s\n", now.Format(time.RFC3339))
	fmt.Printf("Kitchen: %s\n", now.Format(time.Kitchen))

	// 3. Parse time from string
	timeStr := "2024-12-25 15:30:00"
	parsed, _ := time.Parse("2006-01-02 15:04:05", timeStr)
	fmt.Printf("Parsed: %v\n", parsed)

	// 4. Add/subtract time
	future := now.Add(2 * time.Hour)
	past := now.Add(-30 * time.Minute)
	fmt.Printf("Future: %s\n", future.Format("15:04:05"))
	fmt.Printf("Past: %s\n", past.Format("15:04:05"))

	// 5. Duration
	duration := time.Since(parsed)
	fmt.Printf("Duration since Christmas: %.0f days\n", duration.Hours()/24)

	// 6. Compare times
	fmt.Printf("Is now after parsed? %t\n", now.After(parsed))

	// 7. Unix timestamp
	fmt.Printf("Unix timestamp: %d\n", now.Unix())

	// 8. Sleep
	fmt.Println("Sleeping for 1 second...")
	time.Sleep(1 * time.Second)
	fmt.Println("Done!")

	// 9. Timer
	timer := time.NewTimer(500 * time.Millisecond)
	<-timer.C
	fmt.Println("Timer finished!")

	// 10. Time components
	fmt.Printf("Year: %d, Month: %s, Day: %d\n",
		now.Year(), now.Month(), now.Day())
	fmt.Printf("Hour: %d, Minute: %d, Second: %d\n",
		now.Hour(), now.Minute(), now.Second())
}

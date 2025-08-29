package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {
	fmt.Println("⏰ Go Time Package Demo")
	fmt.Println("=======================")

	// Basic time operations
	fmt.Println("\n1. 📅 Basic Time Operations")
	basicTimeOperations()

	// Time formatting and parsing
	fmt.Println("\n2. 📝 Time Formatting and Parsing")
	timeFormattingAndParsing()

	// Time zones
	fmt.Println("\n3. 🌍 Time Zones")
	timeZoneOperations()

	// Duration operations
	fmt.Println("\n4. ⏱️ Duration Operations")
	durationOperations()

	// Time comparisons
	fmt.Println("\n5. ⚖️ Time Comparisons")
	timeComparisons()

	// Practical examples
	fmt.Println("\n6. 🛠️ Practical Examples")
	practicalExamples()

	// Performance timing
	fmt.Println("\n7. 📊 Performance Timing")
	performanceTiming()

	// Timers and tickers
	fmt.Println("\n8. ⏲️ Timers and Tickers")
	timersAndTickers()

	// Prevent terminal window from closing on Windows
	if runtime.GOOS == "windows" {
		fmt.Println("\nPress Enter to exit...")
		bufio.NewScanner(os.Stdin).Scan()
	}
}

// 1. Basic Time Operations
func basicTimeOperations() {
	// Current time
	now := time.Now()
	fmt.Printf("   📅 Current time: %v\n", now)
	fmt.Printf("   📅 Current UTC time: %v\n", now.UTC())

	// Specific time creation
	birthday := time.Date(1990, time.May, 15, 14, 30, 0, 0, time.UTC)
	fmt.Printf("   🎂 Birthday: %v\n", birthday)

	// Unix timestamps
	unixTime := now.Unix()
	unixNano := now.UnixNano()
	fmt.Printf("   🕐 Unix timestamp: %d\n", unixTime)
	fmt.Printf("   🕐 Unix nano: %d\n", unixNano)

	// From Unix timestamp
	fromUnix := time.Unix(unixTime, 0)
	fmt.Printf("   🔄 From Unix: %v\n", fromUnix)

	// Time components
	fmt.Printf("   📊 Year: %d, Month: %s, Day: %d\n", now.Year(), now.Month(), now.Day())
	fmt.Printf("   🕒 Hour: %d, Minute: %d, Second: %d\n", now.Hour(), now.Minute(), now.Second())
	fmt.Printf("   📆 Weekday: %s, Yearday: %d\n", now.Weekday(), now.YearDay())
}

// 2. Time Formatting and Parsing
func timeFormattingAndParsing() {
	now := time.Now()

	// Go's reference time: Mon Jan 2 15:04:05 MST 2006 (Unix time: 1136239445)
	fmt.Println("   📝 Various formats:")
	fmt.Printf("   • RFC3339: %s\n", now.Format(time.RFC3339))
	fmt.Printf("   • Kitchen: %s\n", now.Format(time.Kitchen))
	fmt.Printf("   • Stamp: %s\n", now.Format(time.Stamp))
	fmt.Printf("   • Custom: %s\n", now.Format("2006-01-02 15:04:05"))
	fmt.Printf("   • Custom: %s\n", now.Format("Jan 2, 2006 at 3:04 PM"))
	fmt.Printf("   • Custom: %s\n", now.Format("Monday, January 2, 2006"))
	fmt.Printf("   • ISO 8601: %s\n", now.Format("2006-01-02T15:04:05Z07:00"))

	// Parsing strings to time
	fmt.Println("\n   🔍 Parsing time strings:")
	timeString := "2023-12-25 15:30:45"
	parsed, err := time.Parse("2006-01-02 15:04:05", timeString)
	if err != nil {
		fmt.Printf("   ❌ Parse error: %v\n", err)
	} else {
		fmt.Printf("   ✅ Parsed: %v\n", parsed)
	}

	// Parse with timezone
	timeWithTZ := "2023-12-25T15:30:45Z"
	parsedTZ, err := time.Parse(time.RFC3339, timeWithTZ)
	if err != nil {
		fmt.Printf("   ❌ Parse TZ error: %v\n", err)
	} else {
		fmt.Printf("   ✅ Parsed with TZ: %v\n", parsedTZ)
	}
}

// 3. Time Zone Operations
func timeZoneOperations() {
	now := time.Now()

	// Load specific time zones
	locations := []string{"UTC", "America/New_York", "Europe/London", "Asia/Tokyo", "Australia/Sydney"}

	for _, locName := range locations {
		loc, err := time.LoadLocation(locName)
		if err != nil {
			fmt.Printf("   ❌ Error loading %s: %v\n", locName, err)
			continue
		}

		localTime := now.In(loc)
		fmt.Printf("   🌍 %-20s: %s\n", locName, localTime.Format("2006-01-02 15:04:05 MST"))
	}

	// Get timezone offset
	_, offset := now.Zone()
	offsetHours := offset / 3600
	fmt.Printf("   🕐 Current timezone offset: %+d hours\n", offsetHours)

	// Fixed timezone (custom offset)
	fixedZone := time.FixedZone("CUSTOM", 5*3600) // +5 hours
	customTime := now.In(fixedZone)
	fmt.Printf("   🎯 Custom timezone (+5): %s\n", customTime.Format("2006-01-02 15:04:05 MST"))
}

// 4. Duration Operations
func durationOperations() {
	// Creating durations
	fmt.Println("   ⏱️ Duration examples:")
	fmt.Printf("   • 1 second: %v\n", time.Second)
	fmt.Printf("   • 5 minutes: %v\n", 5*time.Minute)
	fmt.Printf("   • 2 hours: %v\n", 2*time.Hour)
	fmt.Printf("   • 30 days: %v\n", 24*30*time.Hour)

	// Parse duration from string
	duration, err := time.ParseDuration("2h30m15s")
	if err != nil {
		fmt.Printf("   ❌ Duration parse error: %v\n", err)
	} else {
		fmt.Printf("   ✅ Parsed duration: %v\n", duration)
		fmt.Printf("   📊 In seconds: %.0f\n", duration.Seconds())
		fmt.Printf("   📊 In minutes: %.2f\n", duration.Minutes())
		fmt.Printf("   📊 In hours: %.2f\n", duration.Hours())
	}

	// Duration arithmetic
	now := time.Now()
	future := now.Add(duration)
	past := now.Add(-duration)

	fmt.Printf("   📅 Now: %s\n", now.Format("2006-01-02 15:04:05"))
	fmt.Printf("   📅 Future (+%v): %s\n", duration, future.Format("2006-01-02 15:04:05"))
	fmt.Printf("   📅 Past (-%v): %s\n", duration, past.Format("2006-01-02 15:04:05"))

	// Time since/until
	birthDate := time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC)
	age := time.Since(birthDate)
	fmt.Printf("   🎂 Age since 1990-01-01: %.0f days\n", age.Hours()/24)

	newYear := time.Date(2026, time.January, 1, 0, 0, 0, 0, time.UTC)
	untilNewYear := time.Until(newYear)
	if untilNewYear > 0 {
		fmt.Printf("   🎊 Days until 2026: %.0f days\n", untilNewYear.Hours()/24)
	}
}

// 5. Time Comparisons
func timeComparisons() {
	now := time.Now()
	past := now.Add(-1 * time.Hour)
	future := now.Add(1 * time.Hour)

	fmt.Printf("   📅 Now: %s\n", now.Format("15:04:05"))
	fmt.Printf("   📅 Past: %s\n", past.Format("15:04:05"))
	fmt.Printf("   📅 Future: %s\n", future.Format("15:04:05"))

	// Comparisons
	fmt.Printf("   ⚖️ Now.After(Past): %t\n", now.After(past))
	fmt.Printf("   ⚖️ Now.Before(Future): %t\n", now.Before(future))
	fmt.Printf("   ⚖️ Now.Equal(Now): %t\n", now.Equal(now))

	// Is zero
	var zeroTime time.Time
	fmt.Printf("   ⚖️ Zero time.IsZero(): %t\n", zeroTime.IsZero())
	fmt.Printf("   ⚖️ Now.IsZero(): %t\n", now.IsZero())

	// Time difference
	diff := now.Sub(past)
	fmt.Printf("   📏 Difference (now - past): %v\n", diff)
}

// 6. Practical Examples
func practicalExamples() {
	// Age calculation
	birthDate := time.Date(1990, time.May, 15, 0, 0, 0, 0, time.UTC)
	age := calculateAge(birthDate)
	fmt.Printf("   🎂 Age from 1990-05-15: %d years\n", age)

	// Business days calculation
	start := time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, time.January, 15, 0, 0, 0, 0, time.UTC)
	businessDays := calculateBusinessDays(start, end)
	fmt.Printf("   💼 Business days (Jan 1-15, 2024): %d\n", businessDays)

	// Time until next Friday
	nextFriday := getNextWeekday(time.Now(), time.Friday)
	untilFriday := time.Until(nextFriday)
	fmt.Printf("   📅 Next Friday: %s (in %.1f days)\n",
		nextFriday.Format("2006-01-02"), untilFriday.Hours()/24)

	// Start/end of day
	now := time.Now()
	startOfDay := getStartOfDay(now)
	endOfDay := getEndOfDay(now)
	fmt.Printf("   🌅 Start of day: %s\n", startOfDay.Format("2006-01-02 15:04:05"))
	fmt.Printf("   🌇 End of day: %s\n", endOfDay.Format("2006-01-02 15:04:05"))

	// Time ranges
	if isWeekend(now) {
		fmt.Printf("   🎉 It's weekend!\n")
	} else {
		fmt.Printf("   💼 It's a weekday\n")
	}

	if isBusinessHours(now) {
		fmt.Printf("   🏢 Currently business hours\n")
	} else {
		fmt.Printf("   🏠 Outside business hours\n")
	}
}

// 7. Performance Timing
func performanceTiming() {
	// Simple timing
	start := time.Now()

	// Simulate some work
	total := 0
	for i := 0; i < 1000000; i++ {
		total += i
	}

	elapsed := time.Since(start)
	fmt.Printf("   ⚡ Simple loop took: %v\n", elapsed)

	// More precise timing with multiple measurements
	fmt.Println("   📊 Performance measurements:")

	measurements := make([]time.Duration, 5)
	for i := 0; i < 5; i++ {
		start := time.Now()

		// Simulate work
		time.Sleep(10 * time.Millisecond)

		measurements[i] = time.Since(start)
	}

	var total_time time.Duration
	for i, duration := range measurements {
		fmt.Printf("   • Run %d: %v\n", i+1, duration)
		total_time += duration
	}

	average := total_time / time.Duration(len(measurements))
	fmt.Printf("   📊 Average: %v\n", average)
}

// 8. Timers and Tickers
func timersAndTickers() {
	fmt.Println("   ⏲️ Timer example (3 seconds):")

	// Timer example
	timer := time.NewTimer(3 * time.Second)

	go func() {
		<-timer.C
		fmt.Println("   ⏰ Timer fired!")
	}()

	// Ticker example (limited to avoid long running)
	fmt.Println("   🎯 Ticker example (5 ticks every 500ms):")

	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		count := 0
		for t := range ticker.C {
			count++
			fmt.Printf("   🔄 Tick %d at %s\n", count, t.Format("15:04:05.000"))
			if count >= 5 {
				ticker.Stop()
				return
			}
		}
	}()

	// Wait for timer and ticker to complete
	time.Sleep(4 * time.Second)

	// After function
	fmt.Println("   ⏰ After function example:")
	done := make(chan bool)

	time.AfterFunc(1*time.Second, func() {
		fmt.Println("   ✅ AfterFunc executed after 1 second")
		done <- true
	})

	<-done // Wait for completion
}

// Helper functions

func calculateAge(birthDate time.Time) int {
	now := time.Now()
	age := now.Year() - birthDate.Year()

	// Check if birthday hasn't occurred this year
	if now.YearDay() < birthDate.YearDay() {
		age--
	}

	return age
}

func calculateBusinessDays(start, end time.Time) int {
	count := 0
	for d := start; d.Before(end) || d.Equal(end); d = d.AddDate(0, 0, 1) {
		if d.Weekday() != time.Saturday && d.Weekday() != time.Sunday {
			count++
		}
	}
	return count
}

func getNextWeekday(from time.Time, weekday time.Weekday) time.Time {
	daysUntil := (int(weekday) - int(from.Weekday()) + 7) % 7
	if daysUntil == 0 {
		daysUntil = 7 // Next week if today is the target day
	}
	return from.AddDate(0, 0, daysUntil)
}

func getStartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func getEndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
}

func isWeekend(t time.Time) bool {
	return t.Weekday() == time.Saturday || t.Weekday() == time.Sunday
}

func isBusinessHours(t time.Time) bool {
	hour := t.Hour()
	return hour >= 9 && hour < 17 && !isWeekend(t)
}

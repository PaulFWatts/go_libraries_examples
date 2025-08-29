# Go Time Package Demo

A comprehensive demonstration of Go's built-in `time` package, showcasing all major features and practical usage patterns.

## 🚀 Features

This demo covers all aspects of Go's time package:
- **Basic time operations** (creation, components, timestamps)
- **Time formatting and parsing** (custom formats, RFC standards)
- **Time zone handling** (different locations, offsets)
- **Duration operations** (arithmetic, parsing, conversions)
- **Time comparisons** (before, after, equal)
- **Practical examples** (age calculation, business days)
- **Performance timing** (benchmarking, measurements)
- **Timers and tickers** (scheduled operations)

## 🔧 Setup

1. **Run the complete demo:**
   ```bash
   go run main.go
   ```

2. **Run quick examples:**
   ```bash
   cd examples
   go run quick_examples.go
   ```

3. **Build executable:**
   ```bash
   go build -o time-demo main.go
   ```

## 📋 What It Demonstrates

### 1. Basic Time Operations
```go
// Current time
now := time.Now()

// Specific time creation
birthday := time.Date(1990, time.May, 15, 14, 30, 0, 0, time.UTC)

// Unix timestamps
unixTime := now.Unix()
fromUnix := time.Unix(unixTime, 0)

// Time components
year := now.Year()
month := now.Month()
day := now.Day()
```

### 2. Time Formatting and Parsing
Go uses a specific reference time for formatting: `Mon Jan 2 15:04:05 MST 2006`

```go
// Formatting
now.Format("2006-01-02 15:04:05")          // 2024-08-29 14:30:45
now.Format("Jan 2, 2006 at 3:04 PM")       // Aug 29, 2024 at 2:30 PM
now.Format(time.RFC3339)                    // 2024-08-29T14:30:45Z07:00

// Parsing
parsed, err := time.Parse("2006-01-02", "2024-12-25")
```

### 3. Time Zone Operations
```go
// Load timezone
loc, err := time.LoadLocation("America/New_York")
localTime := now.In(loc)

// Fixed timezone
fixedZone := time.FixedZone("CUSTOM", 5*3600) // +5 hours
```

### 4. Duration Operations
```go
// Creating durations
duration := 2*time.Hour + 30*time.Minute + 15*time.Second

// Parsing duration strings
d, err := time.ParseDuration("2h30m15s")

// Duration arithmetic
future := now.Add(duration)
past := now.Add(-duration)

// Time since/until
age := time.Since(birthDate)
countdown := time.Until(deadline)
```

### 5. Time Comparisons
```go
// Comparison methods
now.After(past)     // true
now.Before(future)  // true
now.Equal(now)      // true

// Time difference
diff := now.Sub(past)
```

### 6. Practical Examples
```go
// Age calculation
func calculateAge(birthDate time.Time) int {
    now := time.Now()
    age := now.Year() - birthDate.Year()
    if now.YearDay() < birthDate.YearDay() {
        age--
    }
    return age
}

// Business days
func calculateBusinessDays(start, end time.Time) int {
    count := 0
    for d := start; d.Before(end) || d.Equal(end); d = d.AddDate(0, 0, 1) {
        if d.Weekday() != time.Saturday && d.Weekday() != time.Sunday {
            count++
        }
    }
    return count
}
```

### 7. Performance Timing
```go
// Simple timing
start := time.Now()
// ... do work ...
elapsed := time.Since(start)

// Multiple measurements
measurements := make([]time.Duration, 5)
for i := 0; i < 5; i++ {
    start := time.Now()
    // ... do work ...
    measurements[i] = time.Since(start)
}
```

### 8. Timers and Tickers
```go
// One-time timer
timer := time.NewTimer(5 * time.Second)
<-timer.C // Block until timer fires

// Repeating ticker
ticker := time.NewTicker(1 * time.Second)
for t := range ticker.C {
    fmt.Println("Tick at", t)
}

// Delayed function execution
time.AfterFunc(2*time.Second, func() {
    fmt.Println("Executed after 2 seconds")
})
```

## 🎯 Sample Output

```
⏰ Go Time Package Demo
=======================

1. 📅 Basic Time Operations
   📅 Current time: 2024-08-29 14:30:45.123456 +0000 UTC
   📅 Current UTC time: 2024-08-29 14:30:45.123456 +0000 UTC
   🎂 Birthday: 1990-05-15 14:30:00 +0000 UTC
   🕐 Unix timestamp: 1724941845
   📊 Year: 2024, Month: August, Day: 29
   🕒 Hour: 14, Minute: 30, Second: 45

2. 📝 Time Formatting and Parsing
   📝 Various formats:
   • RFC3339: 2024-08-29T14:30:45Z
   • Kitchen: 2:30PM
   • Custom: 2024-08-29 14:30:45
   • Custom: Aug 29, 2024 at 2:30 PM
   ✅ Parsed: 2023-12-25 15:30:45 +0000 UTC

3. 🌍 Time Zones
   🌍 UTC                 : 2024-08-29 14:30:45 UTC
   🌍 America/New_York    : 2024-08-29 10:30:45 EDT
   🌍 Europe/London       : 2024-08-29 15:30:45 BST
   🌍 Asia/Tokyo          : 2024-08-29 23:30:45 JST

4. ⏱️ Duration Operations
   ⏱️ Duration examples:
   • 1 second: 1s
   • 5 minutes: 5m0s
   • 2 hours: 2h0m0s
   ✅ Parsed duration: 2h30m15s
   📊 In seconds: 9015
   📊 In minutes: 150.25
```

## 🛠️ Common Use Cases

### Date/Time Manipulation
- **Birthday/age calculations**
- **Deadline tracking**
- **Schedule management**
- **Log timestamps**

### Business Applications
- **Working hours validation**
- **Business day calculations**
- **Meeting scheduling**
- **Time-based billing**

### Performance & Monitoring
- **Function execution timing**
- **API response time tracking**
- **Background task scheduling**
- **Rate limiting**

### Data Processing
- **Log file parsing**
- **Time series data**
- **Event timestamping**
- **Data archiving**

## 📚 Key Constants and Types

### Time Constants
```go
time.Nanosecond   = 1
time.Microsecond  = 1000 * Nanosecond
time.Millisecond  = 1000 * Microsecond
time.Second       = 1000 * Millisecond
time.Minute       = 60 * Second
time.Hour         = 60 * Minute
```

### Layout Constants
```go
time.Layout      = "01/02 03:04:05PM '06 -0700"
time.ANSIC       = "Mon Jan _2 15:04:05 2006"
time.UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
time.RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
time.RFC822      = "02 Jan 06 15:04 MST"
time.RFC822Z     = "02 Jan 06 15:04 -0700"
time.RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
time.RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
time.RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700"
time.RFC3339     = "2006-01-02T15:04:05Z07:00"
time.RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
time.Kitchen     = "3:04PM"
time.Stamp       = "Jan _2 15:04:05"
time.StampMilli  = "Jan _2 15:04:05.000"
time.StampMicro  = "Jan _2 15:04:05.000000"
time.StampNano   = "Jan _2 15:04:05.000000000"
```

### Weekdays
```go
time.Sunday = 0
time.Monday = 1
time.Tuesday = 2
time.Wednesday = 3
time.Thursday = 4
time.Friday = 5
time.Saturday = 6
```

## 🎯 Best Practices

1. **Always handle timezone properly** in global applications
2. **Use UTC for storage** and convert for display
3. **Parse durations from config** using `time.ParseDuration()`
4. **Use appropriate precision** (don't use nanoseconds for long durations)
5. **Stop tickers when done** to prevent goroutine leaks
6. **Use context for timeouts** in concurrent operations

## 🔍 Advanced Topics

### Custom Time Types
```go
type CustomTime struct {
    time.Time
}

func (ct CustomTime) MarshalJSON() ([]byte, error) {
    return []byte(`"` + ct.Format("2006-01-02") + `"`), nil
}
```

### Time-based Caching
```go
type TimedCache struct {
    data      interface{}
    timestamp time.Time
    ttl       time.Duration
}

func (tc *TimedCache) IsExpired() bool {
    return time.Since(tc.timestamp) > tc.ttl
}
```

### Rate Limiting
```go
type RateLimiter struct {
    ticker   *time.Ticker
    requests chan struct{}
}

func NewRateLimiter(rate time.Duration, burst int) *RateLimiter {
    rl := &RateLimiter{
        ticker:   time.NewTicker(rate),
        requests: make(chan struct{}, burst),
    }
    
    go func() {
        for range rl.ticker.C {
            select {
            case rl.requests <- struct{}{}:
            default:
            }
        }
    }()
    
    return rl
}
```

---

*This demo provides a complete overview of Go's time package capabilities for both beginners and advanced users.*

## 🌟 Why the Time Package is Important

- **Universal need** - every application deals with time
- **Built-in reliability** - no external dependencies
- **Timezone awareness** - crucial for global applications
- **Performance timing** - essential for optimization
- **Rich functionality** - covers all time-related needs

*Happy timing! ⏰*

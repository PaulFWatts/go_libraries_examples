# GoQuery Web Scraping Demo

This project demonstrates web scraping in Go using the powerful `goquery` library, which provides jQuery-like syntax for HTML document traversal and manipulation.

## 🚀 Features

- **Simple web scraping** from example.com
- **Custom HTTP headers** and user agents
- **Robust error handling** with timeouts
- **Multiple scraping examples**
- **Helper functions** for common scraping tasks
- **Cross-platform compatibility**

## 📦 Dependencies

- `github.com/PuerkitoBio/goquery` - jQuery-like HTML parsing

## 🔧 Setup

1. **Install dependencies:**
   ```bash
   go mod tidy
   ```

2. **Run the demo:**
   ```bash
   go run main.go
   ```

3. **Build executable:**
   ```bash
   go build -o goquery-demo main.go
   ```

## 📋 What It Demonstrates

### 1. Basic Web Scraping
- Fetching HTML content from websites
- Parsing HTML with goquery
- Extracting titles, headings, and paragraphs
- Finding and processing links

### 2. HTTP Best Practices
- Setting request timeouts
- Custom user agents
- Proper error handling
- Response status checking

### 3. GoQuery Selectors
```go
// CSS selectors
doc.Find("h1, h2, h3")        // Multiple elements
doc.Find("a[href]")           // Attributes
doc.Find("p")                 // All paragraphs

// Iteration
doc.Find("li").Each(func(i int, s *goquery.Selection) {
    fmt.Println(s.Text())
})

// Attribute extraction
href, exists := s.Attr("href")
```

### 4. Helper Functions
- `ExtractMetaTags()` - Extract all meta tags
- `ExtractImages()` - Get all images with alt text
- `truncateText()` - Limit text length for display

## 📝 Code Examples

### Basic Scraping
```go
// Create HTTP client with timeout
client := &http.Client{
    Timeout: 10 * time.Second,
}

// Make request
req, err := http.NewRequest("GET", url, nil)
resp, err := client.Do(req)
defer resp.Body.Close()

// Parse HTML
doc, err := goquery.NewDocumentFromReader(resp.Body)

// Extract data
title := doc.Find("title").Text()
```

### Custom Headers
```go
req.Header.Set("User-Agent", "GoQuery-Demo/1.0")
req.Header.Set("Accept", "text/html,application/xhtml+xml")
```

### Advanced Selection
```go
// Find all links with text
doc.Find("a").Each(func(i int, s *goquery.Selection) {
    href, _ := s.Attr("href")
    text := s.Text()
    fmt.Printf("Link: %s -> %s\n", text, href)
})

// Extract meta tags
meta := make(map[string]string)
doc.Find("meta[name]").Each(func(i int, s *goquery.Selection) {
    name, _ := s.Attr("name")
    content, _ := s.Attr("content")
    meta[name] = content
})
```

## ⚖️ Legal and Ethical Considerations

### Always Remember:
1. **Check robots.txt** - Respect the site's scraping policy
2. **Read Terms of Service** - Some sites prohibit scraping
3. **Rate limiting** - Don't overwhelm servers
4. **User-Agent** - Identify your scraper properly
5. **Respect copyright** - Don't steal content

### Good Practices:
```go
// Add delays between requests
time.Sleep(1 * time.Second)

// Use proper user agent
req.Header.Set("User-Agent", "YourBot/1.0 (contact@example.com)")

// Respect rate limits
rateLimiter := time.Tick(1 * time.Second)
<-rateLimiter // Wait before each request
```

## 🛠️ Common Use Cases

- **Data extraction** from websites
- **Price monitoring** for e-commerce
- **News aggregation**
- **SEO analysis** (meta tags, headings)
- **Social media monitoring**
- **Research and data collection**

## 🎯 Sample Output

```
🔍 GoQuery Web Scraping Demo
=============================

1. 📄 Scraping example.com...
   📋 Title: Example Domain
   📌 h1: Example Domain
   📝 Paragraph 1: This domain is for use in illustrative...
   🔗 Link: More information... -> https://www.iana.org/domains/example

2. 📰 Scraping news headlines...
   📋 Page title: Test Page
   📰 Headline: Herman Melville - Moby-Dick

3. 🕵️ Scraping with custom headers...
   🌐 Response Status: 200 OK
   📊 Content-Type: application/json
   📤 Headers sent (truncated):
      {
        "headers": {
          "User-Agent": "GoQuery-Demo/1.0 (Educational Purpose)"
        }
      }
```

## 🔍 Advanced Features

### Error Handling
- Network timeout handling
- HTTP status code checking
- HTML parsing error recovery
- Graceful degradation

### Performance
- Connection pooling
- Request timeouts
- Memory-efficient parsing
- Selective content extraction

### Security
- Input validation
- Safe header handling
- Proper URL encoding
- HTTPS verification

---

**Note:** This demo is for educational purposes. Always respect website terms of service and robots.txt when scraping real websites.

## 🌐 Useful Resources

- [GoQuery Documentation](https://github.com/PuerkitoBio/goquery)
- [CSS Selectors Reference](https://www.w3schools.com/cssref/css_selectors.asp)
- [HTTP Status Codes](https://httpstatuses.com/)
- [robots.txt Specification](https://www.robotstxt.org/)

*Happy scraping! 🕷️*

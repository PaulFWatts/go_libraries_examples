package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	fmt.Println("🔍 GoQuery Web Scraping Demo")
	fmt.Println("=============================")

	// Example 1: Simple web scraping
	fmt.Println("\n1. 📄 Scraping example.com...")
	scrapeExample()

	// Example 2: Scraping news headlines (example with BBC)
	fmt.Println("\n2. 📰 Scraping news headlines...")
	scrapeNews()

	// Example 3: Scraping with custom user agent
	fmt.Println("\n3. 🕵️ Scraping with custom headers...")
	scrapeWithHeaders()

	// Prevent terminal window from closing on Windows
	if runtime.GOOS == "windows" {
		fmt.Println("\nPress Enter to exit...")
		bufio.NewScanner(os.Stdin).Scan()
	}
}

func scrapeExample() {
	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Make request
	req, err := http.NewRequest("GET", "https://example.com", nil)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error making request: %v", err)
		return
	}
	defer resp.Body.Close()

	// Check status code
	if resp.StatusCode != 200 {
		log.Printf("Status code error: %d %s", resp.StatusCode, resp.Status)
		return
	}

	// Parse HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("Error parsing HTML: %v", err)
		return
	}

	// Extract various elements
	fmt.Printf("   📋 Title: %s\n", doc.Find("title").Text())

	// Find all headings
	doc.Find("h1, h2, h3").Each(func(i int, s *goquery.Selection) {
		heading := strings.TrimSpace(s.Text())
		if heading != "" {
			fmt.Printf("   📌 %s: %s\n", goquery.NodeName(s), heading)
		}
	})

	// Find all paragraphs
	paragraphCount := 0
	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		if text != "" && len(text) > 10 {
			paragraphCount++
			if paragraphCount <= 3 { // Show first 3 paragraphs
				fmt.Printf("   📝 Paragraph %d: %s\n", paragraphCount, truncateText(text, 100))
			}
		}
	})

	// Find all links
	linkCount := 0
	doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists && linkCount < 5 { // Show first 5 links
			linkText := strings.TrimSpace(s.Text())
			if linkText == "" {
				linkText = "No text"
			}
			fmt.Printf("   🔗 Link: %s -> %s\n", truncateText(linkText, 30), href)
			linkCount++
		}
	})
}

func scrapeNews() {
	// Note: This is an example - always check robots.txt and terms of service
	// before scraping real websites

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Using HTTPBin for demonstration (it returns JSON, but we'll scrape it as HTML)
	req, err := http.NewRequest("GET", "https://httpbin.org/html", nil)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error making request: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Printf("Status code error: %d %s", resp.StatusCode, resp.Status)
		return
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("Error parsing HTML: %v", err)
		return
	}

	fmt.Printf("   📋 Page title: %s\n", doc.Find("title").Text())

	// Find all headings in the test page
	doc.Find("h1").Each(func(i int, s *goquery.Selection) {
		fmt.Printf("   📰 Headline: %s\n", s.Text())
	})

	// Find all list items
	doc.Find("li").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		if text != "" {
			fmt.Printf("   • %s\n", text)
		}
	})
}

func scrapeWithHeaders() {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Create request with custom headers
	req, err := http.NewRequest("GET", "https://httpbin.org/headers", nil)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return
	}

	// Add custom headers
	req.Header.Set("User-Agent", "GoQuery-Demo/1.0 (Educational Purpose)")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error making request: %v", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("   🌐 Response Status: %s\n", resp.Status)
	fmt.Printf("   📊 Content-Type: %s\n", resp.Header.Get("Content-Type"))
	fmt.Printf("   📏 Content-Length: %s\n", resp.Header.Get("Content-Length"))

	// This endpoint returns JSON showing the headers we sent
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("Error parsing response: %v", err)
		return
	}

	// Find pre tags (JSON content)
	doc.Find("pre").Each(func(i int, s *goquery.Selection) {
		content := s.Text()
		if strings.Contains(content, "User-Agent") {
			fmt.Printf("   📤 Headers sent (truncated):\n")
			lines := strings.Split(content, "\n")
			for _, line := range lines[:5] { // Show first 5 lines
				if strings.TrimSpace(line) != "" {
					fmt.Printf("      %s\n", line)
				}
			}
		}
	})
}

// Helper function to truncate text
func truncateText(text string, maxLen int) string {
	if len(text) <= maxLen {
		return text
	}
	return text[:maxLen-3] + "..."
}

// Additional helper functions for common scraping tasks

// ExtractMetaTags extracts meta tags from a document
func ExtractMetaTags(doc *goquery.Document) map[string]string {
	meta := make(map[string]string)

	doc.Find("meta").Each(func(i int, s *goquery.Selection) {
		if name, exists := s.Attr("name"); exists {
			if content, exists := s.Attr("content"); exists {
				meta[name] = content
			}
		}
		if property, exists := s.Attr("property"); exists {
			if content, exists := s.Attr("content"); exists {
				meta[property] = content
			}
		}
	})

	return meta
}

// ExtractImages extracts all images with their alt text and src
func ExtractImages(doc *goquery.Document) []map[string]string {
	var images []map[string]string

	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		img := make(map[string]string)
		if src, exists := s.Attr("src"); exists {
			img["src"] = src
		}
		if alt, exists := s.Attr("alt"); exists {
			img["alt"] = alt
		} else {
			img["alt"] = "No alt text"
		}
		images = append(images, img)
	})

	return images
}

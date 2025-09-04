// Package veritas provides URL validation functions.
package veritas

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// ValidateURL validates a URL format.
func ValidateURL(urlStr interface{}) error {
	urlStr, ok := urlStr.(string)
	if !ok {
		return fmt.Errorf("URL must be a string")
	}

	urlStr = cleanString(urlStr.(string), false)
	if isEmpty(urlStr.(string)) {
		return fmt.Errorf("URL cannot be empty")
	}

	// Parse the URL
	parsedURL, err := url.Parse(urlStr.(string))
	if err != nil {
		return fmt.Errorf("invalid URL format: %w", err)
	}

	// Check if scheme is present
	if parsedURL.Scheme == "" {
		return fmt.Errorf("URL must include a scheme (http:// or https://)")
	}

	// Check if host is present
	if parsedURL.Host == "" {
		return fmt.Errorf("URL must include a host")
	}

	// Check if URL returns 200 status code
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Head(urlStr.(string))
	if err != nil {
		return fmt.Errorf("URL is not accessible: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("URL returned status %d, expected 200", resp.StatusCode)
	}

	return nil
}

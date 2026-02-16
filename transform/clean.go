package transform

import (
	"regexp"
	"strings"
	_ "unicode"
)

// Trim removes leading and trailing whitespace
func Trim(s string) string {
	return strings.TrimSpace(s)
}

// CollapseSpaces replaces multiple consecutive spaces with a single space
func CollapseSpaces(s string) string {
	re := regexp.MustCompile(`\s+`)
	result := re.ReplaceAll([]byte(s), []byte(" "))
	return string(result)
}

// RemoveSpecialChars removes all characters except alphanumeric and spaces
func RemoveSpecialChars(s string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9\s]+`)
	result := re.ReplaceAll([]byte(s), []byte(""))
	return string(result)
}

// SplitIntoWords splits a string into words by detecting multiple boundary types:
// - Spaces, underscores, hyphens
// - camelCase transitions (lowercase to uppercase)
// - PascalCase transitions
// Returns a slice of words in lowercase
func SplitIntoWords(s string) []string {
	// NOTE: handle the detection logic for:
	// - Splitting on `_`, `-`, and spaces
	// - Detecting camelCase boundaries (e.g., `helloWorld` -> `["hello", "world"]`)
	// - Filtering out empty strings from the result
}

package transform

import (
	"regexp"
	"strings"
	"unicode"
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

// SplitIntoWords splits a string into words by detecting multiple boundary types
// (1. spaces, underscores, hyphens, 2. camelCase transitions, 3. PascalCase transitions)
// and returns a slice of words in lowercase
func SplitIntoWords(s string) []string {
	runes := []rune(s)

	var words []string
	var currentWord []rune

	for i := 0; i < len(runes); i++ {
		currentRune := runes[i]
		var prevRune, nextRune rune

		if i > 0 {
			prevRune = runes[i-1]
		}

		if i < len(runes)-1 {
			nextRune = runes[i+1]
		}

		// If delimiter, save current word and skip the delimiter
		if isDelimiter(currentRune) {
			currentWord, words = saveWord(currentWord, words)
			continue // Skip the delimiter itself
		}

		// If camelCase boundary, save current word and start new one
		if i > 0 && unicode.IsLower(prevRune) && unicode.IsUpper(currentRune) {
			currentWord, words = saveWord(currentWord, words)
		}

		// If PascalCase, save word and start new
		if i > 0 && i < len(runes)-1 && unicode.IsUpper(prevRune) && unicode.IsUpper(currentRune) && unicode.IsLower(nextRune) {
			currentWord, words = saveWord(currentWord, words)
		}

		// Otherwise, append currentRune to currentWord
		currentWord = append(currentWord, currentRune)
	}
	// Add the last word if currentWord is not empty
	currentWord, words = saveWord(currentWord, words)

	return words
}

// isDelimiter is a helper to check if a rune is a delimiter
func isDelimiter(r rune) bool {
	return r == ' ' || r == '_' || r == '-'
}

// saveWord appends the current word to the words slice if non-empty,
// converts it to lowercase, and returns a fresh empty slice for building the next word.
func saveWord(currentWord []rune, words []string) ([]rune, []string) {
	if len(currentWord) > 0 {
		words = append(words, strings.ToLower(string(currentWord)))
	}
	return []rune{}, words
}

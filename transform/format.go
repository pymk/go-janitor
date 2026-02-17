package transform

import (
	"strings"
	"unicode"
)

// ToSnake converts a string to snake_case
// Expects the string to be pre-processed with SplitIntoWords if needed
func ToSnake(s []string) string {
	return strings.ToLower(strings.Join(s, "_"))
}

// ToKebab converts a string to kebab-case
// Expects the string to be pre-processed with SplitIntoWords if needed
func ToKebab(s []string) string {
	return strings.ToLower(strings.Join(s, "-"))
}

// ToCamel converts a string to camelCase
// Expects the string to be pre-processed with SplitIntoWords if needed
func ToCamel(s []string) string {
	var w []string
	for i, v := range s {
		if i == 0 {
			w = append(w, strings.ToLower(v))
		} else {
			w = append(w, capitalize(v))
		}
	}
	return strings.Join(w, "")
}

// ToPascal converts a string to PascalCase
// Expects the string to be pre-processed with SplitIntoWords if needed
func ToPascal(s []string) string {
	var w []string
	for _, v := range s {
		w = append(w, capitalize(v))
	}
	return strings.Join(w, "")
}

// capitalize returns a string with the first letter uppercase and the rest lowercase
func capitalize(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

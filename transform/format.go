package transform

import (
	_ "strings"
)

// ToSnake converts a string to snake_case
// Expects the string to be pre-processed with SplitIntoWords if needed
func ToSnake(s string) string {}

// ToKebab converts a string to kebab-case
// Expects the string to be pre-processed with SplitIntoWords if needed
func ToKebab(s string) string {}

// ToCamel converts a string to camelCase
// Expects the string to be pre-processed with SplitIntoWords if needed
func ToCamel(s string) string {}

// ToPascal converts a string to PascalCase
// Expects the string to be pre-processed with SplitIntoWords if needed
func ToPascal(s string) string {}

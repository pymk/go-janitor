package cli

import (
	_ "bufio"
	_ "fmt"
	_ "io"
	_ "os"

	_ "github.com/go-janitor/transform"
)

// CLI defines the command-line interface structure
type CLI struct {
	// Case conversion flags (mutually exclusive)
	Lower bool `help:"Convert to lowercase."`
	Upper bool `help:"Convert to UPPERCASE."`

	// Format conversion flags (mutually exclusive)
	Snake  bool `help:"Convert to snake_case."`
	Kebab  bool `help:"Convert to kebab-case."`
	Camel  bool `help:"Convert to camelCase."`
	Pascal bool `help:"Convert to PascalCase."`

	// Cleaning flags
	RemoveSpecial  bool `help:"Keep only alphanumeric and spaces."`
	CollapseSpaces bool `help:"Collapse multiple spaces into one."`
}

// Run executes the transformation pipeline
// Reads from stdin, applies transformations, writes to stdout
func (c *CLI) Run() error {
	// NOTE: Pipeline will be:
	// result := Trim(input)
	// if CollapseSpaces flag: result = CollapseSpaces(result)
	// if RemoveSpecial flag: result = RemoveSpecialChars(result)
	// if format flag is set: result = SplitIntoWords(result) then result = apply format
	// if case flag is set: result = apply case
}

// GetFormatFunc returns the appropriate format conversion function based on flags
// Returns nil if no format flag is set
// If multiple format flags are set, precedence: pascal -> camel -> kebab -> snake
func (c *CLI) GetFormatFunc() func(string) string {}

// GetCaseFunc returns the appropriate case conversion function based on flags
// Returns nil if no case flag is set
// If multiple case flags are set, precedence: upper -> lower
func (c *CLI) GetCaseFunc() func(string) string {}

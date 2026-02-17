package cli

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/pymk/go-janitor/transform"
)

// CLI defines the command-line interface structure
type CLI struct {
	// Case conversion flags
	Lower bool `help:"Convert to lowercase." short:"l"`
	Upper bool `help:"Convert to UPPERCASE." short:"u"`

	// Format conversion flags
	Snake  bool `help:"Convert to snake_case." short:"s"`
	Kebab  bool `help:"Convert to kebab-case." short:"k"`
	Camel  bool `help:"Convert to camelCase." short:"c"`
	Pascal bool `help:"Convert to PascalCase." short:"p"`

	// Cleaning flags
	RemoveSpecial  bool `help:"Keep only alphanumeric and spaces." short:"C"`
	CollapseSpaces bool `help:"Collapse multiple spaces into one." short:"S"`
}

// Run executes the transformation pipeline
// Reads from stdin, applies transformations, writes to stdout
func (c *CLI) Run() error {
	// Read from stdin
	reader := bufio.NewReader(os.Stdin)
	input, err := io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("error reading stdin: %w", err)
	}

	result := string(input)

	// Step 1: Always trim
	result = transform.Trim(result)

	// Step 2: Collapse spaces if flag is set
	if c.CollapseSpaces {
		result = transform.CollapseSpaces(result)
	}

	// Step 3: Remove special chars if flag is set
	if c.RemoveSpecial {
		result = transform.RemoveSpecialChars(result)
	}

	// Step 4: Apply format conversion if flag is set
	formatFunc := c.GetFormatFunc()
	if formatFunc != nil {
		words := transform.SplitIntoWords(result)
		result = formatFunc(words)
	}

	// Step 5: Apply case conversion if flag set
	caseFunc := c.GetCaseFunc()
	if caseFunc != nil {
		result = caseFunc(result)
	}

	// Write to stdout
	fmt.Println(result)
	return nil
}

// GetFormatFunc returns the appropriate format conversion function based on flags
// Returns nil if no format flag is set
// If multiple format flags are set, precedence: pascal -> camel -> kebab -> snake
func (c *CLI) GetFormatFunc() func([]string) string {
	if c.Pascal {
		return transform.ToPascal
	}
	if c.Camel {
		return transform.ToCamel
	}
	if c.Kebab {
		return transform.ToKebab
	}
	if c.Snake {
		return transform.ToSnake
	}
	return nil
}

// GetCaseFunc returns the appropriate case conversion function based on flags
// Returns nil if no case flag is set
// If multiple case flags are set, precedence: upper -> lower
func (c *CLI) GetCaseFunc() func(string) string {
	if c.Upper {
		return transform.ToUpper
	}
	if c.Lower {
		return transform.ToLower
	}
	return nil
}

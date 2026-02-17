# README

String clean-up CLI tool.

## How It Works

Given a string in stdin, the string pipeline runs in this order:

1. Trim (leading and trailing whitespaces)
2. Collapse spaces
3. Remove special characters (only keep alphanumeric)
4. Normalize to words (if format flag is given)
5. Apply format
  - If multiple format flags are set, precedence: pascal -> camel -> kebab -> snake
6. Apply case
  - If multiple case flags are set, precedence: upper -> lower

## Examples

```bash
echo "  hello @#$%^&* world " | go run . --snake --upper --remove-special --collapse-spaces
# HELLO_WORLD

# Combine using short flags
echo "  hello @#$%^&* world " | go run . -su -CS
# HELLO_WORLD

go run . --help
# Usage: janitor [flags]
# String clean-up CLI tool for transforming stdin.
# Flags:
#   -h, --help               Show context-sensitive help.
#   -l, --lower              Convert to lowercase.
#   -u, --upper              Convert to UPPERCASE.
#   -s, --snake              Convert to snake_case.
#   -k, --kebab              Convert to kebab-case.
#   -c, --camel              Convert to camelCase.
#   -p, --pascal             Convert to PascalCase.
#   -C, --remove-special     Keep only alphanumeric and spaces.
#   -S, --collapse-spaces    Collapse multiple spaces into one.
```

## Installation

1. Build the binary:

```bash
# Build the binary
go build -o janitor

# Move to local bin directory
mkdir -p ~/.local/bin && mv janitor ~/.local/bin/

# NOTE: Ensure `~/.local/bin` is in your PATH

# Reload shell
source ~/.zshrc  # or source ~/.bashrc

# Test
echo "  hello @#$%^&* world " | janitor -slCS
# hello_world
```

## Structure

```
go-janitor/
├── main.go
├── cli/
│   └── cli.go          # CLI logic
├── transform/
│   ├── case.go         # Casing
│   ├── format.go       # Formatting
│   └── clean.go        # Cleaning
└── go.mod
```

## Fun Fact

The name is a nod to the [janitor](https://sfirke.github.io/janitor/index.html) package in R.

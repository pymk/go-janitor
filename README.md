# README

String clean-up CLI tool.

How it works: Given a string in stdin, the string pipeline runs in this order:

1. Trim (leading and trailing whitespaces)
2. Collapse spaces
3. Remove special characters (only keep alphanumeric)
4. Normalize to words (if format flag is given)
5. Apply format
  - If multiple format flags are set, precedence: pascal -> camel -> kebab -> snake
6. Apply case
  - If multiple case flags are set, precedence: upper -> lower

The name is a reference to the [janitor](https://sfirke.github.io/janitor/index.html) package in R.

## Structure

```
go-janitor/
├── main.go
├── cli/
│   └── cli.go          # CLI logic
├── transform/
│   ├── case.go         # Casing (ToLower, ToUpper)
│   ├── format.go       # Formatting (ToSnake, ToKebab, ToCamel, ToPascal)
│   └── clean.go        # Cleaning (RemoveSpecial, CollapseSpaces)
└── go.mod
```

## TODO

- [X] Plan and design
- [X] Create directories and files
- [X] Add boilerplate code
- [ ] ?

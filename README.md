# Go Markdownify

A Go library for converting HTML to Markdown, designed as a direct port of
[python-markdownify](https://github.com/matthewwithanm/python-markdownify) by
[Matthew Tretter](https://github.com/matthewwithanm),
[AlexVonB](https://github.com/AlexVonB), [Chris
Papademetrious](https://github.com/chrispitude), and many other valuable
contributors. This implementation aims to match the Python package as closely as
possible in both functionality and API.

## Features

- Convert HTML to Markdown
- Customize the output with various options
- Support for most HTML tags:
  - Headings (h1-h6)
  - Paragraphs
  - Lists (ordered and unordered)
  - Links
  - Images
  - Blockquotes
  - Code blocks
  - Tables
  - Inline formatting (bold, italic, code, etc.)
- Configurable options:
  - Heading style (ATX, ATX_CLOSED, or UNDERLINED)
  - Strong/emphasis symbol (asterisk or underscore)
  - Newline style (spaces or backslash)
  - Code language
  - And more...

## Installation

```bash
go get github.com/mrjoshuak/go-markdownify
```

## Usage

### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/mrjoshuak/go-markdownify"
)

func main() {
    html := "<h1>Hello, World!</h1><p>This is a <strong>test</strong>.</p>"

    // Use default options
    markdown, err := gomarkdownify.Convert(html)
    if err != nil {
        panic(err)
    }
    fmt.Println(markdown)
}
```

### Custom Options

```go
package main

import (
    "fmt"
    "github.com/mrjoshuak/go-markdownify"
)

func main() {
    html := "<h1>Hello, World!</h1><p>This is a <strong>test</strong>.</p>"

    // Use custom options
    options := gomarkdownify.DefaultOptions()
    options.HeadingStyle = gomarkdownify.ATX
    options.StrongEmSymbol = gomarkdownify.UNDERSCORE
    options.NewlineStyle = gomarkdownify.BACKSLASH

    markdown, err := gomarkdownify.Convert(html, options)
    if err != nil {
        panic(err)
    }
    fmt.Println(markdown)
}
```

## Options

| Option               | Type     | Default    | Description                                                           |
| -------------------- | -------- | ---------- | --------------------------------------------------------------------- |
| Autolinks            | bool     | true       | Use `<url>` syntax for URLs that match their link text                |
| Bullets              | string   | "*+-"      | String of bullet characters to use for unordered lists                |
| CodeLanguage         | string   | ""         | Default language for code blocks                                      |
| CodeLanguageCallback | func     | nil        | Function to determine code language from node                         |
| Convert              | []string | nil        | List of tags to convert (if nil, convert all)                         |
| DefaultTitle         | bool     | false      | Use href as title for links when no title is provided                 |
| EscapeAsterisks      | bool     | true       | Escape * in text                                                      |
| EscapeUnderscores    | bool     | true       | Escape _ in text                                                      |
| EscapeMisc           | bool     | false      | Escape other special characters                                       |
| HeadingStyle         | string   | UNDERLINED | Style for headings (ATX, ATX_CLOSED, or UNDERLINED)                   |
| KeepInlineImagesIn   | []string | []         | List of tags to keep inline images in                                 |
| NewlineStyle         | string   | SPACES     | Style for line breaks (SPACES or BACKSLASH)                           |
| NormalizeNewlines    | bool     | true       | Normalize multiple consecutive newlines to a maximum of 2             |
| Strip                | []string | nil        | List of tags to strip (if nil, strip none)                            |
| StripDocument        | string   | LSTRIP     | How to strip document-level whitespace (LSTRIP, RSTRIP, STRIP, or "") |
| StrongEmSymbol       | string   | ASTERISK   | Symbol for strong and emphasis (ASTERISK or UNDERSCORE)               |
| SubSymbol            | string   | ""         | Symbol for subscript                                                  |
| SupSymbol            | string   | ""         | Symbol for superscript                                                |
| TableInferHeader     | bool     | true       | Infer table headers when not explicitly defined                       |
| Wrap                 | bool     | false      | Wrap text at specified width                                          |
| WrapWidth            | int      | 80         | Width to wrap text at                                                 |

## License

This project is licensed under the [MIT License](LICENSE).

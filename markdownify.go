// Package gomarkdownify provides functionality to convert HTML to Markdown.
// It is designed as a direct port of the Python markdownify package
// (https://github.com/matthewwithanm/python-markdownify) and aims to match its
// behavior as closely as possible.
package gomarkdownify

// Convert transforms HTML content into Markdown format.
//
// This is the main entry point for the package. It creates a converter with
// either the default options or the provided custom options, and then uses
// that converter to transform the HTML into Markdown.
//
// The conversion process handles most common HTML elements, including headings,
// paragraphs, lists, links, images, blockquotes, code blocks, tables, and inline
// formatting elements like bold, italic, and code.
//
// Parameters:
//   - html: The HTML content to convert to Markdown.
//   - options: Optional configuration options. If not provided, default options are used.
//
// Returns:
//   - A string containing the Markdown representation of the HTML.
//   - An error if the conversion process fails.
//
// Example:
//
//	html := "<h1>Hello, World!</h1><p>This is a <strong>test</strong>.</p>"
//	markdown, err := gomarkdownify.Convert(html)
//	if err != nil {
//	    // handle error
//	}
//	fmt.Println(markdown)
//
// With custom options:
//
//	html := "<h1>Hello, World!</h1><p>This is a <strong>test</strong>.</p>"
//	options := gomarkdownify.DefaultOptions()
//	options.HeadingStyle = gomarkdownify.ATX
//	options.StrongEmSymbol = gomarkdownify.UNDERSCORE
//	markdown, err := gomarkdownify.Convert(html, options)
//	if err != nil {
//	    // handle error
//	}
//	fmt.Println(markdown)
func Convert(html string, options ...Options) (string, error) {
	opts := DefaultOptions()
	if len(options) > 0 {
		opts = options[0]
	}
	converter := NewConverter(opts)
	return converter.Convert(html)
}

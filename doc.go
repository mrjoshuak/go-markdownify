/*
Package gomarkdownify provides functionality to convert HTML to Markdown.

This package is designed as a direct port of the Python markdownify package
(https://github.com/matthewwithanm/python-markdownify) and aims to match its
behavior as closely as possible. The default behavior is now fully compatible
with the Python implementation.

Basic Usage:

	html := "<h1>Hello, World!</h1><p>This is a <strong>test</strong>.</p>"
	markdown, err := gomarkdownify.Convert(html)
	if err != nil {
	    // handle error
	}
	fmt.Println(markdown)

Custom Options:

	html := "<h1>Hello, World!</h1><p>This is a <strong>test</strong>.</p>"
	options := gomarkdownify.DefaultOptions()
	options.HeadingStyle = gomarkdownify.ATX
	options.StrongEmSymbol = gomarkdownify.UNDERSCORE
	markdown, err := gomarkdownify.Convert(html, options)
	if err != nil {
	    // handle error
	}
	fmt.Println(markdown)

Supported HTML Tags:

The package supports conversion of most common HTML tags to their Markdown equivalents:

  - Headings (h1-h6)
  - Paragraphs (p)
  - Lists (ul, ol, li)
  - Links (a)
  - Images (img)
  - Blockquotes (blockquote)
  - Code blocks (pre, code)
  - Tables (table, tr, th, td)
  - Inline formatting (b, strong, i, em, code, del, s, sub, sup)
  - Horizontal rules (hr)
  - Line breaks (br)

Configuration Options:

The package provides a wide range of configuration options to customize the
output. See the Options struct documentation for details.
*/
package gomarkdownify

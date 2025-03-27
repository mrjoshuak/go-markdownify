package gomarkdownify

import (
	"golang.org/x/net/html"
)

// Options defines the configuration options for the markdown converter.
// These options control how HTML is converted to Markdown, allowing for
// customization of the output format.
type Options struct {
	// Autolinks determines whether to use <url> syntax for URLs that match their link text.
	// When true, links like <a href="https://example.com">https://example.com</a> will be
	// converted to <https://example.com> instead of [https://example.com](https://example.com).
	Autolinks bool

	// Bullets specifies the characters to use for unordered list items.
	// The characters are used in the order provided, with different levels of nesting
	// using different characters. For example, "*+-" would use * for the first level,
	// + for the second level, and - for the third level.
	Bullets string

	// CodeLanguage specifies the default language for code blocks.
	// This is used when a code block doesn't have a language specified.
	CodeLanguage string

	// CodeLanguageCallback is a function that determines the language for a code block
	// based on the HTML node. This allows for custom logic to extract language information
	// from class attributes or other node properties.
	CodeLanguageCallback func(n *html.Node) string

	// Convert is a list of tags to convert. If nil, all supported tags are converted.
	// This can be used to limit which HTML tags are processed.
	Convert []string

	// DefaultTitle determines whether to use the href as the title for links
	// when no title attribute is provided.
	DefaultTitle bool
	
	// StripLinkTitles determines whether to strip all title attributes from links,
	// regardless of whether they were provided in the HTML.
	// This ensures consistent output with the Python markdownify package.
	StripLinkTitles bool

	// EscapeAsterisks determines whether to escape * characters in the text.
	// This prevents them from being interpreted as Markdown formatting.
	EscapeAsterisks bool

	// EscapeUnderscores determines whether to escape _ characters in the text.
	// This prevents them from being interpreted as Markdown formatting.
	EscapeUnderscores bool

	// EscapeMisc determines whether to escape other special characters in the text.
	// This includes characters like #, >, -, +, etc. that have special meaning in Markdown.
	EscapeMisc bool

	// HeadingStyle specifies the style to use for headings.
	// Valid values are ATX (# Heading), ATX_CLOSED (# Heading #), and UNDERLINED (Heading\n=====).
	HeadingStyle string

	// KeepInlineImagesIn is a list of tags in which to keep inline images.
	// By default, images are converted to Markdown image syntax, but this option
	// allows for keeping the original HTML for images within specified tags.
	KeepInlineImagesIn []string

	// NewlineStyle specifies the style to use for line breaks.
	// Valid values are SPACES (two spaces at end of line) and BACKSLASH (backslash at end of line).
	NewlineStyle string

	// NormalizeNewlines determines whether to normalize multiple consecutive newlines
	// to a maximum of 2. This helps maintain consistent spacing in the output.
	NormalizeNewlines bool

	// Strip is a list of tags to strip from the output. If nil, no tags are stripped.
	// Stripped tags are removed completely, including their content.
	Strip []string

	// StripDocument specifies how to handle whitespace at the document level.
	// Valid values are LSTRIP (remove leading newlines), RSTRIP (remove trailing newlines),
	// STRIP (remove both), or "" (don't strip).
	StripDocument string

	// StrongEmSymbol specifies the symbol to use for strong and emphasis formatting.
	// Valid values are ASTERISK (*emphasis* and **strong**) and UNDERSCORE (_emphasis_ and __strong__).
	StrongEmSymbol string

	// SubSymbol specifies the symbol to use for subscript text.
	// If empty, subscript is not converted to Markdown.
	SubSymbol string

	// SupSymbol specifies the symbol to use for superscript text.
	// If empty, superscript is not converted to Markdown.
	SupSymbol string

	// TableInferHeader determines whether to infer table headers when not explicitly defined.
	// When true, the first row of a table is treated as a header row if no <th> tags are present.
	TableInferHeader bool
	
	// DeduplicateHeadings determines whether to remove duplicate headings.
	// When true, subsequent identical headings will be removed from the output.
	// This helps match the behavior of the Python markdownify package.
	DeduplicateHeadings bool

	// Wrap determines whether to wrap text at a specified width.
	// When true, long lines are wrapped to improve readability.
	Wrap bool

	// WrapWidth specifies the width at which to wrap text when Wrap is true.
	// This is measured in characters.
	WrapWidth int
}

// DefaultOptions returns the default options for the markdown converter.
// These defaults are designed to match the behavior of the Python markdownify package.
func DefaultOptions() Options {
	return Options{
		Autolinks:           true,
		Bullets:             "*+-",
		CodeLanguage:        "",
		Convert:             nil,
		DefaultTitle:        false,
		StripLinkTitles:     true,  // Strip titles to match Python markdownify behavior
		EscapeAsterisks:     true,
		EscapeUnderscores:   true,
		EscapeMisc:          false,
		HeadingStyle:        UNDERLINED,
		KeepInlineImagesIn:  []string{},
		NewlineStyle:        SPACES,
		NormalizeNewlines:   true,
		Strip:               nil,
		StripDocument:       LSTRIP,
		StrongEmSymbol:      ASTERISK,
		SubSymbol:           "",
		SupSymbol:           "",
		TableInferHeader:    true, // Match Python markdownify behavior
		DeduplicateHeadings: true, // Match Python markdownify behavior
		Wrap:                false,
		WrapWidth:           80,
	}
}

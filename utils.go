package gomarkdownify

import (
	"strings"

	"golang.org/x/net/html"
)

// chomp removes leading and trailing whitespace from text and returns the
// prefix (leading whitespace), suffix (trailing whitespace), and the trimmed text.
//
// This function is used to preserve whitespace when applying Markdown formatting
// to text. For example, when converting "<strong> text </strong>" to "** text **",
// we want to keep the spaces around the text.
//
// Parameters:
//   - text: The text to process.
//
// Returns:
//   - prefix: The leading whitespace (if any).
//   - suffix: The trailing whitespace (if any).
//   - middle: The text with leading and trailing whitespace removed.
func chomp(text string) (string, string, string) {
	// Handle empty text
	if text == "" {
		return "", "", ""
	}

	// Find the first non-space character
	start := 0
	for start < len(text) && text[start] == ' ' {
		start++
	}

	// Find the last non-space character
	end := len(text) - 1
	for end >= 0 && text[end] == ' ' {
		end--
	}

	// Extract the parts
	var prefix, suffix, middle string

	if start > 0 {
		prefix = strings.Repeat(" ", start)
	}

	if end >= start {
		middle = text[start : end+1]
	}

	if end < len(text)-1 {
		suffix = strings.Repeat(" ", len(text)-end-1)
	}

	return prefix, suffix, middle
}

// contains checks if a slice contains a string.
//
// This is a utility function used throughout the package to check if a tag
// is in a list of tags, or if a parent tag is in a list of parent tags.
//
// Parameters:
//   - slice: The slice of strings to search.
//   - s: The string to look for.
//
// Returns:
//   - true if the string is found in the slice, false otherwise.
func contains(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

// shouldRemoveWhitespaceInside returns true if whitespace should be removed
// inside a block-level element.
//
// This function is used to determine if leading and trailing whitespace should
// be removed from the content of a block-level element. This helps maintain
// clean Markdown output without excessive whitespace.
//
// Parameters:
//   - n: The HTML node to check.
//
// Returns:
//   - true if whitespace should be removed inside the node, false otherwise.
func shouldRemoveWhitespaceInside(n *html.Node) bool {
	if n == nil || n.Type != html.ElementNode {
		return false
	}

	if reHTMLHeading.MatchString(n.Data) {
		return true
	}

	switch n.Data {
	case "p", "blockquote", "article", "div", "section", "ol", "ul", "li",
		"dl", "dt", "dd", "table", "thead", "tbody", "tfoot", "tr", "td", "th":
		return true
	}

	return false
}

// shouldRemoveWhitespaceOutside returns true if whitespace should be removed
// outside a block-level element.
//
// This function is used to determine if whitespace should be removed before
// and after a block-level element. This helps maintain clean Markdown output
// without excessive whitespace.
//
// Parameters:
//   - n: The HTML node to check.
//
// Returns:
//   - true if whitespace should be removed outside the node, false otherwise.
func shouldRemoveWhitespaceOutside(n *html.Node) bool {
	if n == nil || n.Type != html.ElementNode {
		return false
	}

	return shouldRemoveWhitespaceInside(n) || n.Data == "pre"
}

// getAttr gets an attribute value from an HTML node.
//
// This function is used to retrieve attribute values from HTML nodes, such as
// href from <a> tags, src from <img> tags, etc.
//
// Parameters:
//   - n: The HTML node to get the attribute from.
//   - key: The name of the attribute to get.
//
// Returns:
//   - The value of the attribute, or an empty string if the attribute is not found.
func getAttr(n *html.Node, key string) string {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}

// abstractInlineConversion handles simple inline tags like b, em, del, etc.
//
// This function provides a common implementation for converting inline HTML
// elements to their Markdown equivalents. It handles whitespace preservation
// and applies the appropriate Markdown syntax.
//
// Parameters:
//   - n: The HTML node being converted.
//   - text: The text content of the node.
//   - parentTags: A list of parent tags to check for special handling.
//   - markup: The Markdown syntax to apply (e.g., "*" for emphasis).
//
// Returns:
//   - The text with Markdown formatting applied.
func (c *Converter) abstractInlineConversion(n *html.Node, text string, parentTags []string, markup string) string {
	if contains(parentTags, "_noformat") {
		return text
	}

	prefix, suffix, text := chomp(text)
	if text == "" {
		return prefix + suffix
	}

	return prefix + markup + text + markup + suffix
}

// min returns the smaller of two integers.
//
// Parameters:
//   - a: The first integer.
//   - b: The second integer.
//
// Returns:
//   - The smaller of a and b.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// max returns the larger of two integers.
//
// Parameters:
//   - a: The first integer.
//   - b: The second integer.
//
// Returns:
//   - The larger of a and b.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

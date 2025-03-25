package gomarkdownify

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestLinkFormatting(t *testing.T) {
	tests := []struct {
		name     string
		html     string
		expected string
	}{
		{
			name:     "Simple link",
			html:     `<a href="https://example.com">Example</a>`,
			expected: `[Example](https://example.com)`,
		},
		{
			name:     "Link with title",
			html:     `<a href="https://example.com" title="Example Website">Example</a>`,
			expected: `[Example](https://example.com "Example Website")`,
		},
		{
			name:     "Link in paragraph",
			html:     `<p>This is a <a href="https://example.com">link</a> in a paragraph.</p>`,
			expected: `This is a [link](https://example.com) in a paragraph.`,
		},
		{
			name:     "Multiple links in paragraph",
			html:     `<p>Here are <a href="https://example.com">two</a> <a href="https://example.org">links</a>.</p>`,
			expected: `Here are [two](https://example.com) [links](https://example.org).`,
		},
		{
			name:     "Link in heading",
			html:     `<h1>Heading with <a href="https://example.com">link</a></h1>`,
			expected: `# Heading with [link](https://example.com)`,
		},
		{
			name:     "Link in list item",
			html:     `<ul><li>Item with <a href="https://example.com">link</a></li></ul>`,
			expected: `* Item with [link](https://example.com)`,
		},
		{
			name:     "Link in table cell",
			html:     `<table><tr><td>Cell with <a href="https://example.com">link</a></td></tr></table>`,
			expected: `| Cell with [link](https://example.com) |`,
		},
		{
			name:     "Standalone link",
			html:     `<div><a href="https://example.com">Example</a></div>`,
			expected: `[Example](https://example.com)`,
		},
		{
			name:     "Link with nested elements",
			html:     `<a href="https://example.com">Example <strong>with</strong> formatting</a>`,
			expected: `[Example **with** formatting](https://example.com)`,
		},
		{
			name:     "Link that matches href (autolink)",
			html:     `<a href="https://example.com">https://example.com</a>`,
			expected: `<https://example.com>`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test with both Convert and ConvertWithFix
			options := DefaultOptions()
			options.Autolinks = true // Enable autolinks for the autolink test case

			// Test with Convert
			result, err := Convert(tt.html, options)
			if err != nil {
				t.Errorf("Convert error: %v", err)
				return
			}
			result = strings.TrimSpace(result)
			if result != tt.expected {
				t.Errorf("Convert result\nGot:  %q\nWant: %q", result, tt.expected)
			}

			// Test with ConvertWithFix
			resultWithFix, err := ConvertWithFix(tt.html, options)
			if err != nil {
				t.Errorf("ConvertWithFix error: %v", err)
				return
			}
			resultWithFix = strings.TrimSpace(resultWithFix)
			if resultWithFix != tt.expected {
				t.Errorf("ConvertWithFix result\nGot:  %q\nWant: %q", resultWithFix, tt.expected)
			}
		})
	}
}

func TestRealWorldLinks(t *testing.T) {
	tests := []struct {
		name     string
		html     string
		contains string // Instead of exact match, check if output contains this string
	}{
		{
			name: "Example.com more info link",
			html: `<p>This domain is for use in illustrative examples in documents. You may use this
			domain in literature without prior coordination or asking for permission.</p>
			<p><a href="https://www.iana.org/domains/example">More information...</a></p>`,
			contains: `[More information...](https://www.iana.org/domains/example)`,
		},
		{
			name: "IANA example domain RFC links",
			html: `<p>As described in <a href="/go/rfc2606">RFC 2606</a> and <a href="/go/rfc6761">RFC 6761</a>,
			a number of domains such as example.com and example.org are maintained for documentation purposes.</p>`,
			contains: `[RFC 2606](/go/rfc2606)`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			options := DefaultOptions()

			// Test with ConvertWithFix
			resultWithFix, err := ConvertWithFix(tt.html, options)
			if err != nil {
				t.Errorf("ConvertWithFix error: %v", err)
				return
			}

			if !strings.Contains(resultWithFix, tt.contains) {
				t.Errorf("ConvertWithFix result does not contain expected string\nGot:  %q\nShould contain: %q", resultWithFix, tt.contains)
			}
		})
	}
}

// TestLinkFormattingWithDebug is a more detailed test that prints intermediate values
// to help debug link formatting issues
func TestLinkFormattingWithDebug(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping debug test in short mode")
	}

	html := `<p>This is a <a href="https://example.com">link</a> in a paragraph.</p>`
	expected := `This is a [link](https://example.com) in a paragraph.`

	options := DefaultOptions()

	// Create a converter directly to inspect internal state
	converter := NewConverter(options)

	// Parse the HTML
	doc, err := parseHTML(html)
	if err != nil {
		t.Fatalf("Failed to parse HTML: %v", err)
	}

	// Print the HTML structure
	t.Logf("HTML structure:")
	printNode(t, doc, 0)

	// Convert with the fixed method
	result, err := converter.ConvertFixed(html)
	if err != nil {
		t.Fatalf("ConvertFixed error: %v", err)
	}

	result = strings.TrimSpace(result)
	t.Logf("Result: %q", result)
	t.Logf("Expected: %q", expected)

	if result != expected {
		t.Errorf("ConvertFixed result does not match expected\nGot:  %q\nWant: %q", result, expected)
	}
}

// Helper function to parse HTML
func parseHTML(htmlContent string) (*html.Node, error) {
	return html.Parse(strings.NewReader(htmlContent))
}

// Helper function to print the HTML node structure
func printNode(t *testing.T, n *html.Node, level int) {
	indent := strings.Repeat("  ", level)

	switch n.Type {
	case html.ElementNode:
		t.Logf("%sElement: <%s>", indent, n.Data)
		for _, attr := range n.Attr {
			t.Logf("%s  Attr: %s=%q", indent, attr.Key, attr.Val)
		}
	case html.TextNode:
		text := strings.TrimSpace(n.Data)
		if text != "" {
			t.Logf("%sText: %q", indent, text)
		}
	case html.CommentNode:
		t.Logf("%sComment: %q", indent, n.Data)
	case html.DocumentNode:
		t.Logf("%sDocument", indent)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		printNode(t, c, level+1)
	}

	if n.Type == html.ElementNode {
		t.Logf("%sElement: </%s>", indent, n.Data)
	}
}

package gomarkdownify

import (
	"strings"
	"testing"
)

// TestWhitespaceHandling tests the whitespace and spacing issues reported
// by downstream users.
func TestWhitespaceHandling(t *testing.T) {
	// Test case from the reported issues document
	html := `
<h1>Test Heading</h1>
<p>First paragraph.</p>
<p>Second paragraph.</p>
<ul>
  <li>Item 1</li>
  <li>Item 2</li>
</ul>
`

	// Expected output from the issues document
	expected := `# Test Heading

First paragraph.

Second paragraph.

* Item 1
* Item 2`

	options := DefaultOptions()
	options.HeadingStyle = ATX
	options.StripDocument = STRIP // Match the expected output (no trailing newlines)

	result, err := Convert(html, options)
	if err != nil {
		t.Fatalf("Failed to convert HTML: %v", err)
	}

	// Normalize newlines for comparison
	result = strings.TrimSpace(result)
	expected = strings.TrimSpace(expected)

	if result != expected {
		t.Errorf("Expected:\n%s\n\nGot:\n%s", expected, result)
	}
}

// TestHeadingDuplication tests the heading duplication issue reported
// by downstream users.
func TestHeadingDuplication(t *testing.T) {
	// Test case from the reported issues document
	html := `
<div>
  <h1>Example Domain</h1>
  <div>
    <h1>Example Domain</h1>
    <p>This domain is for use in illustrative examples in documents.</p>
  </div>
</div>
`

	// Expected output from the issues document
	expected := `# Example Domain

This domain is for use in illustrative examples in documents.`

	options := DefaultOptions()
	options.HeadingStyle = ATX
	options.DeduplicateHeadings = true
	options.StripDocument = STRIP // Match the expected output (no trailing newlines)

	result, err := Convert(html, options)
	if err != nil {
		t.Fatalf("Failed to convert HTML: %v", err)
	}

	// Normalize newlines for comparison
	result = strings.TrimSpace(result)
	expected = strings.TrimSpace(expected)

	if result != expected {
		t.Errorf("Expected:\n%s\n\nGot:\n%s", expected, result)
	}
}

// TestLinkTitleStripping tests the link title handling issue reported
// by downstream users.
func TestLinkTitleStripping(t *testing.T) {
	// Test case from the reported issues document
	html := `<p>This is a test with a <a href="https://example.com">link</a> and another <a href="https://example.org" title="Example">link with title</a>.</p>`

	// Expected output from the issues document
	expected := `This is a test with a [link](https://example.com) and another [link with title](https://example.org).`

	options := DefaultOptions()
	options.StripLinkTitles = true
	options.StripDocument = STRIP // Match the expected output (no trailing newlines)

	result, err := Convert(html, options)
	if err != nil {
		t.Fatalf("Failed to convert HTML: %v", err)
	}

	// Normalize newlines for comparison
	result = strings.TrimSpace(result)
	expected = strings.TrimSpace(expected)

	if result != expected {
		t.Errorf("Expected:\n%s\n\nGot:\n%s", expected, result)
	}
}
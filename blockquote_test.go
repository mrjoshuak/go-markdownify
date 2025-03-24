package gomarkdownify

import (
	"strings"
	"testing"
)

// TestBlockquoteAdvanced tests advanced blockquote conversion scenarios
func TestBlockquoteAdvanced(t *testing.T) {
	// Test blockquote with multiple paragraphs
	html := `<blockquote>
		<p>First paragraph in blockquote</p>
		<p>Second paragraph in blockquote</p>
	</blockquote>`

	result, err := Convert(html)
	if err != nil {
		t.Fatalf("Error converting HTML: %v", err)
	}

	if !strings.Contains(result, "> First paragraph") &&
		!strings.Contains(result, "> Second paragraph") {
		t.Errorf("Expected blockquote with multiple paragraphs, got %q", result)
	}

	// Test blockquote with nested lists
	html = `<blockquote>
		<p>Blockquote with a list:</p>
		<ul>
			<li>Item 1</li>
			<li>Item 2</li>
		</ul>
	</blockquote>`

	result, err = Convert(html)
	if err != nil {
		t.Fatalf("Error converting HTML: %v", err)
	}

	if !strings.Contains(result, "> Blockquote with a list") &&
		!strings.Contains(result, "> * Item 1") &&
		!strings.Contains(result, "> * Item 2") {
		t.Errorf("Expected blockquote with list, got %q", result)
	}

	// Test deeply nested blockquotes
	html = `<blockquote>
		<p>Level 1</p>
		<blockquote>
			<p>Level 2</p>
			<blockquote>
				<p>Level 3</p>
			</blockquote>
		</blockquote>
	</blockquote>`

	result, err = Convert(html)
	if err != nil {
		t.Fatalf("Error converting HTML: %v", err)
	}

	if !strings.Contains(result, "> Level 1") &&
		!strings.Contains(result, "> > Level 2") &&
		!strings.Contains(result, "> > > Level 3") {
		t.Errorf("Expected deeply nested blockquotes, got %q", result)
	}

	// Test blockquote with code blocks
	html = `<blockquote>
		<p>Blockquote with code:</p>
		<pre><code>function example() {
  return true;
}</code></pre>
	</blockquote>`

	result, err = Convert(html)
	if err != nil {
		t.Fatalf("Error converting HTML: %v", err)
	}

	if !strings.Contains(result, "> Blockquote with code") &&
		!strings.Contains(result, "function example()") {
		t.Errorf("Expected blockquote with code block, got %q", result)
	}

	// Test blockquote with inline formatting
	html = `<blockquote>
		<p>Blockquote with <strong>bold</strong>, <em>italic</em>, and <code>code</code> formatting.</p>
	</blockquote>`

	result, err = Convert(html)
	if err != nil {
		t.Fatalf("Error converting HTML: %v", err)
	}

	if !strings.Contains(result, "> Blockquote with **bold**, *italic*, and `code` formatting") {
		t.Errorf("Expected blockquote with inline formatting, got %q", result)
	}
}

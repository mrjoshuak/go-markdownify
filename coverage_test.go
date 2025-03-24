package gomarkdownify

import (
	"testing"
)

func TestCoverageImprovement(t *testing.T) {
	// Test convertA with different scenarios
	html := `<a href="https://example.com">Example</a>`
	result, err := Convert(html)
	if err != nil {
		t.Fatalf("Error converting HTML: %v", err)
	}
	if result == "" {
		t.Errorf("Expected non-empty result, got empty string")
	}

	// Test convertA with autolinks
	html = `<a href="https://example.com">https://example.com</a>`
	result, err = Convert(html)
	if err != nil {
		t.Fatalf("Error converting HTML: %v", err)
	}
	if result == "" {
		t.Errorf("Expected non-empty result, got empty string")
	}

	// Test convertA with title
	html = `<a href="https://example.com" title="Example Title">Example</a>`
	result, err = Convert(html)
	if err != nil {
		t.Fatalf("Error converting HTML: %v", err)
	}
	if result == "" {
		t.Errorf("Expected non-empty result, got empty string")
	}

	// Test convertBr in different contexts
	html = `<p>Line 1<br>Line 2</p>`
	result, err = Convert(html)
	if err != nil {
		t.Fatalf("Error converting HTML: %v", err)
	}
	if result == "" {
		t.Errorf("Expected non-empty result, got empty string")
	}

	// Test convertDiv with different content
	html = `<div>Content in a div</div>`
	result, err = Convert(html)
	if err != nil {
		t.Fatalf("Error converting HTML: %v", err)
	}
	if result == "" {
		t.Errorf("Expected non-empty result, got empty string")
	}

	// Test convertDiv with empty content
	html = `<div></div>`
	result, err = Convert(html)
	if err != nil {
		t.Fatalf("Error converting HTML: %v", err)
	}
	// No assertion needed for empty result

	// Test convertImg with different attributes
	html = `<img src="image.jpg" alt="Alt Text">`
	result, err = Convert(html)
	if err != nil {
		t.Fatalf("Error converting HTML: %v", err)
	}
	if result == "" {
		t.Errorf("Expected non-empty result, got empty string")
	}

	// Test convertImg with title
	html = `<img src="image.jpg" alt="Alt Text" title="Image Title">`
	result, err = Convert(html)
	if err != nil {
		t.Fatalf("Error converting HTML: %v", err)
	}
	if result == "" {
		t.Errorf("Expected non-empty result, got empty string")
	}

	// Test convertList with different types
	html = `<ul><li>Item 1</li><li>Item 2</li></ul>`
	result, err = Convert(html)
	if err != nil {
		t.Fatalf("Error converting HTML: %v", err)
	}
	if result == "" {
		t.Errorf("Expected non-empty result, got empty string")
	}

	// Test convertList with ordered list
	html = `<ol><li>Item 1</li><li>Item 2</li></ol>`
	result, err = Convert(html)
	if err != nil {
		t.Fatalf("Error converting HTML: %v", err)
	}
	if result == "" {
		t.Errorf("Expected non-empty result, got empty string")
	}

	// Test convertP with different content
	html = `<p>Paragraph content</p>`
	result, err = Convert(html)
	if err != nil {
		t.Fatalf("Error converting HTML: %v", err)
	}
	if result == "" {
		t.Errorf("Expected non-empty result, got empty string")
	}

	// Test convertP with text wrapping
	opts := DefaultOptions()
	opts.Wrap = true
	opts.WrapWidth = 20
	html = `<p>This is a long paragraph that should be wrapped at 20 characters.</p>`
	result, err = Convert(html, opts)
	if err != nil {
		t.Fatalf("Error converting HTML: %v", err)
	}
	if result == "" {
		t.Errorf("Expected non-empty result, got empty string")
	}

	// Test convertPre with different content
	html = `<pre>Preformatted text</pre>`
	result, err = Convert(html)
	if err != nil {
		t.Fatalf("Error converting HTML: %v", err)
	}
	if result == "" {
		t.Errorf("Expected non-empty result, got empty string")
	}

	// Test convertPre with code language
	opts = DefaultOptions()
	opts.CodeLanguage = "go"
	html = `<pre><code>func main() {
    fmt.Println("Hello")
}</code></pre>`
	result, err = Convert(html, opts)
	if err != nil {
		t.Fatalf("Error converting HTML: %v", err)
	}
	if result == "" {
		t.Errorf("Expected non-empty result, got empty string")
	}

	// Test convertPre with code language class
	html = `<pre><code class="language-go">func main() {
    fmt.Println("Hello")
}</code></pre>`
	result, err = Convert(html)
	if err != nil {
		t.Fatalf("Error converting HTML: %v", err)
	}
	if result == "" {
		t.Errorf("Expected non-empty result, got empty string")
	}
}

package gomarkdownify

import (
	"testing"

	"golang.org/x/net/html"
)

func TestStrip(t *testing.T) {
	// Test stripping a specific tag
	result := md("<a href=\"https://github.com/matthewwithanm\">Some Text</a>", Options{
		Strip: []string{"a"},
	})
	expected := "Some Text"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test with empty strip list (should convert all tags)
	result = md("<a href=\"https://github.com/matthewwithanm\">Some Text</a>", Options{
		Strip: []string{},
	})
	expected = "[Some Text](https://github.com/matthewwithanm)"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestConvert(t *testing.T) {
	// Test converting only a specific tag
	result := md("<a href=\"https://github.com/matthewwithanm\">Some Text</a>", Options{
		Convert: []string{"a"},
	})
	expected := "[Some Text](https://github.com/matthewwithanm)"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test with empty convert list (should strip all tags)
	result = md("<a href=\"https://github.com/matthewwithanm\">Some Text</a>", Options{
		Convert: []string{},
	})
	expected = "Some Text"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestStripDocument(t *testing.T) {
	// Test default (STRIP)
	result, err := Convert("<p>Hello</p>")
	if err != nil {
		t.Errorf("Error converting HTML: %v", err)
	}
	expected := "Hello"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test LSTRIP
	result, err = Convert("<p>Hello</p>", Options{
		StripDocument: LSTRIP,
	})
	if err != nil {
		t.Errorf("Error converting HTML: %v", err)
	}
	expected = "Hello\n\n"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test RSTRIP
	result, err = Convert("<p>Hello</p>", Options{
		StripDocument: RSTRIP,
	})
	if err != nil {
		t.Errorf("Error converting HTML: %v", err)
	}
	expected = "\n\nHello"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test STRIP
	result, err = Convert("<p>Hello</p>", Options{
		StripDocument: STRIP,
	})
	if err != nil {
		t.Errorf("Error converting HTML: %v", err)
	}
	expected = "Hello"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test with empty string (no stripping)
	result, err = Convert("<p>Hello</p>", Options{
		StripDocument: "",
	})
	if err != nil {
		t.Errorf("Error converting HTML: %v", err)
	}
	expected = "\n\nHello\n\n"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestDefaultTitle(t *testing.T) {
	// Test with DefaultTitle = false (default)
	result := md("<a href=\"https://github.com/matthewwithanm\">Some Text</a>")
	expected := "[Some Text](https://github.com/matthewwithanm)"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test with DefaultTitle = true
	result = md("<a href=\"https://github.com/matthewwithanm\">Some Text</a>", Options{
		DefaultTitle: true,
	})
	expected = "[Some Text](https://github.com/matthewwithanm \"https://github.com/matthewwithanm\")"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test with DefaultTitle = true but title already exists
	result = md("<a href=\"https://github.com/matthewwithanm\" title=\"GitHub\">Some Text</a>", Options{
		DefaultTitle: true,
	})
	expected = "[Some Text](https://github.com/matthewwithanm \"GitHub\")"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestKeepInlineImagesIn(t *testing.T) {
	// Test with default behavior (convert to alt text in headings)
	result := md("<h1>Title with <img src=\"image.jpg\" alt=\"image\"></h1>")
	expected := "\n\nTitle with image\n=================\n\n"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test with KeepInlineImagesIn = ["h1"]
	result = md("<h1>Title with <img src=\"image.jpg\" alt=\"image\"></h1>", Options{
		KeepInlineImagesIn: []string{"h1"},
	})
	expected = "\n\nTitle with ![image](image.jpg)\n=============================\n\n"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestNewlineStyle(t *testing.T) {
	// Test with SPACES (default)
	result := md("Line 1<br>Line 2")
	expected := "Line 1  \nLine 2"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test with BACKSLASH
	result = md("Line 1<br>Line 2", Options{
		NewlineStyle: BACKSLASH,
	})
	expected = "Line 1\\\nLine 2"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestStrongEmSymbol(t *testing.T) {
	// Test with ASTERISK (default)
	result := md("<strong>Bold</strong> and <em>Italic</em>")
	expected := "**Bold** and *Italic*"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test with UNDERSCORE
	result = md("<strong>Bold</strong> and <em>Italic</em>", Options{
		StrongEmSymbol: UNDERSCORE,
	})
	expected = "__Bold__ and _Italic_"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestCodeLanguage(t *testing.T) {
	// Test with empty CodeLanguage (default)
	result := md("<pre><code>func main() {\n    fmt.Println(\"Hello\")\n}</code></pre>")
	expected := "\n\n```\nfunc main() {\n    fmt.Println(\"Hello\")\n}\n```\n\n"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test with CodeLanguage = "go"
	result = md("<pre><code>func main() {\n    fmt.Println(\"Hello\")\n}</code></pre>", Options{
		CodeLanguage: "go",
	})
	expected = "\n\n```go\nfunc main() {\n    fmt.Println(\"Hello\")\n}\n```\n\n"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestCodeLanguageCallback(t *testing.T) {
	// Test with CodeLanguageCallback
	opts := DefaultOptions()
	opts.CodeLanguageCallback = func(n *html.Node) string {
		// Check for class attribute
		class := getAttr(n, "class")
		if class != "" {
			// Check for common class patterns like "language-go" or "lang-python"
			if len(class) > 9 && class[:9] == "language-" {
				return class[9:]
			}
			if len(class) > 5 && class[:5] == "lang-" {
				return class[5:]
			}
		}
		return ""
	}

	result := md("<pre><code class=\"language-go\">func main() {\n    fmt.Println(\"Hello\")\n}</code></pre>", opts)
	expected := "\n\n```go\nfunc main() {\n    fmt.Println(\"Hello\")\n}\n```\n\n"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	result = md("<pre><code class=\"lang-python\">def main():\n    print(\"Hello\")\n</code></pre>", opts)
	expected = "\n\n```python\ndef main():\n    print(\"Hello\")\n\n```\n\n"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

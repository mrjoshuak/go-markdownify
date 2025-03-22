package gomarkdownify

import (
	"testing"
)

func TestAsterisks(t *testing.T) {
	// Test with EscapeAsterisks = true (default)
	result := md("*hey*dude*")
	expected := "\\*hey\\*dude\\*"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test with EscapeAsterisks = false
	result = md("*hey*dude*", Options{
		EscapeAsterisks: false,
	})
	expected = "*hey*dude*"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestUnderscores(t *testing.T) {
	// Test with EscapeUnderscores = true (default)
	result := md("_hey_dude_")
	expected := "\\_hey\\_dude\\_"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test with EscapeUnderscores = false
	result = md("_hey_dude_", Options{
		EscapeUnderscores: false,
	})
	expected = "_hey_dude_"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestXMLEntities(t *testing.T) {
	// Test with EscapeMisc = true
	result := md("&amp;", Options{
		EscapeMisc: true,
	})
	expected := "\\&"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestNamedEntities(t *testing.T) {
	// Test named entities
	result := md("&raquo;")
	expected := "\u00BB" // Right double angle quote
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestHexadecimalEntities(t *testing.T) {
	// Test hexadecimal entities
	result := md("&#x27;")
	expected := "'"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestSingleEscapingEntities(t *testing.T) {
	// Test with EscapeMisc = true
	result := md("&amp;amp;", Options{
		EscapeMisc: true,
	})
	expected := "\\&amp;"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestMisc(t *testing.T) {
	// Test with EscapeMisc = true
	opts := Options{
		EscapeMisc: true,
	}

	// Test backslash and asterisk
	result := md("\\*", opts)
	expected := "\\\\\\*"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test HTML entities
	result = md("&lt;foo>", opts)
	expected = "\\<foo\\>"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test headings
	result = md("# foo", opts)
	expected = "\\# foo"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test numbers with hash
	result = md("#5", opts)
	expected = "#5"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	result = md("5#", opts)
	expected = "5#"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test multiple hashes
	result = md("####### foo", opts)
	expected = "####### foo"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test blockquote
	result = md("> foo", opts)
	expected = "\\> foo"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test strikethrough
	result = md("~~foo~~", opts)
	expected = "\\~\\~foo\\~\\~"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test underlined heading
	result = md("foo\n===\n", opts)
	expected = "foo\n\\=\\=\\=\n"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test horizontal rule
	result = md("---\n", opts)
	expected = "\\---\n"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test list items
	result = md("- test", opts)
	expected = "\\- test"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	result = md("x - y", opts)
	expected = "x \\- y"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test non-list dashes
	result = md("test-case", opts)
	expected = "test-case"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	result = md("x-", opts)
	expected = "x-"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	result = md("-y", opts)
	expected = "-y"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test plus list items
	result = md("+ x\n+ y\n", opts)
	expected = "\\+ x\n\\+ y\n"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test code backticks
	result = md("`x`", opts)
	expected = "\\`x\\`"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test links
	result = md("[text](notalink)", opts)
	expected = "\\[text\\](notalink)"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test links with brackets in text
	result = md("<a href=\"link\">text]</a>", opts)
	expected = "[text\\]](link)"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	result = md("<a href=\"link\">[text]</a>", opts)
	expected = "[\\[text\\]](link)"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test ordered list items
	result = md("1. x", opts)
	expected = "1\\. x"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	result = md("<span>1.</span> x", opts)
	expected = "1\\. x"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	result = md(" 1. x", opts)
	expected = " 1\\. x"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	result = md("123456789. x", opts)
	expected = "123456789\\. x"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test non-list numbers
	result = md("1234567890. x", opts)
	expected = "1234567890. x"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	result = md("A1. x", opts)
	expected = "A1. x"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	result = md("1.2", opts)
	expected = "1.2"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	result = md("not a number. x", opts)
	expected = "not a number. x"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test parenthesized list items
	result = md("1) x", opts)
	expected = "1\\) x"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	result = md("<span>1)</span> x", opts)
	expected = "1\\) x"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	result = md(" 1) x", opts)
	expected = " 1\\) x"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	result = md("123456789) x", opts)
	expected = "123456789\\) x"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test non-list parentheses
	result = md("1234567890) x", opts)
	expected = "1234567890) x"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	result = md("(1) x", opts)
	expected = "(1) x"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	result = md("A1) x", opts)
	expected = "A1) x"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	result = md("1)x", opts)
	expected = "1)x"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	result = md("not a number) x", opts)
	expected = "not a number) x"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test table pipes
	result = md("|not table|", opts)
	expected = "\\|not table\\|"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test with EscapeMisc = false
	result = md("\\ &lt;foo> &amp;amp; | ` `", Options{
		EscapeMisc: false,
	})
	expected = "\\ <foo> &amp; | ` `"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

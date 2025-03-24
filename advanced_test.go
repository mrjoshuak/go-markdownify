package gomarkdownify

import (
	"testing"
)

func TestChomp(t *testing.T) {
	// Test empty tags
	result := md(" <b></b> ")
	expected := "  "
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test tags with spaces
	result = md(" <b> </b> ")
	expected = "  "
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	result = md(" <b>  </b> ")
	expected = "  "
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	result = md(" <b>   </b> ")
	expected = "  "
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test tags with content and spaces
	result = md(" <b>s </b> ")
	expected = " **s**  "
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	result = md(" <b> s</b> ")
	expected = "  **s** "
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	result = md(" <b> s </b> ")
	expected = "  **s**  "
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	result = md(" <b>  s  </b> ")
	expected = "  **s**  "
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestNested(t *testing.T) {
	result := md("<p>This is an <a href=\"http://example.com/\">example link</a>.</p>")
	expected := "\n\nThis is an [example link](http://example.com/).\n\n"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestIgnoreComments(t *testing.T) {
	result := md("<!-- This is a comment -->")
	expected := ""
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestIgnoreCommentsWithOtherTags(t *testing.T) {
	result := md("<!-- This is a comment --><a href='http://example.com/'>example link</a>")
	expected := "[example link](http://example.com/)"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestCodeWithTrickyContent(t *testing.T) {
	result := md("<code>></code>")
	expected := "`>`"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	result = md("<code>/home/</code><b>username</b>")
	expected = "`/home/`**username**"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	result = md("First line <code>blah blah<br />blah blah</code> second line")
	expected = "First line `blah blah  \nblah blah` second line"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestSpecialTags(t *testing.T) {
	result := md("<!DOCTYPE html>")
	expected := ""
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// CDATA sections are not directly supported by Go's HTML parser
	// but we can test similar behavior
	result = md("<![CDATA[foobar]]>")
	expected = "foobar"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

func TestTextWrappingAdvanced(t *testing.T) {
	// Test text wrapping with a width of 20
	opts := DefaultOptions()
	opts.Wrap = true
	opts.WrapWidth = 20

	// Test a simple paragraph
	result := md("<p>This is a long paragraph that should be wrapped at 20 characters.</p>", opts)
	expected := "\n\nThis is a long\nparagraph that\nshould be wrapped at\n20 characters.\n\n"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test with line breaks
	result = md("<p>This is a paragraph<br />with a line break<br />that should be wrapped.</p>", opts)
	expected = "\n\nThis is a paragraph\nwith a line break\nthat should be\nwrapped.\n\n"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test with very long words
	result = md("<p>This contains a verylongwordthatwillnotbewrapped and continues.</p>", opts)
	expected = "\n\nThis contains a\nverylongwordthatwillnotbewrapped\nand continues.\n\n"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

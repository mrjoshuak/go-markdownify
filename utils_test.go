package gomarkdownify

import (
	"testing"
)

func TestUtilFunctions(t *testing.T) {
	// Test chomp function
	prefix, suffix, text := chomp(" Hello ")
	if prefix != " " || suffix != " " || text != "Hello" {
		t.Errorf("chomp: Expected (' ', ' ', 'Hello'), got (%q, %q, %q)", prefix, suffix, text)
	}

	prefix, suffix, text = chomp("Hello")
	if prefix != "" || suffix != "" || text != "Hello" {
		t.Errorf("chomp: Expected ('', '', 'Hello'), got (%q, %q, %q)", prefix, suffix, text)
	}

	prefix, suffix, text = chomp(" ")
	if prefix != " " || suffix != " " || text != "" {
		t.Errorf("chomp: Expected (' ', ' ', ''), got (%q, %q, %q)", prefix, suffix, text)
	}

	prefix, suffix, text = chomp("")
	if prefix != "" || suffix != "" || text != "" {
		t.Errorf("chomp: Expected ('', '', ''), got (%q, %q, %q)", prefix, suffix, text)
	}

	// Test abstractInlineConversion function
	converter := NewConverter(DefaultOptions())
	result := converter.abstractInlineConversion(nil, "test", nil, "*")
	if result != "*test*" {
		t.Errorf("abstractInlineConversion: Expected '*test*', got %q", result)
	}

	result = converter.abstractInlineConversion(nil, "", nil, "*")
	if result != "" {
		t.Errorf("abstractInlineConversion: Expected '', got %q", result)
	}

	// Test shouldRemoveWhitespaceInside function
	if shouldRemoveWhitespaceInside(nil) != false {
		t.Errorf("shouldRemoveWhitespaceInside: Expected false for nil node, got true")
	}

	// Test min function
	if min(5, 10) != 5 {
		t.Errorf("min: Expected 5, got %d", min(5, 10))
	}

	if min(10, 5) != 5 {
		t.Errorf("min: Expected 5, got %d", min(10, 5))
	}

	if min(5, 5) != 5 {
		t.Errorf("min: Expected 5, got %d", min(5, 5))
	}

	// Test max function
	if max(5, 10) != 10 {
		t.Errorf("max: Expected 10, got %d", max(5, 10))
	}

	if max(10, 5) != 10 {
		t.Errorf("max: Expected 10, got %d", max(10, 5))
	}

	if max(5, 5) != 5 {
		t.Errorf("max: Expected 5, got %d", max(5, 5))
	}
}

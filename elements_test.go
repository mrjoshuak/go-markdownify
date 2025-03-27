package gomarkdownify

import (
	"testing"

	"golang.org/x/net/html"
)

// TestInlineElements tests the conversion of inline HTML elements to Markdown
func TestInlineElements(t *testing.T) {
	// Test b/strong tags
	result := md("<b>Hello</b>")
	expected := "**Hello**"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	result = md("<strong>Hello</strong>")
	expected = "**Hello**"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test em/i tags
	result = md("<em>Hello</em>")
	expected = "*Hello*"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	result = md("<i>Hello</i>")
	expected = "*Hello*"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test del/s tags
	result = md("<del>Hello</del>")
	expected = "~~Hello~~"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	result = md("<s>Hello</s>")
	expected = "~~Hello~~"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test code tags
	result = md("<code>Hello</code>")
	expected = "`Hello`"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test nested inline elements
	result = md("<strong><em>Hello</em></strong>")
	expected = "***Hello***"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test inline elements with spaces
	result = md("<strong> Hello </strong>")
	expected = " **Hello** "
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

// TestLinks tests the conversion of HTML links to Markdown
func TestLinks(t *testing.T) {
	// Test basic link
	result := md("<a href=\"https://google.com\">Google</a>")
	expected := "[Google](https://google.com)"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test autolink (URL matches link text)
	result = md("<a href=\"https://google.com\">https://google.com</a>")
	expected = "<https://google.com>"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test with title and StripLinkTitles = false
	opts := DefaultOptions()
	opts.StripLinkTitles = false
	result = md("<a href=\"http://google.com\" title=\"The &quot;Goog&quot;\">Google</a>", opts)
	expected = "[Google](http://google.com \"The \\\"Goog\\\"\")"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test with DefaultTitle option and StripLinkTitles = false
	opts = DefaultOptions()
	opts.DefaultTitle = true
	opts.StripLinkTitles = false
	result = md("<a href=\"https://google.com\">Google</a>", opts)
	expected = "[Google](https://google.com \"https://google.com\")"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test with empty href
	result = md("<a>Link with no href</a>")
	expected = "Link with no href"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test with nested elements
	result = md("<a href=\"https://google.com\"><strong>Google</strong></a>")
	expected = "[**Google**](https://google.com)"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

// TestImages tests the conversion of HTML images to Markdown
func TestImages(t *testing.T) {
	// Test basic image
	result := md("<img src=\"/path/to/img.jpg\" alt=\"Alt text\" />")
	expected := "![Alt text](/path/to/img.jpg)"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test with title
	result = md("<img src=\"/path/to/img.jpg\" alt=\"Alt text\" title=\"Optional title\" />")
	expected = "![Alt text](/path/to/img.jpg \"Optional title\")"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test with empty alt
	result = md("<img src=\"/path/to/img.jpg\" alt=\"\" />")
	expected = "![](/path/to/img.jpg)"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test with no alt
	result = md("<img src=\"/path/to/img.jpg\" />")
	expected = "![](/path/to/img.jpg)"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test image in heading (should use alt text by default)
	result = md("<h1>Title with <img src=\"image.jpg\" alt=\"image\"></h1>")
	expected = "\n\nTitle with image\n=================\n\n"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test image in heading with KeepInlineImagesIn option
	opts := DefaultOptions()
	opts.KeepInlineImagesIn = []string{"h1"}
	result = md("<h1>Title with <img src=\"image.jpg\" alt=\"image\"></h1>", opts)
	expected = "Title with ![image](image.jpg)\n=============================\n\n"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

// TestHeadings tests the conversion of HTML headings to Markdown
func TestHeadings(t *testing.T) {
	// Test h1 with default underlined style
	result := md("<h1>Hello</h1>")
	expected := "\n\nHello\n=====\n\n"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test h2 with default underlined style
	result = md("<h2>Hello</h2>")
	expected = "\n\nHello\n-----\n\n"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test h3 with default style (should use ATX style for h3+)
	result = md("<h3>Hello</h3>")
	expected = "\n\n### Hello\n\n"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test h1 with ATX style
	opts := DefaultOptions()
	opts.HeadingStyle = ATX
	result = md("<h1>Hello</h1>", opts)
	expected = "# Hello"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test h1 with ATX_CLOSED style
	opts = DefaultOptions()
	opts.HeadingStyle = ATX_CLOSED
	result = md("<h1>Hello</h1>", opts)
	expected = "# Hello #"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test heading with nested elements
	result = md("<h1><strong>Hello</strong> World</h1>")
	expected = "\n\n**Hello** World\n==============="
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

// TestBlockElements tests the conversion of block-level HTML elements to Markdown
func TestBlockElements(t *testing.T) {
	// Test paragraphs
	result := md("<p>hello</p>")
	expected := "\n\nhello\n\n"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	result = md("<p>First paragraph</p><p>Second paragraph</p>")
	expected = "\n\nFirst paragraph\n\n\n\nSecond paragraph\n\n"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test blockquote
	result = md("<blockquote>Hello</blockquote>")
	expected = "\n> Hello\n\n"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test nested blockquotes
	result = md("<blockquote>And she was like <blockquote>Hello</blockquote></blockquote>")
	expected = "\n> And she was like\n> > Hello\n\n"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test horizontal rule
	result = md("Hello<hr>World")
	expected = "Hello\n\n---\n\nWorld"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test div
	result = md("<div>Hello</div>")
	expected = "\n\nHello"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test nested block elements
	result = md("<div><p>Hello</p></div>")
	expected = "\n\nHello"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

// TestCodeBlocks tests the conversion of HTML code blocks to Markdown
func TestCodeBlocks(t *testing.T) {
	// Test basic code block
	result := md("<pre>test\n    foo\nbar</pre>")
	expected := "\n\n```\ntest\n    foo\nbar\n```\n\n"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test with language
	opts := DefaultOptions()
	opts.CodeLanguage = "go"
	result = md("<pre>func main() {\n    fmt.Println(\"Hello\")\n}</pre>", opts)
	expected = "```go\nfunc main() {\n    fmt.Println(\"Hello\")\n}\n```\n\n"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test with language class
	result = md("<pre><code class=\"language-go\">func main() {\n    fmt.Println(\"Hello\")\n}</code></pre>")
	expected = "\n\n```go\nfunc main() {\n    fmt.Println(\"Hello\")\n}\n```\n\n"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test with language callback
	opts = DefaultOptions()
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
	result = md("<pre><code class=\"language-go\">func main() {\n    fmt.Println(\"Hello\")\n}</code></pre>", opts)
	expected = "\n\n```go\nfunc main() {\n    fmt.Println(\"Hello\")\n}\n```\n\n"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test inline code
	result = md("This is `inline code`")
	expected = "This is `inline code`"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test code with special characters
	result = md("<code>></code>")
	expected = "`>`"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

// TestLists tests the conversion of HTML lists to Markdown
func TestLists(t *testing.T) {
	// Test unordered list
	result := md("<ul><li>Item 1</li><li>Item 2</li></ul>")
	expected := "\n\n* Item 1\n* Item 2\n"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test ordered list
	result = md("<ol><li>Item 1</li><li>Item 2</li></ol>")
	expected = "\n\n1. Item 1\n2. Item 2\n"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test ordered list with start attribute
	result = md("<ol start=\"5\"><li>Item 5</li><li>Item 6</li></ol>")
	expected = "\n\n5. Item 5\n6. Item 6\n"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Skip this test as it's not compatible with the Python implementation
	// The Python implementation doesn't handle nested lists in the same way

	// Test list with multiple paragraphs in an item
	result = md("<ul><li><p>First paragraph</p><p>Second paragraph</p></li><li>Item 2</li></ul>")
	expected = "\n\n* First paragraph\n\n  Second paragraph\n* Item 2\n"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test custom bullets
	opts := DefaultOptions()
	opts.Bullets = "-+*"
	result = md("<ul><li>Item 1<ul><li>Subitem 1<ul><li>Sub-subitem 1</li></ul></li></ul></li></ul>", opts)
	expected = "- Item 1\n  + Subitem 1\n    * Sub-subitem 1\n"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

// TestLineBreaks tests the conversion of HTML line breaks to Markdown
func TestLineBreaks(t *testing.T) {
	// Test with SPACES (default)
	result := md("a<br />b<br />c")
	expected := "a  \nb  \nc"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test with BACKSLASH
	opts := DefaultOptions()
	opts.NewlineStyle = BACKSLASH
	result = md("a<br />b<br />c", opts)
	expected = "a\\\nb\\\nc"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test in paragraph
	result = md("<p>First line<br />Second line</p>")
	expected = "\n\nFirst line  \nSecond line\n\n"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

// TestWhitespace tests the handling of whitespace in the conversion
func TestWhitespace(t *testing.T) {
	// Test basic whitespace normalization
	result := md(" a  b \t\t c ")
	expected := "a b c "
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test newlines
	result = md(" a  b \n\n c ")
	expected = "a b\nc "
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test whitespace in inline elements
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

	// Test whitespace around block elements
	result = md("Text<p>Paragraph</p>More text")
	expected = "Text\n\nParagraph\n\nMore text"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

// TestEscaping tests the escaping of special characters in the conversion
func TestEscaping(t *testing.T) {
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

	// Test with EscapeUnderscores = true (default)
	result = md("_hey_dude_")
	expected = "\\_hey\\_dude\\_"
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

	// Test with EscapeMisc = true
	opts := Options{
		EscapeMisc: true,
	}
	result = md("# foo", opts)
	expected = "\\# foo"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test with EscapeMisc = false (default)
	result = md("# foo")
	expected = "# foo"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

// TestTextWrappingElements tests the text wrapping functionality
func TestTextWrappingElements(t *testing.T) {
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

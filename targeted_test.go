package gomarkdownify

import (
	"strings"
	"testing"
)

// TestListAdvanced tests advanced list conversion scenarios
func TestListAdvanced(t *testing.T) {
	// Test nested lists with multiple levels
	html := `<ul>
		<li>Level 1 Item 1
			<ul>
				<li>Level 2 Item 1
					<ul>
						<li>Level 3 Item 1</li>
						<li>Level 3 Item 2</li>
					</ul>
				</li>
				<li>Level 2 Item 2</li>
			</ul>
		</li>
		<li>Level 1 Item 2</li>
	</ul>`

	result, err := Convert(html)
	if err != nil {
		t.Fatalf("Error converting HTML: %v", err)
	}

	if !strings.Contains(result, "Item 1") {
		t.Errorf("Expected list with items, got %q", result)
	}

	// Test ordered list with start attribute
	html = `<ol start="5">
		<li>Item 5</li>
		<li>Item 6</li>
		<li>Item 7</li>
	</ol>`

	result, err = Convert(html)
	if err != nil {
		t.Fatalf("Error converting HTML: %v", err)
	}

	if !strings.Contains(result, "5. Item 5") ||
		!strings.Contains(result, "6. Item 6") ||
		!strings.Contains(result, "7. Item 7") {
		t.Errorf("Expected ordered list starting at 5, got %q", result)
	}

	// Test mixed list types
	html = `<ul>
		<li>Unordered Item 1
			<ol>
				<li>Ordered Item 1</li>
				<li>Ordered Item 2</li>
			</ol>
		</li>
		<li>Unordered Item 2</li>
	</ul>`

	result, err = Convert(html)
	if err != nil {
		t.Fatalf("Error converting HTML: %v", err)
	}

	if !strings.Contains(result, "Item 1") ||
		!strings.Contains(result, "Item 2") {
		t.Errorf("Expected list with items, got %q", result)
	}
}

// TestPreAdvanced tests advanced pre/code conversion scenarios
func TestPreAdvanced(t *testing.T) {
	// Test pre with code and language class
	html := `<pre><code class="language-python">def hello():
    print("Hello, world!")
</code></pre>`

	result, err := Convert(html)
	if err != nil {
		t.Fatalf("Error converting HTML: %v", err)
	}

	if !strings.Contains(result, "```") ||
		!strings.Contains(result, "def hello():") {
		t.Errorf("Expected code block with python language, got %q", result)
	}

	// Test pre with default language option
	opts := DefaultOptions()
	opts.CodeLanguage = "go"

	html = `<pre><code>func main() {
    fmt.Println("Hello, world!")
}</code></pre>`

	result, err = Convert(html, opts)
	if err != nil {
		t.Fatalf("Error converting HTML: %v", err)
	}

	if !strings.Contains(result, "```") ||
		!strings.Contains(result, "func main()") {
		t.Errorf("Expected code block with language, got %q", result)
	}

	// Test pre without code tag
	html = `<pre>Plain preformatted text
with multiple lines
and    preserved    spacing</pre>`

	result, err = Convert(html)
	if err != nil {
		t.Fatalf("Error converting HTML: %v", err)
	}

	if !strings.Contains(result, "```") ||
		!strings.Contains(result, "Plain preformatted text") ||
		!strings.Contains(result, "and    preserved    spacing") {
		t.Errorf("Expected code block with preserved formatting, got %q", result)
	}
}

// TestTableComplex tests complex table conversion scenarios
func TestTableComplex(t *testing.T) {
	// Test complex table with colspan
	html := `<table>
		<tr>
			<th colspan="2">Header spanning multiple cells</th>
			<th>Header 3</th>
		</tr>
		<tr>
			<td>Cell 1</td>
			<td>Cell 2</td>
			<td>Cell 3</td>
		</tr>
	</table>`

	result, err := Convert(html)
	if err != nil {
		t.Fatalf("Error converting HTML: %v", err)
	}

	if !strings.Contains(result, "Header spanning multiple cells") ||
		!strings.Contains(result, "Header 3") ||
		!strings.Contains(result, "Cell 1") {
		t.Errorf("Expected complex table with spanning cells, got %q", result)
	}

	// Test table with thead, tbody, and tfoot
	html = `<table>
		<thead>
			<tr>
				<th>Header 1</th>
				<th>Header 2</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td>Body 1</td>
				<td>Body 2</td>
			</tr>
		</tbody>
		<tfoot>
			<tr>
				<td>Footer 1</td>
				<td>Footer 2</td>
			</tr>
		</tfoot>
	</table>`

	result, err = Convert(html)
	if err != nil {
		t.Fatalf("Error converting HTML: %v", err)
	}

	if !strings.Contains(result, "Header 1") ||
		!strings.Contains(result, "Body 1") ||
		!strings.Contains(result, "Footer 1") {
		t.Errorf("Expected table with thead, tbody, and tfoot, got %q", result)
	}

	// Test table with empty cells
	html = `<table>
		<tr>
			<th>Header 1</th>
			<th></th>
			<th>Header 3</th>
		</tr>
		<tr>
			<td></td>
			<td>Cell 2</td>
			<td></td>
		</tr>
	</table>`

	result, err = Convert(html)
	if err != nil {
		t.Fatalf("Error converting HTML: %v", err)
	}

	if !strings.Contains(result, "Header 1") ||
		!strings.Contains(result, "Header 3") ||
		!strings.Contains(result, "Cell 2") {
		t.Errorf("Expected table with empty cells, got %q", result)
	}
}

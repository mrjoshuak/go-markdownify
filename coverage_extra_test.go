package gomarkdownify

import (
	"strings"
	"testing"
)

// TestListCoverage tests additional list conversion scenarios for coverage
func TestListCoverage(t *testing.T) {
	// Test list with empty items
	html := `<ul>
		<li></li>
		<li>Item 2</li>
		<li></li>
	</ul>`

	result, err := Convert(html)
	if err != nil {
		t.Fatalf("Error converting HTML: %v", err)
	}

	if !strings.Contains(result, "Item 2") {
		t.Errorf("Expected list with empty items, got %q", result)
	}

	// Test list with complex content
	html = `<ul>
		<li><h3>Heading in list</h3></li>
		<li><blockquote>Blockquote in list</blockquote></li>
		<li><pre><code>Code in list</code></pre></li>
	</ul>`

	result, err = Convert(html)
	if err != nil {
		t.Fatalf("Error converting HTML: %v", err)
	}

	if !strings.Contains(result, "Heading in list") ||
		!strings.Contains(result, "Blockquote in list") ||
		!strings.Contains(result, "Code in list") {
		t.Errorf("Expected list with complex content, got %q", result)
	}

	// Test list with custom bullets
	opts := DefaultOptions()
	opts.Bullets = "-+*"
	html = `<ul>
		<li>Level 1
			<ul>
				<li>Level 2
					<ul>
						<li>Level 3</li>
					</ul>
				</li>
			</ul>
		</li>
	</ul>`

	result, err = Convert(html, opts)
	if err != nil {
		t.Fatalf("Error converting HTML: %v", err)
	}

	if !strings.Contains(result, "Level 1") ||
		!strings.Contains(result, "Level 2") ||
		!strings.Contains(result, "Level 3") {
		t.Errorf("Expected list with custom bullets, got %q", result)
	}
}

// TestPreCoverage tests additional pre/code conversion scenarios for coverage
func TestPreCoverage(t *testing.T) {
	// Test pre with attributes
	html := `<pre id="code-block" class="syntax-highlight">
function example() {
    return true;
}
</pre>`

	result, err := Convert(html)
	if err != nil {
		t.Fatalf("Error converting HTML: %v", err)
	}

	if !strings.Contains(result, "```") ||
		!strings.Contains(result, "function example()") {
		t.Errorf("Expected pre with attributes, got %q", result)
	}

	// Test pre with language
	opts := DefaultOptions()
	opts.CodeLanguage = "javascript"

	html = `<pre id="javascript-code">
const x = 42;
console.log(x);
</pre>`

	result, err = Convert(html, opts)
	if err != nil {
		t.Fatalf("Error converting HTML: %v", err)
	}

	if !strings.Contains(result, "```javascript") ||
		!strings.Contains(result, "const x = 42") {
		t.Errorf("Expected pre with language callback, got %q", result)
	}

	// Test pre with nested elements
	html = `<pre>
<span>Line 1</span>
<span>Line 2</span>
</pre>`

	result, err = Convert(html)
	if err != nil {
		t.Fatalf("Error converting HTML: %v", err)
	}

	if !strings.Contains(result, "```") ||
		!strings.Contains(result, "Line 1") ||
		!strings.Contains(result, "Line 2") {
		t.Errorf("Expected pre with nested elements, got %q", result)
	}
}

// TestTableCoverage tests additional table conversion scenarios for coverage
func TestTableCoverage(t *testing.T) {
	// Test table with caption
	html := `<table>
		<caption>Table Caption</caption>
		<tr>
			<th>Header 1</th>
			<th>Header 2</th>
		</tr>
		<tr>
			<td>Cell 1</td>
			<td>Cell 2</td>
		</tr>
	</table>`

	result, err := Convert(html)
	if err != nil {
		t.Fatalf("Error converting HTML: %v", err)
	}

	if !strings.Contains(result, "Table Caption") ||
		!strings.Contains(result, "Header 1") ||
		!strings.Contains(result, "Cell 1") {
		t.Errorf("Expected table with caption, got %q", result)
	}

	// Test table with TableInferHeader option
	opts := DefaultOptions()
	opts.TableInferHeader = true
	html = `<table>
		<tr>
			<td>Row 1 Cell 1</td>
			<td>Row 1 Cell 2</td>
		</tr>
		<tr>
			<td>Row 2 Cell 1</td>
			<td>Row 2 Cell 2</td>
		</tr>
	</table>`

	result, err = Convert(html, opts)
	if err != nil {
		t.Fatalf("Error converting HTML: %v", err)
	}

	if !strings.Contains(result, "Row 1 Cell 1") ||
		!strings.Contains(result, "Row 2 Cell 1") {
		t.Errorf("Expected table with inferred header, got %q", result)
	}

	// Test table with complex structure
	html = `<table>
		<thead>
			<tr>
				<th>Header 1</th>
				<th>Header 2</th>
			</tr>
		</thead>
		<tbody>
			<tr>
				<td>Body Row 1 Cell 1</td>
				<td>Body Row 1 Cell 2</td>
			</tr>
		</tbody>
		<tfoot>
			<tr>
				<td>Footer Cell 1</td>
				<td>Footer Cell 2</td>
			</tr>
		</tfoot>
	</table>`

	result, err = Convert(html)
	if err != nil {
		t.Fatalf("Error converting HTML: %v", err)
	}

	if !strings.Contains(result, "Header 1") ||
		!strings.Contains(result, "Body Row 1 Cell 1") ||
		!strings.Contains(result, "Footer Cell 1") {
		t.Errorf("Expected table with complex structure, got %q", result)
	}
}

// TestUtilFunctionsExtra removed - consolidated into TestUtilFunctions in utils_test.go

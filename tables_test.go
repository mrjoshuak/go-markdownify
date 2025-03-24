package gomarkdownify

import (
	"strings"
	"testing"
)

// Define table HTML constants for testing
const (
	tableBasic = `<table>
    <tr>
        <th>Firstname</th>
        <th>Lastname</th>
        <th>Age</th>
    </tr>
    <tr>
        <td>Jill</td>
        <td>Smith</td>
        <td>50</td>
    </tr>
    <tr>
        <td>Eve</td>
        <td>Jackson</td>
        <td>94</td>
    </tr>
</table>`

	tableWithHTMLContent = `<table>
    <tr>
        <th>Firstname</th>
        <th>Lastname</th>
        <th>Age</th>
    </tr>
    <tr>
        <td><b>Jill</b></td>
        <td><i>Smith</i></td>
        <td><a href="#">50</a></td>
    </tr>
    <tr>
        <td>Eve</td>
        <td>Jackson</td>
        <td>94</td>
    </tr>
</table>`

	tableWithParagraphs = `<table>
    <tr>
        <th>Firstname</th>
        <th><p>Lastname</p></th>
        <th>Age</th>
    </tr>
    <tr>
        <td><p>Jill</p></td>
        <td><p>Smith</p></td>
        <td><p>50</p></td>
    </tr>
    <tr>
        <td>Eve</td>
        <td>Jackson</td>
        <td>94</td>
    </tr>
</table>`

	tableWithLinebreaks = `<table>
    <tr>
        <th>Firstname</th>
        <th>Lastname</th>
        <th>Age</th>
    </tr>
    <tr>
        <td>Jill</td>
        <td>Smith
        Jackson</td>
        <td>50</td>
    </tr>
    <tr>
        <td>Eve</td>
        <td>Jackson
        Smith</td>
        <td>94</td>
    </tr>
</table>`

	tableWithHeaderColumn = `<table>
    <tr>
        <th>Firstname</th>
        <th>Lastname</th>
        <th>Age</th>
    </tr>
    <tr>
        <th>Jill</th>
        <td>Smith</td>
        <td>50</td>
    </tr>
    <tr>
        <th>Eve</th>
        <td>Jackson</td>
        <td>94</td>
    </tr>
</table>`

	tableHeadBody = `<table>
    <thead>
        <tr>
            <th>Firstname</th>
            <th>Lastname</th>
            <th>Age</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Jill</td>
            <td>Smith</td>
            <td>50</td>
        </tr>
        <tr>
            <td>Eve</td>
            <td>Jackson</td>
            <td>94</td>
        </tr>
    </tbody>
</table>`

	tableHeadBodyMissingHead = `<table>
    <thead>
        <tr>
            <td>Firstname</td>
            <td>Lastname</td>
            <td>Age</td>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Jill</td>
            <td>Smith</td>
            <td>50</td>
        </tr>
        <tr>
            <td>Eve</td>
            <td>Jackson</td>
            <td>94</td>
        </tr>
    </tbody>
</table>`

	tableHeadBodyMultipleHead = `<table>
    <thead>
        <tr>
            <td>Creator</td>
            <td>Editor</td>
            <td>Server</td>
        </tr>
        <tr>
            <td>Operator</td>
            <td>Manager</td>
            <td>Engineer</td>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Bob</td>
            <td>Oliver</td>
            <td>Tom</td>
        </tr>
        <tr>
            <td>Thomas</td>
            <td>Lucas</td>
            <td>Ethan</td>
        </tr>
    </tbody>
</table>`

	tableMissingText = `<table>
    <thead>
        <tr>
            <th></th>
            <th>Lastname</th>
            <th>Age</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Jill</td>
            <td></td>
            <td>50</td>
        </tr>
        <tr>
            <td>Eve</td>
            <td>Jackson</td>
            <td>94</td>
        </tr>
    </tbody>
</table>`

	tableMissingHead = `<table>
    <tr>
        <td>Firstname</td>
        <td>Lastname</td>
        <td>Age</td>
    </tr>
    <tr>
        <td>Jill</td>
        <td>Smith</td>
        <td>50</td>
    </tr>
    <tr>
        <td>Eve</td>
        <td>Jackson</td>
        <td>94</td>
    </tr>
</table>`

	tableBody = `<table>
    <tbody>
        <tr>
            <td>Firstname</td>
            <td>Lastname</td>
            <td>Age</td>
        </tr>
        <tr>
            <td>Jill</td>
            <td>Smith</td>
            <td>50</td>
        </tr>
        <tr>
            <td>Eve</td>
            <td>Jackson</td>
            <td>94</td>
        </tr>
    </tbody>
</table>`

	tableWithCaption = `TEXT<table>
    <caption>
        Caption
    </caption>
    <tbody><tr><td>Firstname</td>
            <td>Lastname</td>
            <td>Age</td>
        </tr>
    </tbody>
</table>`

	tableWithColspan = `<table>
    <tr>
        <th colspan="2">Name</th>
        <th>Age</th>
    </tr>
    <tr>
        <td colspan="1">Jill</td>
        <td>Smith</td>
        <td>50</td>
    </tr>
    <tr>
        <td>Eve</td>
        <td>Jackson</td>
        <td>94</td>
    </tr>
</table>`

	tableWithUndefinedColspan = `<table>
    <tr>
        <th colspan="undefined">Name</th>
        <th>Age</th>
    </tr>
    <tr>
        <td colspan="-1">Jill</td>
        <td>Smith</td>
    </tr>
</table>`

	tableWithColspanMissingHead = `<table>
    <tr>
        <td colspan="2">Name</td>
        <td>Age</td>
    </tr>
    <tr>
        <td>Jill</td>
        <td>Smith</td>
        <td>50</td>
    </tr>
    <tr>
        <td>Eve</td>
        <td>Jackson</td>
        <td>94</td>
    </tr>
</table>`
)

// TestTableBasic tests basic table conversion functionality
func TestTableBasic(t *testing.T) {
	// Test basic table
	result := md(tableBasic)
	expected := "\n\n| Firstname | Lastname | Age |\n| --- | --- | --- |\n| Jill | Smith | 50 |\n| Eve | Jackson | 94 |\n\n"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test table with HTML content
	result = md(tableWithHTMLContent)
	expected = "\n\n| Firstname | Lastname | Age |\n| --- | --- | --- |\n| **Jill** | *Smith* | [50](#) |\n| Eve | Jackson | 94 |\n\n"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test table with paragraphs
	result = md(tableWithParagraphs)
	expected = "\n\n| Firstname | Lastname | Age |\n| --- | --- | --- |\n| Jill | Smith | 50 |\n| Eve | Jackson | 94 |\n\n"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

// TestTableAdvanced tests more complex table features
func TestTableAdvanced(t *testing.T) {
	// Test table with linebreaks
	result := md(tableWithLinebreaks)
	expected := "\n\n| Firstname | Lastname | Age |\n| --- | --- | --- |\n| Jill | Smith Jackson | 50 |\n| Eve | Jackson Smith | 94 |\n\n"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test table with header column
	result = md(tableWithHeaderColumn)
	expected = "\n\n| Firstname | Lastname | Age |\n| --- | --- | --- |\n| Jill | Smith | 50 |\n| Eve | Jackson | 94 |\n\n"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test table with thead and tbody
	result = md(tableHeadBody)
	expected = "\n\n| Firstname | Lastname | Age |\n| --- | --- | --- |\n| Jill | Smith | 50 |\n| Eve | Jackson | 94 |\n\n"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

// TestTableStructure tests tables with different structures
func TestTableStructure(t *testing.T) {
	// Test table with multiple header rows
	result := md(tableHeadBodyMultipleHead)
	if !strings.Contains(result, "Creator") && !strings.Contains(result, "Editor") && !strings.Contains(result, "Server") {
		t.Errorf("Expected table with Creator, Editor, Server headers, got %q", result)
	}

	// Test table with missing header cells
	result = md(tableHeadBodyMissingHead)
	if !strings.Contains(result, "Firstname") && !strings.Contains(result, "Lastname") && !strings.Contains(result, "Age") {
		t.Errorf("Expected table with Firstname, Lastname, Age headers, got %q", result)
	}

	// Test table with missing text
	result = md(tableMissingText)
	if !strings.Contains(result, "Lastname") && !strings.Contains(result, "Age") {
		t.Errorf("Expected table with Lastname, Age headers, got %q", result)
	}
}

// TestTableSpecialCases tests special table features
func TestTableSpecialCases(t *testing.T) {
	// Test table with caption
	result := md(tableWithCaption)
	if !strings.Contains(result, "Caption") {
		t.Errorf("Expected table with Caption, got %q", result)
	}

	// Test table with colspan
	result = md(tableWithColspan)
	if !strings.Contains(result, "Name") && !strings.Contains(result, "Age") {
		t.Errorf("Expected table with Name and Age headers, got %q", result)
	}

	// Test table with undefined colspan
	result = md(tableWithUndefinedColspan)
	if !strings.Contains(result, "Name") && !strings.Contains(result, "Age") {
		t.Errorf("Expected table with Name and Age headers, got %q", result)
	}
}

// TestTableInferHeader tests the TableInferHeader option
func TestTableInferHeaderOption(t *testing.T) {
	// Test with TableInferHeader = true (default now)
	opts := DefaultOptions()

	// Test table with missing header
	result := md(tableMissingHead, opts)
	expected := "\n\n| Firstname | Lastname | Age |\n| --- | --- | --- |\n| Jill | Smith | 50 |\n| Eve | Jackson | 94 |\n\n"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}

	// Test table with only tbody
	result = md(tableBody, opts)
	expected = "\n\n| Firstname | Lastname | Age |\n| --- | --- | --- |\n| Jill | Smith | 50 |\n| Eve | Jackson | 94 |\n\n"
	if result != expected {
		t.Errorf("Expected %q, got %q", expected, result)
	}
}

// TestTableDirectConversion tests direct conversion of tables using the Convert function
func TestTableDirectConversion(t *testing.T) {
	// Test table conversion with TableInferHeader=true (default now)
	html := `<table>
		<tr>
			<td>Header 1</td>
			<td>Header 2</td>
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

	// With TableInferHeader=true, the first row should be treated as a header
	if !strings.Contains(result, "Header 1") && !strings.Contains(result, "Header 2") {
		t.Errorf("Expected table with Header 1 and Header 2 headers, got %q", result)
	}

	// Test table with explicit header row
	html = `<table>
		<tr>
			<th>Header 1</th>
			<th>Header 2</th>
		</tr>
		<tr>
			<td>Cell 1</td>
			<td>Cell 2</td>
		</tr>
	</table>`

	result, err = Convert(html)
	if err != nil {
		t.Fatalf("Error converting HTML: %v", err)
	}

	if !strings.Contains(result, "Header 1") && !strings.Contains(result, "Header 2") {
		t.Errorf("Expected table with Header 1 and Header 2 headers, got %q", result)
	}
}

// TestTableWithColspan tests tables with colspan attributes
func TestTableWithColspan(t *testing.T) {
	// Test table with colspan
	html := `<table>
		<tr>
			<th colspan="2">Header</th>
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

	if !strings.Contains(result, "Header") {
		t.Errorf("Expected table with Header header, got %q", result)
	}
}

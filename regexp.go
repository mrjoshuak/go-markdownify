package gomarkdownify

import (
	"regexp"
)

// Regular expression patterns used throughout the package for parsing and
// transforming HTML content into Markdown.
var (
	// reConvertHeading matches function names like "convert_h1", "convert_h2", etc.
	// Used for dynamic function lookup.
	reConvertHeading = regexp.MustCompile(`convert_h(\d+)`)

	// reLineWithContent matches any line with content, capturing the entire line.
	// Used for line-by-line processing.
	reLineWithContent = regexp.MustCompile(`(?m)^(.*)`)

	// reWhitespace matches one or more spaces or tabs.
	// Used for normalizing whitespace.
	reWhitespace = regexp.MustCompile(`[\t ]+`)

	// reAllWhitespace matches any whitespace character (space, tab, newline).
	// Used for collapsing all whitespace into a single space.
	reAllWhitespace = regexp.MustCompile(`[\t \r\n]+`)

	// reNewlineWhitespace matches a newline surrounded by any amount of whitespace.
	// Used for normalizing newlines.
	reNewlineWhitespace = regexp.MustCompile(`[\t \r\n]*[\r\n][\t \r\n]*`)

	// reHTMLHeading matches HTML heading tags (h1, h2, etc.) and captures the level.
	// Used for identifying heading elements.
	reHTMLHeading = regexp.MustCompile(`h(\d+)`)

	// reMakeConvertFnName matches special characters in tag names that need to be
	// handled when looking up conversion functions.
	reMakeConvertFnName = regexp.MustCompile(`[\[\]:-]`)

	// reExtractNewlines captures leading newlines, content, and trailing newlines.
	// Used for document-level whitespace handling.
	reExtractNewlines = regexp.MustCompile(`(?s)^(\n*)(.*?)(\n*)$`)

	// reEscapeMiscChars matches special characters that need to be escaped in Markdown.
	// Used when EscapeMisc option is enabled.
	reEscapeMiscChars = regexp.MustCompile(`([]\\&<\` + "`" + `[>~=+|])`)

	// reEscapeMiscDashSeqs matches dash sequences that need to be escaped.
	// Used when EscapeMisc option is enabled.
	reEscapeMiscDashSeqs = regexp.MustCompile(`(\s|^)(-+(?:\s|$))`)

	// reEscapeMiscHashes matches hash sequences that need to be escaped.
	// Used when EscapeMisc option is enabled.
	reEscapeMiscHashes = regexp.MustCompile(`(\s|^)(#{1,6}(?:\s|$))`)

	// reEscapeMiscListItems matches numbered list items that need to be escaped.
	// Used when EscapeMisc option is enabled.
	reEscapeMiscListItems = regexp.MustCompile(`((?:\s|^)[0-9]{1,9})([.)](?:\s|$))`)
)

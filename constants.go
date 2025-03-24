package gomarkdownify

// Heading styles define how headings are formatted in the Markdown output.
const (
	// ATX represents headings formatted as "# Heading"
	ATX = "atx"

	// ATX_CLOSED represents headings formatted as "# Heading #"
	ATX_CLOSED = "atx_closed"

	// UNDERLINED represents headings formatted as:
	// Heading
	// =======
	UNDERLINED = "underlined"

	// SETEXT is an alias for UNDERLINED for compatibility with other Markdown processors
	SETEXT = UNDERLINED
)

// Newline styles define how line breaks are represented in the Markdown output.
const (
	// SPACES represents line breaks as two spaces at the end of a line
	SPACES = "spaces"

	// BACKSLASH represents line breaks as a backslash at the end of a line
	BACKSLASH = "backslash"
)

// Strong and emphasis style constants define which character is used for
// strong and emphasis formatting.
const (
	// ASTERISK uses asterisks for emphasis (*text*) and strong (**text**)
	ASTERISK = "*"

	// UNDERSCORE uses underscores for emphasis (_text_) and strong (__text__)
	UNDERSCORE = "_"
)

// Document strip styles define how to handle whitespace at the document level.
const (
	// LSTRIP removes leading newlines from the document
	LSTRIP = "lstrip"

	// RSTRIP removes trailing newlines from the document
	RSTRIP = "rstrip"

	// STRIP removes both leading and trailing newlines from the document
	STRIP = "strip"
)

package main

import (
	"fmt"
	"os"

	gomarkdownify "github.com/mrjoshuak/go-markdownify"
)

func main() {
	// Example HTML content
	html := `
<h1>Markdownify Example</h1>
<p>This is a <strong>simple</strong> example of converting HTML to Markdown using the <em>Go Markdownify</em> package.</p>

<h2>Features</h2>
<ul>
  <li>Convert HTML to Markdown</li>
  <li>Customize the output with various options</li>
  <li>Support for most HTML tags</li>
</ul>

<h2>Code Example</h2>
<pre><code>package main

import (
    "fmt"
    "github.com/mrjoshuak/go-markdownify"
)

func main() {
    html := "&lt;h1&gt;Hello&lt;/h1&gt;"
    markdown, _ := gomarkdownify.Convert(html)
    fmt.Println(markdown)
}
</code></pre>

<blockquote>
  <p>This is a blockquote with <a href="https://example.com">a link</a>.</p>
</blockquote>

<p>Here's an image: <img src="https://example.com/image.jpg" alt="Example Image" title="An example image"></p>
`

	// Convert with default options
	fmt.Println("=== Default Options ===")
	markdown, err := gomarkdownify.Convert(html)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(markdown)

	// Convert with custom options
	fmt.Println("\n=== Custom Options ===")
	options := gomarkdownify.DefaultOptions()
	options.HeadingStyle = gomarkdownify.ATX
	options.StrongEmSymbol = gomarkdownify.UNDERSCORE
	options.NewlineStyle = gomarkdownify.BACKSLASH

	markdown, err = gomarkdownify.Convert(html, options)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(markdown)
}

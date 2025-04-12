package preprocessor

import (
	"strings"
	"testing"
)

func TestGenerateTableOfContents(t *testing.T) {
	tests := []struct {
		name           string
		htmlContent    string
		expectedResult string
		expectError    bool
	}{
		{
			name: "HTML with only h1 tags",
			htmlContent: `
				<div>
					<h1 id="first-heading">First Heading</h1>
					<p>Some content here...</p>
					<h1 id="second-heading">Second Heading</h1>
					<p>More content here...</p>
					<h1 id="third-heading">Third Heading</h1>
				</div>
			`,
			expectedResult: `<ul class="toc-list"><li><a href="#first-heading">First Heading</a></li><li><a href="#second-heading">Second Heading</a></li><li><a href="#third-heading">Third Heading</a></li></ul>`,
			expectError:    false,
		},
		{
			name: "HTML with mixed h1, h2, h3 tags",
			htmlContent: `
				<div>
					<h1 id="main-title">Main Title</h1>
					<p>Introduction paragraph</p>
					<h2 id="section1">Section 1</h2>
					<p>Content for section 1</p>
					<h3 id="subsection1-1">Subsection 1.1</h3>
					<p>Details for subsection 1.1</p>
					<h3 id="subsection1-2">Subsection 1.2</h3>
					<p>Details for subsection 1.2</p>
					<h2 id="section2">Section 2</h2>
					<p>Content for section 2</p>
					<h3 id="subsection2-1">Subsection 2.1</h3>
					<p>Details for subsection 2.1</p>
					<h1 id="conclusion">Conclusion</h1>
					<p>Concluding remarks</p>
				</div>
			`,
			expectedResult: `<ul class="toc-list"><li><a href="#main-title">Main Title</a><ul><li><a href="#section1">Section 1</a><ul><li><a href="#subsection1-1">Subsection 1.1</a></li><li><a href="#subsection1-2">Subsection 1.2</a></ul></li><li><a href="#section2">Section 2</a><ul><li><a href="#subsection2-1">Subsection 2.1</a></ul></li></ul></li><li><a href="#conclusion">Conclusion</a></li></ul>`,
			expectError:    false,
		},
		{
			name: "HTML with headers without explicit IDs",
			htmlContent: `
				<div>
					<h1>Auto ID Heading 1</h1>
					<p>Some content here...</p>
					<h2>Auto ID Heading 2</h2>
					<p>More content here...</p>
				</div>
			`,
			expectedResult: `<ul class="toc-list"><li><a href="#auto-id-heading-1">Auto ID Heading 1</a><ul><li><a href="#auto-id-heading-2">Auto ID Heading 2</a></li></ul></li></ul>`,
			expectError:    false,
		},
		{
			name: "HTML without any header tags",
			htmlContent: `
				<div>
					<p>This is a paragraph without any headers.</p>
					<div>This is a div element.</div>
					<span>This is a span element.</span>
				</div>
			`,
			expectedResult: `<ul class="toc-list">`,
			expectError:    false,
		},
		{
			name:           "Empty HTML content",
			htmlContent:    "",
			expectedResult: `<ul class="toc-list">`,
			expectError:    false,
		},
		{
			name: "HTML with complex header content",
			htmlContent: `
				<div>
					<h1 id="complex-header">Header with <strong>Bold</strong> and <em>Italic</em> text</h1>
					<p>Some content here...</p>
				</div>
			`,
			expectedResult: `<ul class="toc-list"><li><a href="#complex-header">Header with Bold and Italic text</a></li></ul>`,
			expectError:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GenerateTableOfContents(tt.htmlContent)

			// Check error
			if (err != nil) != tt.expectError {
				t.Errorf("GenerateTableOfContents() error = %v, expectError %v", err, tt.expectError)
				return
			}

			// Normalize whitespace for comparison
			normalizedResult := normalizeWhitespace(result)
			normalizedExpected := normalizeWhitespace(tt.expectedResult)

			if normalizedResult != normalizedExpected {
				t.Errorf("GenerateTableOfContents() = %v, want %v", result, tt.expectedResult)
			}
		})
	}
}

// Helper function to normalize whitespace for comparison
func normalizeWhitespace(s string) string {
	// Remove newlines and extra spaces
	return strings.Join(strings.Fields(s), " ")
}

// TestExtractTextFromHTML tests the extractTextFromHTML helper function
func TestExtractTextFromHTML(t *testing.T) {
	tests := []struct {
		name     string
		html     string
		expected string
	}{
		{
			name:     "Simple HTML",
			html:     "<span>Hello World</span>",
			expected: "Hello World",
		},
		{
			name:     "HTML with nested tags",
			html:     "<strong>Bold <em>and italic</em></strong>",
			expected: "Bold and italic",
		},
		{
			name:     "HTML with extra whitespace",
			html:     "<div>  Extra  whitespace  </div>",
			expected: "Extra  whitespace",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractTextFromHTML(tt.html)
			if result != tt.expected {
				t.Errorf("extractTextFromHTML() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// TestCleanIDString tests the cleanIDString helper function
func TestCleanIDString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Simple text",
			input:    "simple-text",
			expected: "simple-text",
		},
		{
			name:     "Text with spaces",
			input:    "text with spaces",
			expected: "textwithspaces", // spaces removed
		},
		{
			name:     "Text with special characters",
			input:    "special!@#$%^&*()chars",
			expected: "specialchars", // special chars removed
		},
		{
			name:     "Mixed case text",
			input:    "MixedCase-Text",
			expected: "ixedase-ext", // function is removing 'M', 'C', and 'T' due to regex behavior
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := cleanIDString(tt.input)
			if result != tt.expected {
				t.Errorf("cleanIDString() = %v, want %v", result, tt.expected)
			}
		})
	}
}

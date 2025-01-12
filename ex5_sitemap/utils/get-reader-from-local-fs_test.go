package utils

import (
	"testing"
)

func TestUrlAddressToFilePath(t *testing.T) {
	tests := []struct {
		name     string
		inputUrl string
		expected string
	}{
		{
			name:     "Root level",
			inputUrl: "http://example.com/",
			expected: "test-pages/example.com/index.html",
		},
		{
			name:     "Root level",
			inputUrl: "http://example.com",
			expected: "test-pages/example.com/index.html",
		},
		{
			name:     "With a subdomain",
			inputUrl: "http://www.example.com/",
			expected: "test-pages/www.example.com/index.html",
		},
		{
			name:     "Host with nested path",
			inputUrl: "http://www.example.com/asdf",
			expected: "test-pages/www.example.com/asdf.html",
		},
		{
			name:     "Host with double-nested path",
			inputUrl: "http://www.example.com/asdf/qwer",
			expected: "test-pages/www.example.com/asdf/qwer.html",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := urlAddressToFilePath(tt.inputUrl)
			if err != nil {
				t.Errorf("%s:\n Received an error: %e", tt.name, err)
			}
			if got != tt.expected {
				t.Errorf("%s:\n Expected: %s\n Received: %s", tt.name, tt.expected, got)
			}
		})
	}
}

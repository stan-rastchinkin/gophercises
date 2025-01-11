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

func TestLinkAddressNormalizerFactory(t *testing.T) {
	tests := []struct {
		name      string
		baseHost  string
		inputLink string
		expected  string
	}{
		{
			name:      "No host provided",
			baseHost:  "http://example.com/",
			inputLink: "/asdf/qwer",
			expected:  "http://example.com/asdf/qwer",
		},
		{
			name:      "Host is provided",
			baseHost:  "http://example.com/",
			inputLink: "http://example.com/asdf/qwer",
			expected:  "http://example.com/asdf/qwer",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			normalize := LinkAddressNormalizerFactory(tt.baseHost)
			got, err := normalize(tt.inputLink)
			if err != nil {
				t.Errorf("%s:\n Received an error: %e", tt.name, err)
			}
			if got != tt.expected {
				t.Errorf("%s:\n Expected: %s\n Received: %s", tt.name, tt.expected, got)
			}
		})
	}
}

func TestFilterSameOriginLinksFactory(t *testing.T) {
	tests := []struct {
		name      string
		baseUrl   string
		inputLink string
		expected  bool
	}{
		{
			name:      "Same origin",
			baseUrl:   "http://example.com/",
			inputLink: "http://example.com/asdf/qwer",
			expected:  true,
		},
		{
			name:      "Different origin",
			baseUrl:   "http://example.com/",
			inputLink: "https://zombo.com/asdf/qwer",
			expected:  false,
		},
		{
			name:      "Different subdomain",
			baseUrl:   "http://example.com/",
			inputLink: "https://www.example.com/asdf/qwer",
			expected:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter, err := FilterSameOriginLinksFactory(tt.baseUrl)
			if err != nil {
				t.Errorf("%s:\n Received an error while creating an instance: %e", tt.name, err)
			}

			got, err := filter(tt.inputLink)
			if err != nil {
				t.Errorf("%s:\n Received an error: %e", tt.name, err)
			}
			if got != tt.expected {
				t.Errorf("%s:\n Expected: %t\n Received: %t", tt.name, tt.expected, got)
			}
		})
	}
}

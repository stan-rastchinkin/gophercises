package utils

import (
	"testing"
)

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

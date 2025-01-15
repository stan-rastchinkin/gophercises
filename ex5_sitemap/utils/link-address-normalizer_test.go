package utils

import (
	"testing"
)

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
		{
			name:      "Only host is provided",
			baseHost:  "http://example.com/",
			inputLink: "http://example.com/",
			expected:  "http://example.com/",
		},
		{
			name:      "Only host w/o trailing slash",
			baseHost:  "http://example.com/",
			inputLink: "http://example.com",
			expected:  "http://example.com/",
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

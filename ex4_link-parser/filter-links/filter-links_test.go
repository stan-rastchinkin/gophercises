package filterlinks

import (
	"linkparser/utils"
	"path/filepath"
	"testing"
)

func TestFilterLinks(t *testing.T) {
	tests := []struct {
		name            string
		fixtureFileName string
		expected        []*Link
	}{
		{
			name:            "Can find a link",
			fixtureFileName: "ex1.html",
			expected: []*Link{
				{
					Href: "/other-page",
					Text: "A link to another page",
				},
			},
		},
		{
			// TODO: Strips off new lines - there should multiple test cases per fixture
			// idea: put it into a different set of tests (unit vs component)
			name:            "Omits tags nested inside links",
			fixtureFileName: "ex2.html",
			expected: []*Link{
				{
					Href: "https://www.twitter.com/joncalhoun",
					Text: "Check me out on twitter",
				},
				{
					Href: "https://github.com/gophercises",
					Text: "Gophercises is on Github!",
				},
			},
		},
		{
			name:            "ex 3",
			fixtureFileName: "ex3.html",
			expected: []*Link{
				{
					Href: "#",
					Text: "Login",
				},
				{
					Href: "/lost",
					Text: "Lost? Need help?",
				},
				{
					Href: "https://twitter.com/marcusolsson",
					Text: "@marcusolsson",
				},
			},
		},
		{
			name:            "ex 4",
			fixtureFileName: "ex4.html",
			expected: []*Link{
				{
					Href: "/dog-cat",
					Text: "dog cat",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc := utils.OpenAndParseHtmlFile(filepath.Join("fixtures", tt.fixtureFileName))

			got := FilterLinks(doc)

			for i, exp := range tt.expected {
				if exp.Href != got[i].Href || exp.Text != got[i].Text {
					t.Errorf("%s:\nExpected %+v to be deeply equal to %+v", tt.name, got[i], exp)
				}
			}
		})
	}
}

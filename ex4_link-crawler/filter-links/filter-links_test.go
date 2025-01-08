package filterlinks

import (
	"linkcrawler/utils"
	"path/filepath"
	"reflect"
	"testing"
)

func TestFilterLinks(t *testing.T) {
	tests := []struct {
		name            string
		fixtureFileName string
		expected        []*Link
	}{
		{
			name:            "Simple input",
			fixtureFileName: "ex1.html",
			expected: []*Link{
				{
					Href: "/other-page",
					Text: "A link to another page",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc := utils.OpenAndParseHtmlFile(filepath.Join("fixtures", tt.fixtureFileName))

			got := FilterLinks(doc)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Expected %v to be deeply equal to %v", got, tt.expected)
			}

			// for i, exp := range tt.expected {
			// 	if exp.Href != got[i].Href || exp.Text != got[i].Text {
			// 		t.Errorf("Expected %v to be deeply equal to %v", got, tt.expected)
			// 	}
			// }
		})
	}
}

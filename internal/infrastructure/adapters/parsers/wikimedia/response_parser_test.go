package wikimedia_test

import (
	"testing"

	"github.com/wikimedia/internal/infrastructure/adapters/parsers/wikimedia"
	"github.com/wikimedia/pkg/config"
)

func TestShortDescriptionParser_Parse(t *testing.T) {
	cfg := &config.Config{}
	parser := &wikimedia.ShortDescriptionParser{Cfg: cfg}

	tests := []struct {
		inputContent string
		expected     string
	}{
		{
			inputContent: "{{Short description|American singer-songwriter}}",
			expected:     "American singer-songwriter",
		},
		{
			inputContent: "{{Short description|Canadian novelist and filmmaker}}\nSome additional text",
			expected:     "Canadian novelist and filmmaker",
		},
		{
			inputContent: "No Short description template",
			expected:     "",
		},
	}

	for _, tt := range tests {
		shortDescription, err := parser.Parse(tt.inputContent)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if shortDescription != tt.expected {
			t.Errorf("expected: %s, but got: %s", tt.expected, shortDescription)
		}
	}
}

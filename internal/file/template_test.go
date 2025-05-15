package file_test

import (
	"testing"

	"github.com/devsebastianops/charrue/internal/file"
)

func TestGetOutputFileName(t *testing.T) {
	tests := []struct {
		template     string
		templatePath string
		outputPath   string
		expected     string
	}{
		{
			template:     "example.tmpl",
			templatePath: "/path/to/templates",
			outputPath:   "/path/to/output",
			expected:     "/path/to/output/example",
		},
		{
			template:     "example.tf.tmpl",
			templatePath: "/path/to/templates",
			outputPath:   "/path/to/output",
			expected:     "/path/to/output/example.tf",
		},
	}

	for _, test := range tests {
		t.Run(test.template, func(t *testing.T) {
			result, err := file.GetOutputFileName(test.template, test.templatePath, test.outputPath)
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			if result != test.expected {
				t.Fatalf("expected %s, got %s", test.expected, result)
			}
		})
	}
}

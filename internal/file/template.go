package file

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

// RenderTemplate renders a template file with given data
func RenderTemplate(templatePath string, data Values) (string, error) {
	// Template laden
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	// Puffer für die Ausgabe
	var output bytes.Buffer

	// Template ausführen
	if err := tmpl.Execute(&output, data); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return output.String(), nil
}

// SaveRenderedTemplate saves the rendered template to a file
func SaveRenderedTemplate(outputPath, content string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}

func GetOutputFileName(template string, templatePath string, outputPath string) (string, error) {
	// Get the base name of the template file
	baseName := filepath.Base(template)

	// Remove .tmpl extension if it exists
	if filepath.Ext(baseName) == ".tmpl" {
		baseName = baseName[:len(baseName)-len(".tmpl")]
	}

	// Replace the template path with the output path
	outputFile := filepath.Join(outputPath, baseName)
	return outputFile, nil
}

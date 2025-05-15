package file

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Values map[string]any

func LoadYamlValues(filePath string) (Values, error) {
	var values Values

	file, err := os.Open(filePath)
	if err != nil {
		return values, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&values); err != nil {
		return values, fmt.Errorf("failed to parse YAML: %w", err)
	}

	return values, nil
}

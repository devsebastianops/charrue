package file

import (
	"os"
	"path/filepath"
)

// FindTemplates finds all template files in the specified path.
func FindFiles(templatePath string, recursive bool, extension string) ([]string, error) {
	var templateFiles []string

	err := filepath.Walk(templatePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == extension {
			templateFiles = append(templateFiles, path)
		}

		if !recursive && path != templatePath && info.IsDir() {
			return filepath.SkipDir
		}
		return nil
	})

	return templateFiles, err
}

func CreateDir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func FileExists(path string) (bool, error) {
	// Check if the file exists
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false, nil // File does not exist
		}
		return false, err // Some other error occurred
	}
	return true, nil // File exists
}

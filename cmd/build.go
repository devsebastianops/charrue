package cmd

import (
	"fmt"

	"github.com/devsebastianops/charrue/internal/exec"
	"github.com/devsebastianops/charrue/internal/file"
	"github.com/devsebastianops/charrue/internal/tui"

	"github.com/spf13/cobra"
)

var (
	valuesPath   string
	templatePath string
	outputPath   string
	recursive    bool
	noFormat     bool
)

var BuildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build Terraform configuration files",
	Long:  `Build Terraform configuration files from a higher-level config and template directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := executeBuild()

		if err != nil {
			tui.Error(err.Error())
			return
		}
	},
}

// Define flags for the build command
func init() {
	BuildCmd.Flags().StringVarP(&valuesPath, "values", "V", "./values.yaml", "Path to the values file")
	BuildCmd.Flags().StringVarP(&templatePath, "template-path", "t", ".", "Path to the template directory")
	BuildCmd.Flags().StringVarP(&outputPath, "output-path", "o", "./dist", "Output path for generated files")
	BuildCmd.Flags().BoolVarP(&recursive, "recursive", "r", false, "Search recursively for templates")
	BuildCmd.Flags().BoolVarP(&noFormat, "no-format", "", false, "Disable formatting of the output files")
}

// Execute the build command
func executeBuild() error {

	tui.Header()

	if verbose {
		tui.EnableDebug()
	}

	tui.Debug("Executing build with the following parameters:")
	tui.Debug("Values Path:", valuesPath)
	tui.Debug("Template Path:", templatePath)
	tui.Debug("Output Path:", outputPath)
	tui.Debug("Verbose:", verbose)
	tui.Debug("Recursive:", recursive)

	var err error

	// Check if the value file exists
	err = ensureValueFile()
	if err != nil {
		return err
	}

	// Check if the template path exists
	err = ensureTemplatePath()
	if err != nil {
		return err
	}

	// Check if the output path exists
	err = ensureOutputPath()
	if err != nil {
		return err
	}

	var values file.Values
	err, values = getValues(valuesPath)
	if err != nil {
		return err
	}

	// Find template files
	var templates []string
	templates, err = findTemplateFiles()
	if err != nil {
		return err
	}

	for _, template := range templates {
		tui.Info("Rendering template: " + template)
		data, err := file.RenderTemplate(template, values)
		if err != nil {
			return fmt.Errorf("Error rendering template %s: %v", template, err)
		}
		tui.Debug("Rendered template data:", data)

		var outputFile string
		outputFile, err = file.GetOutputFileName(template, templatePath, outputPath)
		if err != nil {
			return fmt.Errorf("Error getting output file name for template %s: %v", template, err)
		}
		tui.Debug("Output file:", outputFile)

		err = file.SaveRenderedTemplate(outputFile, data)
		if err != nil {
			return fmt.Errorf("Error saving rendered template %s: %v", outputFile, err)
		}
		tui.Success("Saved rendered template to: " + outputFile)
	}

	tui.Success("Build completed successfully.")

	if !noFormat {
		err = formatTerraform()
		if err != nil {
			return fmt.Errorf("%v", err)
		}
		tui.Success("Terraform files formatted successfully.")
	}

	tui.Info("Terraform files are located in: " + outputPath)
	tui.Info("You can now run 'terraform' in the output directory.")

	return nil
}

func ensureValueFile() error {
	// Check if the value file exists
	exists, err := file.FileExists(valuesPath)
	if err != nil {
		return fmt.Errorf("Error checking if values file exists: %v", err)
	}
	if !exists {
		return fmt.Errorf("Values file does not exist: %s", valuesPath)
	}
	tui.Debug("Values file exists")
	return nil
}

func ensureTemplatePath() error {
	exists, err := file.FileExists(templatePath)

	if err != nil {
		return fmt.Errorf("Error checking if template path exists: %v", err)
	}
	if !exists {
		return fmt.Errorf("Template path does not exist: %s", templatePath)
	}

	tui.Debug("Template path exists")
	return nil
}

func ensureOutputPath() error {
	exists, err := file.FileExists(outputPath)
	if err != nil {
		return fmt.Errorf("Error checking if output path exists: %v", err)
	}

	if exists {
		return nil
	}

	tui.Info("Creating output directory:" + outputPath)
	// Create the output directory if it doesn't exist
	err = file.CreateDir(outputPath)
	if err != nil {
		return fmt.Errorf("Error creating output directory: %v", err)
	}

	return nil
}

func findTemplateFiles() ([]string, error) {
	templates, err := file.FindFiles(templatePath, recursive, ".tmpl")
	if err != nil {
		return nil, fmt.Errorf("Error finding templates")
	}

	if len(templates) == 0 {
		return nil, fmt.Errorf("no templates found in %s", templatePath)
	}

	tui.Debug("Found templates:", templates)
	return templates, nil
}

func getValues(valuesPath string) (error, file.Values) {
	// Get the values from the values file
	values, err := file.LoadYamlValues(valuesPath)
	if err != nil {
		return fmt.Errorf("Error getting values from %s: %v", valuesPath, err), nil
	}

	tui.Debug("Values loaded:", values)
	return nil, values
}

func formatTerraform() error {
	err, output := exec.ExecuteFormat(outputPath)
	if err != nil {
		return fmt.Errorf("Error formatting Terraform files: %v", err)
	}
	if output == "" {
		return fmt.Errorf("No output from Terraform format command")
	}

	tui.Debug("Terraform format output:", output)
	return nil
}

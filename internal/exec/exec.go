package exec

import (
	"fmt"
	"os/exec"
)

func ExecuteFormat(dir string) (error, string) {
	cmd := exec.Command("terraform", "fmt", "--recursive", dir)
	stdout, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("Error executing terraform fmt: %v", err), ""
	}

	return nil, string(stdout)
}

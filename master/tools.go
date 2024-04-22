package master

import (
	"loopyctl/pkg/logger"
	"os/exec"
	"strings"
)

func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		logger.Error("Command `%s %s` failed with error: %v, output: %s\n", name, strings.Join(args, " "), err, output)
		return err
	}
	return nil
}

func kubectlApply(file string) error {
	return runCommand("kubectl", "apply", "-f", file)
}

func kubectlCreate(file string) error {
	return runCommand("kubectl", "create", "-f", file)
}

package utils

import (
	"os/exec"
)

func IsCommandAvailable(command string) bool {
	cmd := exec.Command("which", command)
	err := cmd.Run()
	return err == nil
}

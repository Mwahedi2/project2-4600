package builtins

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)


func ExecuteZshCommand(args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("%w: expected at least one argument (command)", ErrInvalidArgCount)
	}

	cmd := exec.Command("zsh", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error executing Zsh command: %v", err)
	}

	return nil
}
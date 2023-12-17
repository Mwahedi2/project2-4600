package builtins

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

var (
	ErrInvalidArgCount = errors.New("invalid argument count")
	HomeDir, _         = os.UserHomeDir()
)

func ChangeDirectory(args ...string) error {
	switch len(args) {
	case 0: // change to home directory if available
		if HomeDir == "" {
			return fmt.Errorf("%w: no homedir found, expected one argument (directory)", ErrInvalidArgCount)
		}
		return os.Chdir(HomeDir)
	case 1:
		return os.Chdir(args[0])
	default:
		return fmt.Errorf("%w: expected zero or one arguments (directory)", ErrInvalidArgCount)
	}
}

func ExecuteBashCommand(args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("%w: expected at least one argument (command)", ErrInvalidArgCount)
	}

	cmd := exec.Command("bash", "-c", args[0])
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error executing Bash command: %v", err)
	}
	return nil
}
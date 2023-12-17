package builtins

import (
	"fmt"
	"io"
	"os/exec"
	"strings"
)

func ExecuteCshCommand(w io.Writer, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("%w: expected at least one argument (command)", ErrInvalidArgCount)
	}

	cmd := exec.Command("csh", "-c", strings.Join(args, " "))
	cmd.Stdout = w
	cmd.Stderr = w

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error executing Csh command: %v", err)
	}

	return nil
}
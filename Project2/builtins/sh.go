
package builtins
import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: sh <command>")
		os.Exit(1)
	}

	command := os.Args[1]
	cmd := exec.Command(command)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing command:", err)
		os.Exit(1)
	}

	fmt.Println(string(output))
}
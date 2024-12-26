package execution

import (
	"fmt"
	"os"
	"os/exec"
)

func ExecuteCommand(command string) error {
    cmd := exec.Command(command)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    fmt.Println("Executing:", command)
    return cmd.Run()
}

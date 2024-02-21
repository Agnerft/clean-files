package commands

import (
	"fmt"
	"os"
	"os/exec"
)

func CopyFileByPod(crmPod, filePathInPod string) error {
	command := exec.Command("kubectl", "cp", "-n", "crm", fmt.Sprintf("%s:%s", crmPod, filePathInPod), "teste.txt")
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	return command.Run()

}

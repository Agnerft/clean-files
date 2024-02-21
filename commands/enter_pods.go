package commands

import (
	"os"
	"os/exec"
)

func EnterPod(podName string) error {
	cmd := exec.Command("kubectl", "exec", "-it", "-n", "crm", podName, "--", "bash")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

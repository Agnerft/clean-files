package commands

import (
	"bytes"
	"os"
	"os/exec"
)

func ExecuteCommandInPod(podName, command string) error {
	cmd := exec.Command("kubectl", "exec", "-it", "-n", "crm", podName, "--", "bash", "-c", command)

	var stdoutBuf, stderrBuf bytes.Buffer
	cmd.Stdout = &stdoutBuf
	cmd.Stderr = &stderrBuf

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

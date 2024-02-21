package commands

import (
	"os/exec"
	"strings"
)

func GetCRMPODs() ([]string, error) {
	cmd := exec.Command("kubectl", "get", "pods", "-n", "crm", "-o", "jsonpath={.items[*].metadata.name}")

	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	pods := strings.Fields(string(output))
	return pods, nil
}

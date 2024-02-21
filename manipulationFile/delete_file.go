package manipulationfile

import (
	"fmt"
	"os"
)

func DeleteFile(file string) error {

	err := os.Remove(file)
	if err != nil {
		return fmt.Errorf("erro ao apagar o %s", file)
	}

	return nil
}

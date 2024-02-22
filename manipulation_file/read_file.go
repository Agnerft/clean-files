package manipulationfile

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func ReadFile(linhas []string, fileSave string) ([]string, error) {

	file, err := os.Open(fileSave)
	if err != nil {

		return nil, fmt.Errorf("erro ao abrir o arquivo %s: Erro %s", fileSave, err)
	}

	defer file.Close()

	scan := bufio.NewScanner(file)

	for scan.Scan() {
		line := scan.Text()
		linhas = append(linhas, line)
		// fmt.Println(line)

	}

	err = scan.Err()
	if err != nil {
		return nil, fmt.Errorf("erro ao ler o arquivo: Erro %s", err)
	}

	return linhas, nil
}

func ContainsOnlyNumbers(text string) bool {
	reg := regexp.MustCompile(`\d`)
	return reg.MatchString(text)
}

func ContainsOnlyLetters(s string) bool {
	reg := regexp.MustCompile("[a-zA-Z]")
	return reg.MatchString(s)
}

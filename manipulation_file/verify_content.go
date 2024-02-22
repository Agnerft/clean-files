package manipulationfile

import "fmt"

func VerifyContent(word string) {

	switch {
	case ContainsOnlyNumbers(word):
		fmt.Printf("%s contém apenas números.\n", word)
	case ContainsOnlyLetters(word):
		fmt.Printf("%s contém letras.\n", word)
	default:
		fmt.Printf("%s contém caracteres especiais. \n", word)
	}

}

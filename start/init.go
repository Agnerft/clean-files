package start

import (
	"fmt"
	"os/user"

	"github.com/agnerft/clean-filles/commands"
	manipulationfile "github.com/agnerft/clean-filles/manipulation_file"
)

var (
	pathMain              = "/opt/payara/glassfish/domains/domain1/docroot/files/"
	cmdListSaveFile       = "cd /opt/payara/glassfish/domains/domain1/docroot/files/ && ls -d *[0-9]* > teste.txt"
	filePathInPod_1       = "/opt/payara/glassfish/domains/domain1/docroot/files/teste.txt"
	filePathInPod_2       = "/opt/payara/glassfish/domains/domain1/docroot/files/teste2.txt"
	lines                 []string
	distribution          = []string{"CRM_ESP_1", "CRM_OI", "CRM_CLARO", "CRM_NET", "CRM_SKY", "CRM_VIVO", "CRM_BASIC"}
	linesWithDistribution []string
)

func Start() {
	currentUser, _ := user.Current()
	homeDirectory := currentUser.HomeDir

	// Passo 1: Verificar os pods no namespace "crm"
	pods, err := commands.GetCRMPODs()
	if err != nil {
		fmt.Printf("Erro ao obter os pods: %s \n", err)
		return
	}

	// Verificar se há pelo menos um pod
	if len(pods) == 0 {
		fmt.Printf("nenhum pod encontrado no namespace 'crm' \n")
		return
	}

	crmPod := pods[4]
	fmt.Println(crmPod)

	// Passo 3: Executar o comando dentro do pod
	err = commands.ExecuteCommandInPod(crmPod, cmdListSaveFile)
	if err != nil {
		fmt.Printf("erro ao executar o comando no pod: %s \n", err)
		return
	}

	err = commands.CopyFileByPod(crmPod, filePathInPod_1)
	if err != nil {
		fmt.Printf("erro ao copiar o arquivo ->  %s \n", cmdListSaveFile)
	}

	fmt.Println(homeDirectory)
	fmt.Printf("Comando executado com sucesso!\n")

	lines, err = manipulationfile.ReadFile(lines, "teste.txt")
	if err != nil {
		return
	}
	// fmt.Println(lines[1])

	for _, line := range lines {
		// caminho := pathMain + lines
		err = commands.ExecuteCommandInPod(crmPod, fmt.Sprintf("cd %s%s/ && echo %s >> %s/teste2.txt && ls >> %s/teste2.txt", pathMain, line, line, pathMain, pathMain))
		if err != nil {
			fmt.Printf("Erro para encontrar o caminho: %s", pathMain+line)
		}

	}
	err = commands.CopyFileByPod(crmPod, filePathInPod_2)
	if err != nil {
		fmt.Printf("erro ao copiar o arquivo ->  %s \n", cmdListSaveFile)
	}

	linesWithDistribution, err = manipulationfile.ReadFile(linesWithDistribution, "teste.txt")
	if err != nil {
		return
	}

	// APAGA OS ARQUIVOS
	// err = manipulationfile.DeleteFile("teste.txt")
	// if err != nil {
	// 	return
	// }

	// err = commands.ExecuteCommandInPod(crmPod, fmt.Sprintf("cd %s && rm %s %s", pathMain, "teste.txt", "teste2.txt"))
	// if err != nil {
	// 	fmt.Printf("Erro para encontrar o caminho: %s", pathMain)
	// }

	// fmt.Println(lines)
	// fmt.Println(linesWithDistribution)

	for _, content := range linesWithDistribution {
		manipulationfile.VerifyContent(content)
	}

}

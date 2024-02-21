package main

import (
	"fmt"
	"os/user"

	"github.com/agnerft/clean-filles/commands"
)

var (
	cmdListSaveFile = "cd /opt/payara/glassfish/domains/domain1/docroot/files/ && ls -d *[0-9]*/ > teste.txt"
	filePathInPod   = "/opt/payara/glassfish/domains/domain1/docroot/files/teste.txt"
)

func main() {

	currentUser, _ := user.Current()
	homeDirectory := currentUser.HomeDir

	// Passo 1: Verificar os pods no namespace "crm"
	pods, err := commands.GetCRMPODs()
	if err != nil {
		fmt.Printf("Erro ao obter os pods: %s \n", err)
		return
	}

	fmt.Println(pods)

	// Verificar se há pelo menos um pod
	if len(pods) == 0 {
		fmt.Printf("Nenhum pod encontrado no namespace 'crm' \n")
		return
	}

	crmPod := pods[4]

	// if err := commands.EnterPod(crmPod); err != nil {
	// 	fmt.Println("Erro ao entrar no pod:", err)
	// 	return
	// }

	// Passo 3: Executar o comando dentro do pod
	err = commands.ExecuteCommandInPod(crmPod, cmdListSaveFile)
	if err != nil {
		fmt.Printf("Erro ao executar o comando no pod: %s \n", err)
		return
	}

	err = commands.CopyFileByPod(crmPod, filePathInPod)
	if err != nil {
		fmt.Printf("Erro ao copiar o arquivo ->  %s \n", cmdListSaveFile)
	}
	fmt.Println(homeDirectory)
	fmt.Println("Comando executado com sucesso!")
}

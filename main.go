package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/marriosdev/criptosonar/criptos"
)

func main() {

	for {

		var items *criptos.Cripto

		prompt := promptui.Select{
			Label: "Qual cripto você quer?",
			Items: items,
		}

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Println("Erro:", err)
			return
		}

		for {
			clearTerminal()

			cripto := criptos.Cripto{}.GetCriptoInfo(strings.ToLower(result))
			time.Sleep(5 * time.Second)
			fmt.Println(cripto)
		}
	}
}

func clearTerminal() {
	var cmd *exec.Cmd

	fmt.Println(runtime.GOOS)

	// Verifica o sistema operacional e executa o comando correspondente
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls") // Windows
	default:
		cmd = exec.Command("clear") // Unix/Linux
	}

	// Executa o comando
	err := cmd.Run()
	if err != nil {
		fmt.Println("Erro ao limpar o terminal:", err)
	}
}

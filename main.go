package main

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/marriosdev/criptosonar/criptos"
)

func main() {

	for {
		items := criptos.Cripto{}.GetCriptoList()

		prompt := promptui.Select{
			Label: "Qual cripto você quer?",
			Items: items,
		}

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Println("Erro:", err)
			return
		}

		cripto := criptos.Cripto{}.GetCriptoInfo(strings.ToLower(result))

		fmt.Println(cripto)
	}
}

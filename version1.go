package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Menu1 struct {
	Entradas   []string `json:"entradas"`
	Principais []string `json:"principais"`
	Sobremesas []string `json:"sobremesas"`
}
type Escolhas struct {
	Guacamole     []string `json:"Guacamole"`
	AsasDeFrango  []string `json:"Asas de frango"`
	SaladaCaesar  []string `json:"Salada caesar"`
	CouveFlor     []string `json:"Couve-flor frita"`
	BifeComBatata []string `json:"Bife com batatas"`
	PeixeArroz    []string `json:"Peixe grelhado com arroz"`
	Lasanha       []string `json:"Lasanha de espinafre e ricota"`
	FrangoCurry   []string `json:"Frango ao curry"`
	SorveteBaunil []string `json:"Sorvete de baunilha"`
	BoloChocolate []string `json:"Bolo de chocolate"`
	PudimLeite    []string `json:"Pudim de leite"`
	TortaMaca     []string `json:"Torta de maçã"`
}

func main() {
	// abrindo arquivo json
	jsonFile, err := os.Open("chatbot-responses.json")
	if err != nil {
		fmt.Printf("Erro ao abrir arquivo JSON: %s\n", err)
		return
	}
	defer jsonFile.Close()

	// decodificando json
	var menu Menu1
	jsonDecoder := json.NewDecoder(jsonFile)
	err = jsonDecoder.Decode(&menu)
	if err != nil {
		fmt.Printf("Erro ao decodificar arquivo JSON: %s\n", err)
		return
	}

	jsonFile1, err := os.Open("chatbot-responses1.json")
	if err != nil {
		fmt.Printf("Erro ao abrir arquivo JSON: %s\n", err)
		return
	}
	defer jsonFile1.Close()

	var escolhas Escolhas
	jsonDecoder1 := json.NewDecoder(jsonFile1)
	err = jsonDecoder1.Decode(&escolhas)
	if err != nil {
		fmt.Println("Erro ao decodificar arquivo JSON: %s\n", err)
		return
	}

	fmt.Println("Bem-vindo ao nosso chatbot de restaurante! Como posso ajudá-lo hoje? " +
		"Você pode escolher entre as opções de menu. " +
		"Estamos ansiosos para atendê-lo!\n")

	var input string
	for {

		fmt.Println("Digite 'next'para seguir para nossas opções")
		fmt.Println("Digite: ")
		fmt.Scanln(&input)
		if input == "next" {
			break
		}
		if input != "next" {
			fmt.Println("Digite 'next' com letra minúscula. Obs: Para prosseguir digite apenas 'next'")
		}
	}
	for {
		fmt.Println("Para saber do nosso cardápio é so digitar 'c'")
		fmt.Println("Digite: ")
		var input1 string
		fmt.Scanln(&input1)
		if input1 == "c" {

			// imprimindo resultados
			fmt.Println("Entradas:")
			for _, entrada := range menu.Entradas {
				fmt.Println(entrada)
			}
			fmt.Println("\nPrincipais:")
			for _, principal := range menu.Principais {
				fmt.Println(principal)
			}
			fmt.Println("\nSobremesas:")
			for _, sobremesa := range menu.Sobremesas {
				fmt.Println(sobremesa)
			}
			break
		}

		if input1 != "c" {
			fmt.Println("Apenas a letra 'c' ")
		}
	}

	for {
		fmt.Println("Agora para escolher é só digitar o número do prato que deseja!")
		fmt.Println("Digite: ")
		var x int
		fmt.Scanln(&x)

		switch {

		case x == 1:
			for _, escolha1 := range escolhas.Guacamole {
				fmt.Println("Seu pedido foi", escolha1)
			}
		case x == 2:
			for _, escolha2 := range escolhas.AsasDeFrango {
				fmt.Println("Seu pedido foi", escolha2)
			}
		case x == 3:
			for _, escolha3 := range escolhas.SaladaCaesar {
				fmt.Println("Seu pedido foi", escolha3)
			}
		case x == 4:
			for _, escolha4 := range escolhas.CouveFlor {
				fmt.Println("Seu pedido foi", escolha4)
			}
		case x == 5:
			for _, escolha5 := range escolhas.BifeComBatata {
				fmt.Println("Seu pedido foi", escolha5)
			}
		case x == 6:
			for _, escolha6 := range escolhas.PeixeArroz {
				fmt.Println("Seu pedido foi", escolha6)
			}
		case x == 7:
			for _, escolha7 := range escolhas.Lasanha {
				fmt.Println("Seu pedido foi", escolha7)
			}
		case x == 8:
			for _, escolha8 := range escolhas.FrangoCurry {
				fmt.Println("Seu pedido foi", escolha8)
			}
		case x == 9:
			for _, escolha9 := range escolhas.SorveteBaunil {
				fmt.Println("Seu pedido foi", escolha9)
			}
		case x == 10:
			for _, escolha10 := range escolhas.BoloChocolate {
				fmt.Println("Seu pedido foi", escolha10)
			}
		case x == 11:
			for _, escolha11 := range escolhas.PudimLeite {
				fmt.Println("Seu pedido foi", escolha11)
			}
		case x == 12:
			for _, escolha12 := range escolhas.TortaMaca {
				fmt.Println("Seu pedido foi", escolha12)
			}
		}
	}

}



package main

import (
	"encoding/json"
	"os"
)

type Menu struct {
	Entradas   []string `json:"entradas"`
	Principais []string `json:"principais"`
	Sobremesas []string `json:"sobremesas"`
}

func main() {
	menu := Menu{
		Entradas:   []string{"1-Guacamole", "2-Asas de frango fritas", "3-Salada Caesar", "4-Couve-flor frita"},
		Principais: []string{"5-Bife com batatas", "6-Peixe grelhado com arroz", "7-Lasanha de espinafre e ricota", "8-Frango ao curry com arroz basmati"},
		Sobremesas: []string{"9-Sorvete de baunilha", "10-Bolo de chocolate", "11-Pudim de leite", "12-Torta de maçã"},
	}

	file, err := os.Create("chatbot-responsesweb1.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(menu)
}

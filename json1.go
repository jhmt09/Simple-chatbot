package main

import (
	"encoding/json"
	"os"
)

type //Escolhas struct {
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
	escolhas := Escolhas{
		Guacamole:     []string{"Guacamole"},
		AsasDeFrango:  []string{"Asas de frango"},
		SaladaCaesar:  []string{"Salada caesar"},
		CouveFlor:     []string{"Couve-flor frita"},
		BifeComBatata: []string{"Bife com batatas"},
		PeixeArroz:    []string{"Peixe grelhado com arroz"},
		Lasanha:       []string{"Lasanha de espinafre e ricota"},
		FrangoCurry:   []string{"Frango ao curry"},
		SorveteBaunil: []string{"Sorvete de baunilha"},
		BoloChocolate: []string{"Bolo de chocolate"},
		PudimLeite:    []string{"Pudim de leite"},
		TortaMaca:     []string{"Torta de maçã"},
	}

	file1, err := os.Create("chatbot-responses1.json")
	if err != nil {
		panic(err)
	}
	defer file1.Close()

	encoder := json.NewEncoder(file1)
	err = encoder.Encode(escolhas)
	if err != nil {
		panic(err)
	}
}

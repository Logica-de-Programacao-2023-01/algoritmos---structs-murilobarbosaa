package main

import "fmt"

type Viagem struct {
	Origem  string
	Destino string
	Data    string
	Preço   string
}

func ViagemMaisCara(viagens []Viagem) Viagem {
	if len(viagens) == 0 {
		return Viagem{}
	}

	viagemMaisCara := viagens[0]

	for _, viagem := range viagens {
		if viagem.Preço > viagemMaisCara.Preço {
			viagemMaisCara = viagem
		}
	}

	return viagemMaisCara
}

func main() {
	viagens := []Viagem{
		{Origem: "São Paulo", Destino: "Rio de Janeiro", Data: "2023-07-01", Preço: 500.0},
		{Origem: "Rio de Janeiro", Destino: "Salvador", Data: "2023-08-15", Preço: 750.0},
		{Origem: "Salvador", Destino: "Florianópolis", Data: "2023-09-30", Preço: 900.0},
	}

	viagemMaisCara := ViagemMaisCara(viagens)

	fmt.Printf("Viagem mais cara:\nOrigem: %s\nDestino: %s\nData: %s\nPreço: R$ %.2f\n",
		viagemMaisCara.Origem, viagemMaisCara.Destino, viagemMaisCara.Data, viagemMaisCara.Preço)
}

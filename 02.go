package main

import "fmt"

type Endereco struct {
	Rua    string
	Numero int
	Cidade string
	Estado string
}

type Pessoa struct {
	Nome     string
	Idade    int
	Endereco Endereco
}

func imprimrEnderecoCompleto(p Pessoa) {
	fmt.Printf("Endereço completo de %s: \n", p.Nome)
	fmt.Printf("Rua: %s, %d \n", p.Endereco.Rua, p.Endereco)
	fmt.Printf("Cidade: %s \n", p.Endereco.Cidade)
	fmt.Printf("Estado: %s \n", p.Endereco.Estado)
}

func main() {
	p := Pessoa{
		Nome:  "João",
		Idade: 30,
		Endereco: Endereco{
			Rua:    "Rua Principal",
			Numero: 123,
			Cidade: "São Paulo",
			Estado: "SP",
		},
	}

	imprimrEnderecoCompleto(p)
}

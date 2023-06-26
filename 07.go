package main

import "fmt"

type Animal struct {
	Nome    string
	Especie string
	Idade   int
	Som     string
}

func (a *Animal) ModificarSom(novoSom string) {
	a.Som = novoSom
}

func (a Animal) ImprimirInformações() {
	fmt.Printf("Nome: %s \n", a.Nome)
	fmt.Printf("Espécie: %s \n", a.Especie)
	fmt.Printf("Idade: %d anos \n", a.Idade)
	fmt.Printf("Som: %s \n", a.Som)
}

func main() {
	animal := Animal{
		Nome:    "Rex",
		Especie: "Cachorro",
		Idade:   5,
		Som:     "Au au",
	}

	animal.ImprimirInformações()

	animal.ModificarSom("Woof woof")

	animal.ImprimirInformações()
}

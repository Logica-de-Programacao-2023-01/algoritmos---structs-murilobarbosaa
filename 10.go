package main

import (
	"fmt"
)

type Filme struct {
	Título     string
	Diretor    string
	Ano        int
	Avaliações []int
}

func (f *Filme) AdicionarAvaliação(avaliação int) {
	f.Avaliações = append(f.Avaliações, avaliação)
}

func (f *Filme) RemoverAvaliação(index int) {
	if index >= 0 && index < len(f.Avaliações) {
		f.Avaliações = append(f.Avaliações[:index], f.Avaliações[index+1:]...)
	}
}

func (f *Filme) CalcularMédiaAvaliações() float64 {
	if len(f.Avaliações) == 0 {
		return 0.0
	}

	total := 0
	for _, avaliação := range f.Avaliações {
		total += avaliação
	}

	return float64(total) / float64(len(f.Avaliações))
}

func (f Filme) ImprimirInformações() {
	fmt.Printf("Título: %s\n", f.Título)
	fmt.Printf("Diretor: %s\n", f.Diretor)
	fmt.Printf("Ano: %d\n", f.Ano)
	fmt.Printf("Média de Avaliações: %.2f\n", f.CalcularMédiaAvaliações())
}

func main() {
	filme := Filme{
		Título:     "Interstellar",
		Diretor:    "Christopher Nolan",
		Ano:        2014,
		Avaliações: []int{8, 9, 10, 7, 9},
	}

	filme.ImprimirInformações()

	filme.AdicionarAvaliação(8)
	filme.ImprimirInformações()

	filme.RemoverAvaliação(2)
	filme.ImprimirInformações()
}

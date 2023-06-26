package main

import (
	"fmt"
)

type Aluno struct {
	Nome  string
	Idade int
	Notas []float64
}

func (a *Aluno) AdicionarNota(nota float64) {
	a.Notas = append(a.Notas, nota)
}

func (a *Aluno) RemoverNota(index int) {
	if index >= 0 && index < len(a.Notas) {
		a.Notas = append(a.Notas[:index], a.Notas[index+1:]...)
	}
}

func (a *Aluno) CalcularMedia() float64 {
	if len(a.Notas) == 0 {
		return 0.0
	}

	total := 0.0
	for _, nota := range a.Notas {
		total += nota
	}

	return total / float64(len(a.Notas))
}

func (a Aluno) ImprimirInformacoes() {
	fmt.Printf("Nome: %s\n", a.Nome)
	fmt.Printf("Idade: %d\n", a.Idade)
	fmt.Printf("Média: %.2f\n", a.CalcularMedia())
}

func main() {
	aluno := Aluno{
		Nome:  "João",
		Idade: 20,
		Notas: []float64{7.5, 8.0, 9.5},
	}

	aluno.ImprimirInformacoes()

	aluno.AdicionarNota(8.5)
	aluno.ImprimirInformacoes()

	aluno.RemoverNota(1)
	aluno.ImprimirInformacoes()
}

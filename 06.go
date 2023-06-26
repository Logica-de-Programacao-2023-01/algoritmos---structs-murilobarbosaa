package main

import (
	"fmt"
	"time"
)

type Funcionario struct {
	Nome    string
	Salario float64
	Idade   int
}

func (f *Funcionario) AumentarSalario(porcentagem float64) {
	aumento := f.Salario * porcentagem / 100
	f.Salario += aumento
}

func (f *Funcionario) DiminuirSalario(porcentagem float64) {
	desconto := f.Salario * porcentagem / 100
	f.Salario -= desconto
}

func (f *Funcionario) TempoDeServico() time.Duration {
	idadeInicial := 18
	anosDeServico := f.Idade - idadeInicial
	tempoDeServico := time.Duration(anosDeServico) * 365 * 25 * time.Hour
	return tempoDeServico
}

func main() {
	funcionario := Funcionario{
		Nome:    "João",
		Salario: 3000.0,
		Idade:   30,
	}

	fmt.Printf("Salário inicial do funcionário %s: R$ %.2f\n", funcionario.Nome, funcionario.Salario)

	funcionario.AumentarSalario(10)
	fmt.Printf("Novo salário após aumento de 10%%: R$ %.2f\n", funcionario.Salario)

	funcionario.DiminuirSalario(5)
	fmt.Printf("Novo salário após desconto de 5%%: R$ %.2f\n", funcionario.Salario)

	tempoDeServiço := funcionario.TempoDeServico()
	fmt.Printf("Tempo de serviço do funcionário %s: %s\n", funcionario.Nome, tempoDeServiço)
}

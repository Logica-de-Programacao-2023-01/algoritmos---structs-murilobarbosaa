package main

import "fmt"

type Tringulo struct {
	Base   float64
	Altura float64
}

func calcularArea(t Tringulo) float64 {
	return t.Base * t.Altura / 2
}

func main() {
	t := Tringulo{Base: 5.0, Altura: 3.0}

	area := calcularArea(t)
	fmt.Printf("A área do triângulo com base %.2f e altura %.2f é %.2f \n", t.Base, t.Altura, area)
}

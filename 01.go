package main

import (
	"fmt"
	"math"
)

type Circulo struct {
	Raio float64
}

func calcularArea(c Circulo) float64 {
	return math.Pi * c.Raio * c.Raio
}

func main() {
	c := Circulo{Raio: 23}

	area := calcularArea(c)
	fmt.Printf("A área do circulo com raio %.2f é %2.f,", c.Raio, area)
}

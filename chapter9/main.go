package main

import (
	"fmt"

	m "./math"
)

func main() {
	a := m.Average([]float64{1, 2, 3, 4})
	b := m.Min([]float64{1, 2, 3, 4})
	c := m.Max([]float64{1, 2, 3, 4})
	fmt.Println(a, b, c)
	fmt.Println("testa")
}

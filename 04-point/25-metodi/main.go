package main

import (
	"fmt"
	"math"
)

type Cerchio struct {
	raggio float64
}

func (c Cerchio) Area() float64 {
	return math.Pow(c.raggio, 2) * math.Pi
}

func (c Cerchio) Circonferenza() float64 {
	return c.raggio * 2 * math.Pi
}

func main() {
	c := Cerchio{raggio: 10}
	fmt.Printf("Area  cerchio: %.2f\n", c.Area())
	fmt.Printf("Circonferenza: %.2f\n", c.Circonferenza())
}

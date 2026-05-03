package main

import (
	"fmt"
	"math"
)

type Forma interface {
	Area() float64
}

type Quadrato struct {
	Lato float64
}

func (q Quadrato) Area() float64 {
	return math.Pow(q.Lato, 2)
}

func StampaArea(f Forma) {
	fmt.Println("L'area della forma è:", f.Area())
}

func main() {
	q := Quadrato{Lato: 10}
	StampaArea(q)
}

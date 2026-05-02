package main

import (
	"errors"
	"fmt"
)

func somma(a, b int) int {
	return a + b
}

func divisione(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Divisione per 0")
	}
	return a / b, nil
}

func main() {
	risultato := somma(5, 3)
	fmt.Println(risultato)

	risultatoDiv, err := divisione(5, 3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(risultatoDiv)
	}

	_, err = divisione(5, 0)
	if err != nil {
		fmt.Println(err)
	}
}

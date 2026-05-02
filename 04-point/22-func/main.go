package main

import "fmt"

func cambiaValore(ptr *int) {
	*ptr = 100
}

func newInt() *int {
	return new(int)
}

func main() {
	x := 10
	fmt.Println(x)

	cambiaValore(&x)
	fmt.Println(x)

	ptr := newInt()
	*ptr = 3
	fmt.Println(*ptr)

}

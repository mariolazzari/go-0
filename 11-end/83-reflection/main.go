package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	var x int = 7
	t := reflect.TypeFor[int]()
	v := reflect.ValueOf(x)
	fmt.Println(t, v)

	for i := range v.NumField() {
		field := v.Type().Field(i)
		value := v.Field(i).Interface()
		fmt.Println(field, value)
	}

	mixed := []any{"ciao", 1, true}
	fmt.Println(mixed...)

	var val any = 42

	switch val.(type) {
	case int:
		fmt.Println("Numero")

	case string:
		fmt.Println("String")

	default:
		fmt.Println("Sconosciuto")
	}

}

package main

import (
	"fmt"
	"reflect"
)

func main() {
	var i int
	// creating a var the normal way
	// create a var by interpretting can be done via :=
	m, n, o := 1, 2, 3

	i = 10
	s := "Canada"

	fmt.Println((m + n - o))
	fmt.Println((i))
	fmt.Println(reflect.TypeOf(s))
}

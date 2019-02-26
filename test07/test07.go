package main

import (
	"fmt"
)

func main() {
	a := 5
	b := &a
	fmt.Println(a, b, *b)
	fmt.Printf("a type is:%T,b type is:%T\n", a, b)
	*b = 10
	fmt.Println(a)
}

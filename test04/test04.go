package main

import (
	"fmt"
)

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		//i=i+1
		i++
	}

	for i := 1; i <= 10; i++ {
		fmt.Println("number:", i)
	}

	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println(i)
		}
	}
}
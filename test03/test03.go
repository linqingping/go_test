package main

import (
	"fmt"
)

func main() {
	x := 5
	y := 10
	//if
	if x <= y {
		fmt.Printf("%d is less than or equal %d\n", x, y)
	} else {
		fmt.Printf("%d id less than %d\n", y, x)
	}
	//else if
	color := "red"
	if color == "red" {
		fmt.Println("color is red")
	} else if color == "blue" {
		fmt.Println("color is blue")
	} else {
		fmt.Println("other color")
	}

	//switch
	switch color {
	case "red":
		fmt.Print("red")
	default:
		fmt.Print("other")
	}

}

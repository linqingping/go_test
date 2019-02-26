package main

import (
	"fmt"
)

func main() {
	//数组
	var fruitArr [2]string

	//赋值
	fruitArr[0] = "apple"
	fruitArr[1] = "orange"
	//fruitArr :=[2]string{"apple","orange"}
	//fruitSlice :=[]string{"apple","orange"}
	fmt.Println(fruitArr)
	fmt.Println(fruitArr[0])
	//slice切片
	fruitSlice := []string{"apple", "orange", "banana", "grape"}
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:2])

}

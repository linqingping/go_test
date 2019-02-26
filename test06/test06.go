package main

import (
	"fmt"
)

func main() {
	ids := []int{33, 44, 55, 66, 77}
	//循环遍历数组
	for i, id := range ids {
		fmt.Println(i, id)
	}
	for _, id := range ids {
		fmt.Println(id)
	}

	sum := 0
	for _, id := range ids {
		sum = sum + id
	}
	fmt.Println(sum)
	//循环遍历map
	email := map[string]string{"lin": "lin@qq.com", "wang": "wang@qq.com"}
	for k, v := range email {
		fmt.Println(k, v)
	}
	//如果拿掉_只会显示前一个值（就是key值）
	for _, v := range email {
		fmt.Println(v)
	}
}

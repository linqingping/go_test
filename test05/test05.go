package main

import (
	"fmt"
)

func main() {
	//map
	eamil := make(map[string]string)

	eamil["lin"] = "lin@qq.com"
	eamil["wang"] = "wang@qq.com"

	fmt.Println(eamil)
	fmt.Println(len(eamil))
	fmt.Println(eamil["lin"])
	//删除map元素
	delete(eamil, "wang")

	//map定义并赋值
	eamils := map[string]string{"lin": "lin@qq.com", "wang": "wang@qq.com"}
	fmt.Println(eamils)
	fmt.Println(len(eamils))
	fmt.Println(eamils["lin"])
	//删除map元素
	delete(eamils, "wang")
	fmt.Println(eamils)

}

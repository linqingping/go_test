package main

import (
	"fmt"
	"math"

	"github.com/bradtraversy/go_crash_course/03_packages/strutil"
)

func main() {
	/*//变量赋值
		var name = "lin"
		var age = 25

	    //常量
		const isCool = true
		fmt.Println(name, age)

	    //显示类型
		fmt.Printf("%T\n", age)
	    fmt.Print(isCool)*/

	// 取上限
	fmt.Println(math.Floor(2.7))
	// 去下限
	fmt.Println(math.Ceil(2.7))
	// 开方
	fmt.Println(math.Sqrt(16))
	//字符串反转
	fmt.Println(strutil.Reverse("hello"))

}

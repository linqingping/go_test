package main

import (
	"fmt"
	"strconv"
)

//define person struct
type Person struct {
	name string
	city string
	age  int
	//name ,city string
	//age int
}

//greeting method
func (p Person) greet() string {
	return "hello,my name is " + p.name + " age " + strconv.Itoa(p.age)
} //strconv.Itoa数字转字符串

func (p *Person) hasbirthirty() {
	p.age++
}

func (p *Person) getcity(city string) {
	if p.city == city {
		p.city = "quanzhou"
	} else {
		return
	}
}

func main() {
	person1 := Person{name: "lin", city: "fuzhou", age: 25}
	person1.getcity("fuzhou")
	fmt.Println(person1)
	// person2 := Person{name: "wang", city: "quanzhou", age: 25}
	// fmt.Println(person2)
	person1.age--
	fmt.Println(person1.age)
	person1.hasbirthirty()
	value := person1.greet()
	fmt.Println(value)
}

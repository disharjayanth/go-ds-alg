package main

import "fmt"

type person struct {
	name string
}

func (p *person) printName() {
	fmt.Println("Name:", p.name)
}

func main() {
	p := person{
		name: "Smith",
	}

	p.printName()
}

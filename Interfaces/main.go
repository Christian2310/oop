package main

import "fmt"

type Greeter interface {
	Greet()
}

type Person struct {
	Name string
}

func (p Person) Greet() {
	fmt.Printf("Hello I'm %s", p.Name)
}

type Text string

func (t Text) Greet() {
	fmt.Printf("Hello I'm %s", t)
}

func main() {
	var g Greeter = Person{Name: "Christian"}
	g.Greet()
}

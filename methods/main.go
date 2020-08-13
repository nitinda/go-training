package main

import "fmt"

func main() {
	fmt.Println("---")
	gg := greeter{
		greeting: "Hello Method",
		name:     "Why",
	}

	gg.greet()

	f := greeter{
		greeting: "Second Greeter",
		name:     "Who",
	}

	f.gree()
	fmt.Println(f.name)

}

type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}

func (g *greeter) gree() {
	fmt.Println(g.greeting, g.name)
	g.name = "Noooo"
}

package main

import "fmt"

type Saiyan struct {
	*Person
	Power int
}

type Person struct {
	Name string
}

func main() {
	goku := NewSaiyan(&Person{"ND"}, 9090)
	Super(goku)

	fmt.Println(goku.Person)

	goku.Introduce()
}

func Super(s *Saiyan) {
	// s = &Saiyan{"AASSDF", 121212}
	s.Power += 111111
}

func (p *Person) Introduce() {
	fmt.Printf("Hi, I'm %s\n", p.Name)
}

func NewSaiyan(name *Person, power int) *Saiyan {
	return &Saiyan{
		Person: name,
		Power:  power,
	}
}

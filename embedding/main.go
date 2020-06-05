package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	Name   string `required max:"100"`
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

func main() {
	fmt.Println("My favorite number is")

	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = true

	c := Bird{
		Animal: Animal{
			Name:   "Kimu",
			Origin: "China",
		},
		SpeedKPH: 90,
		CanFly:   false,
	}

	// B Brid
	fmt.Println(b)

	// c Bird
	fmt.Println(c)

	// Tags
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
}

package main

import (
	"fmt"
	"math/rand"
)

type Doctor struct {
	Number     int
	ActorName  string
	Companions []string
}

func main() {
	fmt.Println("My favorite number is", rand.Intn(1000))

	// Arrays
	var students [3]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "Lisa"
	students[1] = "Mona Lisa"
	students[2] = "No Way"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Number of Students: %v\n", len(students))

	// Array and Pointers
	array1 := [...]int{1, 2, 3, 4, 5}
	array2 := &array1 // Referance passing
	array2[1] = 777
	fmt.Printf("Integer Array Elements are: %v\n", array1)

	// Slices
	slice1 := []int{11, 22, 33}
	slice2 := slice1 // Referance passing
	slice2[1] = 7777
	fmt.Printf("Slice1 Elements are: %v\n", slice1)
	fmt.Printf("Slice2 Elements are: %v\n", slice2)
	fmt.Printf("Number of elements of Slice1 : %v\n", len(slice1))

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:]   // slice all elements
	c := a[3:]  // slice from 4th element to all
	d := a[:6]  // slice from 1 to 6
	e := a[3:6] // slice from 4 to 6
	a[5] = 9000
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	// slice with builtin function make
	slice3 := []int{}
	fmt.Println(slice3)
	fmt.Printf("Length of slice3: %v\n", len(slice3))
	fmt.Printf("Capacity of slice3: %v\n", cap(slice3))
	slice3 = append(slice3, 1, 2)
	fmt.Println(slice3)
	fmt.Printf("Length of slice3: %v\n", len(slice3))
	fmt.Printf("Capacity of slice3: %v\n", cap(slice3))

	printSlice(slice3)

	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)

	// Maps
	// m := map[[4]int]int{}
	statePopulation := make(map[string]int)
	statePopulation = map[string]int{
		"California": 141123123,
		"Texas":      56578658758,
		"Ohio":       23412341,
		"New York":   232323232,
	}
	fmt.Println(statePopulation)
	fmt.Printf("Length of map %v\n", len(statePopulation))
	statePopulation["Georgia"] = 12312312321
	fmt.Println(statePopulation)
	fmt.Printf("Length of map %v\n", len(statePopulation))
	delete(statePopulation, "Ohio")
	fmt.Println(statePopulation)
	fmt.Printf("Length of map %v\n", len(statePopulation))

	p, ok := statePopulation["NoState"]
	fmt.Printf("Check key state %v 	%v\n", p, ok)

	// Structs
	aDoctor := Doctor{
		Number:    1,
		ActorName: "Kuttus",
		Companions: []string{
			"Puttus",
			"ABC Kids Tv",
			"ChuChu Tv",
		},
	}

	bDoctor := struct{ name string }{name: "ChuChu and ABC Kids TV"}

	fmt.Println("============= Struct ====================")
	fmt.Println(aDoctor.Companions[1])
	fmt.Println(bDoctor)

	cDoctor := bDoctor
	cDoctor.name = "Kuttus TV"
	fmt.Println(cDoctor)

	dDoctor := &bDoctor
	dDoctor.name = "Why"
	fmt.Println(bDoctor)
	fmt.Println(dDoctor)

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

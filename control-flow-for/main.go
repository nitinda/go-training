package main

import "fmt"

func main() {
	fmt.Println("============== Simple For Loop ==============")
	for i := 0; i < 5; i++ {
		fmt.Printf("For Loop: %v\n", i)
	}
	fmt.Println("===========================")
	for i := 0; i < 10; i = i + 2 {
		fmt.Printf("For Loop: %v\n", i)
	}
	fmt.Println("===========================")
	i := 0
	for i < 10 {
		fmt.Printf("For Loop: %v\n", i)
		i++
	}
	fmt.Println("============== Ranged For Loop ==============")

	// Slice
	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v)
	}
	fmt.Println("===========================")

	// Array
	a := [3]int{1, 2, 3}
	for k, v := range a {
		fmt.Println(k, v)
	}

	fmt.Println("==========================")

	// Map
	m := map[string]int{
		"California": 141123123,
		"Texas":      56578658758,
		"Ohio":       23412341,
		"New York":   232323232,
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
	fmt.Println("==========================")

	for k := range m {
		fmt.Println(k)
	}
	fmt.Println("==========================")
	for _, v := range m {
		fmt.Println(v)
	}

	str := "Hello World"
	for k, v := range str {
		fmt.Println(k, string(v))
	}

}

package main

import "fmt"

func main() {
	var a int = 32
	var b *int = &a
	fmt.Println(a, *b)
	a = 13
	fmt.Println(a, *b)

	var ms *myStruct
	ms = new(myStruct)
	fmt.Println((*ms).foo)

	(*ms).foo = 2000
	fmt.Println((*ms).foo)

	// User friendly
	fmt.Println(ms.foo)
}

type myStruct struct {
	foo int
}

package main

import (
	"fmt"
)

func main() {
	fmt.Println("====================== Control Flow if ======================")

	if true {
		fmt.Println("This test is true")
	}

	if 2 >= 2 {
		fmt.Println("This test : 2 >=2")
	}

	if 2 != 3 {
		fmt.Println("This test : 2 != 2")
	}

	if 2 > 33 {
		fmt.Println("This test is if")
	} else {
		fmt.Println("This test is else")
	}
}

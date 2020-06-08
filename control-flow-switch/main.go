package main

import (
	"fmt"
)

func main() {
	fmt.Println("=================== Switch ====================")
	i := 10
	switch i {
	case 1, 3, 5, 7, 9:
		fmt.Println("Case ---- 1")
	case 2, 4, 6, 8, 10:
		fmt.Println("Case ---- 2")
	case 11, 13, 15, 17, 19:
		fmt.Println("Case ---- 3")
	default:
		fmt.Println("No matching case ... !")
	}

	fmt.Println("=================== Switch with Interface ====================")

	var j interface{} = 1.0
	switch j.(type) {
	case int:
		fmt.Println("Variable is type Interger")
	case string:
		fmt.Println("Varriable is type String")
	case float64:
		fmt.Println("Varriable is type Float64")
	case float32:
		fmt.Println("Varriable is type Float32")
	default:
		fmt.Println("Unknown Type")
	}

}

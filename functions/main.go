package main

import "fmt"

func main() {
	sayMessage("Hello World", 2)

	s := sum(1, 2, 3, 4, 5)
	fmt.Println("The Sum is : ", s)
	s = sum2(1, 2, 3, 4, 5, 6, 7, 8, 9)

	fmt.Println("The Sum is : ", s)

	d, err := divide(5.0, 10.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)

	// Anonymous function
	func() {
		msg := "Anonymous function"
		fmt.Println(msg)
	}()
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot process zero")
	}

	return a / b, nil
}

func sayMessage(msg string, idx int) {
	fmt.Println(msg)
	fmt.Println("Number is : ", idx)
}

func sum(values ...int) int {
	result := 0
	for _, v := range values {
		result += v
	}
	return result
}

func sum2(values ...int) (result int) {
	for _, v := range values {
		result += v
	}
	return
}

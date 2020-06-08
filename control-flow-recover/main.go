package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Staring Function")
	panicker()
	fmt.Println("End")

}

func panicker() {
	fmt.Println("About to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}()
	panic("Something bad happened")
	fmt.Println("Done panicking !!!")
}

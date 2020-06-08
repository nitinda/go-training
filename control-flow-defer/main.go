package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("==================== Start ====================")

	// Exeute after the main function
	defer fmt.Println("==================== Middle ====================")

	fmt.Println("==================== End ====================")

	res, err := http.Get("http://www.google.co.uk/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", robots)
}

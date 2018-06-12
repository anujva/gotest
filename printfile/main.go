package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
func main() {
	filePath := os.Args[1]
	b, err := ioutil.ReadFile(filePath)
	check(err)
	fmt.Println(string(b))
}

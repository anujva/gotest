package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("http://www.google.com")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	google, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Println(string(google))
}

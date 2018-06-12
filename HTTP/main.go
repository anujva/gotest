package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

//StdOutWriter writes to the stdout
type StdOutWriter struct{}

func (stdOutWriter StdOutWriter) Write(bs []byte) (n int, err error) {
	fmt.Println(string(bs))
	return len(bs), nil
}

func main() {
	res, err := http.Get("http://www.google.com")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	// google, err := ioutil.ReadAll(res.Body)
	// res.Body.Close()
	stdOutWriter := StdOutWriter{}
	io.Copy(stdOutWriter, res.Body)
}

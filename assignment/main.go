package main

import (
	"fmt"
	"strconv"
)

func main() {
	sliceOfInts := []int{}
	for i := 0; i <= 10; i++ {
		sliceOfInts = append(sliceOfInts, i)
	}

	for _, val := range sliceOfInts {
		if val%2 == 0 {
			fmt.Println(strconv.Itoa(val) + " is even")
		} else {
			fmt.Println(strconv.Itoa(val) + " is odd")
		}
	}
}

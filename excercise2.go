package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := os.Args[1]
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Not a number:", input)
		os.Exit(1)
	}

	if num%2 == 0 {
		fmt.Println("Even Number:", num)
	} else {
		fmt.Println("Odd Number:", num)
	}

}

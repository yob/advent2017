package main

import (
	"flag"
	"fmt"
	"log"
	"github.com/yob/advent2017/three"
	"strconv"
)

// application entry point
func main() {
	flag.Parse()
	input := flag.Arg(0)

	if input == "" {
		log.Fatal("USAGE: three <input string>")
	}

	number, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal("USAGE: three <input string>")
	}
	fmt.Printf("result: %d\n", three.Distance(number))
}

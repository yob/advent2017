package main

import (
	"flag"
	"fmt"
	"log"
	"github.com/yob/advent2017/one"
)

// application entry point
func main() {
	flag.Parse()
	input := flag.Arg(0)

	if input == "" {
		log.Fatal("USAGE: one <input string>")
	}

	fmt.Printf("result: %d\n", one.Sum(input))
}

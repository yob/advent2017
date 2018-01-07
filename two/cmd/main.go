package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"github.com/yob/advent2017/two"
	"os"
)

// application entry point
func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		log.Fatal("USAGE: cat input.txt | two")
	}
	input := string(bytes)

	fmt.Printf("result: %d\n", two.Checksum(input))
}

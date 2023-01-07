package main

import (
	"fmt"
	"zeenisionlang/packages/parser"
)

func main() {
	if parser.IsLetter('B') {
		fmt.Println("IsLetter work")
	}
}

package main

import (
	"fmt"
	"zeenisionlang/packages/parser"
)

func main() {
	source := "let a = 100\n"
	source += "func test() => println(\"test\")\n"
	source += "blueprint ProductBase {}\n"
	tokens := parser.Lexer(source)
	for _, tokenInLine := range tokens {
		fmt.Println(tokenInLine)
		// for _, token := range tokenInLine {
		// 	fmt.Println(token)
		// }
	}
}

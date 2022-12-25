package parser

import (
	"bufio"
	"regexp"
	"strings"
)

func getWordFromText(text string) []string {
	text += " "
	words := []string{}
	word := ""
	for _, ch := range text {
		if string(ch) != " " {
			word += string(ch)
		} else {
			word = strings.ReplaceAll(word, " ", "")
			if word != "" {
				words = append(words, word)
			}
			word = ""
		}
	}
	return words
}

func isNumber(text string) bool {
	text = strings.ReplaceAll(text, ",", "")
	text = strings.ReplaceAll(text, ".", "")
	var re *regexp.Regexp = regexp.MustCompile("[0-9]+")
	if len(re.FindAllString(text, -1)) > 0 {
		return true
	} else {
		return false
	}
}

func isString(text string) bool {
	runes := []rune(text)
	return string(runes[0]) == "\"" && string(runes[len(runes)-1]) == "\""
}

//lint:ignore U1000 for use in comming soon
func isIdentifier(text string) bool {
	var re *regexp.Regexp = regexp.MustCompile("[a-zA-Z_][a-zA-Z0-9_]")
	if len(re.FindAllString(text, -1)) > 0 {
		return true
	} else {
		return false
	}
}

func isBool(text string) bool {
	return text == "true" || text == "false"
}

func getTokenType(text string) TokenType {
	if isNumber(text) {
		return NUMBER_TOKEN
	} else if isString(text) {
		return STRING_TOKEN
	} else if isIdentifier(text) {
		return IDENTIFIER_TOKEN
	} else if isBool(text) {
		return BOOL_TOKEN
	} else {
		return UNKNOWN_TOKEN
	}
}

func Lexer(source string) [][]Token {
	tokenList := [][]Token{}
	scanner := bufio.NewScanner(strings.NewReader(source))
	for scanner.Scan() {
		tokens := []Token{}
		words := getWordFromText(scanner.Text())
		for _, word := range words {
			tokens = append(tokens, Token{getTokenType(word), word})
		}
		tokenList = append(tokenList, tokens)
	}
	return tokenList
}

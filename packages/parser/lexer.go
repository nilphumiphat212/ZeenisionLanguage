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

func isKeyword(text string) (bool, KeywordType) {
	keywordMap := make(map[string]KeywordType)
	keywordMap["global"] = GLOBAL_KEYWORD
	keywordMap["let"] = LET_KEYWORD
	keywordMap["const"] = CONST_KEYWORD
	keywordMap["if"] = IF_KEYWORD
	keywordMap["else"] = ELSE_KEYWORD
	keywordMap["elseif"] = ELSEIF_KEYWORD
	keywordMap["switch"] = SWITCH_KEYWORD
	keywordMap["loop"] = LOOP_KEYWORD
	keywordMap["continue"] = LOOP_KEYWORD
	keywordMap["return"] = RETURN_KEYWORD
	keywordMap["object"] = OBJECT_KEYWORD
	keywordMap["import"] = IMPORT_KEYWORD
	keywordMap["func"] = FUNCTION_KEYWORD

	if keywordMatch, match := keywordMap[text]; match {
		return true, keywordMatch
	}

	return false, UNKNOWN_KEYWORD
}

func getTokenType(text string) (TokenType, TokenValue) {
	staticMap := make(map[string]TokenType)
	staticMap["+"] = ADD_TOKEN
	staticMap["-"] = MINUS_TOKEN
	staticMap["*"] = MULTIPLY_TOKEN
	staticMap["/"] = DIVIDE_TOKEN

	if tokenMatch, match := staticMap[text]; match {
		return tokenMatch, nil
	}

	if isNumber(text) {
		return NUMBER_TOKEN, nil
	} else if isString(text) {
		return STRING_TOKEN, nil
	} else if isIdentifier(text) {
		return IDENTIFIER_TOKEN, nil
	} else if isBool(text) {
		return BOOL_TOKEN, nil
	} else if isKeyword, keywordType := isKeyword(text); isKeyword {
		return KEYWORD_TOKEN, keywordType
	} else {
		return UNKNOWN_TOKEN, nil
	}
}

func Lexer(source string) [][]Token {
	tokenList := [][]Token{}
	scanner := bufio.NewScanner(strings.NewReader(source))
	for scanner.Scan() {
		tokens := []Token{}
		words := getWordFromText(scanner.Text())
		for _, word := range words {
			var value TokenValue = nil
			tokenType, tokenValue := getTokenType(word)

			if tokenType == KEYWORD_TOKEN {
				value = tokenValue
			} else {
				value = word
			}

			tokens = append(tokens, Token{tokenType, value})
		}
		tokenList = append(tokenList, tokens)
	}
	return tokenList
}

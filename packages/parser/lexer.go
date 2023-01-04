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

func hasLeftParentheses(tokenValue TokenValue) bool {
	runes := []rune(tokenValue.(string))
	return string(runes[0]) == "("
}

func hasRightParentheses(tokenValue TokenValue) bool {
	runes := []rune(tokenValue.(string))
	return string(runes[len(runes)-1]) == ")"
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
	keywordMap["break"] = BREAK_KEYWORD
	keywordMap["continue"] = CONTINUE_KEYWORD
	keywordMap["return"] = RETURN_KEYWORD
	keywordMap["object"] = OBJECT_KEYWORD
	keywordMap["import"] = IMPORT_KEYWORD
	keywordMap["func"] = FUNCTION_KEYWORD
	keywordMap["namespace"] = NAMESPACE_KEYWORD
	keywordMap["using"] = USING_KEYWORD
	keywordMap["class"] = CLASS_KEYWORD
	keywordMap["model"] = MODEL_CLASS_KEYWORD
	keywordMap["db"] = DB_CLASS_KEYWORD
	keywordMap["table"] = TABLE_CLASS_KEYWORD
	keywordMap["primary"] = PRIMARY_KEY_KEYWORD
	keywordMap["foriegn"] = FOREIGN_KEY_KEYWORD
	keywordMap["blueprint"] = BLUEPRINT_KEYWORD
	keywordMap["enum"] = ENUM_KEYWORD
	keywordMap["public"] = PUBLIC_KEYWORD
	keywordMap["private"] = PRIVATE_KEYWORD
	keywordMap["protected"] = PROTECTED_KEYWORD
	keywordMap["abstract"] = ABSTRACT_KEYWORD
	keywordMap["late"] = LATE_KEYWORD
	keywordMap["virtual"] = VIRTUAL_KEYWORD
	keywordMap["internal"] = INTERNAL_KEYWORD
	keywordMap["ref"] = REF_KEYWORD
	keywordMap["embed"] = EMBED_KEYWORD
	keywordMap["extension"] = EXTENSION_KEYWORD
	keywordMap["observe"] = OBSERVE_KEYWORD
	keywordMap["ui"] = UI_KEYWORD

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
	} else if isKeyword, keywordType := isKeyword(text); isKeyword {
		return KEYWORD_TOKEN, keywordType
	} else if isBool(text) {
		return BOOL_TOKEN, nil
	} else if isString(text) {
		return STRING_TOKEN, nil
	} else if isIdentifier(text) {
		return IDENTIFIER_TOKEN, nil
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

			if hasLeftParentheses(value) {
				tokens = append(tokens, Token{LEFT_PARENTHESES_TOKEN, "("})
			}

			tokens = append(tokens, Token{tokenType, value})

			if hasRightParentheses(value) {
				tokens = append(tokens, Token{RIGHT_PARENTHESES_TOEKN, ")"})
			}
		}
		tokenList = append(tokenList, tokens)
	}
	return tokenList
}

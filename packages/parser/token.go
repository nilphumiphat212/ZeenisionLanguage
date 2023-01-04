package parser

type TokenType int
type TokenValue interface{}

const (
	UNKNOWN_TOKEN TokenType = iota
	STRING_TOKEN
	NUMBER_TOKEN
	BOOL_TOKEN
	ADD_TOKEN
	MINUS_TOKEN
	MULTIPLY_TOKEN
	DIVIDE_TOKEN
	KEYWORD_TOKEN
	IDENTIFIER_TOKEN
	LEFT_PARENTHESES_TOKEN
	RIGHT_PARENTHESES_TOEKN
)

type Token struct {
	Type  TokenType
	Value TokenValue
}

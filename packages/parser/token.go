package parser

type TokenType = int
type TokenValue = interface{}

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
)

type Token struct {
	Type  TokenType
	Value TokenValue
}

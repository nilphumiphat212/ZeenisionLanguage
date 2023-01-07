package parser

type Lexer struct {
	source       string
	position     int
	readPosition int
	ch           byte
}

func New(source string) *Lexer {
	lexer := &Lexer{source: source}
	lexer.readChar()
	return lexer
}

func (lexer *Lexer) readChar() {
	if lexer.readPosition >= len(lexer.source) {
		lexer.ch = 0
	} else {
		lexer.ch = lexer.source[lexer.readPosition]
	}
	lexer.position = lexer.readPosition
	lexer.readPosition++
}

//lint:ignore U1000 use in comming soon.
func newToken(tokenType TokenType, tokenValue TokenValue) Token {
	return Token{tokenType, tokenValue}
}

//export for debug
func IsLetter(ch byte) bool {
	return isLetter(ch)
}

func IsDigit(ch byte) bool {
	return isDigit(ch)
}

func isLetter(ch byte) bool {
	return ('a' <= ch && 'z' <= ch) || ('A' <= ch && 'Z' <= ch || ch == '_')
}

func isDigit(ch byte) bool {
	return '0' <= ch && '9' <= ch
}

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

func (lexer *Lexer) NextToken() Token {
	lexer.skipWhitespace()

	return newToken(UNKNOWN_TOKEN, nil)
}

func (lexer *Lexer) skipWhitespace() {
	for lexer.ch == ' ' || lexer.ch == '\t' || lexer.ch == '\r' || lexer.ch == '\n' {
		lexer.readChar()
	}
}

//lint:ignore U1000 use in comming soon.
func newToken(tokenType TokenType, tokenValue TokenValue) Token {
	return Token{tokenType, tokenValue}
}

//lint:ignore U1000 use in comming soon.
func (lexer *Lexer) readIdentifier() string {
	position := lexer.position
	for isLetter(lexer.ch) {
		lexer.readChar()
	}
	return lexer.source[position:lexer.position]
}

//lint:ignore U1000 use in comming soon.
func (lexer *Lexer) readNumber() string {
	position := lexer.position
	for isDigit(lexer.ch) {
		lexer.readChar()
	}
	return lexer.source[position:lexer.position]
}

//export for debugfor
func IsLetter(ch byte) bool {
	return isLetter(ch)
}

func IsDigit(ch byte) bool {
	return isDigit(ch)
}

//end export

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

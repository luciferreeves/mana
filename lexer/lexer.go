package lexer
import "mana/tokens"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

// New returns a new Lexer instance.
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// NextToken returns the next token in the input string.
func (l *Lexer) NextToken() tokens.Token {
	var tok tokens.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = tokens.Token{Type: tokens.EQ, Literal: literal}
		} else {
			tok = newToken(tokens.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(tokens.PLUS, l.ch)
	case '-':
		tok = newToken(tokens.MINUS, l.ch)
	case '/':
		tok = newToken(tokens.SLASH, l.ch)
	case '*':
		tok = newToken(tokens.ASTERISK, l.ch)
	case '<':
		tok = newToken(tokens.LT, l.ch)
	case '>':
		tok = newToken(tokens.GT, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = tokens.Token{Type: tokens.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(tokens.BANG, l.ch)
		}
	case ';':
		tok = newToken(tokens.SEMICOLON, l.ch)
	case '(':
		tok = newToken(tokens.LPAREN, l.ch)
	case ')':
		tok = newToken(tokens.RPAREN, l.ch)
	case ',':
		tok = newToken(tokens.COMMA, l.ch)
	case '{':
		tok = newToken(tokens.LBRACE, l.ch)
	case '}':
		tok = newToken(tokens.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = tokens.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = tokens.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = tokens.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(tokens.ILLEGAL, l.ch)
		}
	}
	
	l.readChar()
	return tok
}

// skipWhitespace skips whitespace characters.
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// newToken returns a new Token instance.
func newToken(tokenType tokens.TokenType, ch byte) tokens.Token {
	return tokens.Token{Type: tokenType, Literal: string(ch)}
}

// isLetter returns true if the given character is a letter.
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// isDigit returns true if the given character is a digit.
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// readChar reads the next character in the input and advances the position in the input string.
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII code for "NUL" character
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// peekChar returns the next character in the input string without advancing the position in the input string.
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0 // ASCII code for "NUL" character
	} else {
		return l.input[l.readPosition]
	}
}

// readIdentifier reads an identifier and advances the position in the input string until it encounters a non-letter character.
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// readNumber reads a number and advances the position in the input string until it encounters a non-digit character.
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

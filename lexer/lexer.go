package lexer

import (
	"monkey/token"
)

type Lexer struct {
	input string
	curr  int  // points to current char
	next  int  // points to next char
	ch    byte // current char
}

func New(input string) *Lexer {
	l := Lexer{
		input: input,
	}
	l.readChar()
	return &l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Kind = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Kind = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Kind = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readChar() {
	if l.next >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.next]
	}
	l.curr = l.next
	l.next += 1
}

func (l *Lexer) readIdentifier() string {
	curr := l.curr
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[curr:l.curr]
}

func (l *Lexer) readNumber() string {
	curr := l.curr
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[curr:l.curr]
}

func newToken(kind string, ch byte) token.Token {
	return token.Token{kind, string(ch)}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

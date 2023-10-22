package lexer

import (
	"monkey/token"
)

type Lexer struct {
	input        string
	position     int // current position in input
	readPosition int // current reading position in input
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() (tok token.Token) {
	switch l.ch {
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok = newToken(token.EOF, l.ch)
		tok.Literal = ""
	}
	l.readChar()
	return
}

func newToken(tp token.TokenType, ch byte) token.Token {
	return token.Token{Type: tp, Literal: string(ch)}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
		l.position = l.readPosition
		l.readPosition++
	}
}

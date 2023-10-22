package lexer

import (
	"monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := "+=(){},;"
	testCases := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{
			expectedType:    token.PLUS,
			expectedLiteral: "+",
		},
		{
			expectedType:    token.ASSIGN,
			expectedLiteral: "=",
		},
		{
			expectedType:    token.LPAREN,
			expectedLiteral: "(",
		},
		{
			expectedType:    token.RPAREN,
			expectedLiteral: ")",
		},
		{
			expectedType:    token.LBRACE,
			expectedLiteral: "{",
		},
		{
			expectedType:    token.RBRACE,
			expectedLiteral: "}",
		},
		{
			expectedType:    token.COMMA,
			expectedLiteral: ",",
		},
		{
			expectedType:    token.SEMICOLON,
			expectedLiteral: ";",
		},
		{
			expectedType:    token.EOF,
			expectedLiteral: "",
		},
	}
	l := New(input)
	for i, tC := range testCases {
		t.Run("simple", func(t *testing.T) {
			tok := l.NextToken()
			if tok.Type != tC.expectedType {
				t.Fatalf("test[%d] - tokenType wrong. expected=%q, got=%q",
					i, tC.expectedType, tok.Type,
				)
			}
			if tok.Literal != tC.expectedLiteral {
				t.Fatalf("test[%d] - Literal wrong. expected=%q, got=%q",
					i, tC.expectedLiteral, tok.Literal,
				)
			}
		})
	}
}

func TestNextToken2(t *testing.T) {
	input := `let five = 5;
let ten = 10;
let add = fn(x, y) {
	x + y;
};
let result = add(five, ten);`
	testCases := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
	}
	l := New(input)
	for i, tC := range testCases {
		t.Run("sentence", func(t *testing.T) {
			tok := l.NextToken()
			if tok.Type != tC.expectedType {
				t.Fatalf("test[%d] - tokenType wrong. expected=%q, got=%q",
					i, tC.expectedType, tok.Type,
				)
			}
			if tok.Literal != tC.expectedLiteral {
				t.Fatalf("test[%d] - Literal wrong. expected=%q, got=%q",
					i, tC.expectedLiteral, tok.Literal,
				)
			}
		})
	}
}

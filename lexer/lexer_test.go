package lexer

import (
	"go_int/token"
	"testing"
)

func TestNextToken(t *Testing.T) {
	input := `;,+=({})`

	tests := []struct {
		expectedType token.TokenType
		expectedLiteral string
	} {
		{token.SEMICOLON, ";"},
		{token.COMMA, ","},
		{token.PLUS, "+"},
		{token.ASSIGN, "="},
		{token.LPAREN, "("},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.RPAREN, ")"}
		{token.EOF, ""},
	}

	l := New(input)
	
	for i, tt := range tests {
		tok := l.NextToken()
	}

	if tok.Type != tt.expectedType {
	
		t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, 
					tt.expectedType, tok.Type)
	}

	if tok.Literal != tt.expectedLiteral {

		t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i,
					tt.expectedLiteral, tok.Literal)

	}
}
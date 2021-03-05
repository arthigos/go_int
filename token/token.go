package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	EOF       = "EOF"
	IDENT     = "IDENT"
	INT       = "INT"
	ASSING    = "="
	PLUS      = "+"
	ILLEGAL   = "ILLEGAL"
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	FUNCTION  = "FUNCTION"
	LET       = "LET"
)

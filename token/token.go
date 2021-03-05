package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	EOF          = "EOF"
	IDENTIFIER   = "IDENTIFIER"
	INT          = "INT"
	ASSING       = "="
	PLUS         = "+"
	ILLEGAL      = "ILLEGAL"
	COMMA        = ","
	SEMICOLON    = ";"
	LPARENTHESIS = "("
	RPARENTHESIS = ")"
	LCURLY       = "{"
	RCURLY       = "}"
	FUNCTION     = "FUNCTION"
	LET          = "LET"
)

package token

import "fmt"

type TokenType int

type Token struct {
	Type		TokenType
	Lexeme		string
	Start		int
	End			int
	Line		int
	Column		int
	FilePath	*string
}

func (t *Token) String() string {
	filepath := ""
	if (t.FilePath != nil) {
		filepath = *t.FilePath
	}
	return fmt.Sprintf(
		"%d %q %d..%d (%s:%d:%d)",
		t.Type,
		t.Lexeme,
		t.Start, t.Column,
		filepath,
		t.Line, t.Column,
	)
}

const (
	ILLEGAL TokenType = iota
	EOI

	INT
	STRING
	IDENTIFIER

	PLUS
	MINUS
	SLASH
	STAR
	LESS
	GREATER
	LESS_EQUAL
	GREATER_EQUAL
	EQUAL

	ASSIGN

	VAR
	FN
	IF
	ELSE
	LOOP
	RETURN

	LPAREN
	RPAREN
	LBRACE
	RBRACE
	LBRACK
	RBRACK
)
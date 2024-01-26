package token

import (
	"fmt"
)

type TokenType int

type Token struct {
	Type		TokenType	`json:"type"`
	Lexeme		string		`json:"lexeme"`
	Start		int			`json:"start"`
	End			int			`json:"end"`
	Line		int			`json:"line"`
	Column		int			`json:"column"`
	FilePath	*string		`json:"filepath"`
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
		t.Start, t.End,
		filepath,
		t.Line, t.Column,
	)
}

const (
	_ TokenType = iota
	ILLEGAL
	EOI

	INT
	STRING
	CHAR
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
	DECLARE

	FN
	IF
	ELSE
	LOOP
	RETURN
	PRINTLN

	LPAREN
	RPAREN
	LBRACE
	RBRACE
	LBRACK
	RBRACK
	SEMICOLON
	COMMA
	COLON

	// MUST BE FINAL VALUE
	MAX_TYPE 
)


var keywords = map[string] TokenType {
	"println": 	PRINTLN,
	"fn": 		FN,
	"if": 		IF,
	"else": 	ELSE,
	"loop": 	LOOP,
	"return": 	RETURN,
} 

func LookupKeyword(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENTIFIER
}


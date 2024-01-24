package main

import (
	"fmt"
	"tinylang/pkg/token"
)

func main() {
	filepath := "Hello.tl"
	tok := token.Token{
		Type: token.ASSIGN,
		Lexeme: "=",
		Start: 0,
		End: 1,
		Line: 1,
		Column: 1,
		FilePath: &filepath,
	}

	fmt.Println(tok.String())
}
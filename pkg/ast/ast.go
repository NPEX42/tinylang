package ast

import (
	"bytes"
	"tinylang/pkg/token"
)

type Node interface {
	String() string
}

type Statement interface {
	Node
	stmt()
}

type Expression interface {
	Node
	expr()
}

type Program struct {
	Statements []Statement
}

type PrintStmt struct {
	Token token.Token 	`json:"token"`
	Value Expression	`json:"value"`
}

func (stmt *PrintStmt) stmt() {}
func (stmt *PrintStmt) String() string { 
	var s bytes.Buffer

	s.WriteString("println")

	if (stmt.Value != nil) {
		s.WriteString(" " + stmt.Value.String())
	}

	s.WriteString(";")

	return s.String()
}




package ast

import "tinylang/pkg/token"

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

type PrintStmt struct {
	Token token.Token 	`json:"token"`
	Value Expression	`json:"value"`
}




package parser

import (
	"tinylang/pkg/lexer"
	"tinylang/pkg/token"
)

type Parser struct {
	CurrToken	token.Token
	NextToken	token.Token

	l			*lexer.Lexer
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l: l,
	}

	p.pop()
	p.pop()

	return p
}

// func (p *Parser) Program() ast.Program {
// 	for (p.l.HasNext()) {

// 	}
// }

func (p *Parser) matchCurrent(tt token.TokenType) bool {
	return p.CurrToken.Type == tt;
}

func (p *Parser) matchNext(tt token.TokenType) bool {
	return p.NextToken.Type == tt;
}

func (p *Parser) pop() {
	p.CurrToken = p.NextToken
	p.NextToken = p.l.NextToken()
}


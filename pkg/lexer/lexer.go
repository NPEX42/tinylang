package lexer

import (
	"tinylang/pkg/token"
	"unicode"
)

type Lexer struct {
	start		int
	current 	int
	source 		string
	line 		int
	column		int
	filename 	*string
}

func New(source string, filename *string) *Lexer {
	l := &Lexer{
		source: 	source,
		start: 		0,
		current: 	0,
		line: 		1,
		column: 	1,
		filename: 	filename,
	}

	return l
}

func (l *Lexer) NextToken() token.Token {
	l.skipWhitespace()
	l.start = l.current
	var tok token.Token
	c := l.pop()
	switch c {
	case ';':
		tok = l.newToken(token.SEMICOLON)
	case '=':
		if (l.match('=')) { 
			tok = l.newToken(token.EQUAL)
		} else {
			tok = l.newToken(token.ASSIGN)
		}
	case '+':
		tok = l.newToken(token.PLUS)
	case '-':
		tok = l.newToken(token.MINUS)
	case '/':
		if (l.match('/')) { 
			l.skipComment()
		} else {
			tok = l.newToken(token.SLASH)
		}
	case '"':
		tok = l.stringLiteral()
	case '*':
		tok = l.newToken(token.STAR)
	case '{':
		tok = l.newToken(token.LBRACE)
	case '}':
		tok = l.newToken(token.RBRACE)
	case '(':
		tok = l.newToken(token.LPAREN)
	case ')':
		tok = l.newToken(token.RPAREN)
	case '[':
		tok = l.newToken(token.LBRACK)
	case ']':
		tok = l.newToken(token.RBRACK)
	case 0:
		tok = l.newToken(token.EOI)
	case '\'':
		l.pop()
		l.pop()
		tok = l.newToken(token.CHAR)

	case '>':
		if l.match('=') {
			tok = l.newToken(token.GREATER_EQUAL) 
		} else {
			tok = l.newToken(token.GREATER)
		}

	case '<':
		if l.match('=') {
			tok = l.newToken(token.LESS_EQUAL) 
		} else {
			tok = l.newToken(token.LESS)
		}

	case ':':
		if l.match('=') {
			tok = l.newToken(token.DECLARE)
		} else {
			tok = l.newToken(token.ILLEGAL)
		}

	default:
		if (unicode.IsLetter(rune(c))) {
			tok = l.identifier()
			return tok;
		}

		if (unicode.IsDigit(rune(c))) {
			for unicode.IsDigit(rune(l.peek())) {
				l.pop()
			}
			tok = l.newToken(token.INT)
			return tok;
		}

		tok = l.newToken(token.ILLEGAL)
	}
	return tok
}

func (l *Lexer) HasNext() bool {
	return l.current < len(l.source)
}

func (l *Lexer) newToken(Type token.TokenType) token.Token {
	return token.Token{
		Type: Type,
		Lexeme: l.lexeme(),
		Start: l.start,
		End: l.current,
		Line: l.line,
		Column: l.column - len(l.lexeme()), 
		FilePath: l.filename,
	}
}

func (l *Lexer) lexeme() string {
	return l.source[l.start:l.current]
}

func (l *Lexer) skipComment() {
	for l.peek() != '\n' && l.HasNext() {
		l.pop()
	}
}

func (l *Lexer) skipWhitespace() {
	for (l.peek() == ' ' || l.peek() == '\t' || l.peek() == '\r' || l.peek() == '\n') && l.HasNext() {
		if l.peek() == '\n' {
			l.line++
			l.column = 1
		}

		l.pop()
	}
}

func (l *Lexer) stringLiteral() token.Token {
	for !l.match('"') && l.HasNext() {
		if l.peek() == '\n' {
			l.line++
			l.column = 1
		}

		l.pop()
	}

	return token.Token {
		Type: token.STRING,
		Lexeme: l.source[l.start+1:l.current-1],
		Start: l.start,
		End: l.current,
		Line: l.line,
		Column: l.column - len(l.lexeme()), 
		FilePath: l.filename,
	}
}

func (l *Lexer) identifier() token.Token {
	for unicode.IsLetter(rune(l.peek())) || unicode.IsDigit(rune(l.peek())) {
		l.pop()
	}
	return l.newToken(token.LookupKeyword(l.lexeme()))
}


func (l *Lexer) match(c byte) bool {
	if (l.peek() == c) {
		l.pop()
		return true
	}
	return false
}

func (l *Lexer) peek() byte {
	if (l.HasNext()) {
		c := l.source[l.current]
		return c
	} else {
		return 0
	}
}

func (l *Lexer) pop() byte {
	if (l.current >= len(l.source)) {
		return 0
	} else {
		c := l.source[l.current]
		l.current += 1
		l.column += 1
		return c
	}
}
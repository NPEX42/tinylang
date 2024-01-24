package lexer

type Lexer struct {
	start		int
	current 	int
	source 		string
	line 		int
	column		int
	filename 	*string
}

func New(source string) *Lexer {
	l := &Lexer{
		source: source,
		start: 0,
		current: 0,
		line: 0,
		column: 0,
		filename: nil,
	}

	return l
}
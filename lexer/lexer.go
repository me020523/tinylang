package lexer

import "io"

type Lexer struct {
	source io.Reader
}

func NewLexer(source io.Reader) *Lexer {
	return &Lexer{
		source,
	}
}

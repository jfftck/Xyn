package lexer

import (
	"fmt"
	"xyn/token"
)

type Lexer struct {
	input   []rune
	pos     int
	readPos int
	ch      rune
}

// New returns a new lexer for the given input string.
func New(input string) *Lexer {
	return &Lexer{input: []rune(input)}
}

// NextToken returns the next token from the input.
func (l *Lexer) NextToken() token.Token {
	l.skipWhitespace()
	l.readPos = l.pos
	l.readRune()
}

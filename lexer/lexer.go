package lexer

import "xyn/token"

type Lexer struct {
	input   []rune
	pos     int
	readPos int
	rn      rune
}

// New returns a new lexer for the given input string.
func New(input string) *Lexer {
	l := &Lexer{input: []rune(input)}

	l.readRune()

	return l
}
func (l *Lexer) readRune() {
	l.rn = l.peekRune()
	l.pos = l.readPos
	l.readPos++
}

// NextToken returns the next token from the input.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.rn {
	case '=':
		if l.peekRune() == '=' {
			rn := l.rn

			l.readRune()
			tok = l.newToken(token.EQUAL, string(rn)+string(l.rn))
		} else {
			tok = l.newToken(token.REASSIGN, string(l.rn))
		}
	case '!':
		if l.peekRune() == '=' {
			rn := l.rn

			l.readRune()
			tok = l.newToken(token.NOT_EQUAL, string(rn)+string(l.rn))
		} else {
			tok = l.newToken(token.BANG, string(l.rn))
		}
	case ':':
		if l.peekRune() == '=' {
			rn := l.rn

			l.readRune()
			tok = l.newToken(token.ASSIGN, string(rn)+string(l.rn))
		} else {
			tok = l.newToken(token.COLON, string(l.rn))
		}
	case ';':
		tok = l.newToken(token.SEMICOLON, string(l.rn))
	case '(':
		tok = l.newToken(token.LPAREN, string(l.rn))
	case ')':
		tok = l.newToken(token.RPAREN, string(l.rn))
	case ',':
		tok = l.newToken(token.COMMA, string(l.rn))
	case '+':
		tok = l.newToken(token.PLUS, string(l.rn))
	case '-':
		if l.peekRune() == '>' {
			rn := l.rn

			l.readRune()
			tok = l.newToken(token.RIGHT_SIGNAL, string(rn)+string(l.rn))
		} else {
			tok = l.newToken(token.MINUS, string(l.rn))
		}
	case '*':
		tok = l.newToken(token.ASTERISK, string(l.rn))
	case '/':
		tok = l.newToken(token.SLASH, string(l.rn))
	case '<':
		if l.peekRune() == '=' {
			rn := l.rn

			l.readRune()
			tok = l.newToken(token.LESS_EQUAL, string(rn)+string(l.rn))
		} else if l.peekRune() == '<' {
			rn := l.rn

			l.readRune()
			tok = l.newToken(token.LEFT_SHIFT, string(rn)+string(l.rn))
		} else if l.peekRune() == '-' {
			rn := l.rn

			l.readRune()
			if l.peekRune() == '>' {
				rn2 := l.rn

				l.readRune()
				tok = l.newToken(token.BIDIRECTIONAL_SIGNAL, string(rn)+string(rn2)+string(l.rn))
			} else {
				tok = l.newToken(token.LEFT_SIGNAL, string(rn)+string(l.rn))
			}
		} else if l.peekRune() == ':' {
			rn := l.rn

			l.readRune()
			if l.peekRune() == '=' {
				rn2 := l.rn

				l.readRune()
				tok = l.newToken(token.ASSIGN_SIGNAL, string(rn)+string(rn2)+string(l.rn))
			} else {
				tok = l.newToken(token.ILLEGAL, string(rn)+string(l.rn))
			}
		} else {
			tok = l.newToken(token.LESS, string(l.rn))
		}
	case '>':
		if l.peekRune() == '=' {
			rn := l.rn

			l.readRune()
			tok = l.newToken(token.GREATER_EQUAL, string(rn)+string(l.rn))
		} else if l.peekRune() == '>' {
			rn := l.rn

			l.readRune()
			tok = l.newToken(token.RIGHT_SHIFT, string(rn)+string(l.rn))
		} else {
			tok = l.newToken(token.GREATER, string(l.rn))
		}
	case '{':
		tok = l.newToken(token.LBRACE, string(l.rn))
	case '}':
		tok = l.newToken(token.RBRACE, string(l.rn))
	case 0:
		tok = l.newToken(token.EOF, "")
	default:
		if isLetter(l.rn) {
			tok = l.readIdentifier()

			return tok
		} else if isDigit(l.rn) {
			tok = l.readNumber()

			return tok
		} else {
			tok = l.newToken(token.ILLEGAL, string(l.rn))
		}
	}

	l.readRune()

	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.rn == ' ' || l.rn == '\t' || l.rn == '\n' || l.rn == '\r' {
		l.readRune()
	}
}

func (l *Lexer) peekRune() rune {
	if l.readPos >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPos]
	}
}

func (l *Lexer) newToken(t token.TokenType, lit string) token.Token {
	return token.Token{Type: t, Literal: lit}
}

func (l *Lexer) readIdentifier() token.Token {
	pos := l.pos
	for isLetter(l.rn) {
		l.readRune()
	}

	id := token.LookupIdent(string(l.input[pos:l.pos]))

	return l.newToken(id, string(l.input[pos:l.pos]))
}

func (l *Lexer) readNumber() token.Token {
	pos := l.pos
	for isDigit(l.rn) {
		l.readRune()
	}

	return l.newToken(token.INT, string(l.input[pos:l.pos]))
}

func isLetter(rn rune) bool {
	return 'a' <= rn && rn <= 'z' || 'A' <= rn && rn <= 'Z' || rn == '_'
}

func isDigit(rn rune) bool {
	return '0' <= rn && rn <= '9'
}
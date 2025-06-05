package lexer

import (
	"testing"
	"xyn/token"
)

func TestNextToken(t *testing.T) {
	input := `five := 5;
		ten := 10;

		add := fn(x, y) {
			x + y;
		}

		result := add(five, ten);
		<- -> <:= - < > = == != [ ] | \ / * % $ # @ ! " ' & ^ ~

		shared move copy borrow
		fn as in where true false if else return
		`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IDENT, "five"},
		{token.ASSIGN, ":="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "ten"},
		{token.ASSIGN, ":="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "add"},
		{token.ASSIGN, ":="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.IDENT, "result"},
		{token.ASSIGN, ":="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.LEFT_SIGNAL, "<-"},
		{token.RIGHT_SIGNAL, "->"},
		{token.ASSIGN_SIGNAL, "<:="},
		{token.MINUS, "-"},
		{token.LESS, "<"},
		{token.GREATER, ">"},
		{token.REASSIGN, "="},
		{token.EQUAL, "=="},
		{token.NOT_EQUAL, "!="},
		{token.LBRACKET, "["},
		{token.RBRACKET, "]"},
		{token.PIPE, "|"},
		{token.BACKSLASH, "\\"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.PERCENT, "%"},
		{token.DOLLAR, "$"},
		{token.HASH, "#"},
		{token.AT, "@"},
		{token.BANG, "!"},
		{token.DOUBLE_QUOTE, "\""},
		{token.QUOTE, "'"},
		{token.AMPERSAND, "&"},
		{token.CARET, "^"},
		{token.TILDE, "~"},
		{token.SHARED, "shared"},
		{token.MOVE, "move"},
		{token.COPY, "copy"},
		{token.BORROW, "borrow"},
		{token.FUNCTION, "fn"},
		{token.AS, "as"},
		{token.IN, "in"},
		{token.WHERE, "where"},
		{token.TRUE, "true"},
		{token.FALSE, "false"},
		{token.IF, "if"},
		{token.ELSE, "else"},
		{token.RETURN, "return"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType || tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q (%q), got=%q (%q)", i, token.LookupTypeName(tt.expectedType), tt.expectedLiteral, token.LookupTypeName(tok.Type), tok.Literal)
		}
	}
}

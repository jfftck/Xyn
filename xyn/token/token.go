package token

type TokenType int

type Token struct {
	Type    TokenType
	Literal string
}

const (
    ILLEGAL TokenType = iota
    EOF

    // Identifiers and literals
    IDENT
    INT
    STRING

    // Operators
    ASSIGN
    PLUS
    MINUS
    BANG
    ASTERISK
    SLASH
    BACKSLASH
    HASH
    EQUAL
    REASSIGN
    
    // Delimiters
    COMMA
    SEMICOLON
    COLON

    LPAREN
    RPAREN
    LBRACE
    RBRACE
    LBRACKET
    RBRACKET

    // Keywords
    FUNCTION
    WHERE
    RETURN
    IF
    ELSE
    TRUE
    FALSE
    AS
    IN
)
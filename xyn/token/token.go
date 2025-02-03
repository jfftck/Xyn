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

	// Comparison operators
	LESS
	LESS_EQUAL
	GREATER
	GREATER_EQUAL
	EQUAL
	NOT_EQUAL

	// Logical operators
	AND
	OR
	XOR
	LEFT_SHIFT
	RIGHT_SHIFT

	// Signals
	LEFT_SIGNAL
	RIGHT_SIGNAL
	BIDIRECTIONAL_SIGNAL
	ASSIGN_SIGNAL // <:= (Squid)
)

var typeNames = map[TokenType]string {
    ILLEGAL: "ILLEGAL",
    EOF: "EOF",
    IDENT: "IDENT",
    INT: "INT",
    STRING: "STRING",
    ASSIGN: "ASSIGN",
    PLUS: "PLUS",
    MINUS: "MINUS",
    BANG: "BANG",
    ASTERISK: "ASTERISK",
    SLASH: "SLASH",
    BACKSLASH: "BACKSLASH",
    HASH: "HASH",
    REASSIGN: "REASSIGN",
    COMMA: "COMMA",
    SEMICOLON: "SEMICOLON",
    COLON: "COLON",
    LPAREN: "LPAREN",
    RPAREN: "RPAREN",
    LBRACE: "LBRACE",
    RBRACE: "RBRACE",
    LBRACKET: "LBRACKET",
    RBRACKET: "RBRACKET",
    FUNCTION: "FUNCTION",
    WHERE: "WHERE",
    RETURN: "RETURN",
    IF: "IF",
    ELSE: "ELSE",
    TRUE: "TRUE",
    FALSE: "FALSE",
    AS: "AS",
    IN: "IN",
    LESS: "LESS",
    LESS_EQUAL: "LESS_EQUAL",
    GREATER: "GREATER",
    GREATER_EQUAL: "GREATER_EQUAL",
    EQUAL: "EQUAL",
    NOT_EQUAL: "NOT_EQUAL",
    AND: "AND",
    OR: "OR",
    XOR: "XOR",
    LEFT_SHIFT: "LEFT_SHIFT",
    RIGHT_SHIFT: "RIGHT_SHIFT",
    LEFT_SIGNAL: "LEFT_SIGNAL",
    RIGHT_SIGNAL: "RIGHT_SIGNAL",
    BIDIRECTIONAL_SIGNAL: "BIDIRECTIONAL_SIGNAL",
    ASSIGN_SIGNAL: "ASSIGN_SIGNAL",
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"as":     AS,
	"in":     IN,
	"where":  WHERE,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

func LookupTypeName(tokenType TokenType) string {
    return typeNames[tokenType]
}

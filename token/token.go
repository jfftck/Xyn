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
	LET
	IDENT
	INT
	LONG
	FLOAT
	DECIMAL
	STRING
	TEMPLATE_STRING
	CHAR
	BYTE
	BOOL

	// Operators
	ASSIGN
	PLUS
	MINUS
	ASTERISK
	SLASH
	BACKSLASH
	TILDE

	// Sybols
	BANG
	QUESTION
	DOT
	AT
	DOLLAR
	PERCENT
	AMPERSAND
	QUOTE
	UNDERSCORE
	PIPE
	CARET
	DOUBLE_QUOTE
	BACK_QUOTE
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

	// Ownership modifiers
	SHARED // Reference counted
	MOVE   // Transfer ownership
	COPY   // Create independent copy
	BORROW // Temporary reference

	// Special keywords
	IMPORT
	EXPORT
	EXIT
	MAIN
)

var typeNames = map[TokenType]string{
	ILLEGAL:              "ILLEGAL",
	EOF:                  "EOF",
	LET:                  "LET",
	IDENT:                "IDENT",
	INT:                  "INT",
	LONG:                 "LONG",
	FLOAT:                "FLOAT",
	DECIMAL:              "DECIMAL",
	STRING:               "STRING",
	TEMPLATE_STRING:      "TEMPLATE_STRING",
	CHAR:                 "CHAR",
	BYTE:                 "BYTE",
	BOOL:                 "BOOL",
	ASSIGN:               "ASSIGN",
	PLUS:                 "PLUS",
	MINUS:                "MINUS",
	BANG:                 "BANG",
	ASTERISK:             "ASTERISK",
	SLASH:                "SLASH",
	BACKSLASH:            "BACKSLASH",
	HASH:                 "HASH",
	REASSIGN:             "REASSIGN",
	COMMA:                "COMMA",
	SEMICOLON:            "SEMICOLON",
	COLON:                "COLON",
	LPAREN:               "LPAREN",
	RPAREN:               "RPAREN",
	LBRACE:               "LBRACE",
	RBRACE:               "RBRACE",
	LBRACKET:             "LBRACKET",
	RBRACKET:             "RBRACKET",
	FUNCTION:             "FUNCTION",
	WHERE:                "WHERE",
	RETURN:               "RETURN",
	IF:                   "IF",
	ELSE:                 "ELSE",
	TRUE:                 "TRUE",
	FALSE:                "FALSE",
	AS:                   "AS",
	IN:                   "IN",
	LESS:                 "LESS",
	LESS_EQUAL:           "LESS_EQUAL",
	GREATER:              "GREATER",
	GREATER_EQUAL:        "GREATER_EQUAL",
	EQUAL:                "EQUAL",
	NOT_EQUAL:            "NOT_EQUAL",
	AND:                  "AND",
	OR:                   "OR",
	XOR:                  "XOR",
	LEFT_SHIFT:           "LEFT_SHIFT",
	RIGHT_SHIFT:          "RIGHT_SHIFT",
	LEFT_SIGNAL:          "LEFT_SIGNAL",
	RIGHT_SIGNAL:         "RIGHT_SIGNAL",
	BIDIRECTIONAL_SIGNAL: "BIDIRECTIONAL_SIGNAL",
	ASSIGN_SIGNAL:        "ASSIGN_SIGNAL",
	QUESTION:             "QUESTION",
	DOT:                  "DOT",
	AT:                   "AT",
	DOLLAR:               "DOLLAR",
	PERCENT:              "PERCENT",
	AMPERSAND:            "AMPERSAND",
	QUOTE:                "QUOTE",
	UNDERSCORE:           "UNDERSCORE",
	PIPE:                 "PIPE",
	CARET:                "CARET",
	DOUBLE_QUOTE:         "DOUBLE_QUOTE",
	BACK_QUOTE:           "BACK_QUOTE",
	TILDE:                "TILDE",
	SHARED:               "SHARED",
	MOVE:                 "MOVE",
	COPY:                 "COPY",
	BORROW:               "BORROW",
	EXIT:                 "EXIT",
	IMPORT:               "IMPORT",
	EXPORT:               "EXPORT",
	MAIN:                 "MAIN",
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
	"shared": SHARED,
	"move":   MOVE,
	"copy":   COPY,
	"borrow": BORROW,
}

var specialKeywords = map[string]TokenType{
	"@exit":   EXIT,
	"@import": IMPORT,
	"@export": EXPORT,
	"@main":   MAIN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	if tok, ok := specialKeywords[ident]; ok {
		return tok
	}
	return IDENT
}

func LookupTypeName(tokenType TokenType) string {
	return typeNames[tokenType]
}

func (t Token) TypeAsString() struct {
	Type    string
	Literal string
} {
	return struct {
		Type    string
		Literal string
	}{LookupTypeName(t.Type), t.Literal}
}

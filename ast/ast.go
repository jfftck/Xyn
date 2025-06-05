package ast

import "xyn/token"

type Node interface {
    TokenLiteral() string
}

type Statement interface {
    Node
    statementNode()
}

type Expression interface {
    Node
    expressionNode()
}

type Program struct {
    Statements []Statement
}

type ExpressionStatement struct {
    Token token.Token
    Expression Expression
}

func (p *Program) TokenLiteral() string {
    if len(p.Statements) > 0 {
        return p.Statements[0].TokenLiteral()
    } else {
        return ""
    }
}

type LetStatement struct {
    Token token.Token
    Name  *Identifier
    Value Expression
}

func (ae *LetStatement) statementNode() {
    
}
func (ae *LetStatement) TokenLiteral() string {
    return ae.Token.Literal
}

type Identifier struct {
    Token token.Token
    Value string
}

func (i *Identifier) expressionNode() {
    
}
func (i *Identifier) TokenLiteral() string {
    return i.Token.Literal
}


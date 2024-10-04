package parser

import (
	"github.com/itsmeberwyn/go-interpreter/ast"
	"github.com/itsmeberwyn/go-interpreter/lexer"
	"github.com/itsmeberwyn/go-interpreter/token"
)

type Parser struct {
	l *lexer.Lexer

	currentToken token.Token
	peekToken    token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.NextToken()
	p.NextToken()
	return p
}

func (p *Parser) NextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
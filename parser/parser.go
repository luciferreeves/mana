package parser

import (
	"fmt"
	"mana/ast"
	"mana/lexer"
	"mana/tokens"
)

// Parser represents a parser.
type Parser struct {
	l *lexer.Lexer

	curToken  tokens.Token
	peekToken tokens.Token

	errors []string
}

// New returns a new Parser.
func New(l *lexer.Lexer) *Parser {
	var p *Parser = &Parser{
		l: l,
		errors: []string{},
	}

	// Read two tokens, so curToken and peekToken are both set.
	p.nextToken()
	p.nextToken()

	return p
}

// nextToken advances the tokens.
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram parses a program.
func (p *Parser) ParseProgram() *ast.Program {

	// Initialize the program with an empty slice of statements.
	var program *ast.Program = &ast.Program{}
	program.Statements = []ast.Statement{}

	// Iterate over the tokens until we reach the end of file.
	for p.curToken.Type != tokens.EOF {
		// Parse the statement and append it to the program.
		var stmt ast.Statement = p.parseStatement()

		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

		// Advance to the next token.
		p.nextToken()
	}

	// Return the program.
	return program
}

// parseStatement parses a statement.
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case tokens.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

// parseLetStatement parses a let statement.
func (p *Parser) parseLetStatement() *ast.LetStatement {
	var stmt *ast.LetStatement = &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(tokens.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(tokens.ASSIGN) {
		return nil
	}

	// TODO: We're skipping the expressions until we
	// encounter a semicolon.

	for !p.curTokenIs(tokens.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// curTokenIs returns true if the current token is of the given type.
func (p *Parser) curTokenIs(t tokens.TokenType) bool {
	return p.curToken.Type == t
}

// peekTokenIs returns true if the next token is of the given type.
func (p *Parser) peekTokenIs(t tokens.TokenType) bool {
	return p.peekToken.Type == t
}

// expectPeek expects the next token to be of the given type.
func (p *Parser) expectPeek(t tokens.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()

		return true
	} else {
		p.peekError(t)
		
		return false
	}
}

// Errors returns the parser errors.
func (p *Parser) Errors() []string {
	return p.errors
}

// peekError returns an error message.
func (p *Parser) peekError(t tokens.TokenType) {
	var msg string = fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)

	p.errors = append(p.errors, msg)
}


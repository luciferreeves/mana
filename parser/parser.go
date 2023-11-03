package parser

import (
	"fmt"
	"mana/ast"
	"mana/lexer"
	"mana/tokens"
	"strconv"
)

// Define the precedence of the operators.
const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

// Parser represents a parser.
type Parser struct {
	l      *lexer.Lexer
	errors []string

	curToken  tokens.Token
	peekToken tokens.Token

	prefixParseFns map[tokens.TokenType]prefixParseFn
	infixParseFns  map[tokens.TokenType]infixParseFn
}

// New returns a new Parser.
func New(l *lexer.Lexer) *Parser {
	var p *Parser = &Parser{
		l:      l,
		errors: []string{},
	}

	// Read two tokens, so curToken and peekToken are both set.
	p.nextToken()
	p.nextToken()

	// Initialize the prefix parse functions.
	p.prefixParseFns = make(map[tokens.TokenType]prefixParseFn)
	p.registerPrefix(tokens.IDENT, p.parseIdentifier)
	p.registerPrefix(tokens.INT, p.parseIntegerLiteral)

	return p
}

// nextToken advances the tokens.
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// registerPrefix registers a prefix parse function.
func (p *Parser) registerPrefix(tokenType tokens.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

// registerInfix registers an infix parse function.
func (p *Parser) registerInfix(tokenType tokens.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
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
	case tokens.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

// parseIdentifier parses an identifier.
func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

// parseIntegerLiteral parses an integer literal.
func (p *Parser) parseIntegerLiteral() ast.Expression {
	var lit *ast.IntegerLiteral = &ast.IntegerLiteral{ Token: p.curToken }

	var value, err = strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		var msg string = fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = value

	return lit
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

// parseReturnStatement parses a return statement.
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	var stmt *ast.ReturnStatement = &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	// TODO: We're skipping the expressions until we
	// encounter a semicolon.

	for !p.curTokenIs(tokens.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// parseExpressionStatement parses an expression statement.
func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	var stmt *ast.ExpressionStatement = &ast.ExpressionStatement{Token: p.curToken}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(tokens.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// parseExpression parses an expression.
func (p *Parser) parseExpression(precedence int) ast.Expression {
	var prefix = p.prefixParseFns[p.curToken.Type]

	if prefix == nil {
		return nil
	}

	var leftExp ast.Expression = prefix()

	return leftExp
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

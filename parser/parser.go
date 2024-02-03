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

var precedences = map[tokens.TokenType]int{
	tokens.EQ:       EQUALS,
	tokens.NOT_EQ:   EQUALS,
	tokens.LT:       LESSGREATER,
	tokens.GT:       LESSGREATER,
	tokens.PLUS:     SUM,
	tokens.MINUS:    SUM,
	tokens.SLASH:    PRODUCT,
	tokens.ASTERISK: PRODUCT,
	tokens.LPAREN:   CALL,
}

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
	p.registerPrefix(tokens.BANG, p.parsePrefixExpression)
	p.registerPrefix(tokens.MINUS, p.parsePrefixExpression)
	p.registerPrefix(tokens.TRUE, p.parseBoolean)
	p.registerPrefix(tokens.FALSE, p.parseBoolean)
	p.registerPrefix(tokens.IF, p.parseIfExpression)
	p.registerPrefix(tokens.FUNCTION, p.parseFunctionLiteral)
	p.registerPrefix(tokens.LPAREN, p.parseGroupedExpression)

	// Initialize the infix parse functions.
	p.infixParseFns = make(map[tokens.TokenType]infixParseFn)
	p.registerInfix(tokens.PLUS, p.parseInfixExpression)
	p.registerInfix(tokens.MINUS, p.parseInfixExpression)
	p.registerInfix(tokens.SLASH, p.parseInfixExpression)
	p.registerInfix(tokens.ASTERISK, p.parseInfixExpression)
	p.registerInfix(tokens.EQ, p.parseInfixExpression)
	p.registerInfix(tokens.NOT_EQ, p.parseInfixExpression)
	p.registerInfix(tokens.LT, p.parseInfixExpression)
	p.registerInfix(tokens.GT, p.parseInfixExpression)
	p.registerInfix(tokens.LPAREN, p.parseCallExpression)

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

// parseBoolean parses a boolean.
func (p *Parser) parseBoolean() ast.Expression {
	return &ast.Boolean{Token: p.curToken, Value: p.curTokenIs(tokens.TRUE)}
}

// parseIntegerLiteral parses an integer literal.
func (p *Parser) parseIntegerLiteral() ast.Expression {
	// defer untrace(trace("parseIntegerLiteral"))

	var lit *ast.IntegerLiteral = &ast.IntegerLiteral{Token: p.curToken}

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
	// defer untrace(trace("parseExpressionStatement"))

	var stmt *ast.ExpressionStatement = &ast.ExpressionStatement{Token: p.curToken}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(tokens.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// noPrefixParseFnError returns an error message.
func (p *Parser) noPrefixParseFnError(t tokens.TokenType) {
	var msg string = fmt.Sprintf("no prefix parse function for %s found", t)

	p.errors = append(p.errors, msg)
}

// parseExpression parses an expression.
func (p *Parser) parseExpression(precedence int) ast.Expression {
	// defer untrace(trace("parseExpression"))

	prefix := p.prefixParseFns[p.curToken.Type]

	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}

	leftExp := prefix()

	for !p.peekTokenIs(tokens.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]

		if infix == nil {
			return leftExp
		}

		p.nextToken()

		leftExp = infix(leftExp)
		
	}

	return leftExp
}

// parsePrefixExpression parses a prefix expression.
func (p *Parser) parsePrefixExpression() ast.Expression {
	// defer untrace(trace("parsePrefixExpression"))

	var expression *ast.PrefixExpression = &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}

	p.nextToken()

	expression.Right = p.parseExpression(PREFIX)

	return expression
}

// parseInfixExpression parses an infix expression.
func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	// defer untrace(trace("parseInfixExpression"))

	expression := &ast.InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}

	precedence := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)

	return expression
}

func (p *Parser) parseIfExpression() ast.Expression {
	expression := &ast.IfExpression{Token: p.curToken}

	if !p.expectPeek(tokens.LPAREN) { return nil }

	p.nextToken()
	expression.Condition = p.parseExpression(LOWEST)

	if !p.expectPeek(tokens.RPAREN) { return nil }

	if !p.expectPeek(tokens.LBRACE) { return nil }

	expression.Consequence = p.parseBlockStatement()

	if p.peekTokenIs(tokens.ELSE) {
		p.nextToken()

		if !p.expectPeek(tokens.LBRACE) { return nil }

		expression.Alternative = p.parseBlockStatement()
	}

	return expression
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: p.curToken}
	block.Statements = []ast.Statement{}

	p.nextToken()

	for !p.curTokenIs(tokens.RBRACE) && !p.curTokenIs(tokens.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}

		p.nextToken()
	}

	return block
}

// parseFunctionLiteral parses a function literal.
func (p *Parser) parseFunctionLiteral() ast.Expression {
	lit := &ast.FunctionLiteral{Token: p.curToken}

	if !p.expectPeek(tokens.LPAREN) { return nil }

	lit.Parameters = p.parseFunctionParameters()

	if !p.expectPeek(tokens.LBRACE) { return nil }

	lit.Body = p.parseBlockStatement()

	return lit

}

// parseFunctionParameters parses function parameters.
func (p *Parser) parseFunctionParameters() []*ast.Identifier {
	identifiers := []*ast.Identifier{}

	if p.peekTokenIs(tokens.RPAREN) {
		p.nextToken()
		return identifiers
	}

	p.nextToken()

	ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	identifiers = append(identifiers, ident)

	for p.peekTokenIs(tokens.COMMA) {
		p.nextToken()
		p.nextToken()
		ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		identifiers = append(identifiers, ident)
	}

	if !p.expectPeek(tokens.RPAREN) { return nil }

	return identifiers
}

// parseGroupedExpression parses a grouped expression.
func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()

	exp := p.parseExpression(LOWEST)

	if !p.expectPeek(tokens.RPAREN) { return nil }

	return exp
}

// parseCallExpression parses a call expression.
func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	exp := &ast.CallExpression{Token: p.curToken, Function: function}
	exp.Arguments = p.parseCallArguments()

	return exp
}

// parseCallArguments parses call arguments.
func (p *Parser) parseCallArguments() []ast.Expression {
	args := []ast.Expression{}

	if p.peekTokenIs(tokens.RPAREN) {
		p.nextToken()
		return args
	}

	p.nextToken()
	args = append(args, p.parseExpression(LOWEST))

	for p.peekTokenIs(tokens.COMMA) {
		p.nextToken()
		p.nextToken()
		args = append(args, p.parseExpression(LOWEST))
	}

	if !p.expectPeek(tokens.RPAREN) { return nil }

	return args
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

// peek and cur precedences 

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}

	return LOWEST
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}

	return LOWEST
}

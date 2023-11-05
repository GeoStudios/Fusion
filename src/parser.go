package main

import (
	"log"
	"strconv"
)

type Parser struct {
	ptr    int
	tokens []Token
}

func New_Parser(tokens []Token) *Parser { return &Parser{ptr: 0, tokens: tokens} }

func (p *Parser) NotEof() bool { return p.ptr < len(p.tokens) }

func (p *Parser) At() Token { return p.tokens[p.ptr] }

func (p *Parser) Peek(ahead int) Token {
	if p.NotEof() && p.ptr+ahead < len(p.tokens) {
		return p.tokens[p.ptr+ahead]
	}
	return Token{Type: EOF}
}

func (p *Parser) Next() Token { p.ptr++; return p.tokens[p.ptr] }
func (p *Parser) Past() Token {
	if p.NotEof() && p.ptr-1 < len(p.tokens) {
		return p.tokens[p.ptr-1]
	}
	return Token{Type: EOF}
}

func (p *Parser) Expect(ExpectedType TokenType) Token {
	p.ptr++
	if p.tokens[p.ptr-1].Type == ExpectedType {
		return p.tokens[p.ptr-1]
	}
	log.Fatalf("Unexpected Type -> \"%v\" Expected -> \"%v\""+" Line: %v",
		string(p.tokens[p.ptr-1].Type),
		string(ExpectedType),
		strconv.Itoa(p.At().Line+1))
	return Token{Type: EOF}
}

func (p *Parser) ExpectMsg(ExpectedType TokenType, Msg string) Token {
	p.ptr++
	if p.tokens[p.ptr-1].Type == ExpectedType {
		return p.tokens[p.ptr-1]
	}
	log.Fatalf("Unexpected Type -> \"%v\" Expected -> \"%v\""+" Line: %v\n%v",
		string(p.tokens[p.ptr-1].Type),
		string(ExpectedType),
		strconv.Itoa(p.At().Line+1),
		Msg)
	return Token{Type: EOF}
}

func (p *Parser) ProduceAst() Stmt {
	body := []Stmt{}
	p.ExpectMsg(Package, "Expected Package On First Line")
	pkgName := p.Expect(Identifer).Literal
	for p.NotEof() {
		body = append(body, p.ParseStmt())
	}
	return &Program{Body: body, PackageName: pkgName}
}

func (p *Parser) ParseStmt() Stmt {
	switch p.At().Type {
	case Constant, Variable:
	}
}
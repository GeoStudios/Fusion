package main

import (
	"log"
	"strings"
)

type Parser struct {
	ptr    int
	tokens []Token
}

func New_Parser(tokens []Token) *Parser { return &Parser{ptr: 0, tokens: tokens} }

func (p *Parser) NotEof() bool { return p.tokens[p.ptr].Type != EOF }

func (p *Parser) At() Token { return p.tokens[p.ptr] }

func (p *Parser) Peek(ahead int) Token {
	if p.NotEof() && p.ptr+ahead < len(p.tokens) {
		return p.tokens[p.ptr+ahead]
	}
	return Token{Type: EOF}
}

func (p *Parser) Next() Token { p.ptr++; return p.tokens[p.ptr-1] }
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
		p.tokens[p.ptr-1],
		ExpectedType,
		p.At().Line+1)
	return Token{Type: EOF}
}

func (p *Parser) ExpectMsg(ExpectedType TokenType, Msg string) Token {
	p.ptr++
	if p.tokens[p.ptr-1].Type == ExpectedType {
		return p.tokens[p.ptr-1]
	}
	log.Fatalf("Unexpected Type -> \"%v\" Expected -> \"%v\""+" Line: %v\n%v",
		p.tokens[p.ptr-1],
		ExpectedType,
		p.At().Line+1,
		Msg)
	return Token{Type: EOF}
}

func (p *Parser) ProduceAst() Stmt {
	body := []Stmt{}
	p.ExpectMsg(Package, "Expected Package On First Line")
	pkgName := p.Expect(Identifer).Literal
	p.Expect(SemiColon)
	for p.NotEof() {
		body = append(body, p.ParseStmt_PRG())
	}
	return &Program{Body: body, PackageName: pkgName}
}

func (p *Parser) ParseStmt_PRG() Stmt {
	switch p.At().Type {
	case Return: log.Fatalln("ILLEGAL RETURN STMT, CANNOT BE OUTSIDE BLOCKSTMT"); return &ExprStmt{ Expression: p.ParseExpr() }
	case Import: return p.ParseImportStmt()
	case If: return p.ParseIfStmt()
	case Function: return p.ParseFunctionStmt()
	case Constant, Variable: return p.ParseVarStmt()
	default:
		stmt := Stmt(&ExprStmt{ Expression: p.ParseExpr() })
		p.Expect(SemiColon)
		return stmt
	// default:
	// 	log.Fatalf("Unexpected Token: %v Line: %v",
	// 	fmt.Sprint(p.At()),
	// 	strconv.Itoa(p.At().Line+1))
	// 	return &IntLiteral{Value: 1}
	}
}

func (p *Parser) ParseStmt() Stmt {
	switch p.At().Type {
	case Return: return p.ParseReturnStmt()
	case Import: log.Fatalln("ILLEGAL IMPORT STMT, IMPORT IS HIGH LEVEL"); return &ExprStmt{ Expression: p.ParseExpr() }
	case If: return p.ParseIfStmt()
	case Function: return p.ParseFunctionStmt()
	case Constant, Variable, Integer, Float, String, Boolean, Void:
		return p.ParseVarStmt()
	default:
		stmt := Stmt(&ExprStmt{ Expression: p.ParseExpr() })
		p.Expect(SemiColon)
		return stmt
	}
}

func (p *Parser) ParseReturnStmt() Stmt {
	var expr Expr
	p.Expect(Return)
	if p.At().Type == SemiColon { expr = &NullLiteral{} } else { expr = p.ParseExpr() }
	p.Expect(SemiColon)
	return &ReturnStmt{
		Expression: expr,
	}
}

func (p *Parser) ParseImportStmt() Stmt {
	p.Expect(Import)
	value := p.Expect(String).Literal
	IsBuiltIn := strings.HasPrefix(value, "@/")
	if IsBuiltIn {
		value = strings.TrimPrefix(value, "@/")
		value = "core/" + value
	}
	p.Expect(SemiColon)
	return &ImportStmt{
		IsBuiltIn: IsBuiltIn,
		PackageName: value,
	}
}

func (p *Parser) ParseVarStmt() Stmt {
	// IsConst := p.Next().Type == Constant
	Type := GetTypeFromToken(p.Next().Type)
	Name := p.Expect(Identifer).Literal
	p.Expect(Basic_Assign)
	Value := p.ParseExpr()
	p.Expect(SemiColon)
	return &VarStmt{
		// IsConst: IsConst,
		Name: Name,
		Value: Value,
		ObjType: Type,
	}
}

func (p *Parser) ParseIfStmt() Stmt {
	p.Expect(If)
	p.Expect(OpenParen)
	Condition := p.ParseExpr()
	p.Expect(CloseParen)
	var OnTrue Stmt
	var OnFalse Stmt
    if p.At().Type == OpenBrace {
        OnTrue = p.ParseBlockStmt()
    } else {
        OnTrue = p.ParseStmt()
    }
	if p.At().Type == Else {
		p.Expect(Else)
		if p.At().Type == OpenBrace {
        	OnFalse = p.ParseBlockStmt()
    	} else {
        	OnFalse = p.ParseStmt()
    	}
	}
	if p.At().Type == SemiColon { p.Expect(SemiColon) }
	return &IfStmt{
		Condition: Condition,
		OnTrue: OnTrue,
		OnFalse: OnFalse,
	}
}

func (p *Parser) ParseBlockStmt() Stmt {
	body := []Stmt{}
	p.Expect(OpenBrace)
	for p.At().Type != CloseBrace && p.NotEof() {
		body = append(body, p.ParseStmt())
	}
	p.Expect(CloseBrace)
	return &BlockStmt{Body: body}
}

func (p *Parser) ParseFunctionStmt() Stmt {
	p.Expect(Function)
	Name := p.Expect(Identifer).Literal
	p.Expect(OpenParen)
	Args := p.ParseTypedArguments(CloseParen)
	p.Expect(Arrow)
	Type := GetTypeFromToken(p.Next().Type)
	Body := p.ParseBlockStmt()
	if p.At().Type == SemiColon { p.Expect(SemiColon) }
	return &FunctionStmt{
		Name: Name,
		Args: Args,
		Body: Body,
		ObjType: Type,
	}
}

type Arg struct {
	Value string
	Type ObjectTypes
}

func (p *Parser) ParseTypedArguments(end TokenType) []Arg {
	if p.At().Type == end { p.Expect(end); return []Arg{} }
	args := []Arg{}
	args = append(args, Arg{
		Type: GetTypeFromToken(p.Next().Type),
		Value: p.Expect(Identifer).Literal,
	})
	for p.At().Type == Comma {
		p.Expect(Comma)
		args = append(args, Arg{
			Type: GetTypeFromToken(p.Next().Type),
			Value: p.Expect(Identifer).Literal,
		})
	}
	p.Expect(end)
	return args
}

func (p *Parser) ParseArguments(end TokenType) []Expr {
	if p.At().Type == end { p.Expect(end); return []Expr{} }
	args := []Expr{}
	args = append(args, p.ParseExpr())
	for p.At().Type == Comma {
		p.Expect(Comma)
		args = append(args, p.ParseExpr())
	}
	p.Expect(end)
	return args
}
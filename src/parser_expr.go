package main

import (
	"fmt"
	"log"
	"strconv"
)

func (p *Parser) ParseExpr() Expr {
	return p.ParseTernaryExpr()
}

func (p *Parser) ParseTernaryExpr() Expr {
	left := p.ParseAssignExpr()
	if p.At().Type == QuesionMark {
		p.Expect(QuesionMark)
		iif := p.ParseTernaryExpr()
		p.Expect(Colon)
		other := p.ParseTernaryExpr()
		return &TernaryExpr{
			Condition: left,
			OnTrue: iif,
			OnFalse: other,
		}
	}
	return left
}

func (p *Parser) ParseAssignExpr() Expr {
	left := p.ParseLogicalOrExpr()
	switch t := p.At().Type; t {
	case Basic_Assign, Plus_Assign, Minus_Assign, Multiply_Assign,
		Divide_Assign, Modulo_Assign:
		return &AssignExpr{
			Left:  left,
			Op:    p.Next().Type,
			Right: p.ParseAssignExpr(),
		}
	}
	return left
}

func (p *Parser) ParseLogicalOrExpr() Expr {
	left := p.ParseLogicalAndExpr()
	for p.At().Type == Logic_Or {
		op := p.Next().Type
		right := p.ParseLogicalAndExpr()
		left = &BinaryExpr{
			Left: left,
			Op: op,
			Right: right,
		}
	}
	return left
}

func (p *Parser) ParseLogicalAndExpr() Expr {
	left := p.ParseBitwiseOrExpr()
	for p.At().Type == Logic_And {
		op := p.Next().Type
		right := p.ParseBitwiseOrExpr()
		left = &BinaryExpr{
			Left: left,
			Op: op,
			Right: right,
		}
	}
	return left
}

func (p *Parser) ParseBitwiseOrExpr() Expr {
	left := p.ParseBitwiseXorExpr()
	for p.At().Type == Or {
		op := p.Next().Type
		right := p.ParseBitwiseXorExpr()
		left = &BinaryExpr{
			Left: left,
			Op: op,
			Right: right,
		}
	}
	return left
}


func (p *Parser) ParseBitwiseXorExpr() Expr {
	left := p.ParseBitwiseAndExpr()
	for p.At().Type == Xor {
		op := p.Next().Type
		right := p.ParseBitwiseAndExpr()
		left = &BinaryExpr{
			Left: left,
			Op: op,
			Right: right,
		}
	}
	return left
}


func (p *Parser) ParseBitwiseAndExpr() Expr {
	left := p.ParseEqualityExpr()
	for p.At().Type == And {
		op := p.Next().Type
		right := p.ParseEqualityExpr()
		left = &BinaryExpr{
			Left: left,
			Op: op,
			Right: right,
		}
	}
	return left
}

func (p *Parser) ParseEqualityExpr() Expr {
	left := p.ParseRelationalExpr()
	for p.At().Type == Equals || p.At().Type == NotEquals {
		op := p.Next().Type
		right := p.ParseRelationalExpr()
		left = &BinaryExpr{
			Left:  left,
			Op:    op,
			Right: right,
		}
	}
	return left
}

func (p *Parser) ParseRelationalExpr() Expr {
	left := p.ParseAdditiveExpr()
	for p.At().Type == LessThan || p.At().Type == LessThanEqualTo ||
		p.At().Type == GreaterThan || p.At().Type == GreaterThanEqualTo {
		op := p.Next().Type
		right := p.ParseAdditiveExpr()
		left = &BinaryExpr{
			Left:  left,
			Op:    op,
			Right: right,
		}
	}
	return left
}

func (p *Parser) ParseAdditiveExpr() Expr {
	left := p.ParseMultiplicativeExpr()
	for p.At().Type == Plus || p.At().Type == Minus {
		op := p.Next().Type
		right := p.ParseMultiplicativeExpr()
		left = &BinaryExpr{
			Left: left,
			Op: op,
			Right: right,
		}
	}
	return left
}

func (p *Parser) ParseMultiplicativeExpr() Expr {
	left := p.ParsePrefixExpr()
	for p.At().Type == Multiply || p.At().Type == Divide || p.At().Type == Modulo {
		op := p.Next().Type
		right := p.ParsePrefixExpr()
		left = &BinaryExpr{
			Left: left,
			Op: op,
			Right: right,
		}
	}
	return left
}

func (p *Parser) ParsePrefixExpr() Expr {
	if p.At().Type == Increment ||
	p.At().Type == Decrement ||
	p.At().Type == Logic_Not ||
	p.At().Type == Not ||
	p.At().Type == Plus ||
	p.At().Type == Minus {
		op := p.Next().Type
		right := p.ParseAssignExpr()
		return &PrefixExpr{
			Op:   op,
			Right: right,
		}
	}

	return p.ParsePostfixExpr()
}

func (p *Parser) ParsePostfixExpr() Expr {
	left := p.ParseCallMemberExpr()
	if left.Type() != _Identifier { return left }
	if p.At().Type == Increment || p.At().Type == Decrement {
		op := p.Next().Type
		return &PostfixExpr{
			Op: op,
			Left: left,
		}
	}

	return left

}

func (p *Parser) ParseCallMemberExpr() Expr {
	member := p.ParseMemberExpr()
	if p.At().Type == OpenParen && p.Past().Line == p.At().Line { 
		return p.ParseCallExpr(member)
	}
	return member
}

func (p *Parser) ParseMemberExpr() Expr {

	obj := p.ParsePrimaryExpr()

	for (p.At().Type == Dot || p.At().Type == OpenBracket){
		op := p.Next()
		var property Expr
		var computed bool

		if (op.Type == Dot) {
			computed = false;
			// get identifier
			property = p.ParsePrimaryExpr();
			if (property.Type() != _Identifier) {
			  log.Fatalf("Cannot use dot operator without right hand side being a identifier -> %v\n", property)
			}
		  } else { // this allows obj[computedValue]
			computed = true;
			property = p.ParseExpr();
			p.ExpectMsg(
			  CloseBracket,
			  "Missing closing bracket in computed value.",
			);
		}
		obj = &MemberExpr{
			Obj: obj,
			Property: property,
			Computed: computed,
		}
	}


	return obj

}

func (p *Parser) ParseCallExpr(caller Expr) Expr {
	p.Expect(OpenParen)
	var callExpr Expr = &CallExpr{
		Caller: caller,
		Args: p.ParseArguments(CloseParen),
	}
	if p.At().Type == OpenParen { callExpr = p.ParseCallExpr(callExpr) }
	return callExpr
}

func (p *Parser) ParsePrimaryExpr() Expr {
	switch p.At().Type {
	case Identifer: return &Identifier{Value: p.Next().Literal }
	case String: return &StringLiteral{Value: p.Next().Literal }
	case Null: p.Next(); return &NullLiteral{}
	case Boolean: v := p.Next().Literal; value, _ := strconv.ParseBool(v); return &BoolLiteral{Value: value}
	case Integer: v := p.Next().Literal; value, _ := strconv.ParseInt(v, 10, 64); return &IntLiteral{Value: int(value)}
	case Float: value, _ := strconv.ParseFloat(p.Next().Literal, 64); return &FloatLiteral{Value: value}
	case OpenBracket:
		p.Expect(OpenBracket)
		return &ArrayLiteral{Elements: p.ParseArguments(CloseBracket)}

	case OpenParen:
		p.Expect(OpenParen)
		expr := p.ParseExpr()
		p.Expect(CloseParen)
		return expr
	case OpenBrace:
		pairs := make(map[string]Expr)
		p.Expect(OpenBrace)
		for p.At().Type != CloseBrace {
			key := p.Expect(Identifer).Literal
			p.Expect(Colon)
			value := p.ParseExpr()
			pairs[key] = value
			if p.At().Type == Comma {
				p.Expect(Comma)
			}
		}
		p.Expect(CloseBrace)
		return &HashLiteral{ Pairs: pairs }	
	default:
		log.Fatalf("Unexpected Token: %v Line: %v",
		fmt.Sprint(p.At()),
		strconv.Itoa(p.At().Line+1))
		return &IntLiteral{Value: 1}
	}
}
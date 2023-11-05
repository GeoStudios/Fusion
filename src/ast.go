package main

import (
	"bytes"
	"fmt"
)

type AstNodeType int

const (
	_Program AstNodeType = iota
	_ExprStmt
	_PackageStmt
	_ImportStmt

	_IfStmt
	_WhileStmt
	_BlockStmt
	_VarStmt
	_FunctionStmt

	_StringLiteral
	_FloatLiteral64
	_FloatLiteral32
	_IntLiteral
	_BoolLiteral

	_AssginmentExpr
	_TernaryExpr
	_BinaryExpr
	_PrefixExpr
	_PostfixExpr
)

type AstNode interface {
	String() string
	Type() AstNodeType
}

type Stmt interface{ AstNode }
type Expr interface{ AstNode }

type Program struct {
	Stmt
	Body []Stmt
	PackageName string
}

func (t *Program) Type() AstNodeType { return _Program }
func (t *Program) String() string {
	var str bytes.Buffer
	for i, v := range t.Body {
		str.WriteString(v.String())
		if i < len(t.Body) { str.WriteString(";") }
	}
	return str.String()
}

type ExprStmt struct {
	Stmt
	Expression Expr
}

func (t *ExprStmt) Type() AstNodeType { return _ExprStmt }
func (t *ExprStmt) String() string    { return t.Expression.String() }

type PackageStmt struct {
	Stmt
	PackageName string
}

func (t *PackageStmt) Type() AstNodeType { return _PackageStmt }
func (t *PackageStmt) String() string    { return "package " + t.PackageName + ";" }

type ImportStmt struct {
	Stmt
	PackageName string
	IsBuiltIn   bool
}

func (t *ImportStmt) Type() AstNodeType { return _PackageStmt }
func (t *ImportStmt) String() string    { return "using \"" + t.PackageName + "\";" }

type IfStmt struct {
	Stmt
	Condition Expr
	OnTrue Stmt
	OnFalse Stmt
}

func (t *IfStmt) Type() AstNodeType { return _IfStmt }
func (t *IfStmt) String() string {
	var str bytes.Buffer
	str.WriteString("if (")
	str.WriteString(t.Condition.String())
	str.WriteString(")")
	str.WriteString(t.OnTrue.String())
	if t.OnFalse != nil {
		str.WriteString("else ")
		str.WriteString(t.OnFalse.String())
	}
	return str.String()
}

type WhileStmt struct {
	Stmt
	Condition Expr
	Loop Stmt
}

func (t *WhileStmt) Type() AstNodeType { return _WhileStmt }
func (t *WhileStmt) String() string {
	var str bytes.Buffer
	str.WriteString("while (")
	str.WriteString(t.Condition.String())
	str.WriteString(")")
	str.WriteString(t.Loop.String())
	return str.String()
}

type BlockStmt struct {
	Stmt
	Body []Stmt
}

func (t *BlockStmt) Type() AstNodeType { return _BlockStmt }
func (t *BlockStmt) String() string {
	var str bytes.Buffer
	str.WriteString("{")
	for i, v := range t.Body {
		str.WriteString(v.String())
		if i < len(t.Body) { str.WriteString(";") }
	}
	return str.String()
}

type VarStmt struct {
	Stmt
	Name string
	IsConst bool
	Value Expr
}

func (t *VarStmt) Type() AstNodeType { return _VarStmt }
func (t *VarStmt) String() string {
	var str bytes.Buffer
	if t.IsConst { str.WriteString("const ") } else { str.WriteString("var ") }
	str.WriteString(t.Name)
	str.WriteString(" = ")
	str.WriteString(t.Value.String())
	return str.String()
}

type FunctionStmt struct {
	Stmt
	Name string
	Args []Expr
	Body Stmt
}

func (t *FunctionStmt) Type() AstNodeType { return _FunctionStmt }
func (t *FunctionStmt) String() string {
	var str bytes.Buffer
	str.WriteString("fn ")
	str.WriteString(t.Name)
	str.WriteString("(")
	for i, v := range t.Args {
		str.WriteString(v.String())
		if i < len(t.Args) { str.WriteString(",") }
	}
	str.WriteString(")")
	str.WriteString(t.Body.String())
	return str.String()
}

type StringLiteral struct {
	Stmt
	Value string
}

func (t *StringLiteral) Type() AstNodeType { return _StringLiteral }
func (t *StringLiteral) String() string {
	var str bytes.Buffer
	str.WriteString("\"")
	str.WriteString(t.Value)
	str.WriteString("\"")
	return str.String()
}

type FloatLiteral64 struct {
	Stmt
	Value float64
}

func (t *FloatLiteral64) Type() AstNodeType { return _FloatLiteral64 }
func (t *FloatLiteral64) String() string {
	var str bytes.Buffer
	str.WriteString(fmt.Sprint(t.Value))
	return str.String()
}

type FloatLiteral32 struct {
	Stmt
	Value float32
}

func (t *FloatLiteral32) Type() AstNodeType { return _FloatLiteral32 }
func (t *FloatLiteral32) String() string {
	var str bytes.Buffer
	str.WriteString(fmt.Sprint(t.Value))
	return str.String()
}

type IntLiteral struct {
	Stmt
	Value int
	Unsigned bool
	Arch int
}

func (t *IntLiteral) Type() AstNodeType { return _IntLiteral }
func (t *IntLiteral) String() string {
	var str bytes.Buffer
	str.WriteString(fmt.Sprint(t.Value))
	return str.String()
}

type BoolLiteral struct {
	Stmt
	Value int
}

func (t *BoolLiteral) Type() AstNodeType { return _BoolLiteral }
func (t *BoolLiteral) String() string {
	var str bytes.Buffer
	str.WriteString(fmt.Sprint(t.Value))
	return str.String()
}

type AssignExpr struct {
	Stmt
	Left Expr
	Right Expr
	Op TokenType
}

func (t *AssignExpr) Type() AstNodeType { return _AssginmentExpr }
func (t *AssignExpr) String() string {
	var str bytes.Buffer
	str.WriteString(t.Left.String())
	switch t.Op {
	case Basic_Assign: str.WriteString("=")
	case Plus_Assign: str.WriteString("+=")
	case Minus_Assign: str.WriteString("-=")
	case Multiply_Assign: str.WriteString("*=")
	case Divide_Assign: str.WriteString("/=")
	case Modulo_Assign: str.WriteString("%=")
	}
	str.WriteString(t.Right.String())
	return str.String()
}

type TernaryExpr struct {
	Stmt
	Condition Expr
	OnTrue Expr
	OnFalse Expr
}

func (t *TernaryExpr) Type() AstNodeType { return _TernaryExpr }
func (t *TernaryExpr) String() string {
	var str bytes.Buffer
	str.WriteString(t.Condition.String())
	str.WriteString(" ? ")
	str.WriteString(t.OnTrue.String())
	str.WriteString(" : ")
	str.WriteString(t.OnFalse.String())
	return str.String()
}

type BinaryExpr struct {
	Stmt
	Left Expr
	Right Expr
	Op TokenType
}

func (t *BinaryExpr) Type() AstNodeType { return _BinaryExpr }
func (t *BinaryExpr) String() string {
	var str bytes.Buffer
	str.WriteString(t.Left.String())
	switch t.Op {
	case Plus: str.WriteString("+")
	case Minus: str.WriteString("-")
	case Modulo: str.WriteString("%")
	case Multiply: str.WriteString("*")
	case Divide: str.WriteString("/")
	
	case And: str.WriteString("&")
	case Or: str.WriteString("|")
	case Xor: str.WriteString("^")
	
	case Logic_And: str.WriteString("&&")
	case Logic_Or: str.WriteString("||")

	case Equals: str.WriteString("==")
	case NotEquals: str.WriteString("!=")
	case GreaterThan: str.WriteString(">")
	case GreaterThanEqualTo: str.WriteString(">=")
	case LessThan: str.WriteString("<")
	case LessThanEqualTo: str.WriteString("<=")
	}
	str.WriteString(t.Right.String())
	return str.String()
}

type PrefixExpr struct {
	Stmt
	Right Expr
	Op TokenType
}

func (t *PrefixExpr) Type() AstNodeType { return _PrefixExpr }
func (t *PrefixExpr) String() string {
	var str bytes.Buffer
	switch t.Op {
	case Plus: str.WriteString("+")
	case Minus: str.WriteString("-")
	case Decrement: str.WriteString("--")
	case Increment: str.WriteString("++")
	}
	str.WriteString(t.Right.String())
	return str.String()
}

type PostfixExpr struct {
	Stmt
	Right Expr
	Op TokenType
}

func (t *PostfixExpr) Type() AstNodeType { return _PrefixExpr }
func (t *PostfixExpr) String() string {
	var str bytes.Buffer
	str.WriteString(t.Right.String())
	switch t.Op {
	case Decrement: str.WriteString("--")
	case Increment: str.WriteString("++")
	}
	return str.String()
}
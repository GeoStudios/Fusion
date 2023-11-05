package main

import (
	"bytes"
	"fmt"
	"strings"
)

type Typed_AstNodetype int

const (
	_Typed_Program Typed_AstNodetype = iota
	_Typed_ExprStmt
	_Typed_PackageStmt
	_Typed_ImportStmt

	_Typed_IfStmt
	_Typed_WhileStmt
	_Typed_BlockStmt
	_Typed_VarStmt
	_Typed_FunctionStmt

	_Typed_Identifier
	_Typed_StringLiteral
	_Typed_FloatLiteral64
	_Typed_FloatLiteral32
	_Typed_IntLiteral
	_Typed_BoolLiteral
	_Typed_NullLiteral
	_Typed_HashLiteral
	_Typed_ArrayLiteral

	_Typed_AssginmentExpr
	_Typed_TernaryExpr
	_Typed_BinaryExpr
	_Typed_PrefixExpr
	_Typed_PostfixExpr
	_Typed_CallExpr
	_Typed_MemberExpr
)

type Typed_AstNode interface {
	String() string
	Type() Typed_AstNodetype
}

type Typed_Stmt interface{ AstNode }
type Typed_Expr interface{ AstNode }

type Typed_Program struct {
	Stmt
	Body []Stmt
	PackageName string
}

func (t *Typed_Program) Type() Typed_AstNodetype { return _Typed_Program }
func (t *Typed_Program) String() string {
	var str bytes.Buffer
	str.WriteString("package ")
	str.WriteString(t.PackageName)
	str.WriteString(";")
	for _, v := range t.Body {
		str.WriteString(v.String())
		if v.Type() != _ExprStmt {
			str.WriteString(";")
		}
		// if i < len(t.Body)-1 { str.WriteString(";") }
	}
	return str.String()
}

type Typed_ExprStmt struct {
	Stmt
	Expression Expr
}

func (t *Typed_ExprStmt) Type() Typed_AstNodetype { return _Typed_ExprStmt }
func (t *Typed_ExprStmt) String() string    { return t.Expression.String()+";" }

// type Typed_PackageStmt struct {
// 	Stmt
// 	PackageName string
// }

// func (t *Typed_PackageStmt) Type() Typed_AstNodetype Typed_{ return _PackageStmt }
// func (t *Typed_PackageStmt) String() string    { return "package " + t.PackageName }

type Typed_ImportStmt struct {
	Stmt
	PackageName string
	IsBuiltIn   bool
}

func (t *Typed_ImportStmt) Type() Typed_AstNodetype { return _Typed_PackageStmt }
func (t *Typed_ImportStmt) String() string    {
	if t.IsBuiltIn {
		return "using \"@/" + strings.TrimPrefix(t.PackageName, "core/") + "\""
	}
	return "using \"" + t.PackageName + "\";"
}

type Typed_IfStmt struct {
	Stmt
	Condition Expr
	OnTrue Stmt
	OnFalse Stmt
}

func (t *Typed_IfStmt) Type() Typed_AstNodetype { return _Typed_IfStmt }
func (t *Typed_IfStmt) String() string {
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

type Typed_WhileStmt struct {
	Stmt
	Condition Expr
	Loop Stmt
}

func (t *Typed_WhileStmt) Type() Typed_AstNodetype { return _Typed_WhileStmt }
func (t *Typed_WhileStmt) String() string {
	var str bytes.Buffer
	str.WriteString("while (")
	str.WriteString(t.Condition.String())
	str.WriteString(")")
	str.WriteString(t.Loop.String())
	return str.String()
}

type Typed_BlockStmt struct {
	Stmt
	Body []Stmt
}

func (t *Typed_BlockStmt) Type() Typed_AstNodetype { return _Typed_BlockStmt }
func (t *Typed_BlockStmt) String() string {
	var str bytes.Buffer
	str.WriteString("{")
	for _, v := range t.Body {
		str.WriteString(v.String())
		if v.Type() != _ExprStmt {
			str.WriteString(";")
		}
		// if i < len(t.Body)-1 { str.WriteString(";") }
	}
	str.WriteString("}")
	return str.String()
}

type Typed_VarStmt struct {
	Stmt
	Name string
	IsConst bool
	Value Expr
	Objtype ObjectTypes
}

func (t *Typed_VarStmt) Type() Typed_AstNodetype { return _Typed_VarStmt }
func (t *Typed_VarStmt) String() string {
	var str bytes.Buffer
	// if t.IsConst { str.WriteString("const ") } else { str.WriteString("var ") }
	switch t.Objtype{
		case _StringObject: str.WriteString("string ")
		case _BooleanObject: str.WriteString("boolean ")
		case _IntObject: str.WriteString("int ")
		case _FloatObject: str.WriteString("float ")
		case _NullObject: str.WriteString("void ")
	}
	str.WriteString(t.Name)
	str.WriteString(" = ")
	str.WriteString(t.Value.String())
	return str.String()
}

type Typed_FunctionStmt struct {
	Stmt
	Name string
	Args []Expr
	Body Stmt
	Objtype ObjectTypes
}

func (t *Typed_FunctionStmt) Type() Typed_AstNodetype { return _Typed_FunctionStmt }
func (t *Typed_FunctionStmt) String() string {
	var str bytes.Buffer
	str.WriteString("fn ")
	str.WriteString(t.Name)
	str.WriteString("(")
	for i, v := range t.Args {
		str.WriteString(v.String())
		if i < len(t.Args)-1 { str.WriteString(",") }
	}
	str.WriteString(") ")
	str.WriteString("-> ")
	switch t.Objtype{
		case _StringObject: str.WriteString("string ")
		case _BooleanObject: str.WriteString("boolean ")
		case _IntObject: str.WriteString("int ")
		case _FloatObject: str.WriteString("float ")
		case _NullObject: str.WriteString("void ")
	}
	str.WriteString(t.Body.String())
	return str.String()
}

type Typed_ReturnStmt struct {
	Stmt
	Expression Expr
}

func (t *Typed_ReturnStmt) Type() Typed_AstNodetype { return _Typed_FunctionStmt }
func (t *Typed_ReturnStmt) String() string {
	var str bytes.Buffer
	str.WriteString("return ")
	str.WriteString(t.Expression.String())
	return str.String()
}

type Typed_StringLiteral struct {
	Expr
	Value string
}

func (t *Typed_StringLiteral) Type() Typed_AstNodetype { return _Typed_StringLiteral }
func (t *Typed_StringLiteral) String() string {
	var str bytes.Buffer
	str.WriteString("\"")
	str.WriteString(t.Value)
	str.WriteString("\"")
	return str.String()
}

type Typed_Identifier struct {
	Expr
	Value string
}

func (t *Typed_Identifier) Type() Typed_AstNodetype { return _Typed_Identifier }
func (t *Typed_Identifier) String() string { return t.Value }

type Typed_NullLiteral struct { Expr }

func (t *Typed_NullLiteral) Type() Typed_AstNodetype { return _Typed_NullLiteral }
func (t *Typed_NullLiteral) String() string { return "null" }

type Typed_FloatLiteral struct {
	Expr
	Value float64
}

func (t *Typed_FloatLiteral) Type() Typed_AstNodetype { return _Typed_FloatLiteral64 }
func (t *Typed_FloatLiteral) String() string { return fmt.Sprint(t.Value) }

type Typed_IntLiteral struct {
	Expr
	Value int
}

func (t *Typed_IntLiteral) Type() Typed_AstNodetype { return _Typed_IntLiteral }
func (t *Typed_IntLiteral) String() string { return fmt.Sprint(t.Value) }

type Typed_BoolLiteral struct {
	Expr
	Value bool
}

func (t *Typed_BoolLiteral) Type() Typed_AstNodetype { return _Typed_BoolLiteral }
func (t *Typed_BoolLiteral) String() string { return fmt.Sprint(t.Value) }

type Typed_HashLiteral struct {
	Expr `json:"-"`
	Pairs map[Expr]Expr
}

func (t *Typed_HashLiteral) Type() Typed_AstNodetype { return _Typed_HashLiteral }
func (t *Typed_HashLiteral) String() string {
	var str bytes.Buffer
	str.WriteString("{")
	keys := make([]Expr, 0, len(t.Pairs))
	for e := range t.Pairs {
		keys = append(keys, e)
	}
	i:=0
	for key, value := range t.Pairs {
		str.WriteString(key.String()+":"+value.String())
		if i < len(keys)-1 { str.WriteString(",") }
		i++
	}
	str.WriteString("}")
	return str.String()
}

type Typed_ArrayLiteral struct {
	Expr `json:"-"`
	Elements []Expr
}

func (t *Typed_ArrayLiteral) Type() Typed_AstNodetype { return _Typed_ArrayLiteral }
func (t *Typed_ArrayLiteral) String() string {
	var str bytes.Buffer
	str.WriteString("(")
	for i, v := range t.Elements {
		str.WriteString(v.String())
		if i < len(t.Elements)-1 { str.WriteString(",") }
	}
	str.WriteString(")")
	return str.String()
}

type Typed_AssignExpr struct {
	Expr
	Left Expr
	Right Expr
	Op TokenType
}

func (t *Typed_AssignExpr) Type() Typed_AstNodetype { return _Typed_AssginmentExpr }
func (t *Typed_AssignExpr) String() string {
	var str bytes.Buffer
	str.WriteString("(")
	str.WriteString(t.Left.String())
	switch t.Op {
	case Basic_Assign: str.WriteString("=")
	case Plus_Assign: str.WriteString("+=")
	case Minus_Assign: str.WriteString("-=")
	case Multiply_Assign: str.WriteString("*Typed_=")
	case Divide_Assign: str.WriteString("/=")
	case Modulo_Assign: str.WriteString("%=")
	}
	str.WriteString(t.Right.String())
	str.WriteString(")")
	return str.String()
}

type Typed_TernaryExpr struct {
	Expr
	Condition Expr
	OnTrue Expr
	OnFalse Expr
}

func (t *Typed_TernaryExpr) Type() Typed_AstNodetype { return _Typed_TernaryExpr }
func (t *Typed_TernaryExpr) String() string {
	var str bytes.Buffer
	str.WriteString(t.Condition.String())
	str.WriteString(" ? ")
	str.WriteString(t.OnTrue.String())
	str.WriteString(" : ")
	str.WriteString(t.OnFalse.String())
	return str.String()
}

type Typed_BinaryExpr struct {
	Expr
	Left Expr
	Right Expr
	Op TokenType
}

func (t *Typed_BinaryExpr) Type() Typed_AstNodetype { return _Typed_BinaryExpr }
func (t *Typed_BinaryExpr) String() string {
	var str bytes.Buffer
	str.WriteString("(")
	str.WriteString(t.Left.String())
	switch t.Op {
	case Plus: str.WriteString("+")
	case Minus: str.WriteString("-")
	case Modulo: str.WriteString("%")
	case Multiply: str.WriteString("*Typed_")
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
	str.WriteString(")")
	return str.String()
}

type Typed_PrefixExpr struct {
	Expr
	Right Expr
	Op TokenType
}

func (t *Typed_PrefixExpr) Type() Typed_AstNodetype { return _Typed_PrefixExpr }
func (t *Typed_PrefixExpr) String() string {
	var str bytes.Buffer
	str.WriteString("(")
	switch t.Op {
	case Plus: str.WriteString("+")
	case Minus: str.WriteString("-")
	case Decrement: str.WriteString("--")
	case Increment: str.WriteString("++")
	}
	str.WriteString(t.Right.String())
	str.WriteString(")")
	return str.String()
}

type Typed_PostfixExpr struct {
	Expr
	Left Expr
	Op TokenType
}

func (t *Typed_PostfixExpr) Type() Typed_AstNodetype { return _Typed_PrefixExpr }
func (t *Typed_PostfixExpr) String() string {
	var str bytes.Buffer
	str.WriteString("(")
	str.WriteString(t.Left.String())
	switch t.Op {
	case Decrement: str.WriteString("--")
	case Increment: str.WriteString("++")
	}
	str.WriteString(")")
	return str.String()
}

type Typed_CallExpr struct {
	Expr
	Caller  Expr
	Args []Expr
}

func (t *Typed_CallExpr) Type() Typed_AstNodetype { return _Typed_CallExpr }
func (t *Typed_CallExpr) String() string {
	var str bytes.Buffer
	str.WriteString(t.Caller.String())
	str.WriteString("(")
	for i, v := range t.Args {
		str.WriteString(v.String())
		if i < len(t.Args)-1 { str.WriteString(",") }
	}
	str.WriteString(")")
	return str.String()
}

type Typed_MemberExpr struct {
	Expr
	Obj Expr
	Property Expr
	Computed bool
}

func (t *Typed_MemberExpr) Type() Typed_AstNodetype { return _Typed_MemberExpr }
func (t *Typed_MemberExpr) String() string {
	var str bytes.Buffer
	str.WriteString(t.Obj.String())
	if t.Computed {
		str.WriteString("[")
		str.WriteString(t.Property.String())
		str.WriteString("]")
	} else {
		str.WriteString(".")
		str.WriteString(t.Property.String())
	}
	return str.String()
}

package main

import (
	"bytes"
	"fmt"
	"strings"
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
	_ReturnStmt
	
	_Identifier
	_StringLiteral
	_FloatLiteral64
	_FloatLiteral32
	_IntLiteral
	_BoolLiteral
	_NullLiteral
	_HashLiteral
	_ArrayLiteral

	_AssginmentExpr
	_TernaryExpr
	_BinaryExpr
	_PrefixExpr
	_PostfixExpr
	_CallExpr
	_MemberExpr
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

type ExprStmt struct {
	Stmt
	Expression Expr
}

func (t *ExprStmt) Type() AstNodeType { return _ExprStmt }
func (t *ExprStmt) String() string    { return t.Expression.String()+";" }

// type PackageStmt struct {
// 	Stmt
// 	PackageName string
// }

// func (t *PackageStmt) Type() AstNodeType { return _PackageStmt }
// func (t *PackageStmt) String() string    { return "package " + t.PackageName }

type ImportStmt struct {
	Stmt
	PackageName string
	IsBuiltIn   bool
	IsNative bool
	ChangesName bool
	NewName string
}

func (t *ImportStmt) Type() AstNodeType { return _ImportStmt }
func (t *ImportStmt) String() string    {
	switch {
	case t.IsBuiltIn: return "using \"@/" + strings.TrimPrefix(t.PackageName, "core/") + "\""
	case t.IsNative: return "using \"#/" + t.PackageName + "\""
	default: return "using \"" + t.PackageName + "\";"
	}
}

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
	Loop []Stmt
}

func (t *WhileStmt) Type() AstNodeType { return _WhileStmt }
func (t *WhileStmt) String() string {
	var str bytes.Buffer
	str.WriteString("while (")
	str.WriteString(t.Condition.String())
	str.WriteString(")")
	str.WriteString("{")
	for _, v := range t.Loop {
		str.WriteString(v.String())
		if v.Type() != _ExprStmt {
			str.WriteString(";")
		}
	}
	str.WriteString("}")
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
	for _, v := range t.Body {
		str.WriteString(v.String())
		if v.Type() != _ExprStmt {
			str.WriteString(";")
		}
	}
	str.WriteString("}")
	return str.String()
}

type VarStmt struct {
	Stmt
	Name string
	IsConst bool
	Value Expr
	ObjType ObjectTypes
}

func (t *VarStmt) Type() AstNodeType { return _VarStmt }
func (t *VarStmt) String() string {
	var str bytes.Buffer
	// if t.IsConst { str.WriteString("const ") } else { str.WriteString("var ") }
	switch t.ObjType {
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

type FunctionStmt struct {
	Stmt
	Anonymous bool
	Name string
	Args []Arg
	Body []Stmt
	ObjType ObjectTypes
}

func (t *FunctionStmt) Type() AstNodeType { return _FunctionStmt }
func (t *FunctionStmt) String() string {
	var str bytes.Buffer
	str.WriteString("fn ")
	str.WriteString(t.Name)
	str.WriteString("(")
	for i, v := range t.Args {
		str.WriteString(fmt.Sprint(v.Type))
		str.WriteString(" ")
		str.WriteString(v.Value)
		if i < len(t.Args)-1 { str.WriteString(",") }
	}
	str.WriteString(") ")
	str.WriteString("-> ")
	switch t.ObjType {
		case _StringObject: str.WriteString("string ")
		case _BooleanObject: str.WriteString("boolean ")
		case _IntObject: str.WriteString("int ")
		case _FloatObject: str.WriteString("float ")
		case _NullObject: str.WriteString("void ")
	}
	str.WriteString("{")
	for _, v := range t.Body {
		str.WriteString(v.String())
		if v.Type() != _ExprStmt {
			str.WriteString(";")
		}
	}
	str.WriteString("}")
	return str.String()
}

type ReturnStmt struct {
	Stmt
	Expression Expr
}

func (t *ReturnStmt) Type() AstNodeType { return _ReturnStmt }
func (t *ReturnStmt) String() string {
	var str bytes.Buffer
	str.WriteString("return ")
	str.WriteString(t.Expression.String())
	return str.String()
}

type StringLiteral struct {
	Expr
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

type Identifier struct {
	Expr
	Value string
}

func (t *Identifier) Type() AstNodeType { return _Identifier }
func (t *Identifier) String() string { return t.Value }

type NullLiteral struct { Expr }

func (t *NullLiteral) Type() AstNodeType { return _NullLiteral }
func (t *NullLiteral) String() string { return "null" }

type FloatLiteral struct {
	Expr
	Value float64
}

func (t *FloatLiteral) Type() AstNodeType { return _FloatLiteral64 }
func (t *FloatLiteral) String() string { return fmt.Sprint(t.Value) }

type IntLiteral struct {
	Expr
	Value int
}

func (t *IntLiteral) Type() AstNodeType { return _IntLiteral }
func (t *IntLiteral) String() string { return fmt.Sprint(t.Value) }

type BoolLiteral struct {
	Expr
	Value bool
}

func (t *BoolLiteral) Type() AstNodeType { return _BoolLiteral }
func (t *BoolLiteral) String() string { return fmt.Sprint(t.Value) }

type HashLiteral struct {
	Expr `json:"-"`
	Pairs map[string]Expr
}

func (t *HashLiteral) Type() AstNodeType { return _HashLiteral }
func (t *HashLiteral) String() string {
	var str bytes.Buffer
	str.WriteString("{")
	keys := make([]string, 0, len(t.Pairs))
	for e := range t.Pairs {
		keys = append(keys, e)
	}
	i:=0
	for key, value := range t.Pairs {
		str.WriteString(key+":"+value.String())
		if i < len(keys)-1 { str.WriteString(",") }
		i++
	}
	str.WriteString("}")
	return str.String()
}

type ArrayLiteral struct {
	Expr `json:"-"`
	Elements []Expr
}

func (t *ArrayLiteral) Type() AstNodeType { return _ArrayLiteral }
func (t *ArrayLiteral) String() string {
	var str bytes.Buffer
	str.WriteString("[")
	for i, v := range t.Elements {
		str.WriteString(v.String())
		if i < len(t.Elements)-1 { str.WriteString(", ") }
	}
	str.WriteString("]")
	return str.String()
}

type AssignExpr struct {
	Expr
	Left Expr
	Right Expr
	Op TokenType
}

func (t *AssignExpr) Type() AstNodeType { return _AssginmentExpr }
func (t *AssignExpr) String() string {
	var str bytes.Buffer
	str.WriteString("(")
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
	str.WriteString(")")
	return str.String()
}

type TernaryExpr struct {
	Expr
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
	Expr
	Left Expr
	Right Expr
	Op TokenType
}

func (t *BinaryExpr) Type() AstNodeType { return _BinaryExpr }
func (t *BinaryExpr) String() string {
	var str bytes.Buffer
	str.WriteString("(")
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
	str.WriteString(")")
	return str.String()
}

type PrefixExpr struct {
	Expr
	Right Expr
	Op TokenType
}

func (t *PrefixExpr) Type() AstNodeType { return _PrefixExpr }
func (t *PrefixExpr) String() string {
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

type PostfixExpr struct {
	Expr
	Left Expr
	Op TokenType
}

func (t *PostfixExpr) Type() AstNodeType { return _PrefixExpr }
func (t *PostfixExpr) String() string {
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

type CallExpr struct {
	Expr
	Caller  Expr
	Args []Expr
}

func (t *CallExpr) Type() AstNodeType { return _CallExpr }
func (t *CallExpr) String() string {
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

type MemberExpr struct {
	Expr
	Obj Expr
	Property Expr
	Computed bool
}

func (t *MemberExpr) Type() AstNodeType { return _MemberExpr }
func (t *MemberExpr) String() string {
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

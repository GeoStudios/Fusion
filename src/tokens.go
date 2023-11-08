package main

type TokenType int

const (
	EOF TokenType = iota

	Identifer
	Integer
	Float
	String
	Boolean
	Null
	Void
	NativeFunction
	Function_Type
	HashMap
	Array

	// Comparison and Equality
	Equals
	NotEquals
	GreaterThan
	GreaterThanEqualTo
	LessThan
	LessThanEqualTo

	// Binary Operators
	Plus
	Minus
	Modulo
	Divide
	Multiply

	Exponent
	NthRoot

	And
	Or

	Logic_And
	Logic_Or

	// Assignment
	Basic_Assign
	Plus_Assign
	Minus_Assign
	Multiply_Assign
	Divide_Assign
	Modulo_Assign

	// Prefix/Postfix
	Decrement
	Increment
	Logic_Not

	Xor
	Not

	// Symbols
	Comma
	Dot
	QuesionMark
	Colon
	SemiColon
	Arrow
	Spread

	OpenParen
	CloseParen

	OpenBrace
	CloseBrace

	OpenBracket
	CloseBracket

	// Keywords
	Function
	Import
	Package
	Variable
	Constant
	If
	Else
	While
	Return
	As
)

type Token struct {
	Line     int
	Position int
	Literal  string
	Type     TokenType
}

func GetKeyword(value string) TokenType {
	m := map[string]TokenType{
		"string":         String,
		"bool":           Boolean,
		"int":            Integer,
		"float":          Float,
		"void":           Void,
		"Function":       Function_Type,
		"NativeFunction": NativeFunction,
		"HashMap":        HashMap,
		"Array":          Array,

		"fn":      Function,
		"using":   Import,
		"package": Package,
		"const":   Constant,
		"var":     Variable,
		"true":    Boolean,
		"false":   Boolean,
		"null":    Null,
		"if":      If,
		"else":    Else,
		"while":   While,
		"return":  Return,
		"as":      As,
	}
	if _, ok := m[value]; ok {
		return m[value]
	}
	return Identifer
}
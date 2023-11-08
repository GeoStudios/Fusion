package main

import (
	"bytes"
	"fmt"
	"log"
)

type ObjectTypes int

const (
	_NullObject       ObjectTypes = iota
	_IntObject        
	_FloatObject      
	_StringObject     
	_BooleanObject    
	_HashObject       
	_MemberObject     
	_ArrayObject      
	_FuncObject       
	_NativeFuncObject 
	_ReturnObject
)

type Object interface {
	Type() ObjectTypes
	String() string
}

func GetTypeFromToken(value TokenType) ObjectTypes {
	m := map[TokenType]ObjectTypes{
		String:  _StringObject,
		Boolean: _BooleanObject,
		Integer: _IntObject,
		Float:   _FloatObject,
		Void:    _NullObject,
		Function: _FuncObject,
		NativeFunction: _NativeFuncObject,
		HashMap: _HashObject,
		Array: _ArrayObject,
	}
	if _, ok := m[value]; ok {
		return m[value]
	}
	log.Fatalf("NOT A TYPE %v\n", value)
	return _NullObject
}

type NullObject struct{ Object }

func (s *NullObject) Type() ObjectTypes { return _NullObject }
func (s *NullObject) String() string { return "null" }

type IntObject struct {
	Object
	Value int
}

func (s *IntObject) Type() ObjectTypes { return _IntObject }
func (s *IntObject) String() string { return fmt.Sprint(s.Value) }

type FloatObject struct {
	Object
	Value float64
}

func (s *FloatObject) Type() ObjectTypes { return _FloatObject }
func (s *FloatObject) String() string { return fmt.Sprint(s.Value) }

type BooleanObject struct {
	Object
	Value bool
}

func (s *BooleanObject) Type() ObjectTypes { return _BooleanObject }
func (s *BooleanObject) String() string { return fmt.Sprint(s.Value) }

type StringObject struct {
	Object
	Value string
}

func (s *StringObject) Type() ObjectTypes { return _StringObject }
func (s *StringObject) String() string { return s.Value }

type FunctionObject struct {
	Object
	Name string
	Args []Arg
	Env *Env
	Body []Stmt
	RetType ObjectTypes
}

func (s *FunctionObject) Type() ObjectTypes { return _FuncObject }
func (s *FunctionObject) String() string { return "\"fn "+s.Name+"\"" }

type NativeCall func(args []Object, env *Env) Object

type NativeFunctionObject struct {
	Object
	Name string
	Call NativeCall
	RetType ObjectTypes
}

func (s *NativeFunctionObject) Type() ObjectTypes { return _NativeFuncObject }
func (s *NativeFunctionObject) String() string { return "\"func "+s.Name+"\"" }

type ReturnObject struct {
	Object
	Value Object
}

func (s *ReturnObject) Type() ObjectTypes { return _ReturnObject }
func (s *ReturnObject) String() string { return s.Value.String() }

type HashObject struct {
	Object
	Expr `json:"-"`
	Pairs map[string]Object
}

func (t *HashObject) Type() ObjectTypes { return _HashObject }
func (t *HashObject) String() string {
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

type ArrayObject struct {
	Expr `json:"-"`
	Elements []Object
}

func (t *ArrayObject) Type() ObjectTypes { return _ArrayObject }
func (t *ArrayObject) String() string {
	var str bytes.Buffer
	str.WriteString("[")
	for i, v := range t.Elements {
		str.WriteString(v.String())
		if i < len(t.Elements)-1 { str.WriteString(", ") }
	}
	str.WriteString("]")
	return str.String()
}

func UnWrapAsInt(o Object) int {
	switch o.Type() {
		case _FloatObject: return int(o.(*FloatObject).Value)
		case _IntObject: return o.(*IntObject).Value
		default: log.Fatalf("Cannot Convert %v To Int", o); return 0
	}
}

func UnWrapAsFloat(o Object) float64 {
	switch o.Type() {
		case _FloatObject: return o.(*FloatObject).Value
		case _IntObject: return float64(o.(*IntObject).Value)
		default: log.Fatalf("Cannot Convert %v To Float", o); return 0
	}
}

func IsTruthy(c Object) bool {
	switch c.Type() {
	case _NullObject: return false
	case _IntObject: return UnWrapAsFloat(c) > 0
	case _FloatObject: return UnWrapAsFloat(c) > 0
	case _StringObject: return c.(*StringObject).Value != ""
	case _BooleanObject: return c.(*BooleanObject).Value
	case _HashObject: return len(c.(*HashObject).Pairs) != 0
	case _MemberObject: return true
	case _ArrayObject: return len(c.(*ArrayObject).Elements) != 0
	case _FuncObject: return true
	case _NativeFuncObject: return true
	case _ReturnObject: return IsTruthy(c.(*ReturnObject).Value)
	default: return false
	}
}

func getValue(c Object) Object {
	switch c.Type() {
	case _ReturnObject: return getValue(c.(*ReturnObject).Value)
	default: return c
	}
}
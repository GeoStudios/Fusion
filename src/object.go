package main

import (
	"fmt"
	"log"
)

type ObjectTypes int

const (
	_NullObject       ObjectTypes = iota
	_IntObject        ObjectTypes = iota
	_FloatObject      ObjectTypes = iota
	_StringObject     ObjectTypes = iota
	_BooleanObject    ObjectTypes = iota
	_HashObject       ObjectTypes = iota
	_MemberObject     ObjectTypes = iota
	_ArrayObject      ObjectTypes = iota
	_FuncObject       ObjectTypes = iota
	_NativeFuncObject ObjectTypes = iota
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
	}
	if _, ok := m[value]; ok {
		return m[value]
	}
	log.Fatalln("NOT A TYPE")
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
func (s *StringObject) String() string { return "\""+s.Value+"\"" }

type FunctionObject struct {
	Object
	Name string
	Args []Arg
	Env *Env
	Body Stmt
	RetType ObjectTypes
}

func (s *FunctionObject) Type() ObjectTypes { return _FuncObject }
func (s *FunctionObject) String() string { return "\""+s.Body.String()+"\"" }

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
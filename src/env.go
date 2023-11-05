package main

import (
	"log"
)

type Env struct {
	Parent *Env
	Vars   map[string]Object
	Consts map[string]bool
}

func New_Env(parent *Env) *Env {
	return &Env{
		Parent: parent,
		Vars:   map[string]Object{},
		Consts: map[string]bool{},
	}
}

func (s *Env) DeclareVar(name string, val Object, cons bool, typ ObjectTypes) Object {
	if _, ok := s.Vars[name]; ok { log.Fatalf("Cannot declare variable. Variable is already declared %v\n", name) }
	if val.Type() != typ { log.Fatalf("Value Is Different Type Then Variable %v\n", name) }
	s.Vars[name] = val
	if (cons) { s.Consts[name] = true }
	return val
}

func (s *Env) AssignVar(name string, val Object) Object {
	env := s.Resolve(name)
	if s.Consts[name] { log.Fatalf("Cannot Reassign Constant %v", name) }
	oldVal := s.LookupVar(name)
	if oldVal.Type() != val.Type() { log.Fatalf("Cannot reassign variable with different type, Original %v, New %v", oldVal.Type(), val.Type()) }
	env.Vars[name] = val
	return val
}

func (s *Env) LookupVar(name string) Object { return s.Resolve(name).Vars[name] }

func (s *Env) Resolve(name string) *Env {
	if _, ok := s.Vars[name]; ok { return s }
	if s.Parent == nil { log.Fatalf("Variable %v Does Not Exist.", name)}
	return s.Parent.Resolve(name)
}
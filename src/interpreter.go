package main

import (
	"log"
)

type Interpreter struct{}

func New_Interpreter() *Interpreter { return &Interpreter{} }

func (s *Interpreter) Eval(node AstNode, env *Env) Object {
	switch node.Type() {
	case _NullLiteral: return &NullObject{}
	case _IntLiteral: return &IntObject{Value: node.(*IntLiteral).Value}
	case _FloatLiteral64: return &FloatObject{Value: node.(*FloatLiteral).Value}
	case _BoolLiteral: return &BooleanObject{Value: node.(*BoolLiteral).Value}
	case _StringLiteral: return &StringObject{Value: node.(*StringLiteral).Value}
	case _Program: {
		var lastEval Object
		for _, v := range node.(*Program).Body { lastEval = s.Eval(v, env) }
		return lastEval
	} 
	case _BlockStmt: {
		var lastEval Object
		for _, v := range node.(*BlockStmt).Body { lastEval = s.Eval(v, env) }
		return lastEval
	}
	case _FunctionStmt: return s.EvaluateFunction(node.(*FunctionStmt), env)
	case _CallExpr: return s.EvaluateCallExpr(node.(*CallExpr), env)
	case _Identifier: return env.LookupVar(node.(*Identifier).Value)
	case _ExprStmt: return s.Eval(node.(*ExprStmt).Expression, env)
	case _VarStmt: return s.EvaluateVarStmt(node.(*VarStmt), env)
	case _PrefixExpr: return s.EvaluatePrefixExpr(node.(*PrefixExpr), env)
	case _BinaryExpr: return s.EvaluateBinaryExpr(node.(*BinaryExpr), env)
	default:
		log.Fatalln("Could Not Execute Node:", node)
		return &NullObject{}
	}
}

func (s *Interpreter) EvaluateFunction(node *FunctionStmt, env *Env) Object {
	f := &FunctionObject{
		Name: node.Name,
		Args: node.Args,
		Env: env,
		Body: node.Body,
		RetType: node.ObjType,
	}
	env.DeclareVar(node.Name, f, false, _FuncObject)
	return f
}

func (s *Interpreter) EvaluateCallExpr(node *CallExpr, env *Env) Object {
	args := make([]Object, 0)
	for _, arg := range node.Args {
		args = append(args, s.Eval(arg, env))
	}
	fn := s.Eval(node.Caller, env)
	if fn.Type() == _FuncObject {
		fnn := fn.(*FunctionObject)
		scope := New_Env(fnn.Env)
		if len(fnn.Args) != len(args) {
			log.Fatalf("Too many/little arguments %v/%v arguments filled in func %v",
				len(args),
				len(fnn.Args),
				fnn.Name,
			)
		}
		for i, v := range fnn.Args {
			if args[i].Type() != v.Type {
				log.Fatalf("Arg %v does not have the same type as %v\n", args[i], v)
			}
			scope.DeclareVar(v.Value, args[i], false, v.Type)
		}
		return s.Eval(fnn.Body, scope)
	}
	return nil
}

func (s *Interpreter) EvaluateVarStmt(node *VarStmt, env *Env) Object {
	env.DeclareVar(node.Name, s.Eval(node.Value, env), node.IsConst, node.ObjType)
	return &NullObject{}
}

func (s *Interpreter) EvaluatePrefixExpr(node *PrefixExpr, env *Env) Object {
	right := s.Eval(node.Right, env)
	switch node.Op {
		case Minus:
			switch right.Type() {
				case _IntObject: return &IntObject{Value: -UnWrapAsInt(right)}
				case _FloatObject: return &FloatObject{Value: -UnWrapAsFloat(right)}
				default: log.Fatalf("Type Not Allowed In Prefix Expr %v", right.Type())
			}
		case Plus: return right
	}
	return &NullObject{};
}

func (s *Interpreter) EvaluateBinaryExpr(node *BinaryExpr, env *Env) Object {
	left := s.Eval(node.Left, env)
	right := s.Eval(node.Right, env)
	l := UnWrapAsFloat(left)
	r := UnWrapAsFloat(right)
	op := (left.Type() == _IntObject) && (right.Type() == _IntObject)
	switch node.Op {
		case Plus:
			if op { return &IntObject{Value: int(l + r)} }
			return &FloatObject{Value: l + r}
		case Minus: 
			if op { return &IntObject{Value: int(l - r)} }
			return &FloatObject{Value: l - r}
		case Multiply:
			if op { return &IntObject{Value: int(l * r)} }
			return &FloatObject{Value: l * r}
		case Divide:
			if op { return &IntObject{Value: int(l / r)} }
			return &FloatObject{Value: l / r}
	}
	return nil;
}
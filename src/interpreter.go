package main

import (
	"log"
	"math"
	"os"
	"path"
	"strings"
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
	case _HashLiteral:
		pairs := map[string]Object{}
		for k, e := range node.(*HashLiteral).Pairs { pairs[k] = s.Eval(e, env) }
		return &HashObject{Pairs: pairs}
	case _Program: {
		var lastEval Object
		for _, v := range node.(*Program).Body { lastEval = s.Eval(v, env) }
		return lastEval
	} 
	case _ReturnStmt: return s.Eval(node.(*ReturnStmt).Expression, env)
	case _BlockStmt: {
		ee := New_Env(env)
		var lastEval Object
		for _, v := range node.(*BlockStmt).Body {
			q := s.Eval(v, ee)
			if v.Type() == _ReturnStmt {
				lastEval = q
				break
			}
		}
		return lastEval
	}
	case _ImportStmt: {
		im := node.(*ImportStmt)
		switch {
		case im.IsNative: DeclareNatives(im.PackageName, env, im.NewName, im.ChangesName)
		case im.IsBuiltIn: s.EvaluateImportBuiltIn(im, env)
		default:
			s.EvaluateImport(im, env)
		}
		return &NullObject{}
	}
	case _FunctionStmt: return s.EvaluateFunction(node.(*FunctionStmt), env)
	case _CallExpr: return s.EvaluateCallExpr(node.(*CallExpr), env)
	case _Identifier: return env.LookupVar(node.(*Identifier).Value)
	case _ExprStmt: return s.Eval(node.(*ExprStmt).Expression, env)
	case _VarStmt: return s.EvaluateVarStmt(node.(*VarStmt), env)
	case _PrefixExpr: return s.EvaluatePrefixExpr(node.(*PrefixExpr), env)
	case _BinaryExpr: return s.EvaluateBinaryExpr(node.(*BinaryExpr), env)
	case _MemberExpr: return s.EvaluateMemberExpr(node.(*MemberExpr), env)
	case _IfStmt: return s.EvaluateIfStmt(node.(*IfStmt), env)
	case _WhileStmt: return s.EvaluateWhileStmt(node.(*WhileStmt), env)
	case _ArrayLiteral: {
		ele := []Object{}
		for _, v := range node.(*ArrayLiteral).Elements { ele = append(ele, s.Eval(v, env)) }
		return &ArrayObject{ Elements: ele }
	}
	case _AssginmentExpr: return s.EvaluateAssignExpr(node.(*AssignExpr), env)
	default:
		log.Fatalln("Could Not Execute Node:", node)
		return &NullObject{}
	}
}

func (s *Interpreter) EvaluateImportBuiltIn(node *ImportStmt, env *Env) Object {
	lex := New_Lexer(ReadEmbedFile(node.PackageName))
	_, fileName := path.Split(node.PackageName)
	fileName = strings.Split(fileName, ".")[0]
	parse := New_Parser(lex.Tokenize(), fileName)
	nenv := New_Env(env)
	ast := parse.ProduceAst()
	s.Eval(ast, nenv)
	name := ast.(*Program).PackageName
	if node.ChangesName { name = node.NewName }
	env.DeclareVar(name, s.MapEnvAsHash(nenv), true, _HashObject)
	return &NullObject{}
}

func (s *Interpreter) EvaluateImport(node *ImportStmt, env *Env) Object {
	lex := New_Lexer(os.ReadFile(path.Join(CurrentFilePath, node.PackageName)))
	_, fileName := path.Split(node.PackageName)
	fileName = strings.Split(fileName, ".")[0]
	parse := New_Parser(lex.Tokenize(), fileName)
	nenv := New_Env(env)
	ast := parse.ProduceAst()
	s.Eval(ast, nenv)
	name := ast.(*Program).PackageName
	if node.ChangesName { name = node.NewName }
	env.DeclareVar(name, s.MapEnvAsHash(nenv), true, _HashObject)
	return &NullObject{}
}

func (s *Interpreter) MapEnvAsHash(env *Env) Object {
	m := &HashObject{Pairs: map[string]Object{}}
	for name, value := range env.Vars { m.Pairs[name] = value }
	return m
}

func (s *Interpreter) EvaluateAssignExpr(node *AssignExpr, env *Env) Object {
	if node.Op == Basic_Assign { return s.Eval(node.Right, env) }
	l := UnWrapAsFloat(s.Eval(node.Left, env))
	r := UnWrapAsFloat(s.Eval(node.Right, env))
	name := node.Left.(*Identifier).Value
	var out Object
	switch node.Op {
	case Plus_Assign:
		switch s.Eval(node.Left, env).Type() {
			case _FloatObject: out = &FloatObject{Value: l+r}
			case _IntObject: out = &IntObject{Value: int(l+r)}
		}	
	case Minus_Assign:
		switch s.Eval(node.Left, env).Type() {
			case _FloatObject: out = &FloatObject{Value: l-r}
			case _IntObject: out = &IntObject{Value: int(l-r)}
		}
	case Multiply_Assign:
		switch s.Eval(node.Left, env).Type() {
			case _FloatObject: out = &FloatObject{Value: l*r}
			case _IntObject: out = &IntObject{Value: int(l*r)}
		}
	case Divide_Assign:
		switch s.Eval(node.Left, env).Type() {
			case _FloatObject: out = &FloatObject{Value: l/r}
			case _IntObject: out = &IntObject{Value: int(l/r)}
		}
	case Modulo_Assign:
		switch s.Eval(node.Left, env).Type() {
			case _FloatObject: out = &FloatObject{Value: math.Remainder(l, r)}
			case _IntObject: out = &IntObject{Value: int(math.Remainder(l, r))}
		}
	}
	return env.AssignVar(name, out)
}

func (s *Interpreter) EvaluateIfStmt(node *IfStmt, env *Env) Object {
	privEnv := New_Env(env)
	condition := s.Eval(node.Condition, privEnv)
	if IsTruthy(condition) {
		s.Eval(node.OnTrue, privEnv)
	} else if node.OnFalse != nil {
		s.Eval(node.OnFalse, privEnv)
	}
	return &NullObject{}
}

func (s *Interpreter) EvaluateWhileStmt(node *WhileStmt, env *Env) Object {
	for {
		privEnv := New_Env(env)
		if !IsTruthy(s.Eval(node.Condition, privEnv)) { break }
		for _, v := range node.Loop {
			s.Eval(v, privEnv)
		}
	}
	return &NullObject{}
}

func (s *Interpreter) EvaluateMemberExpr(node *MemberExpr, env *Env) Object {
	var expr Expr = node
	keys := []Object{}
	for {
		if expr.Type() == _MemberExpr {
			switch expr.(*MemberExpr).Property.Type() {
				case _Identifier: keys = append(keys, &StringObject{Value: expr.(*MemberExpr).Property.(*Identifier).Value})
				case _StringLiteral: keys = append(keys, s.Eval(expr.(*MemberExpr).Property.(*StringLiteral), env))
				case _IntLiteral: keys = append(keys, s.Eval(expr.(*MemberExpr).Property.(*IntLiteral), env))
			}
			expr = expr.(*MemberExpr).Obj
		} else {
			switch expr.Type() {
				case _Identifier: keys = append(keys, &StringObject{Value: expr.(*Identifier).Value})
				case _StringLiteral: keys = append(keys, s.Eval(expr.(*StringLiteral), env))
				case _IntLiteral: keys = append(keys, s.Eval(expr.(*IntLiteral), env))
			}
			break
		}
	}
	for i, j := 0, len(keys)-1; i < j; i, j = i+1, j-1 {
		keys[i], keys[j] = keys[j], keys[i]
	}
	keys = keys[1:]
	var nest Object
	parent := s.Eval(expr, env)
	nest = parent
	for _, v := range keys {
		if nest.Type() == _ArrayObject && v.Type() == _IntObject {
			nest = nest.(*ArrayObject).Elements[v.(*IntObject).Value]
		} else {
			nest = nest.(*HashObject).Pairs[v.(*StringObject).Value]
		}

		if nest == nil {
			log.Fatalf("KEY %v DOES NOT EXIST ON NODE %v", v, node)
		}
	}
	return nest
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
	switch fn.Type() {
	case _NativeFuncObject: {
		return fn.(*NativeFunctionObject).Call(args, env)
	}
	case _FuncObject: {
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
		var lastEval Object = &NullObject{}
		for _, v := range fnn.Body {
			q := s.Eval(v, scope)
			if v.Type() == _ReturnStmt {
				lastEval = q
				if getValue(lastEval).Type() != fnn.RetType {
					log.Fatalf("Function %v is returning type %v when it needs to return type %v",
						fnn.Name,
						lastEval.(*ReturnObject).Value.Type(),
						fnn.RetType,
					)
				}
				break
			}
		}
		return lastEval
	}
	}
	return nil
}

func (s *Interpreter) EvaluateVarStmt(node *VarStmt, env *Env) Object {
	env.DeclareVar(node.Name, s.Eval(node.Value, env), node.IsConst, node.ObjType)
	return &NullObject{}
}

func (s *Interpreter) EvaluatePrefixExpr(node *PrefixExpr, env *Env) Object {
	right := s.Eval(node.Right, env)
	var value float64
	switch node.Op {
		case Logic_Not: return &BooleanObject{Value: !IsTruthy(right)}
		// case Minus: value = -UnWrapAsFloat(right)
		// case Plus: value = UnWrapAsFloat(right)
	}
	switch right.Type() {
		case _IntObject: return &IntObject{Value: int(value)}
		case _FloatObject: return &FloatObject{Value: value}
		default: log.Fatalf("Type Not Allowed In Prefix Expr %v", right.Type()); return &NullObject{}
	}
}

func (s *Interpreter) EvaluateBinaryExpr(node *BinaryExpr, env *Env) Object {
	left := s.Eval(node.Left, env)
	right := s.Eval(node.Right, env)
	l := UnWrapAsFloat(left)
	r := UnWrapAsFloat(right)
	inn := (left.Type() == _IntObject) && (right.Type() == _IntObject)
	var res float64
	switch node.Op {
		case Plus: res = l + r
		case Minus: res = l - r
		case Multiply: res = l * r
		case Divide: res = l / r
		case Modulo: res = math.Remainder(l,r)
		case Exponent: res = math.Pow(l, r)
		case NthRoot: res = math.Pow(l, 1/r)

		case Equals: return &BooleanObject{Value: l == r}
		case NotEquals: return &BooleanObject{Value: l != r}

		case GreaterThan: return &BooleanObject{Value: l > r}
		case GreaterThanEqualTo: return &BooleanObject{Value: l >= r}
		case LessThan: return &BooleanObject{Value: l < r}
		case LessThanEqualTo: return &BooleanObject{Value: l <= r}
	}
	if inn {
		return &IntObject{Value: int(res)}
	}; return &FloatObject{Value: res}
}
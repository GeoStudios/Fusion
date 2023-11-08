package main

import "math"

func Declare_StrongMath(env *Env, name string) {
	if env.Vars[name] != nil {
		return
	}
	std := &HashObject{Pairs: map[string]Object{}}
	std.Pairs["Sin"] = &NativeFunctionObject{
		Name: "Sin",
		Call: func(args []Object, env *Env) Object {
			return &FloatObject{Value: math.Sin(UnWrapAsFloat(args[0]))}
		},
		RetType: _FloatObject,
	}
	std.Pairs["Sinh"] = &NativeFunctionObject{
		Name: "Sinh",
		Call: func(args []Object, env *Env) Object {
			return &FloatObject{Value: math.Sinh(UnWrapAsFloat(args[0]))}
		},
		RetType: _FloatObject,
	}
	std.Pairs["ASin"] = &NativeFunctionObject{
		Name: "ASin",
		Call: func(args []Object, env *Env) Object {
			return &FloatObject{Value: math.Asin(UnWrapAsFloat(args[0]))}
		},
		RetType: _FloatObject,
	}
	std.Pairs["ASinh"] = &NativeFunctionObject{
		Name: "ASinh",
		Call: func(args []Object, env *Env) Object {
			return &FloatObject{Value: math.Asinh(UnWrapAsFloat(args[0]))}
		},
		RetType: _FloatObject,
	}

	std.Pairs["Cosine"] = &NativeFunctionObject{
		Name: "Cosine",
		Call: func(args []Object, env *Env) Object {
			return &FloatObject{Value: math.Cos(UnWrapAsFloat(args[0]))}
		},
		RetType: _FloatObject,
	}
	std.Pairs["Cosineh"] = &NativeFunctionObject{
		Name: "Cosineh",
		Call: func(args []Object, env *Env) Object {
			return &FloatObject{Value: math.Cosh(UnWrapAsFloat(args[0]))}
		},
		RetType: _FloatObject,
	}
	std.Pairs["ACosine"] = &NativeFunctionObject{
		Name: "ACosine",
		Call: func(args []Object, env *Env) Object {
			return &FloatObject{Value: math.Acos(UnWrapAsFloat(args[0]))}
		},
		RetType: _FloatObject,
	}
	std.Pairs["ACosineh"] = &NativeFunctionObject{
		Name: "ACosineh",
		Call: func(args []Object, env *Env) Object {
			return &FloatObject{Value: math.Acosh(UnWrapAsFloat(args[0]))}
		},
		RetType: _FloatObject,
	}

	std.Pairs["Tan"] = &NativeFunctionObject{
		Name: "Tan",
		Call: func(args []Object, env *Env) Object {
			return &FloatObject{Value: math.Tan(UnWrapAsFloat(args[0]))}
		},
		RetType: _FloatObject,
	}
	std.Pairs["Tanh"] = &NativeFunctionObject{
		Name: "Tanh",
		Call: func(args []Object, env *Env) Object {
			return &FloatObject{Value: math.Tanh(UnWrapAsFloat(args[0]))}
		},
		RetType: _FloatObject,
	}
	std.Pairs["ATan"] = &NativeFunctionObject{
		Name: "ATan",
		Call: func(args []Object, env *Env) Object {
			return &FloatObject{Value: math.Atan(UnWrapAsFloat(args[0]))}
		},
		RetType: _FloatObject,
	}
	std.Pairs["ATanh"] = &NativeFunctionObject{
		Name: "ATanh",
		Call: func(args []Object, env *Env) Object {
			return &FloatObject{Value: math.Atanh(UnWrapAsFloat(args[0]))}
		},
		RetType: _FloatObject,
	}
	std.Pairs["ATan2"] = &NativeFunctionObject{
		Name: "ATan2",
		Call: func(args []Object, env *Env) Object {
			return &FloatObject{Value: math.Atan2(UnWrapAsFloat(args[0]),
			UnWrapAsFloat(args[1]))}
		},
		RetType: _FloatObject,
	}

	env.DeclareVar(name, std, true, _HashObject)
}
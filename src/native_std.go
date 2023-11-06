package main

import "fmt"

func Declare_STD(env *Env) {
	if env.Vars["std"] != nil { return }
	std := &HashObject{Pairs: map[string]Object{}}
	std.Pairs["print"] = &NativeFunctionObject{
		Name: "print",
		Call: func(args []Object, env *Env) Object {
			for _, v := range args {
				fmt.Println(v)
			}
			return &NullObject{}
		},
		RetType: _NullObject,
	}
	env.DeclareVar("std", std, true, _HashObject)
}
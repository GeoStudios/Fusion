package main

import (
	"fmt"
	"log"
)

func Declare_STD(env *Env, name string) {
	if env.Vars[name] != nil { return }
	std := &HashObject{Pairs: map[string]Object{}}
	std.Pairs["print"] = &NativeFunctionObject{
		Name: "print",
		Call: func(args []Object, env *Env) Object {
			for _, v := range args {
				fmt.Print(v)
			}
			return &NullObject{}
		},
		RetType: _NullObject,
	}
	std.Pairs["println"] = &NativeFunctionObject{
		Name: "println",
		Call: func(args []Object, env *Env) Object {
			for _, v := range args {
				fmt.Println(v)
			}
			return &NullObject{}
		},
		RetType: _NullObject,
	}
	std.Pairs["concat"] = &NativeFunctionObject{
		Name: "concat",
		Call: func(args []Object, env *Env) Object {
			newStr := ""
			for _, v := range args {
				if v.Type() != _StringObject {
					log.Fatalf("Value %v is not of type string in std.concat.", v)
				}
				newStr += v.(*StringObject).Value
			}
			return &StringObject{Value: newStr}
		},
		RetType: _StringObject,
	}

	std.Pairs["Argv"] = func() *ArrayObject {
		a := &ArrayObject{Elements: []Object{}}
		for _, v := range argv {
			a.Elements = append(a.Elements, &StringObject{Value: v})
		}
		return a
	}()
	std.Pairs["LineSeparator"] = &StringObject{Value: "\n"}
	std.Pairs["CarriageRetrun"] = &StringObject{Value: "\r"}
	std.Pairs["LineFeed"] = &StringObject{Value: "\f"}
	std.Pairs["BackSpace"] = &StringObject{Value: "\b"}
	std.Pairs["Tab"] = &StringObject{Value: "\t"}
	env.DeclareVar(name, std, true, _HashObject)
}
package main

func DeclareNatives(pkg string, env *Env) {
	switch pkg {
	case "std":
		Declare_STD(env)
	}
}

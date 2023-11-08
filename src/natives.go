package main

func DeclareNatives(pkg string, env *Env, newName string, changename bool) {
	name := pkg
	if changename {
		name = newName
	}
	switch pkg {
	case "std":
		Declare_STD(env, name)
	case "StrongMath":
		Declare_StrongMath(env, name)
	}
}

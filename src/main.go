package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path"
	"strings"
)

// go:embed lib
var lib_e embed.FS

func ReadEmbedFile(filePath string) string {
	fileSys, _ := fs.Sub(lib_e, path.Join("lib_e", path.Dir(filePath)))
	_, fileName := path.Split(filePath)
	content, _ := fs.ReadFile(fileSys, fileName)
	return string(content)
}

func main() {
	argv := os.Args[1:]
	file, args := argv[0], argv[1:]

	fmt.Println(file)
	fmt.Println(args)

	_, fileName := path.Split(file)
	fileName = strings.Split(fileName, ".")[0]
	lex := New_Lexer(os.ReadFile(file))
	par := New_Parser(lex.Tokenize(), fileName)
	inp := New_Interpreter()
	inp.Eval(par.ProduceAst(), New_Env(nil))
}

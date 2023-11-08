package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
)

//go:embed lib
var lib embed.FS

func ReadEmbedFile(filePath string) ([]byte, error) {
	fileSys, _ := fs.Sub(lib, path.Join("lib", path.Dir(filePath)))
	_, fileName := path.Split(filePath)
	content, err := fs.ReadFile(fileSys, fileName)
	return content, err
}

var CurrentFilePath string
var argv = os.Args[1:]
func main() {
	
	file, args := argv[0], argv[1:]
	argv = args
	fmt.Println(file)
	fmt.Println(args)
	CurrentFilePath, _ = filepath.Abs(file)
	CurrentFilePath = filepath.Dir(CurrentFilePath)
	
	_, fileName := path.Split(file)
	fileName = strings.Split(fileName, ".")[0]
	lex := New_Lexer(os.ReadFile(file))
	par := New_Parser(lex.Tokenize(), fileName)
	inp := New_Interpreter()
	inp.Eval(par.ProduceAst(), New_Env(nil))
}

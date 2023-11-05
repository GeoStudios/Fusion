package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path"
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

	lex := New_Lexer(os.ReadFile(file))
	par := New_Parser(lex.Tokenize())
	fmt.Println(par.ProduceAst().String())	
}

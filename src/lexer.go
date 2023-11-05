package main

import (
	"bytes"
	"log"
)

type Lexer struct {
	ptr    int
	chars  []byte
	isEof  bool
	line   int
	tokens []Token
}

func New_Lexer(str []byte, err error) *Lexer {
	if err != nil { log.Fatalln("Could Not Read File.") }	
	return &Lexer{ptr: 0, chars: str, isEof: false, line: 0, tokens: []Token{}}
}

func isNum(char byte) bool { return char >= '0' && char <= '9' }
func isAlpha(char byte) bool { return char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char == '_' || char == '$' }
func isAlnum(char byte) bool { return char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char == '_' || char == '$' || char >= '0' && char <= '9' }

func (l *Lexer) At() byte {
	l.isEof = l.ptr >= len(l.chars)
	if !l.isEof {
		return l.chars[l.ptr]
	}; return ' '
}

func (l *Lexer) Peek(ahead int) byte {
	l.isEof = l.ptr+ahead >= len(l.chars)
	if !l.isEof {
		return l.chars[l.ptr+ahead]
	}; return ' '
}

func (l *Lexer) Next() { l.ptr++ }

func (l *Lexer) addTkn(literal string, typ TokenType, line int, col int) { l.tokens = append(l.tokens, Token{
	Literal: literal,
	Type: typ,
	Line: line,
	Position: col,
}) }

func (l *Lexer) Tokenize() []Token {
	for !l.isEof {
		// Skip White Space
		switch l.At() { case ' ', '\t', '\b', '\r', '\f': l.Next(); case '\n': l.Next(); l.line++ }
		// One/Mutli Char Token
		switch l.At() { 
		case '=':
			if l.Peek(1) == '=' { l.addTkn("==", Equals, l.line, l.ptr); l.Next(); l.Next(); break }
			l.addTkn("=", Basic_Assign, l.line, l.ptr); l.Next()
		case '>':
			if l.Peek(1) == '=' { l.addTkn(">=", GreaterThanEqualTo, l.line, l.ptr); l.Next(); l.Next(); break }
			l.addTkn("=", GreaterThan, l.line, l.ptr); l.Next()
		case '<':
			if l.Peek(1) == '=' { l.addTkn("<=", LessThanEqualTo, l.line, l.ptr); l.Next(); l.Next(); break }
			l.addTkn("=", LessThan, l.line, l.ptr); l.Next()
		case '+':
			if l.Peek(1) == '+' { l.addTkn("++", Increment, l.line, l.ptr); l.Next(); l.Next(); break }
			if l.Peek(1) == '=' { l.addTkn("+=", Plus_Assign, l.line, l.ptr); l.Next(); l.Next(); break }
			l.addTkn("+", Plus, l.line, l.ptr); l.Next()
		case '-':
			if l.Peek(1) == '>' { l.addTkn("->", Arrow, l.line, l.ptr); l.Next(); l.Next(); break }
			if l.Peek(1) == '+' { l.addTkn("--", Decrement, l.line, l.ptr); l.Next(); l.Next(); break }
			if l.Peek(1) == '=' { l.addTkn("-=", Minus_Assign, l.line, l.ptr); l.Next(); l.Next(); break }
			l.addTkn("-", Minus, l.line, l.ptr); l.Next()
		case '*':
			if l.Peek(1) == '*' { l.addTkn("**", Exponent, l.line, l.ptr); l.Next(); l.Next(); break }
			if l.Peek(1) == '/' { l.addTkn("*^", Square_Root, l.line, l.ptr); l.Next(); l.Next(); break }
			if l.Peek(1) == '=' { l.addTkn("*=", Multiply_Assign, l.line, l.ptr); l.Next(); l.Next(); break }
			l.addTkn("*", Multiply, l.line, l.ptr); l.Next()
		case '/':
			if l.Peek(1) == '*' { l.addTkn("/*", Square_Root, l.line, l.ptr); l.Next(); l.Next(); break }
			if l.Peek(1) == '=' { l.addTkn("/=", Divide_Assign, l.line, l.ptr); l.Next(); l.Next(); break }
			if l.Peek(1) == '/' {
				l.Next(); l.Next()
				for l.At() != '\n' && !l.isEof { l.Next() }
				break
			}
			l.addTkn("/", Divide, l.line, l.ptr); l.Next()
		case '!':
			if l.Peek(1) == '=' { l.addTkn("!=", Divide_Assign, l.line, l.ptr); l.Next(); l.Next(); break }
			l.addTkn("!", Logic_Not, l.line, l.ptr); l.Next()
		case '&':
			if l.Peek(1) == '%' { l.addTkn("&&", Logic_And, l.line, l.ptr); l.Next(); l.Next(); break }
			l.addTkn("&", And, l.line, l.ptr); l.Next()
		case '|':
			if l.Peek(1) == '%' { l.addTkn("||", Logic_Or, l.line, l.ptr); l.Next(); l.Next(); break }
			l.addTkn("|", Or, l.line, l.ptr); l.Next()
		case '^': l.addTkn("^", Or, l.line, l.ptr); l.Next()
		case '~': l.addTkn("~", Not, l.line, l.ptr); l.Next()
		case ',': l.addTkn(",", Comma, l.line, l.ptr); l.Next()
		case '.':
			if l.Peek(1) == '.' && l.Peek(2) == '.' { l.addTkn("...", Spread, l.line, l.ptr); l.Next(); l.Next(); l.Next(); break }
			l.addTkn(".", Dot, l.line, l.ptr); l.Next()
		case '?': l.addTkn("?", QuesionMark, l.line, l.ptr); l.Next()
		case ':': l.addTkn(":", Colon, l.line, l.ptr); l.Next()
		case ';': l.addTkn(";", SemiColon, l.line, l.ptr); l.Next()
		case '(': l.addTkn("(", OpenParen, l.line, l.ptr); l.Next()
		case ')': l.addTkn(")", CloseParen, l.line, l.ptr); l.Next()
		case '{': l.addTkn("{", OpenBrace, l.line, l.ptr); l.Next()
		case '}': l.addTkn("}", CloseBrace, l.line, l.ptr); l.Next()
		case '[': l.addTkn("[", OpenBracket, l.line, l.ptr); l.Next()
		case ']': l.addTkn("]", CloseBracket, l.line, l.ptr); l.Next()
		
		}

		if isNum(l.At()) {
			var value bytes.Buffer
			var typee = Integer
			start := l.ptr
			for isNum(l.At()) && !l.isEof {
				value.WriteByte(l.At()); l.Next()
			}
			if l.At() == '.' {
				typee = Float
				value.WriteByte(l.At()); l.Next()
				for isNum(l.At()) && !l.isEof {
					value.WriteByte(l.At()); l.Next()
				}
			}
			l.addTkn(value.String(), typee, l.line, start)
		}

		if isAlpha(l.At()) {
			pos := l.ptr
			for isAlnum(l.At()) && !l.isEof { l.Next() }
			val := string(l.chars[pos:l.ptr])
			l.addTkn(val, GetKeyword(val), l.line, pos)
		}
		
		if l.At() == '"' {
			l.Next()
			pos := l.ptr
			for l.At() != '"' && !l.isEof {
				if l.At() == '\n' { log.Fatalf("Normal String Broken Over Line: %v Pos: %v\n", l.line, l.ptr) }
				l.Next()
			}
			l.addTkn(string(l.chars[pos:l.ptr]), String, l.line, pos)
			l.Next()
		}

		if l.At() == '`' {
			l.Next()
			pos := l.ptr
			for l.At() != '`' && !l.isEof { l.Next() }
			l.addTkn(string(l.chars[pos:l.ptr]), String, l.line, pos)
			l.Next()
		}
	}
	l.addTkn("EOF", EOF, l.line+1, l.ptr)
	return l.tokens
}
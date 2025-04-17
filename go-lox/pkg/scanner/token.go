package scanner

import "fmt"

type TokenType int

const (
	// Single character tokens

	TOKENTYPE_LEFTPAREN TokenType = iota
	TOKENTYPE_RIGHTPAREN
	TOKENTYPE_LEFTBRACE
	TOKENTYPE_RIGHTBRACE
	TOKENTYPE_COMMA
	TOKENTYPE_DOT
	TOKENTYPE_MINUS
	TOKENTYPE_PLUS
	TOKENTYPE_SEMICOLON
	TOKENTYPE_SLASH
	TOKENTYPE_STAR

	// One or two character tokens

	TOKENTYPE_BANG
	TOKENTYPE_BANGEQUAL
	TOKENTYPE_EQUAL
	TOKENTYPE_EQUALEQUAL
	TOKENTYPE_GREATER
	TOKENTYPE_GREATEREQUAL
	TOKENTYPE_LESS
	TOKENTYPE_LESSEQUAL

	// Literals

	TOKENTYPE_IDENTIFIER
	TOKENTYPE_STRING
	TOKENTYPE_NUMBER

	// Keywords

	TOKENTYPE_AND
	TOKENTYPE_CLASS
	TOKENTYPE_ELSE
	TOKENTYPE_FALSE
	TOKENTYPE_FUN
	TOKENTYPE_FOR
	TOKENTYPE_IF
	TOKENTYPE_NIL
	TOKENTYPE_OR
	TOKENTYPE_PRINT
	TOKENTYPE_RETURN
	TOKENTYPE_SUPER
	TOKENTYPE_THIS
	TOKENTYPE_TRUE
	TOKENTYPE_VAR
	TOKENTYPE_WHILE

	TOKENTYPE_EOF
)

type Token struct {
	Type    TokenType
	Lexeme  string
	Literal any
	Line    int
}

func (t Token) String() string {
	return fmt.Sprintf("%v %s %v", t.Type, t.Lexeme, t.Literal)
}

func eof(line int) Token {
	return Token{
		Type:    TOKENTYPE_EOF,
		Lexeme:  "",
		Literal: nil,
		Line:    line,
	}
}

type Tokens []Token

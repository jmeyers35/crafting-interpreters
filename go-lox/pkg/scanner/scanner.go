package scanner

import (
	"errors"

	goloxerrors "github.com/jmeyers35/golox/pkg/errors"
)

type Scanner interface {
	Scan() Tokens
}

type scannerImpl struct {
	start  int
	cursor int
	line   int
	source []byte
}

// Scan implements Scanner.
func (s *scannerImpl) Scan() Tokens {
	tokens := Tokens{}
	for !s.done() {
		s.start = s.cursor
		next, err := s.scanToken()
		if err != nil {
			goloxerrors.LogError(s.line, err.Error())
		}
		tokens = append(tokens, next)
	}
	tokens = append(tokens, eof(s.line))
	return tokens
}

func (s *scannerImpl) scanToken() (Token, error) {
	b := s.advance()

	switch b {
	case ')':
		return s.tokenAt(TOKENTYPE_RIGHTPAREN, nil), nil
	case '(':
		return s.tokenAt(TOKENTYPE_LEFTPAREN, nil), nil
	case '{':
		return s.tokenAt(TOKENTYPE_LEFTBRACE, nil), nil
	case '}':
		return s.tokenAt(TOKENTYPE_RIGHTBRACE, nil), nil
	case ',':
		return s.tokenAt(TOKENTYPE_COMMA, nil), nil
	case '.':
		return s.tokenAt(TOKENTYPE_DOT, nil), nil
	case '-':
		return s.tokenAt(TOKENTYPE_MINUS, nil), nil
	case '+':
		return s.tokenAt(TOKENTYPE_PLUS, nil), nil
	case ';':
		return s.tokenAt(TOKENTYPE_SEMICOLON, nil), nil
	case '*':
		return s.tokenAt(TOKENTYPE_STAR, nil), nil
	default:
		return Token{}, errors.New("Unexpected character")

	}
}

func (s *scannerImpl) advance() byte {
	nextByte := s.source[s.cursor]
	s.cursor += 1
	return nextByte
}

func (s *scannerImpl) tokenAt(tokenType TokenType, literal any) Token {
	text := s.source[s.start : s.cursor+1]
	return Token{
		Type:    tokenType,
		Literal: literal,
		Line:    s.line,
		Lexeme:  string(text),
	}
}

func (s *scannerImpl) done() bool {
	return s.cursor >= len(s.source)
}

func New(source string) Scanner {
	return &scannerImpl{
		line:   1,
		source: []byte(source),
	}
}

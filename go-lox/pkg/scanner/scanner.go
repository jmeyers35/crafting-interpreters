package scanner

import (
	"errors"

	goloxerrors "github.com/jmeyers35/golox/pkg/errors"
)

var tokenNone = Token{}

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
		if next != tokenNone {
			tokens = append(tokens, next)
		}
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
	case '!':
		return s.matchOr('=', TOKENTYPE_BANGEQUAL, TOKENTYPE_BANG), nil
	case '<':
		return s.matchOr('=', TOKENTYPE_LESSEQUAL, TOKENTYPE_LESS), nil
	case '>':
		return s.matchOr('=', TOKENTYPE_GREATEREQUAL, TOKENTYPE_GREATER), nil
	case '=':
		return s.matchOr('=', TOKENTYPE_EQUALEQUAL, TOKENTYPE_EQUAL), nil
	case '/':
		if s.match('/') {
			for s.peek() != '\n' && !s.done() {
				s.advance()
			}
			return tokenNone, nil
		} else {
			return s.tokenAt(TOKENTYPE_SLASH, nil), nil
		}

	case ' ', '\r', '\t':
		return tokenNone, nil

	case '\n':
		s.line += 1
		return tokenNone, nil

	case '"':
		return s.string()

	default:
		return Token{}, errors.New("Unexpected character")

	}
}

func (s *scannerImpl) string() (Token, error) {
	for s.peek() != '"' && !s.done() {
		if s.peek() == '\n' {
			s.line += 1
		}
		s.advance()
	}

	if s.done() {
		return tokenNone, errors.New("Unterminated string")
	}

	// consume closing quote
	s.advance()

	stringContents := s.source[s.start+1 : s.cursor]
	return s.tokenAt(TOKENTYPE_STRING, stringContents), nil
}

// returns the byte at the current cursor, consuming it.
func (s *scannerImpl) advance() byte {
	nextByte := s.source[s.cursor]
	s.cursor += 1
	return nextByte
}

// returns the byte at the current cursor without consuming it.
func (s *scannerImpl) peek() byte {
	if s.done() {
		return 0
	}
	return s.source[s.cursor]
}

func (s *scannerImpl) matchOr(expected byte, match, or TokenType) Token {
	if s.match(expected) {
		return s.tokenAt(match, nil)
	}
	return s.tokenAt(or, nil)
}

func (s *scannerImpl) match(expected byte) bool {
	if s.done() {
		return false
	}
	if s.source[s.cursor] != expected {
		return false
	}
	s.cursor += 1
	return true
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

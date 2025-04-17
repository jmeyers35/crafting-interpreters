package scanner

type Scanner interface {
	Scan() Tokens
}

type scannerImpl struct {
	cursor int
	line   int
	source []byte
}

// Scan implements Scanner.
func (s *scannerImpl) Scan() Tokens {
	tokens := Tokens{}
	for !s.done() {
		// TODO
	}
	tokens = append(tokens, eof(s.line))
	return tokens
}

func (s *scannerImpl) scanToken() Token {
	// TODO
	return Token{}
}

func (s *scannerImpl) done() bool {
	// TODO
	return false
}

func New(source string) Scanner {
	return &scannerImpl{
		line:   1,
		source: []byte(source),
	}
}

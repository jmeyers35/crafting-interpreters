package scanner

type Scanner interface {
	Scan([]byte) Tokens
}

type scannerImpl struct{}

// Scan implements Scanner.
func (s *scannerImpl) Scan([]byte) []Token {
	panic("unimplemented")
}

func New() Scanner {
	return &scannerImpl{}
}

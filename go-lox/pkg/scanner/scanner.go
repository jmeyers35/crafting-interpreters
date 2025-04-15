package scanner

type Token struct{}

type Scanner interface {
	Scan([]byte) []Token
}

type scannerImpl struct{}

// Scan implements Scanner.
func (s *scannerImpl) Scan([]byte) []Token {
	panic("unimplemented")
}

func New() Scanner {
	return &scannerImpl{}
}

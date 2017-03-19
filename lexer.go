package flexo

type Lexer interface {
	Run(tokens chan<- Token)
}

type TokenType int

const (
	TextToken TokenType = iota
	EmptyToken
	LinkStartToken
	LinkEndToken
	ListStartToken
	ListEndToken
	ListItemStartToken
	ListItemEndToken
)

type Token struct {
	OriginalFormat string
	Type           TokenType
	Content        string
	Attributes     map[string]string
}

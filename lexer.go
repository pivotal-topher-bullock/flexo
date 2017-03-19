package flexo

type Lexer interface {
	Run(tokens chan<- Token)
}

type Config struct {
	LinkPrefix string
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
	TerminatorToken
)

type Token struct {
	OriginalFormat string
	Type           TokenType
	Content        string
	Attributes     map[string]string
}

// type TokenFormatter interface {
// 	fmt.Stringer
// 	Type() TokenType
// 	Attr(string) (string, bool)
// }

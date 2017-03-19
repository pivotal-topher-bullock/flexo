package html

import (
	"github.com/pivotal-topher-bullock/flexo"
	"golang.org/x/net/html"
)

const HtmlFormat string = "html"

func NewHtmlLexer(tokenizer *html.Tokenizer) flexo.Lexer {
	return &htmlLexer{
		tokenizer: tokenizer,
	}
}

type htmlLexer struct {
	tokgienizer *html.Tokenizer
}

func (h *htmlLexer) Run(tokens chan<- flexo.Token) {
	for h.tokenizer.Next() != html.ErrorToken {
		tokens <- h.token(h.tokenizer.Token())
	}

	close(tokens)
}

func (h *htmlLexer) token(fromToken html.Token) flexo.Token {
	return flexo.Token{
		OriginalFormat: HtmlFormat,
		Type:           h.tokenType(fromToken),
		Content:        fromToken.String(),
		Attributes:     h.flattenAttrs(fromToken.Attr),
	}
}

var startTagTypes = map[string]flexo.TokenType{
	"ul": flexo.ListStartToken,
	"a":  flexo.LinkStartToken,
	"li": flexo.ListItemStartToken,
}

var endTagTypes = map[string]flexo.TokenType{
	"ul": flexo.ListEndToken,
	"a":  flexo.LinkEndToken,
	"li": flexo.ListItemEndToken,
}

func (h *htmlLexer) tokenType(token html.Token) flexo.TokenType {
	switch token.Type {
	case html.TextToken:
		return flexo.TextToken
	case html.StartTagToken:
		return startTagTypes[token.DataAtom.String()]
	case html.EndTagToken:
		return endTagTypes[token.DataAtom.String()]
	default:
		return flexo.EmptyToken
	}
}

func (h *htmlLexer) flattenAttrs(fromAttrs []html.Attribute) map[string]string {
	var attrs = make(map[string]string)
	for _, attr := range fromAttrs {
		attrs[attr.Key] = attr.Val
	}
	return attrs
}

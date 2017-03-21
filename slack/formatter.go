package slack

import (
	"fmt"

	"strings"

	"github.com/pivotal-topher-bullock/flexo"
)

func NewFormatter(config flexo.FormatterConfig) flexo.Formatter {
	return &slackFormatter{
		config: config,
	}
}

type slackFormatter struct {
	config flexo.FormatterConfig
}

func (sf *slackFormatter) Format(tokens <-chan flexo.Token) []string {
	out := []string{}
	for token := range tokens {
		out = append(out, sf.formatToken(token))
	}
	return out
}

func (sf *slackFormatter) formatToken(token flexo.Token) string {
	switch token.Type {
	case flexo.TextToken:
		if token.Content == "\n" {
			return ""
		}

		return strings.Replace(token.Content, "**", "*", -1)
	case flexo.LinkStartToken:
		return fmt.Sprintf("<%s%s|", sf.config.LinkPrefix, token.Attributes["href"])
	case flexo.LinkEndToken:
		return ">"
	case flexo.ListItemStartToken:
		return "- "
	case flexo.ListStartToken:
		fallthrough
	case flexo.ListEndToken:
		fallthrough
	case flexo.ListItemEndToken:
		return "\n"
	default:
		return ""
	}
	return ""
}

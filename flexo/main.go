package main

import (
	"fmt"
	"os"

	"strings"

	"github.com/jessevdk/go-flags"
	"github.com/pivotal-topher-bullock/flexo"
	"github.com/pivotal-topher-bullock/flexo/html"
	"github.com/pivotal-topher-bullock/flexo/slack"
)

type FlexoCommand struct {
	// InFormat  string `long:"in-fmt" default:"html" choice:"html" description:"Format of the input"`
	// OutFormat string `long:"out-fmt" default:"slack" choice:"slack" description:"Format of the output"`
	Config flexo.FormatterConfig
}

func (cmd *FlexoCommand) Run(args []string) error {
	var (
		lexer     flexo.Lexer
		formatter flexo.Formatter
	)

	//TODO: Break these out into a separate package that determines lexer / formatter type
	lexer = html.NewLexerFromStdin()
	formatter = slack.NewFormatter(cmd.Config)

	tokens := make(chan flexo.Token)

	go func() {
		lexer.Run(tokens)
	}()

	message := formatter.Format(tokens)
	<-tokens

	os.Stdout.Write([]byte(strings.Join(message, "")))
	os.Exit(0)
	return nil
}

func main() {
	cmd := &FlexoCommand{}

	parser := flags.NewParser(cmd, flags.Default)
	parser.NamespaceDelimiter = "-"

	args, err := parser.Parse()
	if err != nil {
		os.Exit(1)
	}

	err = cmd.Run(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

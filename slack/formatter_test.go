package slack_test

import (
	"github.com/pivotal-topher-bullock/flexo"
	. "github.com/pivotal-topher-bullock/flexo/slack"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Formatter", func() {
	var (
		formatter flexo.Formatter
		config    flexo.FormatterConfig
	)

	BeforeEach(func() {
		formatter = NewSlackFormatter(config)
	})

	Describe("Run", func() {

		var (
			tokenChannel chan flexo.Token
		)

		JustBeforeEach(func() {
			tokenChannel = make(chan flexo.Token)
			lexer.Format(tokenChannel)
		})
	})
})

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
		formatter = NewFormatter(config)
	})

	Describe("Run", func() {
		var (
			tokenChannel chan flexo.Token
			tokens       []flexo.Token
			output       []string
		)

		JustBeforeEach(func() {
			tokenChannel = make(chan flexo.Token)
			go func() {
				for _, token := range tokens {
					tokenChannel <- token
				}
				close(tokenChannel)
			}()

			output = formatter.Format(tokenChannel)
		})

		Context("parsing link tags", func() {
			BeforeEach(func() {
				tokens = []flexo.Token{
					{
						Type:       flexo.LinkStartToken,
						Attributes: map[string]string{"href": "example.com"},
					},
					{
						Type:    flexo.TextToken,
						Content: "a slack link",
					},
					{
						Type: flexo.LinkEndToken,
					},
				}
			})

			It("builds the proper string for a slack message link", func() {
				Expect(len(output)).To(Equal(3))
				Expect(output).To(BeEquivalentTo([]string{"<example.com|", "a slack link", ">"}))
			})
		})

		Context("parsing list items", func() {
			BeforeEach(func() {
				tokens = []flexo.Token{
					{
						Type: flexo.ListStartToken,
					},
					{
						Type: flexo.ListItemStartToken,
					},
					{
						Type:    flexo.TextToken,
						Content: "one",
					},
					{
						Type: flexo.ListItemEndToken,
					},
					{
						Type: flexo.ListItemStartToken,
					},
					{
						Type:    flexo.TextToken,
						Content: "two",
					},
					{
						Type: flexo.ListItemEndToken,
					},
					{
						Type: flexo.ListEndToken,
					},
				}
			})

			It("builds the proper string for a slack message link", func() {
				Expect(len(output)).To(Equal(8))
				Expect(output).To(BeEquivalentTo([]string{"\n", "- ", "one", "\n", "- ", "two", "\n", "\n"}))
			})
		})
	})
})

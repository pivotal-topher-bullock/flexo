package html_test

import (
	"github.com/pivotal-topher-bullock/flexo"
	. "github.com/pivotal-topher-bullock/flexo/html"
	"golang.org/x/net/html"

	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("lexer", func() {

	var (
		lexer     flexo.Lexer
		tokenizer *html.Tokenizer
		inputHtml string
	)

	JustBeforeEach(func() {
		reader := strings.NewReader(inputHtml)
		tokenizer = html.NewTokenizer(reader)
		lexer = NewHtmlLexer(tokenizer)
	})

	Describe("Run", func() {

		var (
			tokenChannel chan flexo.Token
		)

		JustBeforeEach(func() {
			tokenChannel = make(chan flexo.Token)
			go func() {
				lexer.Run(tokenChannel)
			}()
		})

		Context("with a single text token", func() {
			BeforeEach(func() {
				inputHtml = "beep boop"
			})

			It("sends a single flexo.TextToken to the channel", func() {
				var receivedToken flexo.Token
				Eventually(tokenChannel).Should(Receive(&receivedToken))
				Expect(receivedToken.Type).To(Equal(flexo.TextToken))
			})
			It("closes the channel", func() {
				Eventually(tokenChannel).Should(BeClosed())
			})
		})

		Context("with a link tag containing text", func() {
			BeforeEach(func() {
				inputHtml = "<a href='http://www.zombo.com/'>welcome to zombocom</a>"
			})

			It("sends a sequence of flexo tokens to the channel", func() {
				var receivedToken flexo.Token

				By("Sending a LinkStartToken")
				Eventually(tokenChannel).Should(Receive(&receivedToken))
				Expect(receivedToken.Type).Should(Equal(flexo.LinkStartToken))
				Expect(receivedToken.Attributes["href"]).To(Equal("http://www.zombo.com/"))

				By("Sending a TextToken")
				Eventually(tokenChannel).Should(Receive(&receivedToken))
				Expect(receivedToken.Type).Should(Equal(flexo.TextToken))
				Expect(receivedToken.Content).Should(Equal("welcome to zombocom"))

				By("Sending a LinkEndToken")
				Eventually(tokenChannel).Should(Receive(&receivedToken))
				Expect(receivedToken.Type).Should(Equal(flexo.LinkEndToken))

			})

		})
		Context("with a list of text", func() {
			BeforeEach(func() {
				inputHtml = "<ul><li>beep</li><li>boop</li></ul>"
			})

			It("sends a sequence of flexo tokens to the channel", func() {
				var receivedToken flexo.Token

				By("Sending a ListStartToken")
				Eventually(tokenChannel).Should(Receive(&receivedToken))
				Expect(receivedToken.Type).Should(Equal(flexo.ListStartToken))

				By("Sending a ListItemStartToken")
				Eventually(tokenChannel).Should(Receive(&receivedToken))
				Expect(receivedToken.Type).Should(Equal(flexo.ListItemStartToken))

				By("Sending a TextToken")
				Eventually(tokenChannel).Should(Receive(&receivedToken))
				Expect(receivedToken.Type).Should(Equal(flexo.TextToken))
				Expect(receivedToken.Content).Should(Equal("beep"))

				By("Sending a ListItemEndToken")
				Eventually(tokenChannel).Should(Receive(&receivedToken))
				Expect(receivedToken.Type).Should(Equal(flexo.ListItemEndToken))

				By("Sending a ListItemStartToken")
				Eventually(tokenChannel).Should(Receive(&receivedToken))
				Expect(receivedToken.Type).Should(Equal(flexo.ListItemStartToken))

				By("Sending a TextToken")
				Eventually(tokenChannel).Should(Receive(&receivedToken))
				Expect(receivedToken.Type).Should(Equal(flexo.TextToken))
				Expect(receivedToken.Content).Should(Equal("boop"))

				By("Sending a ListItemEndToken")
				Eventually(tokenChannel).Should(Receive(&receivedToken))
				Expect(receivedToken.Type).Should(Equal(flexo.ListItemEndToken))

				By("Sending a ListEndToken")
				Eventually(tokenChannel).Should(Receive(&receivedToken))
				Expect(receivedToken.Type).Should(Equal(flexo.ListEndToken))

			})

		})

	})
})

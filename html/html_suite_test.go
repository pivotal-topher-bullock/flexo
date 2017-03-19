package html_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestHtml(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Html Suite")
}

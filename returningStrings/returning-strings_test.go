package returningStrings

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Example test cases", func() {
	It("should return greeting for Ryan", func() {
		Expect(Greet("Ryan")).To(Equal("Hello, Ryan how are you doing today?"))
	})
})

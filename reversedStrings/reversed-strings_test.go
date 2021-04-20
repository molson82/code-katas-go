package reversedStrings

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Basic Tests", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(Solution("world")).To(Equal("dlrow"))
	})
})

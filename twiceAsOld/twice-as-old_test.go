package twiceAsOld

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Basic Tests", func() {
	It("should test that the solution returns the correct value", func() {
		//     Expect(Solution(1)).To(Equal(2))
		Expect(TwiceAsOld(36, 7)).To(Equal(22))
		Expect(TwiceAsOld(55, 30)).To(Equal(5))
		Expect(TwiceAsOld(42, 21)).To(Equal(0))
		Expect(TwiceAsOld(22, 1)).To(Equal(20))
		Expect(TwiceAsOld(29, 0)).To(Equal(29))
	})
})

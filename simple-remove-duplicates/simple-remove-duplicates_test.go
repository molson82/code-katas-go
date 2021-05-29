package simpleRemoveDuplicates

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func dotest(arr, exp []int) {
	var ans = Solve(arr)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Example tests", func() {
	It("It should work for basic tests", func() {
		dotest([]int{3, 4, 4, 3, 6, 3}, []int{4, 6, 3})
		dotest([]int{1, 2, 1, 2, 1, 2, 3}, []int{1, 2, 3})
		dotest([]int{1, 2, 3, 4}, []int{1, 2, 3, 4})
		dotest([]int{1, 1, 4, 5, 1, 2, 1}, []int{4, 5, 2, 1})
		dotest([]int{1, 2, 1, 2, 1, 1, 3}, []int{2, 1, 3})
		dotest([]int{0, 4, 4, 3, 0, 3}, []int{4, 0, 3})
	})
})

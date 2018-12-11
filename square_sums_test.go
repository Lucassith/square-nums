package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SquareSums", func() {
	It("Should calculate length exact to n", func() {
		Expect(getPairs(3)).To(HaveLen(3))
	})

	It("Should generate valid pairs", func() {
		pairs := getPairs(6)
		Expect(pairs[1]).To(ConsistOf(3))
		Expect(pairs[3]).To(ConsistOf(1, 6))
		Expect(pairs[4]).To(ConsistOf(5))
		Expect(pairs[5]).To(ConsistOf(4))
		Expect(pairs[2]).ToNot(ConsistOf(2))
	})

	It("Should get next value from pair", func() {
		pairs := getPairs(6)
		Expect(getNextFromPair(1, pairs[3])).To(Equal(6))
	})

	It("Should return 0 if no another value", func() {
		pairs := getPairs(6)
		Expect(getNextFromPair(6, pairs[4])).To(Equal(0))
	})

	It("Should return nil for n = 2", func() {
		result := GetSquareSums(2)
		Expect(result).To(HaveLen(0))
	})

	It("Should count for n = 15", func() {
		result := GetSquareSums(15)
		Expect(result).To(HaveLen(15))
	})

	It("Should count for n = 17", func() {
		result := GetSquareSums(17)
		Expect(result).To(HaveLen(17))
	})
})

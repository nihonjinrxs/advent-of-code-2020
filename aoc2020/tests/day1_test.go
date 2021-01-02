package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"codeanddata.app/advent-of-code/aoc2020/day1"
	"codeanddata.app/advent-of-code/aoc2020/tests/helpers"
)

var _ = Describe("Day1", func() {
	Describe("with test data", func() {
		var (
			input []string
		)

		BeforeEach(func() {
			input = []string{"1721", "979", "366", "299", "675", "1456"}
		})

		Context("Part 1: Pair correction", func() {
			It("provides the correct value", func() {
				expected := 514579
				result := day1.FindExpenseReportPairCorrection(input)
				Expect(result).To(Equal(expected))
			})
		})
		Context("Part 2: Trio correction", func() {
			It("provides the correct value", func() {
				expected := 241861950
				result := day1.FindExpenseReportTrioCorrection(input)
				Expect(result).To(Equal(expected))
			})
		})
	})
	Describe("with the provided input data", func() {
		var (
			input []string
		)

		BeforeEach(func() {
			input = helpers.LoadFixture("../testdata/day1.txt")
		})

		Context("Part 1: Pair correction", func() {
			It("provides the correct value", func() {
				expected := 744475
				result := day1.FindExpenseReportPairCorrection(input)
				Expect(result).To(Equal(expected))
			})
		})
		Context("Part 2: Trio correction", func() {
			It("provides the correct value", func() {
				expected := 70276940
				result := day1.FindExpenseReportTrioCorrection(input)
				Expect(result).To(Equal(expected))
			})
		})
	})
})

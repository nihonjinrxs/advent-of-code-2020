package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"codeanddata.app/advent-of-code/aoc2020/day2"
	"codeanddata.app/advent-of-code/aoc2020/tests/helpers"
)

var _ = Describe("Day2", func() {
	Describe("with test data", func() {
		var (
			input []string
		)

		BeforeEach(func() {
			input = []string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"}
		})

		Context("Part 1: Valid Password Count", func() {
			It("provides the correct value", func() {
				expected := 2
				result := day2.CountValidPasswordsByRuneCount(input)
				Expect(result).To(Equal(expected))
			})
		})
		Context("Part 2: Trio correction", func() {
			It("provides the correct value", func() {
				expected := 1
				result := day2.CountValidPasswordsByRunePosition(input)
				Expect(result).To(Equal(expected))
			})
		})
	})
	Describe("with the provided input data", func() {
		var (
			input []string
		)

		BeforeEach(func() {
			input = helpers.LoadFixture("../testdata/day2.txt")
		})

		Context("Part 1: Valid Password Count", func() {
			It("provides the correct value", func() {
				expected := 458
				result := day2.CountValidPasswordsByRuneCount(input)
				Expect(result).To(Equal(expected))
			})
		})
		Context("Part 2: Trio correction", func() {
			It("provides the correct value", func() {
				expected := 342
				result := day2.CountValidPasswordsByRunePosition(input)
				Expect(result).To(Equal(expected))
			})
		})
	})
})

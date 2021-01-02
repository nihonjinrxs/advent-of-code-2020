package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"codeanddata.app/advent-of-code/aoc2020/day3"
	"codeanddata.app/advent-of-code/aoc2020/tests/helpers"
)

var _ = Describe("Day3", func() {
	Describe("with test data", func() {
		var (
			input []string
		)

		BeforeEach(func() {
			input = []string{
				"..##.......",
				"#...#...#..",
				".#....#..#.",
				"..#.#...#.#",
				".#...##..#.",
				"..#.##.....",
				".#.#.#....#",
				".#........#",
				"#.##...#...",
				"#...##....#",
				".#..#...#.#",
			}
		})

		Context("Part 1: Number of Trees in Toboggan Path", func() {
			It("provides the correct value", func() {
				expected := 7
				result := day3.CountTreesOnSlope3R1D(input)
				Expect(result).To(Equal(expected))
			})
		})
		Context("Part 2: Trio correction", func() {
			It("provides the correct value", func() {
				expected := 336
				result := day3.CountTreesOnSlope(input, 1, 1) *
					day3.CountTreesOnSlope(input, 3, 1) *
					day3.CountTreesOnSlope(input, 5, 1) *
					day3.CountTreesOnSlope(input, 7, 1) *
					day3.CountTreesOnSlope(input, 1, 2)
				Expect(result).To(Equal(expected))
			})
		})
	})
	Describe("with the provided input data", func() {
		var (
			input []string
		)

		BeforeEach(func() {
			input = helpers.LoadFixture("../testdata/day3.txt")
		})

		Context("Part 1: Number of Trees in Toboggan Path", func() {
			It("provides the correct value", func() {
				expected := 162
				result := day3.CountTreesOnSlope3R1D(input)
				Expect(result).To(Equal(expected))
			})
		})
		Context("Part 2: Product of slope results", func() {
			It("provides the correct value", func() {
				expected := 3064612320
				result := day3.CountTreesOnSlope(input, 1, 1) *
					day3.CountTreesOnSlope(input, 3, 1) *
					day3.CountTreesOnSlope(input, 5, 1) *
					day3.CountTreesOnSlope(input, 7, 1) *
					day3.CountTreesOnSlope(input, 1, 2)
				Expect(result).To(Equal(expected))
			})
		})
	})
})

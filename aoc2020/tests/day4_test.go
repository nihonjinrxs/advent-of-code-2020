package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"codeanddata.app/advent-of-code/aoc2020/day4"
	"codeanddata.app/advent-of-code/aoc2020/tests/helpers"
)

var _ = Describe("Day4", func() {
	Describe("with test data", func() {
		var (
			input string
		)

		BeforeEach(func() {
			input = `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`
		})

		Context("Part 1: Count valid passports", func() {
			It("provides the correct value", func() {
				expected := 2
				result := day4.CountValidPassports(input)
				Expect(result).To(Equal(expected))
			})
		})
		// Context("Part 2: Trio correction", func() {
		// 	It("provides the correct value", func() {
		// 		expected := 336
		// 		result := day4.CountValidPassports(input)
		// 		Expect(result).To(Equal(expected))
		// 	})
		// })
	})
	Describe("with the provided input data", func() {
		var (
			input string
		)

		BeforeEach(func() {
			input = helpers.LoadFixtureToString("../testdata/day4.txt")
		})

		Context("Part 1: Count valid passports", func() {
			It("provides the correct value", func() {
				expected := 242
				result := day4.CountValidPassports(input)
				Expect(result).To(Equal(expected))
			})
		})
		// Context("Part 2: Product of slope results", func() {
		// 	It("provides the correct value", func() {
		// 		expected := 3064612320
		// 		result := day4.CountValidPassports(input)
		// 		Expect(result).To(Equal(expected))
		// 	})
		// })
	})
})

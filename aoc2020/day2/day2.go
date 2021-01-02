package day2

import (
	"strconv"
	"strings"
)

// CountValidPasswordsByRuneCount takes a string slice of
// password rules with passwords, and counts the number of
// passwords matching their respective rules, interpreting
// the rule values as a count range
func CountValidPasswordsByRuneCount(in []string) int {
	count := 0
	for _, s := range in {
		rule, password := parseLine(s)
		ruleRange, ruleRune := parseRule(rule)
		runeCount := strings.Count(password, ruleRune)
		if runeCount >= ruleRange[0] && runeCount <= ruleRange[1] {
			count++
		}
	}
	return count
}

// CountValidPasswordsByRunePosition takes a string slice of
// password rules with passwords, and counts the number of
// passwords matching their respective rules, interpreting
// the rule values as 1-indexed positions in the password
func CountValidPasswordsByRunePosition(in []string) int {
	count := 0
	for _, s := range in {
		rule, password := parseLine(s)
		pos, ruleRune := parseRule(rule)
		chars := strings.Split(password, "")
		valid := (chars[pos[0]-1] == ruleRune && chars[pos[1]-1] != ruleRune) ||
			(chars[pos[0]-1] != ruleRune && chars[pos[1]-1] == ruleRune)
		if valid {
			count++
		}
	}
	return count
}

func parseLine(line string) (string, string) {
	a := strings.Split(line, ": ")
	return a[0], a[1]
}

func parseRule(rule string) ([]int, string) {
	a := strings.Split(rule, " ")
	b := strings.Split(a[0], "-")
	rangeMin, _ := strconv.Atoi(b[0])
	rangeMax, _ := strconv.Atoi(b[1])
	return []int{rangeMin, rangeMax}, a[1]
}

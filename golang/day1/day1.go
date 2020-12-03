package day1

import (
	"strconv"
)

// FindExpenseReportPairCorrection searches the list of
// expenses for the two that sum to 2020, then returns
// the product of those expenses.
func FindExpenseReportPairCorrection(expenses []string) int {
	exps := make([]int, len(expenses))
	for i, s := range expenses {
		exps[i], _ = strconv.Atoi(s)
	}
	for i, val1 := range exps {
		matchedVal2 := 2020 - val1
		for _, val2 := range exps[i+1:] {
			if matchedVal2 == val2 {
				return val1 * val2
			}
		}
	}
	return -1
}

// FindExpenseReportTrioCorrection searches the list of
// expenses for the two that sum to 2020, then returns
// the product of those expenses.
func FindExpenseReportTrioCorrection(expenses []string) int {
	exps := make([]int, len(expenses))
	for i, s := range expenses {
		exps[i], _ = strconv.Atoi(s)
	}
	for i, val1 := range exps {
		remaining := 2020 - val1
		if remaining > 0 {
			for j, val2 := range exps[i+1:] {
				matchingVal3 := remaining - val2
				if matchingVal3 > 0 {
					for _, val3 := range exps[j+1:] {
						if matchingVal3 == val3 {
							return val1 * val2 * val3
						}
					}
				}
			}
		}
	}
	return -1
}

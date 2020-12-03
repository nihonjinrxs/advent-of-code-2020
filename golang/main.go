package main

import (
	"fmt"
	"os"

	"codeanddata.app/advent-of-code/aoc2020/common"
	"codeanddata.app/advent-of-code/aoc2020/day1"
)

func main() {
	cmd := os.Args[1]
	switch cmd {
	case "day1":
		fmt.Println("Running AoC 2020 day 1...")
		inval := common.ReadStdinLines()
		result := day1.FindExpenseReportPairCorrection(inval)
		fmt.Println("Corrected expense pair value:", result)
		result = day1.FindExpenseReportTrioCorrection(inval)
		fmt.Println("Corrected expense trio value:", result)
		break
	default:
		fmt.Println("Unknown command:", cmd)
	}
}

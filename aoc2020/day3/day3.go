package day3

import (
	"strings"
)

func pmod(a int, b int) int {
	return (a%b + b) % b
}

// CountTreesOnSlope takes a mapgrid (. for open, # for tree)
// and counts the number of trees that would be hit traversing
// from top to bottom using slope hslope right, vslope down
// assuming a horizontally tiled repeating map
func CountTreesOnSlope(mapgrid []string, hslope int, vslope int) int {
	var grid [][]bool
	var gridrow = []bool{}
	var rowlen int
	var treeCount int
	for _, line := range mapgrid {
		chars := strings.Split(line, "")
		gridrow = []bool{}
		for _, c := range chars {
			gridrow = append(gridrow, c == "#")
		}
		grid = append(grid, gridrow)
		rowlen = len(gridrow)
	}
	currX, currY := 0, 0
	if grid[0][0] {
		treeCount++
	}
	for currY < len(grid)-1 {
		currX += hslope
		currY += vslope
		wrappedX := pmod(currX, rowlen)
		if grid[currY][wrappedX] {
			treeCount++
		}
	}
	return treeCount
}

// CountTreesOnSlope3R1D takes a mapgrid (. for open, # for tree)
// and counts the number of trees that would be hit traversing
// from top to bottom using slope 3 right, 1 down
// assuming a horizontally tiled repeating map
func CountTreesOnSlope3R1D(mapgrid []string) int {
	return CountTreesOnSlope(mapgrid, 3, 1)
}

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(fmt.Sprintf("Failed to open file: %s", err))
	}

	input := string(data)
	lines := strings.Split(input, "\n")
	numRows := len(lines)
	numCols := len(lines[0])
	treeMap := make([][]string, numRows)

	for i, line := range lines {
		treeMap[i] = strings.Split(line, "")
	}

	fmt.Printf("Part1: %d\n", part1(treeMap, numRows, numCols))
	fmt.Printf("Part2: %d\n", part2(treeMap, numRows, numCols))
	return
}

func part1(treeMap [][]string, numRows, numCols int) int {
	return moveDown(treeMap, numRows, numCols, 0, 0, 3, 1)
}

func part2(treeMap [][]string, numRows, numCols int) int {
	return moveDown(treeMap, numRows, numCols, 0, 0, 1, 1) *
		moveDown(treeMap, numRows, numCols, 0, 0, 3, 1) *
		moveDown(treeMap, numRows, numCols, 0, 0, 5, 1) *
		moveDown(treeMap, numRows, numCols, 0, 0, 7, 1) *
		moveDown(treeMap, numRows, numCols, 0, 0, 1, 2)
}

func moveDown(treeMap [][]string, numRows, numCols, x, y, r, d int) int {
	treeBump := 0
	if treeMap[y][x%numCols] == "#" {
		treeBump++
	}

	if numRows > y+1 {
		treeBump += moveDown(treeMap, numRows, numCols, x+r, y+d, r, d)
	}

	return treeBump
}

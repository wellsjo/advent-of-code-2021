package main

import (
	"log"
	"sort"
	"strconv"

	"github.com/wellsjo/advent-of-code-2021/common"
)

func main() {
	m := parseInput("input.txt")
	// part1(m)
	part2(m)
}

type smokeMap struct {
	m      [][]int
	marked map[point]struct{}
}

type point struct {
	i, j int
}

func (m smokeMap) isMarked(i, j int) bool {
	_, ok := m.marked[point{i: i, j: j}]
	return ok
}

func (m smokeMap) mark(i, j int) {
	m.marked[point{i: i, j: j}] = struct{}{}
}

func (m smokeMap) get(i, j int) int {
	if i >= len(m.m) || j >= len(m.m[0]) || i < 0 || j < 0 {
		return -1
	}
	return m.m[i][j]
}

func part2(m smokeMap) {
	basinSizes := []int{}
	for i := range m.m {
		for j := range m.m[i] {
			if m.m[i][j] == 9 {
				continue
			}
			if !m.isMarked(i, j) {
				size := m.getBasin(i, j)
				basinSizes = append(basinSizes, size)
			}
		}
	}
	sort.Ints(basinSizes)
	topBasins := basinSizes[len(basinSizes)-3:]
	log.Println("Top 3 Basins", topBasins)
	log.Println("Sum", topBasins[0]*topBasins[1]*topBasins[2])
}

func (m smokeMap) getBasin(i, j int) int {
	if m.isMarked(i, j) {
		return 0
	}

	val := m.get(i, j)
	if val == 9 || val == -1 {
		return 0
	}

	m.mark(i, j)

	ans := 1 + m.getBasin(i+1, j) + m.getBasin(i, j+1)
	ans += m.getBasin(i-1, j) + m.getBasin(i, j-1)
	return ans
}

func part1(m smokeMap) {
	sum := 0
	for i := range m.m {
		for j := range m.m[i] {
			if m.isBottomPoint(i, j) {
				riskLevel := m.m[i][j] + 1
				sum += riskLevel
			}
		}
	}
	log.Println("Part 1", sum)
}

func (m smokeMap) isBottomPoint(i, j int) bool {
	return m.isLessThanAbove(i, j) && m.isLessThanRight(i, j) && m.isLessThanBelow(i, j) && m.isLessThanLeft(i, j)
}

func (m smokeMap) isLessThanAbove(i, j int) bool {
	if i == 0 {
		return true
	}
	return m.m[i][j] < m.m[i-1][j]
}

func (m smokeMap) isLessThanRight(i, j int) bool {
	if j == (len(m.m[0]) - 1) {
		return true
	}
	return m.m[i][j] < m.m[i][j+1]
}

func (m smokeMap) isLessThanBelow(i, j int) bool {
	if i == (len(m.m) - 1) {
		return true
	}
	return m.m[i][j] < m.m[i+1][j]
}

func (m smokeMap) isLessThanLeft(i, j int) bool {
	if j == 0 {
		return true
	}
	return m.m[i][j] < m.m[i][j-1]
}

func parseInput(path string) smokeMap {
	lines, err := common.ReadLines(path)
	if err != nil {
		panic(err)
	}

	m := [][]int{}
	for i := range lines {
		line := lines[i]
		internalArr := []int{}
		for j := range line {
			num, err := strconv.Atoi(string(line[j]))
			if err != nil {
				panic(err)
			}
			internalArr = append(internalArr, num)
		}
		m = append(m, internalArr)
	}

	return smokeMap{
		m:      m,
		marked: map[point]struct{}{},
	}
}

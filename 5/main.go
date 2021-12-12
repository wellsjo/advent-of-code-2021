package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/wellsjo/advent-of-code-2021/common"
)

func main() {
	lines, board := setupGame("input.txt")

	for i := range lines {
		board.applyLine(lines[i])
	}

	log.Println(board)
	log.Println(board.calculateScore())
}

type point struct {
	x, y int
}

type line struct {
	p1, p2 point
}

type lineRange struct {
	xmin, xmax, ymin, ymax int
}

func (l line) getRange() lineRange {
	r := lineRange{
		xmin: l.p1.x,
		xmax: l.p1.x,
		ymin: l.p1.y,
		ymax: l.p1.y,
	}
	if l.p2.x < r.xmin {
		r.xmin = l.p2.x
	}
	if l.p2.x > r.xmax {
		r.xmax = l.p2.x
	}
	if l.p2.y < r.ymin {
		r.ymin = l.p2.y
	}
	if l.p2.y > r.ymax {
		r.ymax = l.p2.y
	}
	return r
}

// Get y from x
func (l line) fx(x int) int {
	m := (l.p2.y - l.p1.y) / (l.p2.x - l.p1.x)
	if m == 0 {
		return l.p1.y
	}
	b := l.p1.y - (m * l.p1.x)
	return (m * x) + b
}

// Get x from y
func (l line) fy(y int) int {
	if l.p2.x == l.p2.x {
		return l.p1.x
	}
	m := (l.p2.y - l.p1.y) / (l.p2.x - l.p1.x)
	b := l.p1.y / (m * l.p1.x)
	return (y - b) / m
}

type board [][]int

func (b board) calculateScore() int {
	score := 0
	for i := range b {
		for j := range b[i] {
			if b[i][j] >= 2 {
				score++
			}
		}
	}
	return score
}

func (b board) applyLine(l line) {
	r := l.getRange()

	if r.xmin != r.xmax {
		for i := r.xmin; i <= r.xmax; i++ {
			y := l.fx(i)
			b[y][i]++
		}
		return
	}

	if r.ymin != r.ymax {
		for i := r.ymin; i <= r.ymax; i++ {
			x := l.fy(i)
			b[i][x]++
		}
	}
}

func (b board) String() string {
	s := "\n"
	for i := 0; i < len(b); i++ {
		s += fmt.Sprintf("%v\n", b[i])
	}
	return s
}

func newBoard(size int) board {
	b := [][]int{}
	for i := 0; i < size; i++ {
		b = append(b, make([]int, size))
	}
	return b
}

func setupGame(path string) ([]line, board) {
	linesIn, err := common.ReadLines(path)
	if err != nil {
		panic(err)
	}

	maxValue := 0
	lines := []line{}

	for i := range linesIn {
		l := linesIn[i]
		parts := strings.Split(l, " -> ")
		leftParts := strings.Split(parts[0], ",")
		rightParts := strings.Split(parts[1], ",")

		p1x, err := strconv.Atoi(leftParts[0])
		if err != nil {
			panic(err)
		}
		if p1x > maxValue {
			maxValue = p1x
		}

		p1y, err := strconv.Atoi(leftParts[1])
		if err != nil {
			panic(err)
		}
		if p1y > maxValue {
			maxValue = p1y
		}

		p2x, err := strconv.Atoi(rightParts[0])
		if err != nil {
			panic(err)
		}
		if p2x > maxValue {
			maxValue = p2x
		}

		p2y, err := strconv.Atoi(rightParts[1])
		if err != nil {
			panic(err)
		}
		if p2y > maxValue {
			maxValue = p2y
		}

		p1 := point{x: p1x, y: p1y}
		p2 := point{x: p2x, y: p2y}
		lines = append(lines, line{p1: p1, p2: p2})
	}

	return lines, newBoard(maxValue + 1)
}

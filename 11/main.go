package main

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/wellsjo/advent-of-code-2021/common"
)

func main() {
	p := parseInput("input.txt")

	p.print()
	fmt.Println()

	// p.part1()
	p.part2()
}

func (p puzzle) part1() {
	steps := 100
	for i := 0; i < steps; i++ {
		p.step()
	}
	p.print()
	fmt.Println(p.flashes, "Flashes")
}

func (p puzzle) part2() {
	step := 0
	for {
		step++
		allFlashed := p.step()
		if allFlashed {
			break
		}
	}
	fmt.Println("All flashed step", step)
}

type puzzle struct {
	board   [][]int
	flashes int
}

func (p *puzzle) step() bool {
	for i := range p.board {
		for j := range p.board[i] {
			p.incr(i, j)
		}
	}

	flashMap := map[position]struct{}{}
	for i := range p.board {
		for j := range p.board[i] {
			if p.board[i][j] > 9 {
				p.flash(i, j, flashMap)
			}
		}
	}

	for i := range p.board {
		for j := range p.board[i] {
			if p.board[i][j] > 9 {
				p.board[i][j] = 0
			}
		}
	}

	return len(flashMap) == len(p.board)*len(p.board[0])
}

func (p *puzzle) incr(x, y int) bool {
	if x >= len(p.board) || y >= len(p.board[0]) || x < 0 || y < 0 {
		return false
	}
	p.board[x][y]++
	return p.board[x][y] > 9
}

func (p *puzzle) flash(x, y int, flashMap map[position]struct{}) {
	if _, ok := flashMap[position{x: x, y: y}]; ok {
		return
	}

	flashMap[position{x: x, y: y}] = struct{}{}
	p.flashes++

	if shouldFlash := p.incr(x-1, y-1); shouldFlash {
		p.flash(x-1, y-1, flashMap)
	}

	if shouldFlash := p.incr(x-1, y); shouldFlash {
		p.flash(x-1, y, flashMap)
	}

	if shouldFlash := p.incr(x-1, y+1); shouldFlash {
		p.flash(x-1, y+1, flashMap)
	}

	if shouldFlash := p.incr(x, y-1); shouldFlash {
		p.flash(x, y-1, flashMap)
	}

	if shouldFlash := p.incr(x, y+1); shouldFlash {
		p.flash(x, y+1, flashMap)
	}

	if shouldFlash := p.incr(x+1, y-1); shouldFlash {
		p.flash(x+1, y-1, flashMap)
	}

	if shouldFlash := p.incr(x+1, y); shouldFlash {
		p.flash(x+1, y, flashMap)
	}

	if shouldFlash := p.incr(x+1, y+1); shouldFlash {
		p.flash(x+1, y+1, flashMap)
	}
}

func (p *puzzle) print() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight)
	for i := range p.board {
		for j := range p.board[i] {
			fmt.Fprint(w, p.board[i][j], "\t")
		}
		fmt.Fprintln(w)
	}
	w.Flush()
}

type position struct {
	x, y int
}

func parseInput(path string) *puzzle {
	lines, err := common.ReadLines(path)
	if err != nil {
		panic(err)
	}
	b := [][]int{}
	for i := range lines {
		line := lines[i]
		intArr := []int{}
		for j := range line {
			numStr := line[j]
			num, err := strconv.Atoi(string(numStr))
			if err != nil {
				panic(err)
			}
			intArr = append(intArr, num)
		}
		b = append(b, intArr)
	}
	return &puzzle{
		board: b,
	}
}

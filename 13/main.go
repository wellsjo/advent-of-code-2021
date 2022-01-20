package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/wellsjo/advent-of-code-2021/common"
)

func main() {
	p, instructions := parseInput("input.txt")

	for i := range instructions {
		in := instructions[i]
		if in.axis == "x" {
			p.foldLeft(in.amount)
		} else {
			p.foldUp(in.amount)
		}
		log.Println("Dots", p.numDots())
	}

	log.Println()
	p.print()
	log.Println("Answer", p.numDots())
}

type instruction struct {
	axis   string
	amount int
}

type paper [][]string

func (p *paper) foldUp(y int) {
	bottomHalf := (*p)[y+1:]
	newLength := y - 1
	for i := range bottomHalf {
		for j := range bottomHalf[i] {
			if (*p)[newLength-i][j] != "x" {
				(*p)[newLength-i][j] = bottomHalf[i][j]
			}
		}
	}
	*p = (*p)[:y]
}

func (p *paper) foldLeft(x int) {
	c := make(paper, len(*p))
	for i := range c {
		c[i] = (*p)[i][:x]
	}
	for i := 0; i < len(*p); i++ {
		for j := x + 1; j < (x*2)+1; j++ {
			if c[i][(2*x)-j] != "x" {
				c[i][(2*x)-j] = (*p)[i][j]
			}
		}
	}
	*p = c
}

func (p *paper) numDots() int {
	count := 0
	for i := range *p {
		for j := range (*p)[i] {
			if (*p)[i][j] == "x" {
				count++
			}
		}
	}
	return count
}

func (p paper) print() {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight)
	for i := range p {
		for j := range p[i] {
			fmt.Fprint(w, p[i][j], "\t")
		}
		fmt.Fprintln(w)
	}
	w.Flush()
}

type point struct {
	x, y int
}

func parseInput(path string) (paper, []instruction) {
	lines, err := common.ReadLines(path)
	if err != nil {
		panic(err)
	}

	points := []point{}
	instructions := []instruction{}
	xMax := 0
	yMax := 0
	for i := range lines {
		if lines[i] == "" {
			continue
		}
		if lines[i][0] == 'f' {
			parts := strings.Split(lines[i], "fold along ")
			parts2 := strings.Split(parts[1], "=")
			axis := string(parts2[0])
			amount, err := strconv.Atoi(parts2[1])
			if err != nil {
				panic(err)
			}
			instructions = append(instructions, instruction{axis: axis, amount: amount})
		} else {
			lineParts := strings.Split(lines[i], ",")
			xStr := lineParts[0]
			yStr := lineParts[1]
			x, _ := strconv.Atoi(xStr)
			y, _ := strconv.Atoi(yStr)
			if x > xMax {
				xMax = x
			}
			if y > yMax {
				yMax = y
			}
			points = append(points, point{x: x, y: y})
		}
	}

	p := make([][]string, yMax+1)
	for i := range p {
		p[i] = make([]string, xMax+1)
	}
	for i := range p {
		for j := range p[i] {
			p[i][j] = "."
		}
	}
	for i := range points {
		p[points[i].y][points[i].x] = "x"
	}
	return p, instructions
}

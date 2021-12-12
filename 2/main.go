package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/wellsjo/advent-of-code-2021/common"
)

func main() {
	lines, err := common.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}

	p := &position{horizontal: 0, depth: 0}
	for i := range lines {
		in := getInstructionsFromInput(lines[i])
		updatePosition(p, in)
		log.Println("POSITION", p)
	}

	fmt.Println("POSITION", p)
	fmt.Println("ANSWER", p.horizontal*p.depth)
}

func updatePosition(p *position, in instruction) {
	if in.direction == down {
		// p.depth += in.amount
		p.aim += in.amount
		return
	}

	if in.direction == forward {
		p.horizontal += in.amount
		p.depth += (p.aim * in.amount)
		return
	}

	if in.direction == up {
		// p.depth -= in.amount
		p.aim -= in.amount
	}
}

func getInstructionsFromInput(in string) instruction {
	parts := strings.Split(in, " ")
	if len(parts) != 2 {
		panic("invalid input")
	}

	dirStr := parts[0]
	amountStr := parts[1]

	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		panic(err)
	}

	return instruction{
		direction: direction(dirStr),
		amount:    amount,
	}
}

type position struct {
	horizontal, depth, aim int
}

type direction string

const (
	forward direction = "forward"
	down    direction = "down"
	up      direction = "up"
)

type instruction struct {
	direction direction
	amount    int
}

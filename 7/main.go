package main

import (
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/wellsjo/advent-of-code-2021/common"
)

func main() {
	m := parseInput("input.txt")
	log.Println(m.getMinScore())
}

type model struct {
	arrangement map[float64]float64
	min, max    float64
}

func (m model) getMinScore() float64 {
	minScore := math.Inf(1)
	for i := m.min; i <= m.max; i++ {
		s := m.getScoreForPosition(i)
		if s < minScore {
			minScore = s
		}
	}
	return minScore
}

func (m model) getScoreForPosition(pos float64) float64 {
	score := 0.0
	for value, amount := range m.arrangement {
		d := math.Abs(value - pos)
		score += amount * getScoreFromDistance(d)
	}
	return score
}

func getScoreFromDistance(value float64) float64 {
	return ((value * value) / 2) + (value / 2)
}

func parseInput(path string) model {
	lines, err := common.ReadLines(path)
	if err != nil {
		panic(err)
	}

	m := model{
		arrangement: make(map[float64]float64),
		max:         math.Inf(-1),
		min:         math.Inf(1),
	}

	values := strings.Split(lines[0], ",")
	for i := range values {
		in, err := strconv.Atoi(values[i])
		if err != nil {
			panic(err)
		}

		in64 := float64(in)

		if in64 < m.min {
			m.min = float64(in)
		}
		if in64 > m.max {
			m.max = float64(in)
		}

		m.arrangement[in64]++
	}

	return m
}

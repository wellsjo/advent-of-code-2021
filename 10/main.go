package main

import (
	"log"
	"sort"

	"github.com/wellsjo/advent-of-code-2021/common"
)

func main() {
	navSystem := parseInput("input.txt")

	navSystem.leggo()
}

func (n navigationSubsystem) leggo() {
	syntaxScore := 0
	completeScores := []int{}
	for i := range n.lines {
		_, cScore, sScore := n.lines[i].solve()
		syntaxScore += sScore
		completeScores = append(completeScores, cScore)
	}

	log.Println("Syntax Score", syntaxScore)

	sort.Ints(completeScores)
	for i := range completeScores {
		if completeScores[i] != 0 {
			completeScores = completeScores[i:]
			break
		}
	}
	middleScore := completeScores[((len(completeScores)+1)/2)-1]
	log.Println("Middle Complete Score", middleScore)
}

type stack []string

func (s *stack) push(char string) {
	*s = append(*s, char)
}

func (s *stack) pop() string {
	length := len(*s)
	c := (*s)[length-1]
	*s = (*s)[:length-1]
	return c
}

var chunks = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
	"<": ">",
}

var charScores = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

var completeScores = map[string]int{
	")": 1,
	"]": 2,
	"}": 3,
	">": 4,
}

type line []string

func (l line) solve() (bool, int, int) {
	s := stack{}
	for i := range l {
		char := l[i]
		if _, isOpenChar := chunks[char]; isOpenChar {
			s.push(char)
		} else {
			last := s.pop()
			if chunks[last] != char {
				return false, 0, charScores[char]
			}
		}
	}

	score := 0
	for {
		if len(s) == 0 {
			break
		}
		char := chunks[s.pop()]
		score = (score * 5) + completeScores[char]
	}

	return true, score, 0
}

type navigationSubsystem struct {
	lines []line
}

func parseInput(path string) navigationSubsystem {
	ls, err := common.ReadLines(path)
	if err != nil {
		panic(err)
	}

	lines := []line{}
	for i := range ls {
		l := ls[i]
		chars := []string{}
		for j := range l {
			chars = append(chars, string(l[j]))
		}
		lines = append(lines, chars)
	}

	return navigationSubsystem{
		lines: lines,
	}
}

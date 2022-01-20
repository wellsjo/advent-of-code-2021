package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/wellsjo/advent-of-code-2021/common"
)

func main() {
	pt := parseInput("input.txt")
	// part1(pt)
	// part2Recursive(pt)
	part2(pt)
}

type polymerTemplate struct {
	polymer    string
	mappings   map[string]string
	charCounts map[string]int
	pairCounts map[pair]int

	// for logging
	step int
}

type pair struct {
	left, right string
}

func part2(pt polymerTemplate) {
	steps := 40
	pt.setCharCounts()
	pt.setInitialPairs()
	for i := 0; i < steps; i++ {
		pt.growPolymer()
	}
	printAnswer(pt.charCounts)
}

func (pt *polymerTemplate) growPolymer() {
	newPairCounts := map[pair]int{}
	for p, count := range pt.pairCounts {
		newChar := pt.mappings[p.left+p.right]
		newPair1 := pair{left: p.left, right: newChar}
		newPair2 := pair{left: newChar, right: p.right}
		pt.charCounts[newChar] += count
		newPairCounts[newPair1] += count
		newPairCounts[newPair2] += count
	}
	pt.pairCounts = newPairCounts
}

// This takes forever
func part2Recursive(pt polymerTemplate) {
	const maxSteps = 40
	pt.setCharCounts()
	pairs := pt.getInitialPairs()

	for i := range pairs {
		pt.countRecursive(pairs[i], 0, maxSteps)
	}
	log.Println("Done", pt.charCounts)
}

func (pt polymerTemplate) countRecursive(p pair, step, maxSteps int) {
	if step == maxSteps {
		return
	}

	newChar := pt.mappings[p.left+p.right]
	pt.charCounts[newChar]++

	pt.countRecursive(pair{left: p.left, right: newChar}, step+1, maxSteps)
	pt.countRecursive(pair{left: newChar, right: p.right}, step+1, maxSteps)
}

func (pt polymerTemplate) getInitialPairs() []pair {
	ps := []pair{}
	for i := range pt.polymer {
		if i == len(pt.polymer)-1 {
			break
		}
		p := pair{
			left:  string(pt.polymer[i]),
			right: string(pt.polymer[i+1]),
		}
		ps = append(ps, p)
	}
	return ps
}

func (pt polymerTemplate) setInitialPairs() {
	for _, p := range pt.getInitialPairs() {
		log.Println("Pair", p)
		pt.pairCounts[p]++
	}
}

func mergeCharCounts(charCounts ...map[string]int) map[string]int {
	newCharCount := map[string]int{}
	for _, cc := range charCounts {
		for k, v := range cc {
			newCharCount[k] += v
		}
	}
	return newCharCount
}

func part1(pt polymerTemplate) {
	steps := 10
	for i := 0; i < steps; i++ {
		pt.parse()
		log.Println("Step", i+1, "Length", len(pt.polymer))
	}

	counts := pt.getCharCounts()
	printAnswer(counts)
}

func printAnswer(charCounts map[string]int) {
	mostCommonCount := -1
	leastCommonCount := -1
	for _, count := range charCounts {
		if leastCommonCount == -1 {
			mostCommonCount = count
			leastCommonCount = count
		} else if count > mostCommonCount {
			mostCommonCount = count
		} else if count < leastCommonCount {
			leastCommonCount = count
		}
	}
	log.Println("Least", leastCommonCount, "Most", mostCommonCount)
	log.Println("Most - Least", mostCommonCount-leastCommonCount)
}

func (pt *polymerTemplate) parse() {
	newPolymer := ""
	for i := len(pt.polymer) - 1; i > 0; i-- {
		pair := string(pt.polymer[i-1]) + string(pt.polymer[i])
		if insertChar, ok := pt.mappings[pair]; ok {
			if i == len(pt.polymer)-1 {
				newPolymer = fmt.Sprintf("%s%s%s", string(pt.polymer[i-1]), insertChar, string(pt.polymer[i])) + newPolymer
			} else {
				newPolymer = fmt.Sprintf("%s%s", string(pt.polymer[i-1]), insertChar) + newPolymer
			}
		}
	}
	pt.polymer = newPolymer
}

func (pt *polymerTemplate) getCharCounts() map[string]int {
	counts := map[string]int{}
	for _, char := range pt.polymer {
		counts[string(char)]++
	}
	return counts
}

func (pt *polymerTemplate) setCharCounts() {
	counts := map[string]int{}
	for _, char := range pt.polymer {
		counts[string(char)]++
	}
	pt.charCounts = counts
}

func parseInput(path string) polymerTemplate {
	lines, err := common.ReadLines(path)
	if err != nil {
		panic(err)
	}
	pt := polymerTemplate{
		mappings:   make(map[string]string),
		pairCounts: make(map[pair]int),
		polymer:    lines[0],
	}
	for _, line := range lines[2:] {
		parts := strings.Split(line, " -> ")
		pt.mappings[parts[0]] = parts[1]
	}
	log.Println("Mapping length", len(pt.mappings))
	return pt
}

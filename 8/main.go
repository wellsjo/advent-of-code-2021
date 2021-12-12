package main

import (
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/wellsjo/advent-of-code-2021/common"
)

var knownDigitLengths = map[int]int{
	2: 1,
	4: 4,
	3: 7,
	7: 8,
}

func main() {
	models := parseInput("input.txt")

	count := 0
	for i := range models {
		count += models[i].decode()
	}
	log.Println("ANSWER", count)
}

type model struct {
	patterns, digits []string
	m                map[string]int
	rm               map[int]string
}

func (m model) decode() int {
	for _, pattern := range m.patterns {
		patternLength := len(pattern)
		knownDigit, ok := knownDigitLengths[patternLength]
		if ok {
			m.m[pattern] = knownDigit
			m.rm[knownDigit] = pattern
		}
	}

	for _, pattern := range m.patterns {
		patternLength := len(pattern)
		patternSet := newPatternSet(pattern)

		if patternLength == 5 {
			if m.isSubset(1, patternSet) {
				m.m[pattern] = 3
			} else if m.testForDigit2(patternSet) {
				m.m[pattern] = 2
			} else {
				m.m[pattern] = 5
			}
		} else if patternLength == 6 {
			if !m.isSubset(1, patternSet) {
				m.m[pattern] = 6
			} else if m.isSubset(4, patternSet) {
				m.m[pattern] = 9
			} else {
				m.m[pattern] = 0
			}
		}
	}

	intStrArr := []string{}
	for i := range m.digits {
		num := m.m[m.digits[i]]
		numStr := strconv.Itoa(num)
		intStrArr = append(intStrArr, numStr)
	}
	numStr := strings.Join(intStrArr, "")
	answer, err := strconv.Atoi(numStr)
	if err != nil {
		panic("invalid")
	}

	log.Println("DECODED", m.patterns, m.digits, answer)
	return answer
}

// Combine 4 with patternSet and compare with 8 (full set) to deduce 2
// This assumes 1 is already tested for.
func (m model) testForDigit2(patternSet set) bool {
	p := newPatternSet(m.rm[4])
	for k, _ := range patternSet {
		p[k] = struct{}{}
	}

	p2 := newPatternSet(m.rm[8])
	if len(p2) != len(p) {
		return false
	}
	for k, _ := range p {
		if _, ok := p2[k]; !ok {
			return false
		}
	}
	return true
}

func (m model) isSubset(digit int, patternSet set) bool {
	p := newPatternSet(m.rm[digit])
	return p.isSubset(patternSet)
}

type set map[string]struct{}

func (s set) isSubset(compare set) bool {
	if len(s) == 0 {
		return false
	}
	for k, _ := range s {
		if _, ok := compare[k]; !ok {
			return false
		}
	}
	return true
}

func newPatternSet(pattern string) set {
	ret := map[string]struct{}{}
	for i := range pattern {
		ret[string(pattern[i])] = struct{}{}
	}
	return ret
}

func parseInput(path string) []model {
	lines, err := common.ReadLines(path)
	if err != nil {
		panic(err)
	}

	models := []model{}
	for i := 0; i < len(lines); i++ {
		parts := strings.Split(lines[i], " | ")
		in := strings.Split(parts[0], " ")
		out := strings.Split(parts[1], " ")
		models = append(models, newModel(in, out))
	}

	return models
}

func newModel(in, out []string) model {
	in = sortInput(in)
	out = sortInput(out)
	return model{
		patterns: in,
		digits:   out,
		m:        map[string]int{},
		rm:       map[int]string{},
	}
}

func sortInput(in []string) []string {
	for i := range in {
		str := in[i]
		strArr := []string{}
		for j := range str {
			strArr = append(strArr, string(str[j]))
		}
		sort.Strings(strArr)
		in[i] = strings.Join(strArr, "")
	}
	return in
}

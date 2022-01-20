package main

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/wellsjo/advent-of-code-2021/common"
)

func main() {
	b := parseInput("test-input.txt")
	b.print()
}

type Board [][]int

func (p Board) print() {
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

func parseInput(path string) Board {
	lines, err := common.ReadLines(path)
	if err != nil {
		panic(err)
	}

	b := Board{}
	for _, line := range lines {
		intArr := []int{}
		for i := range line {
			numStr := string(line[i])
			num, err := strconv.Atoi(numStr)
			if err != nil {
				panic(err)
			}
			intArr = append(intArr, num)
		}
		b = append(b, intArr)
	}

	return b
}

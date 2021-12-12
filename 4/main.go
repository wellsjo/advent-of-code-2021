package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/wellsjo/advent-of-code-2021/common"
)

// 31542 too high

func main() {
	draws, boards := parseInput("input.txt")
	log.Println(draws)
	log.Println(boards)
	log.Printf("Starting\n\n")

	for i := 0; i < len(boards); i++ {
		log.Println("BOARD\n", boards[i])
		log.Println()
	}

	winners, lastPlayed := playGame(draws, boards)

	log.Println("WINNING BOARDS")
	log.Println(winners.getScore(lastPlayed))

	// log.Println("UNMARKED SUM", unmarkedSum, "* ", lastPlayed)

	// log.Println()
	// log.Println("ANSWER", unmarkedSum*lastPlayed)
}

func playGame(draws []int, boards []board) (board, int) {
	var (
		wonBoards = map[int]struct{}{}
		lastDraw  int
	)

	for i := range draws {
		for j := range boards {
			lastDraw = draws[i]
			matched := boards[j].applyDraw(lastDraw)
			if matched {
				if done, _ := boards[j].isDone(); done {
					wonBoards[j] = struct{}{}
					if len(wonBoards) == len(boards) {
						return boards[j], lastDraw
					}
				}
			}
		}
	}

	panic("invalid game")
}

func (b board) calculateUnmarkedSum() int {
	var unmarkedSum int
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b); j++ {
			if !b[i][j].marked {
				unmarkedSum += b[i][j].value
			}
		}
	}
	log.Println("UNMARKED SUM", unmarkedSum)
	return unmarkedSum
}

type position struct {
	value  int
	marked bool
}

type board [][]position

func (b board) getScore(lastPlayed int) int {
	return b.calculateUnmarkedSum() * lastPlayed
}

func (b board) String() string {
	s := ""
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			s += strconv.Itoa(b[i][j].value) + " "
		}
		s += "\n"
	}
	return s
}

func (b board) applyDraw(in int) bool {
	foundMark := false
	for i := range b {
		for j := range b[i] {
			if b[i][j].value == in {
				b[i][j].marked = true
				foundMark = true
			}
		}
	}
	return foundMark
}

func (b board) isDone() (bool, []position) {
	for i := 0; i < 5; i++ {
		if checkMatch(b.getRow(i)) {
			return true, b.getRow(i)
		}
		if checkMatch(b.getColumn(i)) {
			return true, b.getColumn(i)
		}
	}
	return false, nil
}

func checkMatch(r []position) bool {
	for i := 0; i < len(r); i++ {
		if !r[i].marked {
			return false
		}
	}
	return true
}

func (b board) getRow(p int) []position {
	return b[p]
}

func (b board) getColumn(p int) []position {
	ret := []position{}
	for i := 0; i < 5; i++ {
		ret = append(ret, b[i][p])
	}
	return ret
}

func parseInput(path string) ([]int, []board) {
	lines, err := common.ReadLines(path)
	if err != nil {
		panic(err)
	}

	draws := getDraws(lines[0])
	boards := getBoards(lines[2:])

	return draws, boards
}

func getDraws(s string) []int {
	intStrings := strings.Split(s, ",")
	draws := []int{}
	for i := range intStrings {
		draw, err := strconv.Atoi(intStrings[i])
		if err != nil {
			panic(err)
		}
		draws = append(draws, draw)
	}
	return draws
}

func getBoards(lines []string) []board {
	var (
		b      board
		boards []board
	)

	for i := range lines {
		if lines[i] == "" {
			continue
		}

		b = append(b, parseLine(lines[i]))

		if len(b) == 5 {
			boards = append(boards, b)
			b = board{}
			continue
		}
	}

	return boards
}

func parseLine(line string) []position {
	parts := strings.Split(line, " ")
	ret := []position{}

	for i := range parts {
		valueStr := strings.TrimSpace(parts[i])
		if valueStr == "" {
			continue
		}

		value, err := strconv.Atoi(valueStr)
		if err != nil {
			panic(err)
		}

		log.Println("ADDING VALUE", value)
		ret = append(ret, position{value: value})
	}
	return ret
}

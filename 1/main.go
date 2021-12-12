package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type harness struct {
	increases int
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		panic(err)
	}

	h := harness{}

	for i := range lines {
		if i == 0 {
			continue
		}
		if lines[i] > lines[i-1] {
			h.increases++
		}
	}
	log.Println(h.increases)
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineStr := scanner.Text()
		i, err := strconv.Atoi(lineStr)
		if err != nil {
			return nil, err
		}
		lines = append(lines, i)
	}
	return lines, scanner.Err()
}

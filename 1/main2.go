package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// 1066 too high

type harness struct {
	prevWindow int
	increases  int
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		panic(err)
	}

	h := harness{
		prevWindow: -1,
		increases:  0,
	}

	for i := range lines {
		if i < 2 {
			continue
		}

		var window int
		window += lines[i-2]
		window += lines[i-1]
		window += lines[i]

		if h.prevWindow != -1 {
			log.Println("COMPARING", window, h.prevWindow)
			if window > h.prevWindow {
				h.increases++
			}
		}

		h.prevWindow = window
	}

	fmt.Println(h.increases)
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

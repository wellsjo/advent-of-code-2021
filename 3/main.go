package main

import (
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

	o2rating := getOxygenGeneratorRating(lines)
	scrubRating := getScrubberRating(lines)
	log.Println(o2rating, scrubRating)
	log.Println(o2rating * scrubRating)
}

type report [][]string

func getOxygenGeneratorRating(lines []string) int {
	report := createReport(lines)
	dataWidth := len(report[0])
	for i := 0; i <= dataWidth-1; i++ {
		mostCommon, _ := report.getMostCommonBitForPosition(i, "1")
		report = report.filterReportByPositionWithBit(i, mostCommon)
		log.Println(report)
		if len(report) == 1 {
			ret, err := strconv.ParseInt(strings.Join(report[0], ""), 2, 64)
			if err != nil {
				panic(err)
			}
			return int(ret)
		}
	}
	panic("invalid input")
}

func getScrubberRating(lines []string) int {
	report := createReport(lines)
	dataWidth := len(report[0])
	for i := 0; i <= dataWidth-1; i++ {
		_, leastCommon := report.getMostCommonBitForPosition(i, "0")
		log.Println("LEAST COMMON", i, leastCommon)
		report = report.filterReportByPositionWithBit(i, leastCommon)
		log.Println(report)
		if len(report) == 1 {
			ret, err := strconv.ParseInt(strings.Join(report[0], ""), 2, 64)
			if err != nil {
				panic(err)
			}
			return int(ret)
		}
	}
	panic("invalid input")
}

func (r report) filterReportByPositionWithBit(p int, bit string) report {
	var newReport report
	for i := 0; i < len(r); i++ {
		if r[i][p] == bit {
			newReport = append(newReport, r[i])
		}
	}
	return newReport
}

// func (r report) getRates() (int, int) {
// 	var gammaBinary, epsilonBinary string
// 	dataWidth := len(r[0])

// 	for i := 0; i <= dataWidth-1; i++ {
// 		mostCommon, leastCommon := r.getMostCommonBitForPosition(i, "")
// 		gammaBinary += mostCommon
// 		epsilonBinary += leastCommon
// 	}

// 	gammaDecimal, err := strconv.ParseInt(gammaBinary, 2, 64)
// 	if err != nil {
// 		panic(err)
// 	}

// 	epsilonDecimal, err := strconv.ParseInt(epsilonBinary, 2, 64)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return int(gammaDecimal), int(epsilonDecimal)
// }

func (r report) getMostCommonBitForPosition(p int, preferOnEqual string) (string, string) {
	numZeroBits := 0
	numOneBits := 0
	for i := 0; i < len(r); i++ {
		if r[i][p] == "0" {
			numZeroBits++
		} else {
			numOneBits++
		}
	}
	if numZeroBits == numOneBits {
		log.Println("SAME")
		if preferOnEqual == "0" {
			return "0", "0"
		}
		return "1", "1"
	}
	if numZeroBits > numOneBits {
		return "0", "1"
	}
	return "1", "0"
}

func createReport(lines []string) report {
	outerArr := [][]string{}
	for i := range lines {
		line := lines[i]
		innerArr := []string{}
		for j := range line {
			innerArr = append(innerArr, string(line[j]))
		}
		outerArr = append(outerArr, innerArr)
	}
	return outerArr
}

package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/wellsjo/advent-of-code-2021/common"
)

func main() {
	model := createGame("input.txt")

	day := 256
	m := runModel(model, day)

	log.Println("MODEL", "DAY", day, m.totalFish)
}

func runModel(m model, numDays int) model {
	for i := 0; i < numDays; i++ {
		newFishToday := m.newFishOnDay[0]
		m.totalFish += newFishToday

		m.newFishOnDay[0] = m.newFishOnDay[1]
		m.newFishOnDay[1] = m.newFishOnDay[2]
		m.newFishOnDay[2] = m.newFishOnDay[3]
		m.newFishOnDay[3] = m.newFishOnDay[4]
		m.newFishOnDay[4] = m.newFishOnDay[5]
		m.newFishOnDay[5] = m.newFishOnDay[6]
		m.newFishOnDay[6] = newFishToday + m.newFishOnDay[7]
		m.newFishOnDay[7] = m.newFishOnDay[8]
		m.newFishOnDay[8] = newFishToday
	}
	return m
}

// func runModel(m model, numDays int) model {
// 	for i := 0; i < numDays; i++ {
// 		log.Println("MODEL DAY", i, m.totalFish, m.newFishOnDay, m.waitForFishOnDay)

// 		checkDay := i % 6
// 		log.Println("CHECK DAY", checkDay, m.newFishOnDay[checkDay])
// 		addFish := m.newFishOnDay[checkDay]
// 		waitForFish := m.waitForFishOnDay[checkDay]

// 		if waitForFish != 0 {
// 			log.Println("WAIT FOR FISH", i, waitForFish)
// 		}

// 		addFish -= waitForFish
// 		m.waitForFishOnDay[checkDay] = 0

// 		m.totalFish += addFish
// 		m.newFishOnDay[(i+2)%6] += addFish
// 		m.waitForFishOnDay[(i+2)%6] += addFish

// 		if i == numDays-1 {
// 			log.Println()
// 			log.Println("MODEL", i+1, m.totalFish, m.newFishOnDay)
// 		}
// 		log.Println()
// 	}
// 	return m
// }

type model struct {
	totalFish    int
	newFishOnDay map[int]int
}

func createGame(input string) model {
	lines, err := common.ReadLines(input)
	if err != nil {
		panic(err)
	}

	m := model{
		newFishOnDay: make(map[int]int),
	}

	fishIn := strings.Split(lines[0], ",")
	for i := range fishIn {
		in, err := strconv.Atoi(fishIn[i])
		if err != nil {
			panic(err)
		}
		m.newFishOnDay[in]++
		m.totalFish++
	}

	return m
}

// func runGame(m model, numDays int) model {
// 	for i := 0; i < numDays; i++ {
// 		log.Println("Day", i+1)
// 		newModel := m.progressDay()
// 		m = append(m, newModel...)
// 	}
// 	return m
// }

// type model struct {
// 	day      int
// 	fish     int
// 	upcoming map[int]int
// }

// func (m model) progressDay() model {
// 	appendModel := model{}
// 	for i := 0; i < len(m); i++ {
// 		if m[i].progress() {
// 			appendModel = append(appendModel, &lanternFish{internalTimer: 8})
// 		}
// 	}
// 	return appendModel
// }

// type lanternFish struct {
// 	internalTimer int
// }

// func (l *lanternFish) progress() bool {
// 	l.internalTimer--
// 	if l.internalTimer == -1 {
// 		l.internalTimer = 6
// 		return true
// 	}
// 	return false
// }

// func (l *lanternFish) String() string {
// 	return fmt.Sprintf("{%v}", l.internalTimer)
// }

// func createGame(input string) model {
// 	lines, err := common.ReadLines(input)
// 	if err != nil {
// 		panic(err)
// 	}

// 	m := model{}
// 	fishIn := strings.Split(lines[0], ",")
// 	for i := range fishIn {
// 		in, err := strconv.Atoi(fishIn[i])
// 		if err != nil {
// 			panic(err)
// 		}
// 		m = append(m, &lanternFish{internalTimer: in})
// 	}

// 	return m
// }

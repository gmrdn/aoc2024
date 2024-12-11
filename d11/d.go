package d

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type D struct {
	inputStream io.Reader
}

func NewD() *D {
	return &D{}
}

func (d *D) Input(input io.Reader) {
	d.inputStream = bufio.NewReader(input)
}

func (d *D) Run1() int {
	fileScanner := bufio.NewScanner(d.inputStream)
	fileScanner.Split(bufio.ScanLines)

	input := ""
	for fileScanner.Scan() {
		line := fileScanner.Text()
		input += line
	}

	currentStoneRow := strings.Split(input, " ")
	nextStoneRow := []string{}

	maxBlinks := 25

	for i := 0; i < maxBlinks; i++ {
		for _, stone := range currentStoneRow {

			if stone == "0" {
				nextStoneRow = append(nextStoneRow, "1")
				continue
			}
			if len(stone)%2 == 0 {
				stoneRunes := []rune(stone)
				left := string(stoneRunes[:len(stone)/2])
				right := string(stoneRunes[len(stone)/2:])
				rightAsInt, _ := strconv.Atoi(right)
				right = fmt.Sprintf("%d", rightAsInt)
				nextStoneRow = append(nextStoneRow, left)
				nextStoneRow = append(nextStoneRow, right)
			} else {
				stoneInt, _ := strconv.Atoi(stone)
				nextStoneRow = append(nextStoneRow, fmt.Sprintf("%d", stoneInt*2024))
			}

		}

		currentStoneRow = nextStoneRow
		// fmt.Println(currentStoneRow)
		nextStoneRow = []string{}

	}

	result := len(currentStoneRow)

	return result
}

func (d *D) Run2() int {
	fileScanner := bufio.NewScanner(d.inputStream)
	fileScanner.Split(bufio.ScanLines)

	input := ""
	for fileScanner.Scan() {
		line := fileScanner.Text()
		input += line
	}

	currentStoneRow := strings.Split(input, " ")

	maxBlinks := 75
	currentOccurrences := make(map[string]int)

	for _, stone := range currentStoneRow {
		currentOccurrences[stone]++
	}

	nextOccurrences := make(map[string]int)

	for i := 0; i < maxBlinks; i++ {

		// fmt.Println(currentOccurrences)

		for stone, occ := range currentOccurrences {

			if stone == "0" {
				nextOccurrences["1"] += occ
				continue
			}
			if len(stone)%2 == 0 {
				stoneRunes := []rune(stone)
				left := string(stoneRunes[:len(stone)/2])
				right := string(stoneRunes[len(stone)/2:])
				rightAsInt, _ := strconv.Atoi(right)
				right = fmt.Sprintf("%d", rightAsInt)

				nextOccurrences[left] += occ
				nextOccurrences[right] += occ
			} else {
				stoneInt, _ := strconv.Atoi(stone)
				nextOccurrences[fmt.Sprintf("%d", stoneInt*2024)] += occ
			}

		}

		currentOccurrences = nextOccurrences
		nextOccurrences = make(map[string]int)
		// fmt.Println("after calc", "current", currentOccurrences, "next", nextOccurrences)
	}

	result := 0
	for _, occ := range currentOccurrences {
		result += occ
	}

	return result
}

func (d *D) RunStr1() string {
	return ""
}

func (d *D) RunStr2() string {
	return ""
}

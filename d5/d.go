package d

import (
	"bufio"
	"io"
	"math/big"
	"slices"
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

func doSomething1(fileScanner *bufio.Scanner) int {
	result := 0
	return result
}

func doSomething2(fileScanner *bufio.Scanner) int {
	result := 0
	return result
}

func (d *D) Run1() int {
	fileScanner := bufio.NewScanner(d.inputStream)

	fileScanner.Split(bufio.ScanLines)

	fullText := ""
	for fileScanner.Scan() {
		line := fileScanner.Text()
		fullText += line + "\n"
	}

	sections := strings.Split(fullText, "\n\n")

	firstSection := sections[0]
	secondSection := sections[1]

	shouldBeBefore := make(map[int][]int)
	for _, line := range strings.Split(firstSection, "\n") {
		parts := strings.Split(line, "|")
		first, _ := strconv.Atoi(parts[0])
		second, _ := strconv.Atoi(parts[1])
		shouldBeBefore[first] = append(shouldBeBefore[first], second)
	}

	result := 0

	for _, line := range strings.Split(secondSection, "\n") {
		values := strings.Split(line, ",")
		numbers := []int{}
		for _, value := range values {
			number, _ := strconv.Atoi(value)
			numbers = append(numbers, number)
		}

		numbersBeforeSorting := slices.Clone(numbers)
		sortedNumbers := sortWithRules(numbers, shouldBeBefore)

		if slices.Equal(numbersBeforeSorting, sortedNumbers) {
			result += numbersBeforeSorting[len(numbersBeforeSorting)/2]
		}

	}
	return result
}

func sortWithRules(numbers []int, shouldBeBefore map[int][]int) []int {
	slices.SortFunc(numbers, func(a, b int) int {
		if slices.Contains(shouldBeBefore[a], b) {
			return -1
		}
		return 1
	})
	return numbers
}

func (d *D) Run2() int {
	fileScanner := bufio.NewScanner(d.inputStream)

	fileScanner.Split(bufio.ScanLines)

	fullText := ""
	for fileScanner.Scan() {
		line := fileScanner.Text()
		fullText += line + "\n"
	}

	sections := strings.Split(fullText, "\n\n")

	firstSection := sections[0]
	secondSection := sections[1]

	shouldBeBefore := make(map[int][]int)
	for _, line := range strings.Split(firstSection, "\n") {
		parts := strings.Split(line, "|")
		first, _ := strconv.Atoi(parts[0])
		second, _ := strconv.Atoi(parts[1])
		shouldBeBefore[first] = append(shouldBeBefore[first], second)
	}

	result := 0

	for _, line := range strings.Split(secondSection, "\n") {
		values := strings.Split(line, ",")
		numbers := []int{}
		for _, value := range values {
			number, _ := strconv.Atoi(value)
			numbers = append(numbers, number)
		}

		numbersBeforeSorting := slices.Clone(numbers)
		sortedNumbers := sortWithRules(numbers, shouldBeBefore)

		if !slices.Equal(numbersBeforeSorting, sortedNumbers) {
			result += sortedNumbers[len(sortedNumbers)/2]
		}

	}
	return result
}

func (d *D) RunStr1() string {
	return ""
}

func (d *D) RunStr2() string {
	return ""
}

func (d *D) Run1BigInt() big.Int {
	return big.Int{}
}

func (d *D) Run2BigInt() big.Int {
	return big.Int{}
}

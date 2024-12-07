package d

import (
	"bufio"
	"io"
	"regexp"
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
	for fileScanner.Scan() {
		line := fileScanner.Text()

		mulMatches := regexp.MustCompile(`mul\((\d+),(\d+)\)`).FindAllStringSubmatch(line, -1)
		for _, match := range mulMatches {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			result += num1 * num2
		}
	}
	return result
}

func doSomething2(fileScanner *bufio.Scanner) int {
	result := 0
	bigLine := ""
	for fileScanner.Scan() {
		line := fileScanner.Text()
		bigLine += line
	}

	instructionsStartingWithDo := strings.Split(bigLine, "do()")
	for _, instruction := range instructionsStartingWithDo {
		splitOnDont := strings.Split(instruction, "don't()")
		doPart := splitOnDont[0]

		mulMatches := regexp.MustCompile(`mul\((\d+),(\d+)\)`).FindAllStringSubmatch(doPart, -1)
		for _, match := range mulMatches {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			result += num1 * num2
		}

	}
	return result
}

func (d *D) Run1() int {
	fileScanner := bufio.NewScanner(d.inputStream)

	fileScanner.Split(bufio.ScanLines)

	return doSomething1(fileScanner)
}

func (d *D) Run2() int {
	fileScanner := bufio.NewScanner(d.inputStream)

	fileScanner.Split(bufio.ScanLines)

	return doSomething2(fileScanner)
}

func (d *D) RunStr1() string {
	return ""
}

func (d *D) RunStr2() string {
	return ""
}

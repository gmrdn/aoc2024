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

func fillArrays(fileScanner *bufio.Scanner) ([]int, []int) {
	leftArray := []int{}
	rightArray := []int{}
	for fileScanner.Scan() {
		currentLine := fileScanner.Text()
		allNumbers := strings.Split(currentLine, "   ")

		leftNumber, _ := strconv.Atoi(allNumbers[0])
		rightNumber, _ := strconv.Atoi(allNumbers[1])
		leftArray = append(leftArray, leftNumber)
		rightArray = append(rightArray, rightNumber)

	}
	return leftArray, rightArray
}

func (d *D) Run1() int {
	fileScanner := bufio.NewScanner(d.inputStream)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		currentLine := fileScanner.Text()
		fmt.Println(currentLine)
	}

	return 0
}

func (d *D) Run2() int {
	fileScanner := bufio.NewScanner(d.inputStream)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		currentLine := fileScanner.Text()
		fmt.Println(currentLine)
	}

	return 0
}

func (d *D) RunStr1() string {
	return ""
}

func (d *D) RunStr2() string {
	return ""
}

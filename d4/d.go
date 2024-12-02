package d

import (
	"bufio"
	"io"
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

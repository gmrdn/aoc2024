package d

import (
	"bufio"
	"fmt"
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

func (d *D) Run1() int {
	fileScanner := bufio.NewScanner(d.inputStream)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		fmt.Println(line)
	}

	result := 0

	return result
}

func (d *D) Run2() int {
	fileScanner := bufio.NewScanner(d.inputStream)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		fmt.Println(line)
	}

	result := 0

	return result
}

func (d *D) RunStr1() string {
	return ""
}

func (d *D) RunStr2() string {
	return ""
}

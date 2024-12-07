package d

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
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

func calc(numbers []big.Int, expectedResult big.Int) bool {
	if len(numbers) == 0 {
		return false
	}

	if len(numbers) == 1 {
		return numbers[0].Cmp(&expectedResult) == 0
	}

	for i := 0; i < len(numbers); i++ {
		if numbers[i].Cmp(&expectedResult) == 0 {
			return true
		}

		targetAddition := new(big.Int)
		targetAddition.Sub(&expectedResult, &numbers[i])

		targetMultiplication := new(big.Int)
		targetMultiplication.Div(&expectedResult, &numbers[i])
		// targetMultiplication := expectedResult / numbers[i]
		// fmt.Println("targetMultiplication", numbers[i], numbers[i+1:], targetMultiplication)
		// fmt.Println("targetAddition", numbers[i], numbers[i+1:], targetAddition)

		if calc(numbers[i+1:], *targetMultiplication) || calc(numbers[i+1:], *targetAddition) {
			return true
		}
	}

	return false
}

func (d *D) Run1() int {
	return 0
}

func (d *D) Run2() int {
	result := 0

	return result
}

func (d *D) RunStr1() string {
	return ""
}

func (d *D) RunStr2() string {
	return ""
}

func (d *D) Run1BigInt() big.Int {
	fileScanner := bufio.NewScanner(d.inputStream)
	fileScanner.Split(bufio.ScanLines)

	fmt.Println("running d7 Run1Int64")

	result := new(big.Int)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		elems := strings.Split(line, ": ")
		expectedResult := new(big.Int)
		expectedResult.SetString(elems[0], 10)

		numbersAsString := strings.Split(elems[1], " ")
		numbers := []big.Int{}
		for i := 0; i < len(numbersAsString); i++ {
			number := big.Int{}
			number.SetString(numbersAsString[i], 10)
			numbers = append(numbers, number)
		}

		canBeCalculated := calc(numbers, *expectedResult)

		if canBeCalculated {
			result.Add(result, expectedResult)
		} else {
			fmt.Println("not possible", numbers, expectedResult)
		}

	}

	return *result
}

func (d *D) Run2BigInt() big.Int {
	return big.Int{}
}

package d

import (
	"bufio"
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

func calc2(currentCalc big.Int, remainingNumbers []big.Int, expectedResult big.Int) bool {
	if len(remainingNumbers) == 0 {
		return false
	}

	for i := 0; i < len(remainingNumbers); i++ {
		newCalcAddition := new(big.Int)
		newCalcAddition.Add(&currentCalc, &remainingNumbers[i])

		newCalcMultiplication := new(big.Int)
		newCalcMultiplication.Mul(&currentCalc, &remainingNumbers[i])

		newCalcStringConcatenation := new(big.Int)
		currentCalcString := currentCalc.String()
		remainingNumbersString := remainingNumbers[i].String()
		newCalcStringConcatenation.SetString(currentCalcString+remainingNumbersString, 10)

		if (newCalcAddition.Cmp(&expectedResult) == 0 ||
			newCalcMultiplication.Cmp(&expectedResult) == 0 ||
			newCalcStringConcatenation.Cmp(&expectedResult) == 0) &&
			len(remainingNumbers[i+1:]) == 0 {
			return true
		}

		if calc2(*newCalcAddition, remainingNumbers[i+1:], expectedResult) {
			return true
		}
		if calc2(*newCalcMultiplication, remainingNumbers[i+1:], expectedResult) {
			return true
		}
		if calc2(*newCalcStringConcatenation, remainingNumbers[i+1:], expectedResult) {
			return true
		}
	}

	return false
}

func calc(currentCalc big.Int, remainingNumbers []big.Int, expectedResult big.Int) bool {
	if len(remainingNumbers) == 0 {
		return false
	}

	for i := 0; i < len(remainingNumbers); i++ {
		newCalcAddition := new(big.Int)
		newCalcAddition.Add(&currentCalc, &remainingNumbers[i])

		newCalcMultiplication := new(big.Int)
		newCalcMultiplication.Mul(&currentCalc, &remainingNumbers[i])

		if newCalcAddition.Cmp(&expectedResult) == 0 || newCalcMultiplication.Cmp(&expectedResult) == 0 && len(remainingNumbers[i+1:]) == 0 {
			return true
		}

		if calc(*newCalcAddition, remainingNumbers[i+1:], expectedResult) {
			return true
		}
		if calc(*newCalcMultiplication, remainingNumbers[i+1:], expectedResult) {
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

	result := new(big.Int)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		elems := strings.Split(line, ": ")
		expectedResult := big.Int{}
		expectedResult.SetString(elems[0], 10)

		numbersAsString := strings.Split(elems[1], " ")
		numbers := []big.Int{}
		for i := 0; i < len(numbersAsString); i++ {
			number := big.Int{}
			number.SetString(numbersAsString[i], 10)
			numbers = append(numbers, number)
		}

		canBeCalculated := calc(numbers[0], numbers[1:], expectedResult)

		if canBeCalculated {
			result.Add(result, &expectedResult)
		}
	}

	return *result
}

func (d *D) Run2BigInt() big.Int {
	fileScanner := bufio.NewScanner(d.inputStream)
	fileScanner.Split(bufio.ScanLines)

	result := new(big.Int)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		elems := strings.Split(line, ": ")
		expectedResult := big.Int{}
		expectedResult.SetString(elems[0], 10)

		numbersAsString := strings.Split(elems[1], " ")
		numbers := []big.Int{}
		for i := 0; i < len(numbersAsString); i++ {
			number := big.Int{}
			number.SetString(numbersAsString[i], 10)
			numbers = append(numbers, number)
		}

		canBeCalculated := calc2(numbers[0], numbers[1:], expectedResult)

		if canBeCalculated {
			result.Add(result, &expectedResult)
		}
	}

	return *result
}

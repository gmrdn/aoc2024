package d

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
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

func calc2(currentCalc int, remainingNumbers []int, expectedResult int) bool {
	if len(remainingNumbers) == 0 {
		return expectedResult == currentCalc
	}

	nextNumber := remainingNumbers[0]
	newCalcAddition := currentCalc + nextNumber
	newCalcMultiplication := currentCalc * nextNumber

	currentCalcString := strconv.Itoa(currentCalc)
	nextNumbersString := strconv.Itoa(nextNumber)
	newCalcStringConcatenation, err := strconv.Atoi(currentCalcString + nextNumbersString)
	if err != nil {
		fmt.Println(err)
		return false
	}

	if calc2(newCalcAddition, remainingNumbers[1:], expectedResult) ||
		calc2(newCalcMultiplication, remainingNumbers[1:], expectedResult) ||
		calc2(newCalcStringConcatenation, remainingNumbers[1:], expectedResult) {
		return true
	}
	// for i := 0; i < len(remainingNumbers); i++ {
	// 	newCalcAddition := currentCalc + remainingNumbers[i]
	// 	newCalcMultiplication := currentCalc * remainingNumbers[i]
	// 	currentCalcString := strconv.Itoa(currentCalc)
	// 	remainingNumbersString := strconv.Itoa(remainingNumbers[i])
	// 	newCalcStringConcatenation, err := strconv.Atoi(currentCalcString + remainingNumbersString)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return false
	// 	}
	//
	// 	if calc2(newCalcAddition, remainingNumbers[i+1:], expectedResult) ||
	// 		calc2(newCalcMultiplication, remainingNumbers[i+1:], expectedResult) ||
	// 		calc2(newCalcStringConcatenation, remainingNumbers[i+1:], expectedResult) {
	// 		return true
	// 	}
	// }

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

		if (newCalcAddition.Cmp(&expectedResult) == 0 || newCalcMultiplication.Cmp(&expectedResult) == 0) && len(remainingNumbers[i+1:]) == 0 {
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
	fileScanner := bufio.NewScanner(d.inputStream)
	fileScanner.Split(bufio.ScanLines)

	result := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		elems := strings.Split(line, ": ")
		expectedResult, err := strconv.Atoi(elems[0])
		if err != nil {
			fmt.Println(err)
			return 0
		}

		numbersAsString := strings.Split(elems[1], " ")
		numbers := []int{}
		for i := 0; i < len(numbersAsString); i++ {
			number, _ := strconv.Atoi(numbersAsString[i])
			numbers = append(numbers, number)
		}

		canBeCalculated := calc2(numbers[0], numbers[1:], expectedResult)

		if canBeCalculated {
			result += expectedResult
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
	return *new(big.Int)
}

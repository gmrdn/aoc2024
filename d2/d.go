package d

import (
	"bufio"
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

	nbSafe := 0
	nbUnsafe := 0
	for fileScanner.Scan() {
		currentLine := fileScanner.Text()
		allNumbers := strings.Split(currentLine, " ")
		previousNumber, _ := strconv.Atoi(allNumbers[0])
		secondNumber, _ := strconv.Atoi(allNumbers[1])
		diff := secondNumber - previousNumber
		allIncreasing := true
		if diff < 0 {
			allIncreasing = false
		}

		safeRow := true

		for i := 1; i < len(allNumbers); i++ {
			currentNumber, _ := strconv.Atoi(allNumbers[i])
			diff := currentNumber - previousNumber
			if allIncreasing && diff < 0 {
				safeRow = false
				break
			}

			if !allIncreasing && diff > 0 {
				safeRow = false
				break
			}

			if diff < 0 {
				diff = -diff
			}
			if diff < 1 || diff > 3 {
				safeRow = false
				break
			}
			previousNumber = currentNumber
		}
		if safeRow {
			nbSafe++
		} else {
			nbUnsafe++
		}
	}

	return nbSafe
}

func evaluateRow(numbers []string) (bool, int) {
	previousNumber, _ := strconv.Atoi(numbers[0])

	safeRow := true

	failedAt := -1

	generallyIncreasing := true
	nbIncreasing := 0
	for i := 1; i < len(numbers); i++ {
		currentNumber, _ := strconv.Atoi(numbers[i])
		if currentNumber > previousNumber {
			nbIncreasing++
		}
		previousNumber = currentNumber
	}

	if nbIncreasing < len(numbers)-2 {
		generallyIncreasing = false
	}

	previousNumber, _ = strconv.Atoi(numbers[0])
	for i := 1; i < len(numbers); i++ {
		currentNumber, _ := strconv.Atoi(numbers[i])
		diff := currentNumber - previousNumber

		if generallyIncreasing && diff < 0 {
			safeRow = false
			failedAt = i
			break
		}

		if !generallyIncreasing && diff > 0 {
			safeRow = false
			failedAt = i

			break
		}

		if diff < 0 {
			diff = -diff
		}
		if diff < 1 || diff > 3 {
			safeRow = false
			failedAt = i

			break
		}
		previousNumber = currentNumber
	}

	return safeRow, failedAt
}

func (d *D) Run2() int {
	fileScanner := bufio.NewScanner(d.inputStream)

	fileScanner.Split(bufio.ScanLines)

	nbSafe := 0
	nbUnsafe := 0
	for fileScanner.Scan() {
		currentLine := fileScanner.Text()
		allNumbers := strings.Split(currentLine, " ")
		copyAllNumbers1 := strings.Split(currentLine, " ")
		copyAllNumbers2 := strings.Split(currentLine, " ")

		safeRow, failedAt := evaluateRow(allNumbers)

		if safeRow {
			nbSafe++
		} else {
			if failedAt != -1 {
				safeRowRemoveCurrent := false
				safeRowRemoveBefore := false
				safeRowRemoveAfter := false

				newNumbersRemoveCurrent := append(allNumbers[:failedAt], allNumbers[failedAt+1:]...)
				safeRowRemoveCurrent, _ = evaluateRow(newNumbersRemoveCurrent)

				if failedAt > 0 {
					newNumbersRemoveBefore := append(copyAllNumbers1[:failedAt-1], copyAllNumbers1[failedAt:]...)
					safeRowRemoveBefore, _ = evaluateRow(newNumbersRemoveBefore)
				}
				if failedAt < len(allNumbers)-1 {
					newNumbersRemoveAfter := append(copyAllNumbers2[:failedAt+1], copyAllNumbers2[failedAt+2:]...)
					safeRowRemoveAfter, _ = evaluateRow(newNumbersRemoveAfter)
				}
				if safeRowRemoveCurrent || safeRowRemoveBefore || safeRowRemoveAfter {
					nbSafe++
					continue
				}
			}
			nbUnsafe++
		}
	}

	return nbSafe
}

func (d *D) RunStr1() string {
	return ""
}

func (d *D) RunStr2() string {
	return ""
}

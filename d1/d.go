package d

import (
	"bufio"
	"io"
	"math/big"
	"sort"
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

	leftArray, rightArray := fillArrays(fileScanner)
	sort.Ints(leftArray)
	sort.Ints(rightArray)

	distances := []int{}

	for i := 0; i < len(leftArray); i++ {

		currentDistance := rightArray[i] - leftArray[i]
		if currentDistance < 0 {
			currentDistance = currentDistance * -1
		}
		distances = append(distances, currentDistance)
	}

	sum := 0
	for i := 0; i < len(distances); i++ {
		sum += distances[i]
	}

	return sum
}

func (d *D) Run2() int {
	fileScanner := bufio.NewScanner(d.inputStream)
	fileScanner.Split(bufio.ScanLines)
	leftArray, rightArray := fillArrays(fileScanner)

	rightSideOccurences := make(map[int]int)
	for i := 0; i < len(rightArray); i++ {
		rightSideOccurences[rightArray[i]]++
	}

	similarityScore := 0

	for i := 0; i < len(leftArray); i++ {
		if rightSideOccurences[leftArray[i]] > 0 {
			similarityScore = similarityScore + leftArray[i]*rightSideOccurences[leftArray[i]]
		}
	}

	return similarityScore
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

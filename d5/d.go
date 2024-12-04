package d

import (
	"bufio"
	"fmt"
	"io"
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

	matrix := [][]string{}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		matrix = append(matrix, strings.Split(line, ""))
	}

	nbXMAS := 0
	for i, row := range matrix {
		for j, col := range row {
			if col == "X" {
				// up
				wordUp := []string{}
				if i-3 >= 0 {
					wordUp = append(wordUp, matrix[i-1][j])
					wordUp = append(wordUp, matrix[i-2][j])
					wordUp = append(wordUp, matrix[i-3][j])
					if strings.Join(wordUp, "") == "MAS" {
						nbXMAS++
					}
				}
				wordUpRight := []string{}
				if i-3 >= 0 && j+3 < len(row) {
					wordUpRight = append(wordUpRight, matrix[i-1][j+1])
					wordUpRight = append(wordUpRight, matrix[i-2][j+2])
					wordUpRight = append(wordUpRight, matrix[i-3][j+3])
					if strings.Join(wordUpRight, "") == "MAS" {
						nbXMAS++
					}
				}
				wordRight := []string{}
				if j+3 < len(row) {
					wordRight = append(wordRight, matrix[i][j+1])
					wordRight = append(wordRight, matrix[i][j+2])
					wordRight = append(wordRight, matrix[i][j+3])
					if strings.Join(wordRight, "") == "MAS" {
						nbXMAS++
					}
				}
				wordDownRight := []string{}
				if i+3 < len(matrix) && j+3 < len(row) {
					wordDownRight = append(wordDownRight, matrix[i+1][j+1])
					wordDownRight = append(wordDownRight, matrix[i+2][j+2])
					wordDownRight = append(wordDownRight, matrix[i+3][j+3])
					if strings.Join(wordDownRight, "") == "MAS" {
						nbXMAS++
					}
				}
				wordDown := []string{}
				if i+3 < len(matrix) {
					wordDown = append(wordDown, matrix[i+1][j])
					wordDown = append(wordDown, matrix[i+2][j])
					wordDown = append(wordDown, matrix[i+3][j])
					if strings.Join(wordDown, "") == "MAS" {
						nbXMAS++
					}
				}
				wordDownLeft := []string{}
				if i+3 < len(matrix) && j-3 >= 0 {
					wordDownLeft = append(wordDownLeft, matrix[i+1][j-1])
					wordDownLeft = append(wordDownLeft, matrix[i+2][j-2])
					wordDownLeft = append(wordDownLeft, matrix[i+3][j-3])
					if strings.Join(wordDownLeft, "") == "MAS" {
						nbXMAS++
					}
				}
				wordLeft := []string{}
				if j-3 >= 0 {
					wordLeft = append(wordLeft, matrix[i][j-1])
					wordLeft = append(wordLeft, matrix[i][j-2])
					wordLeft = append(wordLeft, matrix[i][j-3])
					if strings.Join(wordLeft, "") == "MAS" {
						nbXMAS++
					}
				}
				wordUpLeft := []string{}
				if i-3 >= 0 && j-3 >= 0 {
					wordUpLeft = append(wordUpLeft, matrix[i-1][j-1])
					wordUpLeft = append(wordUpLeft, matrix[i-2][j-2])
					wordUpLeft = append(wordUpLeft, matrix[i-3][j-3])
					if strings.Join(wordUpLeft, "") == "MAS" {
						nbXMAS++
					}
				}

			}
		}
	}

	return nbXMAS
}

func (d *D) Run2() int {
	fileScanner := bufio.NewScanner(d.inputStream)

	fileScanner.Split(bufio.ScanLines)

	matrix := [][]string{}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		matrix = append(matrix, strings.Split(line, ""))
	}

	nbXMAS := 0
	for i, row := range matrix {
		for j, col := range row {
			if col == "A" {
				wordDiagonalRight := []string{}
				wordDiagonalLeft := []string{}
				if i-1 >= 0 && i+1 < len(matrix) && j-1 >= 0 && j+1 < len(row) {
					fmt.Println("trying A at", i, j)
					wordDiagonalRight = append(wordDiagonalRight, matrix[i-1][j-1])
					wordDiagonalRight = append(wordDiagonalRight, "A")
					wordDiagonalRight = append(wordDiagonalRight, matrix[i+1][j+1])
					fmt.Println("wordDiagonalRight", wordDiagonalRight)
					wordDiagonalLeft = append(wordDiagonalLeft, matrix[i-1][j+1])
					wordDiagonalLeft = append(wordDiagonalLeft, "A")
					wordDiagonalLeft = append(wordDiagonalLeft, matrix[i+1][j-1])
					fmt.Println("wordDiagonalLeft", wordDiagonalLeft)

					if strings.Join(wordDiagonalRight, "") == "MAS" || strings.Join(wordDiagonalRight, "") == "SAM" {
						if strings.Join(wordDiagonalLeft, "") == "MAS" || strings.Join(wordDiagonalLeft, "") == "SAM" {
							nbXMAS++
						}
					}
				}
			}
		}
	}

	return nbXMAS
}

func (d *D) RunStr1() string {
	return ""
}

func (d *D) RunStr2() string {
	return ""
}

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

type Point struct {
	x int
	y int
}

type Node struct {
	point    Point
	distance int
}

func getMatrix(fileScanner *bufio.Scanner) [][]int {
	matrix := make([][]int, 0)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		numbersAsStrings := strings.Split(line, "")
		numbers := make([]int, 0)
		for _, numberAsString := range numbersAsStrings {
			number := 0
			if numberAsString == "." {
				number = -1
			} else {
				number, _ = strconv.Atoi(numberAsString)
			}

			numbers = append(numbers, number)
		}
		matrix = append(matrix, numbers)
	}
	return matrix
}

func (d *D) Run1() int {
	fileScanner := bufio.NewScanner(d.inputStream)
	fileScanner.Split(bufio.ScanLines)

	matrix := getMatrix(fileScanner)

	result := 0

	for i, row := range matrix {
		for j, col := range row {
			if col == 0 {
				// begin of 0 cell

				nbFullPath := 0

				currentPoint := Point{x: i, y: j}
				queue := []Node{{
					point:    currentPoint,
					distance: 0,
				}}
				visited := make(map[Point]bool)
				visited[currentPoint] = true

				for len(queue) > 0 {
					currentNode := queue[0]
					queue = queue[1:]

					currentAltitude := matrix[currentNode.point.x][currentNode.point.y]

					if currentNode.distance == 9 {
						nbFullPath++
						// break
					}

					for _, next := range d.GetNextPositions(matrix, currentNode) {
						if isValid(matrix, visited, next.point, currentAltitude) {
							visited[next.point] = true
							queue = append(queue, next)
						}
					}

				}

				result += nbFullPath

				// end of 0 cell
			}
		}
	}

	return result
}

func (d *D) Run2() int {
	fileScanner := bufio.NewScanner(d.inputStream)
	fileScanner.Split(bufio.ScanLines)

	matrix := getMatrix(fileScanner)

	result := 0

	for i, row := range matrix {
		for j, col := range row {
			if col == 0 {
				// begin of 0 cell

				nbFullPath := 0

				currentPoint := Point{x: i, y: j}
				queue := []Node{{
					point:    currentPoint,
					distance: 0,
				}}

				for len(queue) > 0 {
					currentNode := queue[0]
					queue = queue[1:]

					currentAltitude := matrix[currentNode.point.x][currentNode.point.y]

					if currentNode.distance == 9 {
						nbFullPath++
						// break
					}

					for _, next := range d.GetNextPositions(matrix, currentNode) {
						if isValid2(matrix, next.point, currentAltitude) {
							queue = append(queue, next)
						}
					}

				}

				result += nbFullPath

				// end of 0 cell
			}
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

func (d *D) GetNextPositions(mat [][]int, current Node) []Node {
	positions := []Node{}
	positions = append(positions, Node{
		point:    Point{x: current.point.x, y: current.point.y - 1},
		distance: current.distance + 1,
	})

	positions = append(positions, Node{
		point:    Point{x: current.point.x, y: current.point.y + 1},
		distance: current.distance + 1,
	})

	positions = append(positions, Node{
		point:    Point{x: current.point.x - 1, y: current.point.y},
		distance: current.distance + 1,
	})

	positions = append(positions, Node{
		point:    Point{x: current.point.x + 1, y: current.point.y},
		distance: current.distance + 1,
	})

	return positions
}

func isValid(mat [][]int, visited map[Point]bool, pos Point, currentAltitude int) bool {
	if pos.x < 0 || pos.x >= len(mat[0]) {
		return false
	}

	if pos.y < 0 || pos.y >= len(mat) {
		return false
	}

	if visited[pos] {
		return false
	}

	altitude := mat[pos.x][pos.y]

	return altitude == currentAltitude+1
}

func isValid2(mat [][]int, pos Point, currentAltitude int) bool {
	if pos.x < 0 || pos.x >= len(mat[0]) {
		return false
	}

	if pos.y < 0 || pos.y >= len(mat) {
		return false
	}

	altitude := mat[pos.x][pos.y]

	return altitude == currentAltitude+1
}

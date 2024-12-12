package d

import (
	"bufio"
	"io"
	"slices"
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

type Position struct {
	x int
	y int
}
type Node struct {
	position       Position
	value          string
	nextNodes      []*Node
	hasBorderUp    bool
	hasBorderDown  bool
	hasBorderLeft  bool
	hasBorderRight bool
}

type Zone struct {
	nodes []*Node
}

func (d *D) Run1() int {
	fileScanner := bufio.NewScanner(d.inputStream)
	fileScanner.Split(bufio.ScanLines)

	matrix := make([][]Node, 0)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		allValues := strings.Split(line, "")
		nodesInRow := make([]Node, 0)
		for i := 0; i < len(allValues); i++ {
			position := Position{x: i, y: len(matrix)}
			node := Node{position: position, value: allValues[i]}
			nodesInRow = append(nodesInRow, node)
		}

		matrix = append(matrix, nodesInRow)
	}

	// fmt.Println(matrix)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {

			currentNode := &matrix[i][j]
			currentValue := currentNode.value

			// find neighbors
			if i > 0 {
				// top
				if matrix[i-1][j].value == currentValue {
					currentNode.nextNodes = append(currentNode.nextNodes, &matrix[i-1][j])
				} else {
					currentNode.hasBorderUp = true
				}
			} else {
				currentNode.hasBorderUp = true
			}
			if i < len(matrix)-1 {
				// bottom
				if matrix[i+1][j].value == currentValue {
					currentNode.nextNodes = append(currentNode.nextNodes, &matrix[i+1][j])
				} else {
					currentNode.hasBorderDown = true
				}
			} else {
				currentNode.hasBorderDown = true
			}
			if j > 0 {
				// left
				if matrix[i][j-1].value == currentValue {
					currentNode.nextNodes = append(currentNode.nextNodes, &matrix[i][j-1])
				} else {
					currentNode.hasBorderLeft = true
				}
			} else {
				currentNode.hasBorderLeft = true
			}
			if j < len(matrix[i])-1 {
				// right
				if matrix[i][j+1].value == currentValue {
					currentNode.nextNodes = append(currentNode.nextNodes, &matrix[i][j+1])
				} else {
					currentNode.hasBorderRight = true
				}
			} else {
				currentNode.hasBorderRight = true
			}
		}
	}

	result := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {

			firstNode := &matrix[i][j]

			if firstNode.value == "" {
				continue
			}

			firstZone := Zone{nodes: []*Node{}}
			addNodesToZone(&firstZone, firstNode)

			// fmt.Println(firstZone)

			nbBordersInZone := 0
			nbNodesInZone := 0
			for i := 0; i < len(firstZone.nodes); i++ {
				node := firstZone.nodes[i]
				if node.hasBorderUp {
					nbBordersInZone++
				}
				if node.hasBorderDown {
					nbBordersInZone++
				}
				if node.hasBorderLeft {
					nbBordersInZone++
				}
				if node.hasBorderRight {
					nbBordersInZone++
				}
				nbNodesInZone++
				matrix[node.position.y][node.position.x].value = ""
			}

			// fmt.Println("nbBordersInZone", nbBordersInZone)
			// fmt.Println("nbNodesInZone", nbNodesInZone)

			result += nbBordersInZone * nbNodesInZone

		}
	}

	return result
}

func addNodesToZone(zone *Zone, node *Node) {
	zone.nodes = append(zone.nodes, node)
	for i := 0; i < len(node.nextNodes); i++ {
		if node.nextNodes[i] == nil {
			continue
		}
		nextNode := node.nextNodes[i]
		if !slices.Contains(zone.nodes, nextNode) {
			addNodesToZone(zone, nextNode)
		}
	}
}

func (d *D) Run2() int {
	fileScanner := bufio.NewScanner(d.inputStream)
	fileScanner.Split(bufio.ScanLines)

	matrix := make([][]Node, 0)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		allValues := strings.Split(line, "")
		nodesInRow := make([]Node, 0)
		for i := 0; i < len(allValues); i++ {
			position := Position{x: i, y: len(matrix)}
			node := Node{position: position, value: allValues[i]}
			nodesInRow = append(nodesInRow, node)
		}

		matrix = append(matrix, nodesInRow)
	}

	// fmt.Println(matrix)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {

			currentNode := &matrix[i][j]
			currentValue := currentNode.value

			// find neighbors
			if i > 0 {
				// top
				if matrix[i-1][j].value == currentValue {
					currentNode.nextNodes = append(currentNode.nextNodes, &matrix[i-1][j])
				} else {
					currentNode.hasBorderUp = true
				}
			} else {
				currentNode.hasBorderUp = true
			}
			if i < len(matrix)-1 {
				// bottom
				if matrix[i+1][j].value == currentValue {
					currentNode.nextNodes = append(currentNode.nextNodes, &matrix[i+1][j])
				} else {
					currentNode.hasBorderDown = true
				}
			} else {
				currentNode.hasBorderDown = true
			}
			if j > 0 {
				// left
				if matrix[i][j-1].value == currentValue {
					currentNode.nextNodes = append(currentNode.nextNodes, &matrix[i][j-1])
				} else {
					currentNode.hasBorderLeft = true
				}
			} else {
				currentNode.hasBorderLeft = true
			}
			if j < len(matrix[i])-1 {
				// right
				if matrix[i][j+1].value == currentValue {
					currentNode.nextNodes = append(currentNode.nextNodes, &matrix[i][j+1])
				} else {
					currentNode.hasBorderRight = true
				}
			} else {
				currentNode.hasBorderRight = true
			}
		}
	}

	result := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {

			firstNode := &matrix[i][j]

			if firstNode.value == "" {
				continue
			}

			firstZone := Zone{nodes: []*Node{}}
			addNodesToZone(&firstZone, firstNode)

			// fmt.Println(firstZone)

			nbBordersInZone := 0
			nbNodesInZone := 0
			bordersUp := make(map[int][]int)
			bordersDown := make(map[int][]int)
			bordersLeft := make(map[int][]int)
			bordersRight := make(map[int][]int)

			for i := 0; i < len(firstZone.nodes); i++ {
				node := firstZone.nodes[i]
				if node.hasBorderUp {
					bordersUp[node.position.y] = append(bordersUp[node.position.y], node.position.x)
					nbBordersInZone++
				}
				if node.hasBorderDown {
					bordersDown[node.position.y] = append(bordersDown[node.position.y], node.position.x)
					nbBordersInZone++
				}
				if node.hasBorderLeft {
					bordersLeft[node.position.x] = append(bordersLeft[node.position.x], node.position.y)
					nbBordersInZone++
				}
				if node.hasBorderRight {
					bordersRight[node.position.x] = append(bordersRight[node.position.x], node.position.y)
					nbBordersInZone++
				}

				nbNodesInZone++
				matrix[node.position.y][node.position.x].value = ""
			}
			// fmt.Println("borders up", bordersUp)
			// fmt.Println("borders down", bordersDown)
			// fmt.Println("borders left", bordersLeft)
			// fmt.Println("borders right", bordersRight)
			nbSides := 0
			for _, border := range bordersUp {
				slices.Sort(border)
				previousPos := border[0]
				nbSides++
				for i := 1; i <= len(border)-1; i++ {
					if border[i] != previousPos+1 {
						nbSides++
					}
					previousPos = border[i]
				}
			}

			for _, border := range bordersDown {
				slices.Sort(border)
				previousPos := border[0]
				nbSides++
				for i := 1; i <= len(border)-1; i++ {
					if border[i] != previousPos+1 {
						nbSides++
					}
					previousPos = border[i]
				}
			}

			for _, border := range bordersRight {
				slices.Sort(border)
				previousPos := border[0]
				nbSides++
				for i := 1; i <= len(border)-1; i++ {
					if border[i] != previousPos+1 {
						nbSides++
					}
					previousPos = border[i]
				}
			}

			for _, border := range bordersLeft {
				slices.Sort(border)
				previousPos := border[0]
				nbSides++
				for i := 1; i <= len(border)-1; i++ {
					if border[i] != previousPos+1 {
						nbSides++
					}
					previousPos = border[i]
				}
			}

			// fmt.Println("nbSides", nbSides)
			//
			// fmt.Println("nbBordersInZone", nbBordersInZone)
			// fmt.Println("nbNodesInZone", nbNodesInZone)

			result += nbSides * nbNodesInZone

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

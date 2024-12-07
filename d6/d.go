package d

import (
	"bufio"
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

type point struct {
	x, y int
}
type state struct {
	grid                     [][]string
	pos                      point
	direction                int // 8 = up, 6 = right, 2 = down, 4 = left
	visited                  []point
	newObstaclePosition      point
	obstaclesMetByDirections map[int][]point
}

var nextDirection = map[int]int{
	8: 6,
	6: 2,
	2: 4,
	4: 8,
}

func (s *state) getNextPosition() (point, bool) {
	nextPosition := s.pos
	switch s.direction {
	case 8:
		nextPosition.x -= 1
	case 6:
		nextPosition.y += 1
	case 2:
		nextPosition.x += 1
	case 4:
		nextPosition.y -= 1
	}

	obstacle := false
	if nextPosition.x < 0 || nextPosition.x >= len(s.grid) || nextPosition.y < 0 || nextPosition.y >= len(s.grid[0]) {
		return nextPosition, false
	}

	if s.grid[nextPosition.x][nextPosition.y] == "#" {
		obstacle = true
		obstaclePosition := point{nextPosition.x, nextPosition.y}

		for _, alreadyMet := range s.obstaclesMetByDirections[s.direction] {
			if alreadyMet == obstaclePosition {
				return nextPosition, true
			}
		}

		s.obstaclesMetByDirections[s.direction] = append(s.obstaclesMetByDirections[s.direction], obstaclePosition)

		newDirection := nextDirection[s.direction]
		s.direction = newDirection
	}

	if obstacle {
		nextPosition = s.pos
		switch s.direction {
		case 8:
			nextPosition.x -= 1
		case 6:
			nextPosition.y += 1
		case 2:
			nextPosition.x += 1
		case 4:
			nextPosition.y -= 1
		}
	}

	if nextPosition.x < 0 || nextPosition.x >= len(s.grid) || nextPosition.y < 0 || nextPosition.y >= len(s.grid[0]) {
		return nextPosition, false
	}

	if s.grid[nextPosition.x][nextPosition.y] == "#" {

		newDirection := nextDirection[s.direction]
		s.direction = newDirection

		nextPosition = s.pos
		switch s.direction {
		case 8:
			nextPosition.x -= 1
		case 6:
			nextPosition.y += 1
		case 2:
			nextPosition.x += 1
		case 4:
			nextPosition.y -= 1
		}
	}

	return nextPosition, false
}

func (d *D) Run1() int {
	fileScanner := bufio.NewScanner(d.inputStream)
	fileScanner.Split(bufio.ScanLines)

	grid := [][]string{}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		grid = append(grid, strings.Split(line, ""))
	}

	startingPosition := point{0, 0}

	for i, row := range grid {
		for j, cell := range row {
			if cell == "^" {
				startingPosition = point{i, j}
			}
		}
	}

	state := state{
		pos:                      startingPosition,
		direction:                8,
		visited:                  []point{},
		grid:                     grid,
		obstaclesMetByDirections: map[int][]point{},
	}

	for {
		nextPosition, _ := state.getNextPosition()

		state.visited = append(state.visited, state.pos)
		state.pos = nextPosition

		if nextPosition.x < 0 || nextPosition.x >= len(grid) || nextPosition.y < 0 || nextPosition.y >= len(grid[0]) {
			break
		}
	}

	for _, visit := range state.visited {
		grid[visit.x][visit.y] = "X"
	}

	uniqueVisited := map[point]bool{}
	for _, visit := range state.visited {
		uniqueVisited[visit] = true
	}

	result := len(uniqueVisited)

	return result
}

func (d *D) Run2() int {
	fileScanner := bufio.NewScanner(d.inputStream)
	fileScanner.Split(bufio.ScanLines)

	grid := [][]string{}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		grid = append(grid, strings.Split(line, ""))
	}

	startingPosition := point{0, 0}

	for i, row := range grid {
		for j, cell := range row {
			if cell == "^" {
				startingPosition = point{i, j}
			}
		}
	}

	// run simulations

	nbInfiniteLoops := 0
	for i, row := range grid {
		for j, cell := range row {
			if cell == "." {

				simulationGrid := make([][]string, len(grid))
				for i := range grid {
					simulationGrid[i] = make([]string, len(grid[i]))
					copy(simulationGrid[i], grid[i])
				}

				obstacleCell := "#"
				simulationGrid[i][j] = obstacleCell

				state := state{
					pos:       startingPosition,
					direction: 8,
					visited:   []point{},
					grid:      simulationGrid,
					obstaclesMetByDirections: map[int][]point{
						8: {},
						6: {},
						2: {},
						4: {},
					},
				}

				for {
					nextPosition, infiniteLoop := state.getNextPosition()

					state.visited = append(state.visited, state.pos)
					state.pos = nextPosition

					if nextPosition.x < 0 || nextPosition.x >= len(grid) || nextPosition.y < 0 || nextPosition.y >= len(grid[0]) {
						break
					}

					if infiniteLoop {
						nbInfiniteLoops++
						break
					}
				}
			}
		}
	}

	result := nbInfiniteLoops

	return result
}

func (d *D) RunStr1() string {
	return ""
}

func (d *D) RunStr2() string {
	return ""
}

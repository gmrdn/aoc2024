package d

import (
	"bufio"
	"io"
	"slices"
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
	x, y int
}

func (d *D) Run1() int {
	fileScanner := bufio.NewScanner(d.inputStream)
	fileScanner.Split(bufio.ScanLines)

	result := 0

	grid := make([][]rune, 0)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		grid = append(grid, []rune(line))
	}

	mapAntennaPositions := make(map[rune][]Point)
	for i, row := range grid {
		for j, r := range row {
			if r != '.' {
				mapAntennaPositions[r] = append(mapAntennaPositions[r], Point{i, j})
			}
		}
	}

	antinodeUniqueLocations := make(map[Point]bool)
	for _, allAntennasForThisFrequency := range mapAntennaPositions {
		for _, antenna := range allAntennasForThisFrequency {
			otherAntennas := slices.Clone(allAntennasForThisFrequency)
			for _, otherAntenna := range otherAntennas {
				if otherAntenna == antenna {
					continue
				}
				vector := Point{otherAntenna.x - antenna.x, otherAntenna.y - antenna.y}
				antionode := Point{antenna.x + 2*vector.x, antenna.y + 2*vector.y}

				if antionode.x < 0 || antionode.x >= len(grid) || antionode.y < 0 || antionode.y >= len(grid[0]) {
					continue
				}

				antinodeUniqueLocations[antionode] = true
			}
		}
	}

	nbAntinodes := len(antinodeUniqueLocations)
	result = nbAntinodes
	return result
}

func (d *D) Run2() int {
	fileScanner := bufio.NewScanner(d.inputStream)
	fileScanner.Split(bufio.ScanLines)

	result := 0

	grid := make([][]rune, 0)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		grid = append(grid, []rune(line))
	}

	mapAntennaPositions := make(map[rune][]Point)
	for i, row := range grid {
		for j, r := range row {
			if r != '.' {
				mapAntennaPositions[r] = append(mapAntennaPositions[r], Point{i, j})
			}
		}
	}

	antinodeUniqueLocations := make(map[Point]bool)
	for _, allAntennasForThisFrequency := range mapAntennaPositions {
		for _, antenna := range allAntennasForThisFrequency {
			otherAntennas := slices.Clone(allAntennasForThisFrequency)
			for _, otherAntenna := range otherAntennas {
				if otherAntenna == antenna {
					continue
				}
				vector := Point{otherAntenna.x - antenna.x, otherAntenna.y - antenna.y}
				antinodes := []Point{}
				for i := 1; i < 100; i++ {
					if antenna.x+i*vector.x < 0 || antenna.x+i*vector.x >= len(grid) || antenna.y+i*vector.y < 0 || antenna.y+i*vector.y >= len(grid[0]) {
						break
					}
					antinodes = append(antinodes, Point{antenna.x + i*vector.x, antenna.y + i*vector.y})
				}

				for _, antionode := range antinodes {
					antinodeUniqueLocations[antionode] = true
				}
			}
		}
	}

	nbAntinodes := len(antinodeUniqueLocations)
	result = nbAntinodes
	return result
}

func (d *D) RunStr1() string {
	return ""
}

func (d *D) RunStr2() string {
	return ""
}

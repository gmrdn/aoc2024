package d

import (
	"bufio"
	"io"
	"slices"
	"strconv"
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

func (d *D) Run1() int {
	fileScanner := bufio.NewScanner(d.inputStream)
	fileScanner.Split(bufio.ScanLines)

	input := ""
	for fileScanner.Scan() {
		input = fileScanner.Text()
	}

	freeSpace := false
	fragArray := make([]int, 0)
	id := 0
	for _, v := range input {

		nb, _ := strconv.Atoi(string(v))
		if !freeSpace {
			for j := 0; j < nb; j++ {
				fragArray = append(fragArray, id)
			}
			id++
		} else {
			for j := 0; j < nb; j++ {
				fragArray = append(fragArray, -1)
			}
		}
		freeSpace = !freeSpace
	}

	defragArray := make([]int, 0)
	for i := 0; i < len(fragArray); i++ {
		if fragArray[i] != -1 {
			defragArray = append(defragArray, fragArray[i])
		}
		if fragArray[i] == -1 {
			endValue := fragArray[len(fragArray)-1]
			defragArray = append(defragArray, endValue)

			for j := 0; j <= 9; j++ {
				fragArray = fragArray[:len(fragArray)-1]
				if fragArray[len(fragArray)-1] == -1 {
					fragArray = fragArray[:len(fragArray)-1]
				}
				if fragArray[len(fragArray)-1] != -1 {
					break
				}

			}
		}
	}

	result := 0
	for i := 0; i < len(defragArray); i++ {
		result += defragArray[i] * i
	}
	return result
}

func (d *D) Run2() int {
	fileScanner := bufio.NewScanner(d.inputStream)
	fileScanner.Split(bufio.ScanLines)

	input := ""
	for fileScanner.Scan() {
		input = fileScanner.Text()
	}

	freeSpace := false
	fragArray := make([]int, 0)
	id := 0
	for _, v := range input {

		nb, _ := strconv.Atoi(string(v))
		if !freeSpace {
			for j := 0; j < nb; j++ {
				fragArray = append(fragArray, id)
			}
			id++
		} else {
			for j := 0; j < nb; j++ {
				fragArray = append(fragArray, -1)
			}
		}
		freeSpace = !freeSpace
	}

	defragArray := make([]int, 0)

	fragmentsGroupsFromTheEnd := make([][]int, 0)
	currentValue := fragArray[len(fragArray)-1]
	currentGroup := make([]int, 0)

	for i := len(fragArray) - 1; i >= 0; i-- {
		if fragArray[i] == currentValue {
			currentGroup = append(currentGroup, fragArray[i])
		} else {

			fragmentsGroupsFromTheEnd = append(fragmentsGroupsFromTheEnd, currentGroup)

			if fragArray[i] == -1 {
				for j := 0; j < 10; j++ {
					if fragArray[i] == -1 {
						i--
					}
				}
			}

			currentValue = fragArray[i]
			currentGroup = make([]int, 0)
			currentGroup = append(currentGroup, fragArray[i])

		}
	}

	defragArray = slices.Clone(fragArray)

	for _, group := range fragmentsGroupsFromTheEnd {
		added := false
		for i := 0; i < len(defragArray); i++ {
			emptySpaceLength := 0
			for j := i; j < len(defragArray); j++ {
				if defragArray[j] == -1 {
					emptySpaceLength++
				} else {
					break
				}
			}

			if emptySpaceLength >= len(group) {
				for j := i; j < i+len(group); j++ {
					defragArray[j] = group[j-i]
				}
				added = true
				break
			}
		}
		if added {
			// remove from defragArray's end
			for i := len(defragArray) - 1; i >= 0; i-- {
				if defragArray[i] == group[0] {
					for j := 0; j < len(group); j++ {
						defragArray[i-j] = -1
					}
					break
				}
			}
		}

	}

	result := 0
	for i := 0; i < len(defragArray); i++ {
		if defragArray[i] == -1 {
			continue
		}
		result += defragArray[i] * i
	}

	return result
}

func (d *D) RunStr1() string {
	return ""
}

func (d *D) RunStr2() string {
	return ""
}

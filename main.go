package main

import (
	"bytes"
	"fmt"
	"os"

	"aoc2024/days"

	d1 "aoc2024/d1"
)

func main() {
	days := []days.Day{
		d1.NewD(),
	}

	for index, day := range days {
		input, err := os.ReadFile(fmt.Sprintf("input_d%d.txt", index+1))
		if err != nil {
			fmt.Println(err)
		}
		day.Input(bytes.NewBuffer(input))

		result := day.Run1()
		if result == 0 {
			fmt.Printf("Day %d - 1: %s\n", index+1, day.RunStr2())
		} else {
			fmt.Printf("Day %d - 1: %d\n", index+1, result)
		}

		day.Input(bytes.NewBuffer(input))
		result2 := day.Run2()
		if result == 0 {
			fmt.Printf("Day %d - 2: %s\n", index+1, day.RunStr2())
		} else {
			fmt.Printf("Day %d - 2: %d\n", index+1, result2)
		}
	}
}

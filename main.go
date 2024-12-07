package main

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"aoc2024/days"

	d1 "aoc2024/d1"
	d2 "aoc2024/d2"
	d3 "aoc2024/d3"
	d4 "aoc2024/d4"
	d5 "aoc2024/d5"
	d6 "aoc2024/d6"
	d7 "aoc2024/d7"
	d8 "aoc2024/d8"
)

func main() {
	days := []days.Day{
		d1.NewD(),
		d2.NewD(),
		d3.NewD(),
		d4.NewD(),
		d5.NewD(),
		d6.NewD(),
		d7.NewD(),
		d8.NewD(),
	}

	for index, day := range days {
		input, err := os.ReadFile(fmt.Sprintf("input_d%d.txt", index+1))
		if err != nil {
			fmt.Println(err)
		}
		day.Input(bytes.NewBuffer(input))

		startExec := time.Now()
		result := day.Run1()
		if result == 0 {
			strResult := day.RunStr1()
			fmt.Printf("Day %d - 1: %s (%v)\n", index+1, strResult, time.Since(startExec))
		} else {
			fmt.Printf("Day %d - 1: %d (%v)\n", index+1, result, time.Since(startExec))
		}

		startExec = time.Now()
		day.Input(bytes.NewBuffer(input))
		result2 := day.Run2()
		if result2 == 0 {
			strResult := day.RunStr2()
			fmt.Printf("Day %d - 2: %s (%v)\n", index+1, strResult, time.Since(startExec))
		} else {
			fmt.Printf("Day %d - 2: %d (%v)\n", index+1, result2, time.Since(startExec))
		}

	}
}

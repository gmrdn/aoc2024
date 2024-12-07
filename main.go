package main

import (
	"bytes"
	"fmt"
	"os"

	"aoc2024/days"

	d1 "aoc2024/d1"
	d2 "aoc2024/d2"
	d3 "aoc2024/d3"
	d4 "aoc2024/d4"
	d5 "aoc2024/d5"
	d6 "aoc2024/d6"
	d7 "aoc2024/d7"
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
	}

	for index, day := range days {
		input, err := os.ReadFile(fmt.Sprintf("input_d%d.txt", index+1))
		if err != nil {
			fmt.Println(err)
		}
		day.Input(bytes.NewBuffer(input))

		result := day.Run1()
		if result == 0 {
			strResult := day.RunStr1()
			if strResult == "" {
				bigIntResult := day.Run1BigInt()
				fmt.Printf("Day %d - 1: %s\n", index+1, bigIntResult.String())
			} else {
				fmt.Printf("Day %d - 1: %s\n", index+1, strResult)
			}
		} else {
			fmt.Printf("Day %d - 1: %d\n", index+1, result)
		}

		day.Input(bytes.NewBuffer(input))
		result2 := day.Run2()
		if result == 0 {
			strResult := day.RunStr2()
			if strResult == "" {
				bigIntResult := day.Run2BigInt()
				fmt.Printf("Day %d - 2: %s\n", index+1, bigIntResult.String())
			} else {
				fmt.Printf("Day %d - 2: %s\n", index+1, strResult)
			}
		} else {
			fmt.Printf("Day %d - 2: %d\n", index+1, result2)
		}

	}
}

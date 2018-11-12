package main

import (
	"fmt"
	"time"
)

func main() {
	// for keeping things easy to read
	type placeholder [5][3]byte

	// we're using arrays because:
	// - the pattern of a placeholder is constant (it doesn't change)
	// - the placeholders in the clock is constant (which is 8)
	//
	// so, whenever values are precomputed or constant, use an array.
	//
	// you can always convert it to a slice easily. you'll learn
	// how to do so in the slices section.
	digits := [10]placeholder{
		// 0
		{
			{1, 1, 1},
			{1, 0, 1},
			{1, 0, 1},
			{1, 0, 1},
			{1, 1, 1},
		},

		// 1
		{
			{1, 1, 0},
			{0, 1, 0},
			{0, 1, 0},
			{0, 1, 0},
			{1, 1, 1},
		},

		// 2
		{
			{1, 1, 1},
			{0, 0, 1},
			{1, 1, 1},
			{1, 0, 0},
			{1, 1, 1},
		},

		// 3
		{
			{1, 1, 1},
			{0, 0, 1},
			{1, 1, 1},
			{0, 0, 1},
			{1, 1, 1},
		},

		// 4
		{
			{1, 0, 1},
			{1, 0, 1},
			{1, 1, 1},
			{0, 0, 1},
			{0, 0, 1},
		},

		// 5
		{
			{1, 1, 1},
			{1, 0, 0},
			{1, 1, 1},
			{0, 0, 1},
			{1, 1, 1},
		},

		// 6
		{
			{1, 1, 1},
			{1, 0, 0},
			{1, 1, 1},
			{1, 0, 1},
			{1, 1, 1},
		},

		// 7
		{
			{1, 1, 1},
			{0, 0, 1},
			{0, 0, 1},
			{0, 0, 1},
			{0, 0, 1},
		},

		// 8
		{
			{1, 1, 1},
			{1, 0, 1},
			{1, 1, 1},
			{1, 0, 1},
			{1, 1, 1},
		},

		// 9
		{
			{1, 1, 1},
			{1, 0, 1},
			{1, 1, 1},
			{0, 0, 1},
			{1, 1, 1},
		},
	}

	colon := placeholder{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}

	for {
		// get the current time: hour, minute, and second
		var (
			now            = time.Now()
			hour, min, sec = now.Hour(), now.Minute(), now.Second()
		)

		// again, an array is used here.
		//
		// because, the number of placeholders in this clock
		// is constant (which is 8 placeholders).
		clock := [8]placeholder{
			// separate the digits: 17 becomes, 1 and 7 respectively
			digits[hour/10], digits[hour%10],
			colon,
			digits[min/10], digits[min%10],
			colon,
			digits[sec/10], digits[sec%10],
		}

		// clears the screen
		// only works on bash command prompts
		// \033 is a control code: [2J clears the screen
		fmt.Print("\033[H\033[2J")

		// print the clock
		for line := range clock[0] {
			// print a line for each letter
			for letter := range clock {
				// print a line
				for _, c := range clock[letter][line] {
					if c == 1 {
						fmt.Print("#")
					} else {
						fmt.Print(" ")
					}
				}

				fmt.Print("  ")
			}
			fmt.Println()
		}

		// wait for 1 second
		time.Sleep(time.Second)
	}
}

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

/*
--- Day 12: Rain Risk ---
Your ferry made decent progress toward the island, but the storm came in faster than anyone expected. The ferry needs to take evasive actions!

Unfortunately, the ship's navigation computer seems to be malfunctioning; rather than giving a route directly to safety, it produced extremely circuitous instructions. When the captain uses the PA system to ask if anyone can help, you quickly volunteer.

The navigation instructions (your puzzle input) consists of a sequence of single-character actions paired with integer input values. After staring at them for a few minutes, you work out what they probably mean:

Action N means to move north by the given value.
Action S means to move south by the given value.
Action E means to move east by the given value.
Action W means to move west by the given value.
Action L means to turn left the given number of degrees.
Action R means to turn right the given number of degrees.
Action F means to move forward by the given value in the direction the ship is currently facing.
The ship starts by facing east. Only the L and R actions change the direction the ship is facing. (That is, if the ship is facing east and the next instruction is N10, the ship would move north 10 units, but would still move east if the following action were F.)

For example:

F10
N3
F7
R90
F11
These instructions would be handled as follows:

F10 would move the ship 10 units east (because the ship starts by facing east) to east 10, north 0.
N3 would move the ship 3 units north to east 10, north 3.
F7 would move the ship another 7 units east (because the ship is still facing east) to east 17, north 3.
R90 would cause the ship to turn right by 90 degrees and face south; it remains at east 17, north 3.
F11 would move the ship 11 units south to east 17, south 8.
At the end of these instructions, the ship's Manhattan distance (sum of the absolute values of its east/west position and its north/south position) from its starting position is 17 + 8 = 25.

Figure out where the navigation instructions lead. What is the Manhattan distance between that location and the ship's starting position?

Your puzzle answer was 1106.

The first half of this puzzle is complete! It provides one gold star: *

--- Part Two ---
Before you can give the destination to the captain, you realize that the actual action meanings were printed on the back of the instructions the whole time.

Almost all of the actions indicate how to move a waypoint which is relative to the ship's position:

Action N means to move the waypoint north by the given value.
Action S means to move the waypoint south by the given value.
Action E means to move the waypoint east by the given value.
Action W means to move the waypoint west by the given value.
Action L means to rotate the waypoint around the ship left (counter-clockwise) the given number of degrees.
Action R means to rotate the waypoint around the ship right (clockwise) the given number of degrees.
Action F means to move forward to the waypoint a number of times equal to the given value.
The waypoint starts 10 units east and 1 unit north relative to the ship. The waypoint is relative to the ship; that is, if the ship moves, the waypoint moves with it.

For example, using the same instructions as above:

F10 moves the ship to the waypoint 10 times (a total of 100 units east and 10 units north), leaving the ship at east 100, north 10. The waypoint stays 10 units east and 1 unit north of the ship.
N3 moves the waypoint 3 units north to 10 units east and 4 units north of the ship. The ship remains at east 100, north 10.
F7 moves the ship to the waypoint 7 times (a total of 70 units east and 28 units north), leaving the ship at east 170, north 38. The waypoint stays 10 units east and 4 units north of the ship.
R90 rotates the waypoint around the ship clockwise 90 degrees, moving it to 4 units east and 10 units south of the ship. The ship remains at east 170, north 38.
F11 moves the ship to the waypoint 11 times (a total of 44 units east and 110 units south), leaving the ship at east 214, south 72. The waypoint stays 4 units east and 10 units south of the ship.
After these operations, the ship's Manhattan distance from its starting position is 214 + 72 = 286.

Figure out where the navigation instructions actually lead. What is the Manhattan distance between that location and the ship's starting position?
*/

func day12() {
	file, err := os.Open("./app/aoc/2020/day12_input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	m := map[rune][2]int{
		'N': {0, 1},
		'S': {0, -1},
		'E': {1, 0},
		'W': {-1, 0},
	}

	d := map[int]rune{
		0: 'E',
		1: 'N',
		2: 'W',
		3: 'S',
	}

	currentLoc := [2]int{0, 0}
	currentDirPos := 0
	for scanner.Scan() {
		text := scanner.Text()
		dir := text[0]
		val, _ := strconv.Atoi(text[1:])

		if nextDir, ok := m[rune(dir)]; ok {
			if nextDir[0] != 0 {
				currentLoc[0] += nextDir[0] * val
			}
			if nextDir[1] != 0 {
				currentLoc[1] += nextDir[1] * val
			}
		} else {
			switch dir {
			case 'F':
				{
					if pos, ok := m[d[currentDirPos]]; ok {
						if pos[0] != 0 {
							currentLoc[0] += pos[0] * val
						}
						if pos[1] != 0 {
							currentLoc[1] += pos[1] * val
						}
					}
				}
			case 'L':
				{
					currentDirPos = (4 + currentDirPos + val/90) % 4
				}
			case 'R':
				{
					currentDirPos = (4 + currentDirPos - val/90) % 4
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%f", math.Abs(float64(currentLoc[0]))+math.Abs(float64(currentLoc[1]))) // 18272118
	}
}

func day12b() {
	file, err := os.Open("./app/aoc/2020/day12_input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	m := map[rune][2]int{
		'N': {0, 1},
		'S': {0, -1},
		'E': {1, 0},
		'W': {-1, 0},
	}

	currentLoc := [2]int{0, 0}
	waypoint := [2]int{10, 1}
	for scanner.Scan() {
		text := scanner.Text()
		dir := text[0]
		val, _ := strconv.Atoi(text[1:])

		if nextDir, ok := m[rune(dir)]; ok {
			if nextDir[0] != 0 {
				waypoint[0] += nextDir[0] * val
			}
			if nextDir[1] != 0 {
				waypoint[1] += nextDir[1] * val
			}
		} else {
			switch dir {
			case 'F':
				{
					currentLoc[0] += val * waypoint[0]
					currentLoc[1] += val * waypoint[1]
				}
			case 'R':
				{
					x, y := waypoint[0], waypoint[1]
					switch val {
					case 90:
						{
							waypoint[0] = y
							waypoint[1] = -1 * x
						}
					case 180:
						{
							waypoint[0] = -1 * x
							waypoint[1] = -1 * y
						}
					case 270:
						{
							waypoint[0] = -1 * y
							waypoint[1] = x
						}
					}
				}
			case 'L':
				{
					x, y := waypoint[0], waypoint[1]
					switch val {
					case 90:
						{
							waypoint[0] = -1 * y
							waypoint[1] = x
						}
					case 180:
						{
							waypoint[0] = -1 * x
							waypoint[1] = -1 * y
						}
					case 270:
						{
							waypoint[0] = y
							waypoint[1] = -1 * x
						}
					}
				}
			}
		}

		fmt.Printf("%s, Loc: [%d, %d] - waypoint: [%d, %d]\n", text, currentLoc[0], currentLoc[1], waypoint[0], waypoint[1])
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%f", math.Abs(float64(currentLoc[0]))+math.Abs(float64(currentLoc[1]))) // 18272118
	}
}

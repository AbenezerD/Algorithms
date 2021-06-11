package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
--- Day 17: Conway Cubes ---
As your flight slowly drifts through the sky, the Elves at the Mythical Information Bureau at the North Pole contact you. They'd like some help debugging a malfunctioning experimental energy source aboard one of their super-secret imaging satellites.

The experimental energy source is based on cutting-edge technology: a set of Conway Cubes contained in a pocket dimension! When you hear it's having problems, you can't help but agree to take a look.

The pocket dimension contains an infinite 3-dimensional grid. At every integer 3-dimensional coordinate (x,y,z), there exists a single cube which is either active or inactive.

In the initial state of the pocket dimension, almost all cubes start inactive. The only exception to this is a small flat region of cubes (your puzzle input); the cubes in this region start in the specified active (#) or inactive (.) state.

The energy source then proceeds to boot up by executing six cycles.

Each cube only ever considers its neighbors: any of the 26 other cubes where any of their coordinates differ by at most 1. For example, given the cube at x=1,y=2,z=3, its neighbors include the cube at x=2,y=2,z=2, the cube at x=0,y=2,z=3, and so on.

During a cycle, all cubes simultaneously change their state according to the following rules:

If a cube is active and exactly 2 or 3 of its neighbors are also active, the cube remains active. Otherwise, the cube becomes inactive.
If a cube is inactive but exactly 3 of its neighbors are active, the cube becomes active. Otherwise, the cube remains inactive.
The engineers responsible for this experimental energy source would like you to simulate the pocket dimension and determine what the configuration of cubes should be at the end of the six-cycle boot process.

For example, consider the following initial state:

.#.
..#
###
Even though the pocket dimension is 3-dimensional, this initial state represents a small 2-dimensional slice of it. (In particular, this initial state defines a 3x3x1 region of the 3-dimensional space.)

Simulating a few cycles from this initial state produces the following configurations, where the result of each cycle is shown layer-by-layer at each given z coordinate (and the frame of view follows the active cells in each cycle):

Before any cycles:

z=0
.#.
..#
###


After 1 cycle:

z=-1
#..
..#
.#.

z=0
#.#
.##
.#.

z=1
#..
..#
.#.


After 2 cycles:

z=-2
.....
.....
..#..
.....
.....

z=-1
..#..
.#..#
....#
.#...
.....

z=0
##...
##...
#....
....#
.###.

z=1
..#..
.#..#
....#
.#...
.....

z=2
.....
.....
..#..
.....
.....


After 3 cycles:

z=-2
.......
.......
..##...
..###..
.......
.......
.......

z=-1
..#....
...#...
#......
.....##
.#...#.
..#.#..
...#...

z=0
...#...
.......
#......
.......
.....##
.##.#..
...#...

z=1
..#....
...#...
#......
.....##
.#...#.
..#.#..
...#...

z=2
.......
.......
..##...
..###..
.......
.......
.......
After the full six-cycle boot process completes, 112 cubes are left in the active state.

Starting with your given initial configuration, simulate six cycles. How many cubes are left in the active state after the sixth cycle?

Your puzzle answer was 324.

The first half of this puzzle is complete! It provides one gold star: *

--- Part Two ---
For some reason, your simulated results don't match what the experimental energy source engineers expected. Apparently, the pocket dimension actually has four spatial dimensions, not three.

The pocket dimension contains an infinite 4-dimensional grid. At every integer 4-dimensional coordinate (x,y,z,w), there exists a single cube (really, a hypercube) which is still either active or inactive.

Each cube only ever considers its neighbors: any of the 80 other cubes where any of their coordinates differ by at most 1. For example, given the cube at x=1,y=2,z=3,w=4, its neighbors include the cube at x=2,y=2,z=3,w=3, the cube at x=0,y=2,z=3,w=4, and so on.

The initial state of the pocket dimension still consists of a small flat region of cubes. Furthermore, the same rules for cycle updating still apply: during each cycle, consider the number of active neighbors of each cube.

For example, consider the same initial state as in the example above. Even though the pocket dimension is 4-dimensional, this initial state represents a small 2-dimensional slice of it. (In particular, this initial state defines a 3x3x1x1 region of the 4-dimensional space.)

Simulating a few cycles from this initial state produces the following configurations, where the result of each cycle is shown layer-by-layer at each given z and w coordinate:

Before any cycles:

z=0, w=0
.#.
..#
###


After 1 cycle:

z=-1, w=-1
#..
..#
.#.

z=0, w=-1
#..
..#
.#.

z=1, w=-1
#..
..#
.#.

z=-1, w=0
#..
..#
.#.

z=0, w=0
#.#
.##
.#.

z=1, w=0
#..
..#
.#.

z=-1, w=1
#..
..#
.#.

z=0, w=1
#..
..#
.#.

z=1, w=1
#..
..#
.#.


After 2 cycles:

z=-2, w=-2
.....
.....
..#..
.....
.....

z=-1, w=-2
.....
.....
.....
.....
.....

z=0, w=-2
###..
##.##
#...#
.#..#
.###.

z=1, w=-2
.....
.....
.....
.....
.....

z=2, w=-2
.....
.....
..#..
.....
.....

z=-2, w=-1
.....
.....
.....
.....
.....

z=-1, w=-1
.....
.....
.....
.....
.....

z=0, w=-1
.....
.....
.....
.....
.....

z=1, w=-1
.....
.....
.....
.....
.....

z=2, w=-1
.....
.....
.....
.....
.....

z=-2, w=0
###..
##.##
#...#
.#..#
.###.

z=-1, w=0
.....
.....
.....
.....
.....

z=0, w=0
.....
.....
.....
.....
.....

z=1, w=0
.....
.....
.....
.....
.....

z=2, w=0
###..
##.##
#...#
.#..#
.###.

z=-2, w=1
.....
.....
.....
.....
.....

z=-1, w=1
.....
.....
.....
.....
.....

z=0, w=1
.....
.....
.....
.....
.....

z=1, w=1
.....
.....
.....
.....
.....

z=2, w=1
.....
.....
.....
.....
.....

z=-2, w=2
.....
.....
..#..
.....
.....

z=-1, w=2
.....
.....
.....
.....
.....

z=0, w=2
###..
##.##
#...#
.#..#
.###.

z=1, w=2
.....
.....
.....
.....
.....

z=2, w=2
.....
.....
..#..
.....
.....
After the full six-cycle boot process completes, 848 cubes are left in the active state.

Starting with your given initial configuration, simulate six cycles in a 4-dimensional space. How many cubes are left in the active state after the sixth cycle?
*/

func day17() {
	file, err := os.Open("./app/aoc/2020/day17_input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var arr []string
	for scanner.Scan() {
		arr = append(arr, scanner.Text())
	}

	prevCount := 0
	arrayOfArrOfArr := [][][]string{{arr}}
	for i := 0; i < 6; i++ {
		var newCount int
		arrayOfArrOfArr, newCount = replace17b(arrayOfArrOfArr)
		prevCount = newCount

		fmt.Printf("RUN #%d\n", i+1)
		for j, arr := range arrayOfArrOfArr {
			fmt.Printf("z = %d,", j-len(arrayOfArrOfArr)/2)
			for k, a := range arr {
				fmt.Printf("z = %d, w = %d\n", j-len(arrayOfArrOfArr)/2, k-len(arrayOfArrOfArr)/2)
				for _, s := range a {
					fmt.Println(s)
				}
			}
		}
	}

	//for _, s := range arr {
	//	fmt.Println(s)
	//}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%d", prevCount) // 18272118
	}
}

func replace17(arrOfArr [][]string) ([][]string, int) {
	occupiedCount := 0
	adjacents := [][3]int{
		// x = -1
		{-1, -1, -1}, {-1, -1, 0}, {-1, -1, 1},
		{-1, 0, -1}, {-1, 0, 0}, {-1, 0, 1},
		{-1, 1, -1}, {-1, 1, 0}, {-1, 1, 1},
		// x = 0
		{0, -1, -1}, {0, -1, 0}, {0, -1, 1},
		{0, 0, -1}, {0, 0, 1},
		{0, 1, -1}, {0, 1, 0}, {0, 1, 1},
		// x = 1
		{1, -1, -1}, {1, -1, 0}, {1, -1, 1},
		{1, 0, -1}, {1, 0, 0}, {1, 0, 1},
		{1, 1, -1}, {1, 1, 0}, {1, 1, 1},
	}
	inactiveRow := make([]rune, len(arrOfArr[0][0]))
	for i := range inactiveRow {
		inactiveRow[i] = '.'
	}
	inactiveArr := make([]string, len(arrOfArr[0]))
	for i := range inactiveArr {
		inactiveArr[i] = string(inactiveRow)
	}

	arrOfArr = append([][]string{inactiveArr}, arrOfArr...)
	arrOfArr = append(arrOfArr, inactiveArr)

	for i := range arrOfArr {
		newArr := make([]string, len(arrOfArr[i])+2)
		for j := 1; j < len(newArr)-1; j++ {
			newArr[j] = "." + arrOfArr[i][j-1] + "."
		}
		newArr[0] = string(inactiveRow) + ".."
		newArr[len(newArr)-1] = string(inactiveRow) + ".."
		arrOfArr[i] = newArr
	}
	newArrOfArr := make([][]string, len(arrOfArr))
	for i := 0; i < len(arrOfArr); i++ {
		arr := arrOfArr[i]
		newArr := make([]string, len(arr))
		for j := 0; j < len(arr); j++ {
			row := arr[j]
			newRow := make([]rune, len(newArr))
			for k := 0; k < len(row); k++ {
				switch row[k] {
				case '.':
					{
						count := 0
						for _, seat := range adjacents {
							x := seat[0] + i
							y := seat[1] + j
							z := seat[2] + k
							if x < len(arrOfArr) && x >= 0 &&
								y < len(arr) && y >= 0 &&
								z < len(row) && z >= 0 && arrOfArr[x][y][z] == '#' {
								count++
							}

							if count > 3 {
								break
							}
						}
						if count == 3 {
							newRow[k] = '#'
						} else {
							newRow[k] = '.'
						}
					}
				case '#':
					{
						count := 0
						for _, seat := range adjacents {
							x := seat[0] + i
							y := seat[1] + j
							z := seat[2] + k
							if x < len(arrOfArr) && x >= 0 &&
								y < len(arr) && y >= 0 &&
								z < len(row) && z >= 0 && arrOfArr[x][y][z] == '#' {
								count++
							}

							if count > 3 {
								break
							}
						}
						if count == 2 || count == 3 {
							newRow[k] = '#'
						} else {
							newRow[k] = '.'
						}
					}
				}
				if newRow[k] == '#' {
					occupiedCount++
				}
			}
			newArr[j] = string(newRow)
		}
		newArrOfArr[i] = newArr
	}

	return newArrOfArr, occupiedCount
}

func replace17b(arrOfArrOfArr [][][]string) ([][][]string, int) {
	occupiedCount := 0
	adjacents := [][3]int{
		// x = -1
		{-1, -1, -1}, {-1, -1, 0}, {-1, -1, 1},
		{-1, 0, -1}, {-1, 0, 0}, {-1, 0, 1},
		{-1, 1, -1}, {-1, 1, 0}, {-1, 1, 1},
		// x = 0
		{0, -1, -1}, {0, -1, 0}, {0, -1, 1},
		{0, 0, -1}, {0, 0, 1}, {0, 0, 0},
		{0, 1, -1}, {0, 1, 0}, {0, 1, 1},
		// x = 1
		{1, -1, -1}, {1, -1, 0}, {1, -1, 1},
		{1, 0, -1}, {1, 0, 0}, {1, 0, 1},
		{1, 1, -1}, {1, 1, 0}, {1, 1, 1},
	}
	fourDimensionalAdjacents := make([][4]int, 0)
	for i := -1; i < 2; i++ {
		for _, adjacent := range adjacents {
			if i != 0 || adjacent[0] != 0 || adjacent[1] != 0 || adjacent[2] != 0 {
				fourDimensionalAdjacents = append(fourDimensionalAdjacents, [4]int{i, adjacent[0], adjacent[1], adjacent[2]})
			}
		}
	}
	inactiveRow := make([]rune, len(arrOfArrOfArr[0][0][0]))
	for i := range inactiveRow {
		inactiveRow[i] = '.'
	}
	inactiveArr := make([]string, len(arrOfArrOfArr[0][0]))
	for i := range inactiveArr {
		inactiveArr[i] = string(inactiveRow) + ".."
	}
	inactiveArr = append(inactiveArr, string(inactiveRow)+"..")
	inactiveArr = append(inactiveArr, string(inactiveRow)+"..")
	inactiveArrOfArr := make([][]string, len(arrOfArrOfArr[0]))
	for i := range inactiveArrOfArr {
		inactiveArrOfArr[i] = inactiveArr
	}
	inactiveArrOfArr = append(inactiveArrOfArr, inactiveArr)
	inactiveArrOfArr = append(inactiveArrOfArr, inactiveArr)

	for i := range arrOfArrOfArr {
		arrOfArr := arrOfArrOfArr[i]
		newArrOfArr := make([][]string, len(arrOfArr)+2)
		for j := 1; j < len(newArrOfArr)-1; j++ {
			newArr := make([]string, len(arrOfArr[j-1])+2)
			for k := 1; k < len(newArr)-1; k++ {
				newArr[k] = "." + arrOfArr[j-1][k-1] + "."
			}
			newArr[0] = string(inactiveRow) + ".."
			newArr[len(newArr)-1] = string(inactiveRow) + ".."
			newArrOfArr[j] = newArr
		}
		newArrOfArr[0] = inactiveArr
		newArrOfArr[len(newArrOfArr)-1] = inactiveArr
		arrOfArrOfArr[i] = newArrOfArr
	}
	arrOfArrOfArr = append([][][]string{inactiveArrOfArr}, arrOfArrOfArr...)
	arrOfArrOfArr = append(arrOfArrOfArr, inactiveArrOfArr)

	newArrOfArrOfArr := make([][][]string, len(arrOfArrOfArr))
	for i := 0; i < len(arrOfArrOfArr); i++ {
		arrOfArr := arrOfArrOfArr[i]
		newArrOfArr := make([][]string, len(arrOfArr))
		for j := 0; j < len(arrOfArr); j++ {
			arr := arrOfArr[j]
			newArr := make([]string, len(arr))
			for k := 0; k < len(arr); k++ {
				row := arr[k]
				newRow := make([]rune, len(row))
				for l := 0; l < len(row); l++ {
					switch row[l] {
					case '.':
						{
							count := 0
							for _, seat := range fourDimensionalAdjacents {
								x := seat[0] + i
								y := seat[1] + j
								z := seat[2] + k
								w := seat[3] + l
								if x < len(arrOfArrOfArr) && x >= 0 &&
									y < len(arrOfArr) && y >= 0 &&
									z < len(arr) && z >= 0 &&
									w < len(row) && w >= 0 && arrOfArrOfArr[x][y][z][w] == '#' {
									count++
								}

								if count > 3 {
									break
								}
							}
							if count == 3 {
								newRow[l] = '#'
							} else {
								newRow[l] = '.'
							}
						}
					case '#':
						{
							count := 0
							for _, seat := range fourDimensionalAdjacents {
								x := seat[0] + i
								y := seat[1] + j
								z := seat[2] + k
								w := seat[3] + l
								if x < len(arrOfArrOfArr) && x >= 0 &&
									y < len(arrOfArr) && y >= 0 &&
									z < len(arr) && z >= 0 &&
									w < len(row) && w >= 0 && arrOfArrOfArr[x][y][z][w] == '#' {
									count++
								}

								if count > 3 {
									break
								}
							}
							if count == 2 || count == 3 {
								newRow[l] = '#'
							} else {
								newRow[l] = '.'
							}
						}
					}
					if newRow[l] == '#' {
						occupiedCount++
					}
				}
				newArr[k] = string(newRow)
			}
			newArrOfArr[j] = newArr
		}
		newArrOfArrOfArr[i] = newArrOfArr
	}
	return newArrOfArrOfArr, occupiedCount
}

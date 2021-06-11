package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
--- Day 3: Toboggan Trajectory ---
With the toboggan login problems resolved, you set off toward the airport. While travel by toboggan might be easy, it's certainly not safe: there's very minimal steering and the area is covered in trees. You'll need to see which angles will take you near the fewest trees.

Due to the local geology, trees in this area only grow on exact integer coordinates in a grid. You make a map (your puzzle input) of the open squares (.) and trees (#) you can see. For example:

..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#
These aren't the only trees, though; due to something you read about once involving arboreal genetics and biome stability, the same pattern repeats to the right many times:

..##.........##.........##.........##.........##.........##.......  --->
#...#...#..#...#...#..#...#...#..#...#...#..#...#...#..#...#...#..
.#....#..#..#....#..#..#....#..#..#....#..#..#....#..#..#....#..#.
..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#
.#...##..#..#...##..#..#...##..#..#...##..#..#...##..#..#...##..#.
..#.##.......#.##.......#.##.......#.##.......#.##.......#.##.....  --->
.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#
.#........#.#........#.#........#.#........#.#........#.#........#
#.##...#...#.##...#...#.##...#...#.##...#...#.##...#...#.##...#...
#...##....##...##....##...##....##...##....##...##....##...##....#
.#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#  --->
You start on the open square (.) in the top-left corner and need to reach the bottom (below the bottom-most row on your map).

The toboggan can only follow a few specific slopes (you opted for a cheaper model that prefers rational numbers); start by counting all the trees you would encounter for the slope right 3, down 1:

From your starting position at the top-left, check the position that is right 3 and down 1. Then, check the position that is right 3 and down 1 from there, and so on until you go past the bottom of the map.

The locations you'd check in the above example are marked here with O where there was an open square and X where there was a tree:

..##.........##.........##.........##.........##.........##.......  --->
#..O#...#..#...#...#..#...#...#..#...#...#..#...#...#..#...#...#..
.#....X..#..#....#..#..#....#..#..#....#..#..#....#..#..#....#..#.
..#.#...#O#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#
.#...##..#..X...##..#..#...##..#..#...##..#..#...##..#..#...##..#.
..#.##.......#.X#.......#.##.......#.##.......#.##.......#.##.....  --->
.#.#.#....#.#.#.#.O..#.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#
.#........#.#........X.#........#.#........#.#........#.#........#
#.##...#...#.##...#...#.X#...#...#.##...#...#.##...#...#.##...#...
#...##....##...##....##...#X....##...##....##...##....##...##....#
.#..#...#.#.#..#...#.#.#..#...X.#.#..#...#.#.#..#...#.#.#..#...#.#  --->
In this example, traversing the map using this slope would cause you to encounter 7 trees.

Starting at the top-left corner of your map and following a slope of right 3 and down 1, how many trees would you encounter?

Your puzzle answer was 247.

--- Part Two ---
Time to check the rest of the slopes - you need to minimize the probability of a sudden arboreal stop, after all.

Determine the number of trees you would encounter if, for each of the following slopes, you start at the top-left corner and traverse the map all the way to the bottom:

Right 1, down 1.
Right 3, down 1. (This is the slope you already checked.)
Right 5, down 1.
Right 7, down 1.
Right 1, down 2.
In the above example, these slopes would find 2, 7, 3, 4, and 2 tree(s) respectively; multiplied together, these produce the answer 336.

What do you get if you multiply together the number of trees encountered on each of the listed slopes?

Your puzzle answer was 2983070376.
*/

func trees(num int) {
	file, err := os.Open("./app/aoc/2020/day3_input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	treeCount := 0
	index := 0
	for ; scanner.Scan(); index++ {
		text := scanner.Text()
		n := len(text)
		if index == 0 {
			fmt.Println(text)
			continue
		}
		replace := "O"
		currentTreeLoc := index * num
		u := text[currentTreeLoc%n]
		if u == '#' {
			treeCount++
			replace = "X"
		}

		for i := 0; i < currentTreeLoc; i += n {
			if currentTreeLoc < (i + n) {
				text = text[:(currentTreeLoc%n)] + replace + text[(currentTreeLoc%n):]
			}
			fmt.Print(text)
		}
		fmt.Print("\n")

	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(treeCount)
	}
}

func trees2() {
	file, err := os.Open("./app/aoc/2020/day3_input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	treeCount1 := 0  // 1 right;  1 down
	treeCount3 := 0  // 3 right;  1 down
	treeCount5 := 0  // 5 right;  1 down
	treeCount7 := 0  // 7 right;  1 down
	treeCount21 := 0 // 1 right;  2 down
	index := 0
	for ; scanner.Scan(); index++ {
		text := scanner.Text()
		n := len(text)
		if index == 0 {
			continue
		}
		if text[index%n] == '#' {
			treeCount1++
		}
		if text[(index*3)%n] == '#' {
			treeCount3++
		}
		if text[(index*5)%n] == '#' {
			treeCount5++
		}
		if text[(index*7)%n] == '#' {
			treeCount7++
		}
		if index%2 == 0 && text[(index/2)%n] == '#' {
			treeCount21++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%d * %d * %d * %d * %d = %d",
			treeCount1, treeCount3, treeCount5, treeCount7, treeCount21,
			treeCount1*treeCount3*treeCount5*treeCount7*treeCount21,
		)
	}
}

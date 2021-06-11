package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
--- Day 11: Seating System ---
Your plane lands with plenty of time to spare. The final leg of your journey is a ferry that goes directly to the tropical island where you can finally start your vacation. As you reach the waiting area to board the ferry, you realize you're so early, nobody else has even arrived yet!

By modeling the process people use to choose (or abandon) their seat in the waiting area, you're pretty sure you can predict the best place to sit. You make a quick map of the seat layout (your puzzle input).

The seat layout fits neatly on a grid. Each position is either floor (.), an empty seat (L), or an occupied seat (#). For example, the initial seat layout might look like this:

L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL
Now, you just need to model the people who will be arriving shortly. Fortunately, people are entirely predictable and always follow a simple set of rules. All decisions are based on the number of occupied seats adjacent to a given seat (one of the eight positions immediately up, down, left, right, or diagonal from the seat). The following rules are applied to every seat simultaneously:

If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
Otherwise, the seat's state does not change.
Floor (.) never changes; seats don't move, and nobody sits on the floor.

After one round of these rules, every seat in the example layout becomes occupied:

#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##
After a second round, the seats with four or more occupied adjacent seats become empty again:

#.LL.L#.##
#LLLLLL.L#
L.L.L..L..
#LLL.LL.L#
#.LL.LL.LL
#.LLLL#.##
..L.L.....
#LLLLLLLL#
#.LLLLLL.L
#.#LLLL.##
This process continues for three more rounds:

#.##.L#.##
#L###LL.L#
L.#.#..#..
#L##.##.L#
#.##.LL.LL
#.###L#.##
..#.#.....
#L######L#
#.LL###L.L
#.#L###.##
#.#L.L#.##
#LLL#LL.L#
L.L.L..#..
#LLL.##.L#
#.LL.LL.LL
#.LL#L#.##
..L.L.....
#L#LLLL#L#
#.LLLLLL.L
#.#L#L#.##
#.#L.L#.##
#LLL#LL.L#
L.#.L..#..
#L##.##.L#
#.#L.LL.LL
#.#L#L#.##
..L.L.....
#L#L##L#L#
#.LLLLLL.L
#.#L#L#.##
At this point, something interesting happens: the chaos stabilizes and further applications of these rules cause no seats to change state! Once people stop moving around, you count 37 occupied seats.

Simulate your seating area by applying the seating rules repeatedly until no seats change state. How many seats end up occupied?

Your puzzle answer was 2275.

The first half of this puzzle is complete! It provides one gold star: *

--- Part Two ---
As soon as people start to arrive, you realize your mistake. People don't just care about adjacent seats - they care about the first seat they can see in each of those eight directions!

Now, instead of considering just the eight immediately adjacent seats, consider the first seat in each of those eight directions. For example, the empty seat below would see eight occupied seats:

.......#.
...#.....
.#.......
.........
..#L....#
....#....
.........
#........
...#.....
The leftmost empty seat below would only see one empty seat, but cannot see any of the occupied ones:

.............
.L.L.#.#.#.#.
.............
The empty seat below would see no occupied seats:

.##.##.
#.#.#.#
##...##
...L...
##...##
#.#.#.#
.##.##.
Also, people seem to be more tolerant than you expected: it now takes five or more visible occupied seats for an occupied seat to become empty (rather than four or more from the previous rules). The other rules still apply: empty seats that see no occupied seats become occupied, seats matching no rule don't change, and floor never changes.

Given the same starting layout as above, these new rules cause the seating area to shift around as follows:

L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL
#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##
#.LL.LL.L#
#LLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLL#
#.LLLLLL.L
#.LLLLL.L#
#.L#.##.L#
#L#####.LL
L.#.#..#..
##L#.##.##
#.##.#L.##
#.#####.#L
..#.#.....
LLL####LL#
#.L#####.L
#.L####.L#
#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##LL.LL.L#
L.LL.LL.L#
#.LLLLL.LL
..L.L.....
LLLLLLLLL#
#.LLLLL#.L
#.L#LL#.L#
#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##L#.#L.L#
L.L#.#L.L#
#.L####.LL
..#.#.....
LLL###LLL#
#.LLLLL#.L
#.L#LL#.L#
#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##L#.#L.L#
L.L#.LL.L#
#.LLLL#.LL
..#.L.....
LLL###LLL#
#.LLLLL#.L
#.L#LL#.L#
Again, at this point, people stop shifting around and the seating area reaches equilibrium. Once this occurs, you count 26 occupied seats.

Given the new visibility method and the rule change for occupied seats becoming empty, once equilibrium is reached, how many seats end up occupied?

*/

func day11() {
	file, err := os.Open("./app/aoc/2020/day11_input.txt")
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
	for {
		var newCount int
		arr, newCount = replace(arr)
		if newCount == prevCount {
			break
		}
		prevCount = newCount
		for _, s := range arr {
			fmt.Println(s)
		}
		fmt.Println("----------------")
	}

	for _, s := range arr {
		fmt.Println(s)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%d", prevCount) // 18272118
	}
}

func day11b() {
	file, err := os.Open("./app/aoc/2020/day11_input.txt")
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
	for {
		var newCount int
		arr, newCount = replaceB(arr)
		if newCount == prevCount {
			break
		}
		prevCount = newCount
		for _, s := range arr {
			fmt.Println(s)
		}
		fmt.Println("----------------")
	}

	for _, s := range arr {
		fmt.Println(s)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%d", prevCount) // 18272118
	}
}

func replace(arr []string) ([]string, int) {
	occupiedCount := 0
	newArr := make([]string, len(arr))
	for i := 0; i < len(arr); i++ {
		row := arr[i]
		newRow := make([]rune, len(row))
		for j := 0; j < len(row); j++ {
			switch row[j] {
			case '.':
				{
					newRow[j] = '.'
					continue
				}
			case 'L':
				{
					broke := false
					adjacentSeats := getAdjacentSeats(i, j, len(arr), len(row))
					for _, seat := range adjacentSeats {
						x := seat[0]
						y := seat[1]
						if arr[x][y] == '#' {
							broke = true
							break
						}
					}
					if broke {
						newRow[j] = 'L'
					} else {
						newRow[j] = '#'
					}
				}
			case '#':
				{
					count := 0
					adjacentSeats := getAdjacentSeats(i, j, len(arr), len(row))
					for _, seat := range adjacentSeats {
						x := seat[0]
						y := seat[1]
						if arr[x][y] == '#' {
							count++
						}
					}
					if count >= 4 {
						newRow[j] = 'L'
					} else {
						newRow[j] = '#'
					}
				}
			}
			if newRow[j] == '#' {
				occupiedCount++
			}
		}
		newArr[i] = string(newRow)
	}

	return newArr, occupiedCount
}

func replaceB(arr []string) ([]string, int) {
	occupiedCount := 0
	newArr := make([]string, len(arr))
	adjacents := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	for i := 0; i < len(arr); i++ {
		row := arr[i]
		newRow := make([]rune, len(row))
		for j := 0; j < len(row); j++ {
			switch row[j] {
			case '.':
				{
					newRow[j] = '.'
					continue
				}
			case 'L':
				{
					broke := false
					for _, seat := range adjacents {
						for x, y := i+seat[0], j+seat[1]; x < len(arr) && x >= 0 && y < len(row) && y >= 0; x, y = x+seat[0], y+seat[1] {
							if arr[x][y] == '.' {
								continue
							}
							if arr[x][y] == '#' {
								broke = true
							}
							break
						}
						if broke {
							break
						}
					}
					if broke {
						newRow[j] = 'L'
					} else {
						newRow[j] = '#'
					}
				}
			case '#':
				{
					count := 0
					for _, seat := range adjacents {
						x, y := i+seat[0], j+seat[1]
						for ; x < len(arr) && x >= 0 && y < len(row) && y >= 0; x, y = x+seat[0], y+seat[1] {
							if arr[x][y] == '.' {
								continue
							}
							if arr[x][y] == '#' {
								count++
							}
							break
						}
					}
					if count >= 5 {
						newRow[j] = 'L'
					} else {
						newRow[j] = '#'
					}
				}
			}
			if newRow[j] == '#' {
				occupiedCount++
			}
		}
		newArr[i] = string(newRow)
	}

	return newArr, occupiedCount
}

func getAdjacentSeats(i, j, lenN, lenM int) [][2]int {
	validIs := []int{i}
	validJs := []int{j}
	if i > 0 {
		validIs = append(validIs, i-1)
	}
	if i < lenN-1 {
		validIs = append(validIs, i+1)
	}
	if j > 0 {
		validJs = append(validJs, j-1)
	}
	if j < lenM-1 {
		validJs = append(validJs, j+1)
	}

	adjacent := make([][2]int, 0)
	for _, validI := range validIs {
		for _, validJ := range validJs {
			if validI != i || validJ != j {
				adjacent = append(adjacent, [2]int{validI, validJ})
			}
		}
	}

	return adjacent
}

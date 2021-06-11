package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

/*
--- Day 24: Lobby Layout ---
Your raft makes it to the tropical island; it turns out that the small crab was an excellent navigator. You make your way to the resort.

As you enter the lobby, you discover a small problem: the floor is being renovated. You can't even reach the check-in desk until they've finished installing the new tile floor.

The tiles are all hexagonal; they need to be arranged in a hex grid with a very specific color pattern. Not in the mood to wait, you offer to help figure out the pattern.

The tiles are all white on one side and black on the other. They start with the white side facing up. The lobby is large enough to fit whatever pattern might need to appear there.

A member of the renovation crew gives you a list of the tiles that need to be flipped over (your puzzle input). Each line in the list identifies a single tile that needs to be flipped by giving a series of steps starting from a reference tile in the very center of the room. (Every line starts from the same reference tile.)

Because the tiles are hexagonal, every tile has six neighbors: east, southeast, southwest, west, northwest, and northeast. These directions are given in your list, respectively, as e, se, sw, w, nw, and ne. A tile is identified by a series of these directions with no delimiters; for example, esenee identifies the tile you land on if you start at the reference tile and then move one tile east, one tile southeast, one tile northeast, and one tile east.

Each time a tile is identified, it flips from white to black or from black to white. Tiles might be flipped more than once. For example, a line like esew flips a tile immediately adjacent to the reference tile, and a line like nwwswee flips the reference tile itself.

Here is a larger example:

sesenwnenenewseeswwswswwnenewsewsw
neeenesenwnwwswnenewnwwsewnenwseswesw
seswneswswsenwwnwse
nwnwneseeswswnenewneswwnewseswneseene
swweswneswnenwsewnwneneseenw
eesenwseswswnenwswnwnwsewwnwsene
sewnenenenesenwsewnenwwwse
wenwwweseeeweswwwnwwe
wsweesenenewnwwnwsenewsenwwsesesenwne
neeswseenwwswnwswswnw
nenwswwsewswnenenewsenwsenwnesesenew
enewnwewneswsewnwswenweswnenwsenwsw
sweneswneswneneenwnewenewwneswswnese
swwesenesewenwneswnwwneseswwne
enesenwswwswneneswsenwnewswseenwsese
wnwnesenesenenwwnenwsewesewsesesew
nenewswnwewswnenesenwnesewesw
eneswnwswnwsenenwnwnwwseeswneewsenese
neswnwewnwnwseenwseesewsenwsweewe
wseweeenwnesenwwwswnew
In the above example, 10 tiles are flipped once (to black), and 5 more are flipped twice (to black, then back to white). After all of these instructions have been followed, a total of 10 tiles are black.

Go through the renovation crew's list and determine which tiles they need to flip. After all of the instructions have been followed, how many tiles are left with the black side up?

Your puzzle answer was 382.

The first half of this puzzle is complete! It provides one gold star: *

--- Part Two ---
The tile floor in the lobby is meant to be a living art exhibit. Every day, the tiles are all flipped according to the following rules:

Any black tile with zero or more than 2 black tiles immediately adjacent to it is flipped to white.
Any white tile with exactly 2 black tiles immediately adjacent to it is flipped to black.
Here, tiles immediately adjacent means the six tiles directly touching the tile in question.

The rules are applied simultaneously to every tile; put another way, it is first determined which tiles need to be flipped, then they are all flipped at the same time.

In the above example, the number of black tiles that are facing up after the given number of days has passed is as follows:

Day 1: 15
Day 2: 12
Day 3: 25
Day 4: 14
Day 5: 23
Day 6: 28
Day 7: 41
Day 8: 37
Day 9: 49
Day 10: 37

Day 20: 132
Day 30: 259
Day 40: 406
Day 50: 566
Day 60: 788
Day 70: 1106
Day 80: 1373
Day 90: 1844
Day 100: 2208
After executing this process a total of 100 times, there would be 2208 black tiles facing up.

How many tiles will be black after 100 days?
*/

type tile struct {
	x int
	y int
}

func day24() {
	file, err := os.Open("./app/aoc/2020/day24_input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	tileColour := make(map[tile]bool, 0)
	for scanner.Scan() {
		text := scanner.Text()

		direction := ""
		tile := tile{
			x: 0,
			y: 0,
		}
		for _, dir := range text {
			direction += string(dir)
			switch direction {
			case "e":
				{
					tile.x += 2
					direction = ""
				}
			case "w":
				{
					tile.x -= 2
					direction = ""
				}
			case "se":
				{
					tile.x += 1
					tile.y -= 1
					direction = ""
				}
			case "sw":
				{
					tile.x -= 1
					tile.y -= 1
					direction = ""
				}
			case "ne":
				{
					tile.x += 1
					tile.y += 1
					direction = ""
				}
			case "nw":
				{
					tile.x -= 1
					tile.y += 1
					direction = ""
				}
			default:
				{
					direction = string(dir)
				}
			}
		}

		if color, ok := tileColour[tile]; ok {
			tileColour[tile] = !color
		} else {
			tileColour[tile] = true
		}
	}

	blackCount := 0
	arr := make([]string, 0)
	for tile, isBlack := range tileColour {
		if isBlack {
			blackCount++
		}
		arr = append(arr, fmt.Sprintf("[%d,%d - %v]", tile.x, tile.y, isBlack))
	}

	sort.Strings(arr)
	for i, s := range arr {
		fmt.Printf("%s,", s)
		if i%10 == 0 {
			fmt.Printf("\n")
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%d", blackCount)
	}

}

func day24b() {
	file, err := os.Open("./app/aoc/2020/day24_input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	tileColour := make(map[tile]bool, 0)
	for scanner.Scan() {
		text := scanner.Text()

		direction := ""
		tile := tile{
			x: 0,
			y: 0,
		}
		for _, dir := range text {
			direction += string(dir)
			switch direction {
			case "e":
				{
					tile.x += 2
					direction = ""
				}
			case "w":
				{
					tile.x -= 2
					direction = ""
				}
			case "se":
				{
					tile.x += 1
					tile.y -= 1
					direction = ""
				}
			case "sw":
				{
					tile.x -= 1
					tile.y -= 1
					direction = ""
				}
			case "ne":
				{
					tile.x += 1
					tile.y += 1
					direction = ""
				}
			case "nw":
				{
					tile.x -= 1
					tile.y += 1
					direction = ""
				}
			default:
				{
					direction = string(dir)
				}
			}
		}

		if color, ok := tileColour[tile]; ok {
			tileColour[tile] = !color
		} else {
			tileColour[tile] = true
		}
	}

	//blackCount := 0
	//arr := make([]string, 0)
	//for tile, isBlack := range tileColour {
	//	if isBlack {
	//		blackCount++
	//	}
	//	arr = append(arr, fmt.Sprintf("[%d,%d - %v]", tile.x, tile.y, isBlack))
	//}
	//
	//sort.Strings(arr)
	//for i, s := range arr {
	//	fmt.Printf("%s,", s)
	//	if i%10 == 0 {
	//		fmt.Printf("\n")
	//	}
	//}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%d", getBlackTileCount(tileColour, 100))
	}
}

func getBlackTileCount(tiles map[tile]bool, days int) int {
	adjacentTiles := []tile{{2, 0}, {-2, 0}, {1, -1}, {1, 1}, {-1, -1}, {-1, 1}}
	for i := 0; i < days; i++ {
		temp := make(map[tile]bool, 0)
		bCount := 0
		edgeTiles := getEdgeTiles(tiles)
		xMin := edgeTiles[0].x
		yMin := edgeTiles[0].y
		xMax := edgeTiles[1].x
		yMax := edgeTiles[1].y
		//fmt.Printf("day %d: ", i+1)
		for x := xMin; x <= xMax; x += 1 {
			for y := yMin; y <= yMax; y += 1 {
				if (x+y)%2 != 0 {
					continue
				}
				t := tile{x: x, y: y}
				b := tiles[t]
				count := 0
				for _, adjacent := range adjacentTiles {
					adjacentT := tile{x: t.x + adjacent.x, y: t.y + adjacent.y}
					if tiles[adjacentT] {
						count++
					}
				}
				if b {
					if count == 0 || count > 2 {
						temp[t] = false
					} else {
						bCount++
						temp[t] = true
						//fmt.Printf("[%d, %d]", t.x, t.y)
					}
				} else {
					if count == 2 {
						bCount++
						temp[t] = true
						//fmt.Printf("[%d, %d]", t.x, t.y)
					} else {
						temp[t] = false
					}
				}
			}
		}
		//fmt.Printf(" -- black tiles: %d\n", bCount)

		tiles = temp
	}

	blackCount := 0
	for _, isBlack := range tiles {
		if isBlack {
			blackCount++
		}
	}

	return blackCount
}

func getEdgeTiles(tiles map[tile]bool) [2]tile {
	xMin := math.MaxInt32
	xMax := math.MinInt32
	yMin := math.MaxInt32
	yMax := math.MinInt32

	for t, _ := range tiles {
		if t.x < xMin {
			xMin = t.x
		}
		if t.y < yMin {
			yMin = t.y
		}
		if t.x > xMax {
			xMax = t.x
		}
		if t.y > yMax {
			yMax = t.y
		}
	}

	return [2]tile{{xMin - 1, yMin - 1}, {xMax + 1, yMax + 1}}
}

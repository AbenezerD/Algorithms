package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*
--- Day 20: Jurassic Jigsaw ---
The high-speed train leaves the forest and quickly carries you south. You can even see a desert in the distance! Since you have some spare time, you might as well see if there was anything interesting in the image the Mythical Information Bureau satellite captured.

After decoding the satellite messages, you discover that the data actually contains many small images created by the satellite's camera array. The camera array consists of many cameras; rather than produce a single square image, they produce many smaller square image tiles that need to be reassembled back into a single image.

Each camera in the camera array returns a single monochrome image tile with a random unique ID number. The tiles (your puzzle input) arrived in a random order.

Worse yet, the camera array appears to be malfunctioning: each image tile has been rotated and flipped to a random orientation. Your first task is to reassemble the original image by orienting the tiles so they fit together.

To show how the tiles should be reassembled, each tile's image data includes a border that should line up exactly with its adjacent tiles. All tiles have this border, and the border lines up exactly when the tiles are both oriented correctly. Tiles at the edge of the image also have this border, but the outermost edges won't line up with any other tiles.

For example, suppose you have the following nine tiles:

Tile 2311:
..##.#..#.
##..#.....
#...##..#.
####.#...#
##.##.###.
##...#.###
.#.#.#..##
..#....#..
###...#.#.
..###..###

Tile 1951:
#.##...##.
#.####...#
.....#..##
#...######
.##.#....#
.###.#####
###.##.##.
.###....#.
..#.#..#.#
#...##.#..

Tile 1171:
####...##.
#..##.#..#
##.#..#.#.
.###.####.
..###.####
.##....##.
.#...####.
#.##.####.
####..#...
.....##...

Tile 1427:
###.##.#..
.#..#.##..
.#.##.#..#
#.#.#.##.#
....#...##
...##..##.
...#.#####
.#.####.#.
..#..###.#
..##.#..#.

Tile 1489:
##.#.#....
..##...#..
.##..##...
..#...#...
#####...#.
#..#.#.#.#
...#.#.#..
##.#...##.
..##.##.##
###.##.#..

Tile 2473:
#....####.
#..#.##...
#.##..#...
######.#.#
.#...#.#.#
.#########
.###.#..#.
########.#
##...##.#.
..###.#.#.

Tile 2971:
..#.#....#
#...###...
#.#.###...
##.##..#..
.#####..##
.#..####.#
#..#.#..#.
..####.###
..#.#.###.
...#.#.#.#

Tile 2729:
...#.#.#.#
####.#....
..#.#.....
....#..#.#
.##..##.#.
.#.####...
####.#.#..
##.####...
##..#.##..
#.##...##.

Tile 3079:
#.#.#####.
.#..######
..#.......
######....
####.#..#.
.#...#.##.
#.#####.##
..#.###...
..#.......
..#.###...
By rotating, flipping, and rearranging them, you can find a square arrangement that causes all adjacent borders to line up:

#...##.#.. ..###..### #.#.#####.
..#.#..#.# ###...#.#. .#..######
.###....#. ..#....#.. ..#.......
###.##.##. .#.#.#..## ######....
.###.##### ##...#.### ####.#..#.
.##.#....# ##.##.###. .#...#.##.
#...###### ####.#...# #.#####.##
.....#..## #...##..#. ..#.###...
#.####...# ##..#..... ..#.......
#.##...##. ..##.#..#. ..#.###...

#.##...##. ..##.#..#. ..#.###...
##..#.##.. ..#..###.# ##.##....#
##.####... .#.####.#. ..#.###..#
####.#.#.. ...#.##### ###.#..###
.#.####... ...##..##. .######.##
.##..##.#. ....#...## #.#.#.#...
....#..#.# #.#.#.##.# #.###.###.
..#.#..... .#.##.#..# #.###.##..
####.#.... .#..#.##.. .######...
...#.#.#.# ###.##.#.. .##...####

...#.#.#.# ###.##.#.. .##...####
..#.#.###. ..##.##.## #..#.##..#
..####.### ##.#...##. .#.#..#.##
#..#.#..#. ...#.#.#.. .####.###.
.#..####.# #..#.#.#.# ####.###..
.#####..## #####...#. .##....##.
##.##..#.. ..#...#... .####...#.
#.#.###... .##..##... .####.##.#
#...###... ..##...#.. ...#..####
..#.#....# ##.#.#.... ...##.....
For reference, the IDs of the above tiles are:

1951    2311    3079
2729    1427    2473
2971    1489    1171
To check that you've assembled the image correctly, multiply the IDs of the four corner tiles together. If you do this with the assembled tiles from the example above, you get 1951 * 3079 * 2971 * 1171 = 20899048083289.

Assemble the tiles into an image. What do you get if you multiply together the IDs of the four corner tiles?

Your puzzle answer was 15006909892229.

The first half of this puzzle is complete! It provides one gold star: *

--- Part Two ---
Now, you're ready to check the image for sea monsters.

The borders of each tile are not part of the actual image; start by removing them.

In the example above, the tiles become:

.#.#..#. ##...#.# #..#####
###....# .#....#. .#......
##.##.## #.#.#..# #####...
###.#### #...#.## ###.#..#
##.#.... #.##.### #...#.##
...##### ###.#... .#####.#
....#..# ...##..# .#.###..
.####... #..#.... .#......

#..#.##. .#..###. #.##....
#.####.. #.####.# .#.###..
###.#.#. ..#.#### ##.#..##
#.####.. ..##..## ######.#
##..##.# ...#...# .#.#.#..
...#..#. .#.#.##. .###.###
.#.#.... #.##.#.. .###.##.
###.#... #..#.##. ######..

.#.#.### .##.##.# ..#.##..
.####.## #.#...## #.#..#.#
..#.#..# ..#.#.#. ####.###
#..####. ..#.#.#. ###.###.
#####..# ####...# ##....##
#.##..#. .#...#.. ####...#
.#.###.. ##..##.. ####.##.
...###.. .##...#. ..#..###
Remove the gaps to form the actual image:

.#.#..#.##...#.##..#####
###....#.#....#..#......
##.##.###.#.#..######...
###.#####...#.#####.#..#
##.#....#.##.####...#.##
...########.#....#####.#
....#..#...##..#.#.###..
.####...#..#.....#......
#..#.##..#..###.#.##....
#.####..#.####.#.#.###..
###.#.#...#.######.#..##
#.####....##..########.#
##..##.#...#...#.#.#.#..
...#..#..#.#.##..###.###
.#.#....#.##.#...###.##.
###.#...#..#.##.######..
.#.#.###.##.##.#..#.##..
.####.###.#...###.#..#.#
..#.#..#..#.#.#.####.###
#..####...#.#.#.###.###.
#####..#####...###....##
#.##..#..#...#..####...#
.#.###..##..##..####.##.
...###...##...#...#..###
Now, you're ready to search for sea monsters! Because your image is monochrome, a sea monster will look like this:

                  #
#    ##    ##    ###
 #  #  #  #  #  #
When looking for this pattern in the image, the spaces can be anything; only the # need to match. Also, you might need to rotate or flip your image before it's oriented correctly to find sea monsters. In the above image, after flipping and rotating it to the appropriate orientation, there are two sea monsters (marked with O):

.####...#####..#...###..
#####..#..#.#.####..#.#.
.#.#...#.###...#.##.O#..
#.O.##.OO#.#.OO.##.OOO##
..#O.#O#.O##O..O.#O##.##
...#.#..##.##...#..#..##
#.##.#..#.#..#..##.#.#..
.###.##.....#...###.#...
#.####.#.#....##.#..#.#.
##...#..#....#..#...####
..#.##...###..#.#####..#
....#.##.#.#####....#...
..##.##.###.....#.##..#.
#...#...###..####....##.
.#.##...#.##.#.#.###...#
#.###.#..####...##..#...
#.###...#.##...#.##O###.
.O##.#OO.###OO##..OOO##.
..O#.O..O..O.#O##O##.###
#.#..##.########..#..##.
#.#####..#.#...##..#....
#....##..#.#########..##
#...#.....#..##...###.##
#..###....##.#...##.##.#
Determine how rough the waters are in the sea monsters' habitat by counting the number of # that are not part of a sea monster. In the above example, the habitat's water roughness is 273.

How many # are not part of a sea monster?
*/

func day20() {
	file, err := os.Open("./app/aoc/2020/day20_input.txt")
	if err != nil {
		//fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var currentTile int
	edges := make(map[int][]string, 0)
	lastTile := ""
	for scanner.Scan() {
		text := scanner.Text()
		if strings.HasPrefix(text, "Tile") {
			text = strings.Replace(text, "Tile ", "", 1)
			text = strings.Replace(text, ":", "", 1)

			tile, _ := strconv.Atoi(text)
			edges[tile] = make([]string, 4)
			currentTile = tile
		} else if text == "" {
			edges[currentTile][3] = lastTile
			lastTile = ""
		} else {
			if lastTile == "" {
				edges[currentTile][0] = text
			}
			edges[currentTile][1] = edges[currentTile][1] + string(text[0])
			edges[currentTile][2] = edges[currentTile][2] + string(text[len(text)-1])

			lastTile = text
		}
	}

	if lastTile != "" {
		edges[currentTile][3] = lastTile
	}

	//for tile, edges := range edges {
	//	//fmt.Printf("Tile: %d\nEdges: \n", tile)
	//	//fmt.Printf("\tTop: %s\n", edges[0])
	//	//fmt.Printf("\tLeft: %s\n", edges[1])
	//	//fmt.Printf("\tRight: %s\n", edges[2])
	//	//fmt.Printf("\tBottom: %s\n", edges[3])
	//}

	multiple := 1
	for t1, edges1 := range edges {
		edgesCount := 0

		for t2, edges2 := range edges {
			if t1 == t2 {
				continue
			}

			if isContained(edges1[0], edges2) {
				edgesCount++
			}
			if isContained(edges1[1], edges2) {
				edgesCount++
			}
			if isContained(edges1[2], edges2) {
				edgesCount++
			}
			if isContained(edges1[3], edges2) {
				edgesCount++
			}
		}

		if edgesCount == 2 {
			fmt.Printf("Corner tile found: [Tile: %d]\n", t1)
			multiple *= t1
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%d", multiple)
	}
}

func day20b() {
	file, err := os.Open("./app/aoc/2020/day20_input.txt")
	if err != nil {
		//fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var currentTile int
	edges := make(map[int][]string, 0)
	lastTile := ""
	for scanner.Scan() {
		text := scanner.Text()
		if strings.HasPrefix(text, "Tile") {
			text = strings.Replace(text, "Tile ", "", 1)
			text = strings.Replace(text, ":", "", 1)

			tile, _ := strconv.Atoi(text)
			edges[tile] = make([]string, 4)
			currentTile = tile
		} else if text == "" {
			edges[currentTile][3] = lastTile
			lastTile = ""
		} else {
			if lastTile == "" {
				edges[currentTile][0] = text
			}
			edges[currentTile][1] = edges[currentTile][1] + string(text[0])
			edges[currentTile][2] = edges[currentTile][2] + string(text[len(text)-1])

			lastTile = text
		}
	}

	if lastTile != "" {
		edges[currentTile][3] = lastTile
	}

	//for tile, edges := range edges {
	//	//fmt.Printf("Tile: %d\nEdges: \n", tile)
	//	//fmt.Printf("\tTop: %s\n", edges[0])
	//	//fmt.Printf("\tLeft: %s\n", edges[1])
	//	//fmt.Printf("\tRight: %s\n", edges[2])
	//	//fmt.Printf("\tBottom: %s\n", edges[3])
	//}

	edgeMap := make(map[int][]int, 0)
	corners := make([]int, 0)
	for t1, edges1 := range edges {
		edgesCount := 0

		for t2, edges2 := range edges {
			if t1 == t2 {
				continue
			}

			if isContained(edges1[0], edges2) {
				edgeMap[t1] = append(edgeMap[t1], t2)
				edgesCount++
			}
			if isContained(edges1[1], edges2) {
				edgeMap[t1] = append(edgeMap[t1], t2)
				edgesCount++
			}
			if isContained(edges1[2], edges2) {
				edgeMap[t1] = append(edgeMap[t1], t2)
				edgesCount++
			}
			if isContained(edges1[3], edges2) {
				edgeMap[t1] = append(edgeMap[t1], t2)
				edgesCount++
			}
		}

		if edgesCount == 2 {
			fmt.Printf("Corner tile found: [Tile: %d]\n", t1)
			corners = append(corners, t1)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%d", 1)
	}
}

func isContained(val string, list []string) bool {
	for _, item := range list {
		if val == item || reverseStr(val) == item {
			return true
		}
	}

	return false
}

func reverseStr(val string) string {
	retStr := ""
	for i := len(val) - 1; i >= 0; i-- {
		retStr += string(val[i])
	}

	return retStr
}

const isize = 12
const tsize = 10

func day20c() {
	input, _ := ioutil.ReadFile("./app/aoc/2020/day20_input.txt")
	split := strings.Split(strings.TrimSpace(string(input)), "\n\n")

	tiles := map[int][]string{}
	edges := map[string][]int{}
	for _, s := range split {
		var id int
		fmt.Sscanf(s, "Tile %d:", &id)
		tiles[id] = strings.Split(s, "\n")[1:]
		for _, e := range getEdges(tiles[id]) {
			edges[canon(e)] = append(edges[canon(e)], id)
		}
	}

	counts := map[int]int{}
	var corner int

	part1 := 1
	for _, v := range edges {
		if len(v) == 1 {
			if counts[v[0]]++; counts[v[0]] == 2 {
				corner = v[0]
				part1 *= corner
			}
		}
	}
	fmt.Println(part1)

	order := [isize][isize][]string{{tiles[corner]}}
	for i := 1; len(edges[canon(order[0][0][0])]) != 1 || len(edges[canon(col(order[0][0], 0))]) != 1; i++ {
		order[0][0] = rotate(tiles[corner], i)
	}
	delete(tiles, corner)

	image := []string{}
	for y := 0; y < isize; y++ {
		for x := 0; x < isize; x++ {
			if y == 0 && x == 0 {
				continue
			}

			for id, tile := range tiles {
				if (y == 0 || hasEdge(tile, order[y-1][x][tsize-1])) && (x == 0 || hasEdge(tile, col(order[y][x-1], tsize-1))) {
					order[y][x] = tile
					for i := 1; y != 0 && order[y][x][0] != order[y-1][x][tsize-1] || x != 0 && col(order[y][x], 0) != col(order[y][x-1], tsize-1); i++ {
						order[y][x] = rotate(tile, i)
					}
					delete(tiles, id)
					break
				}
			}
		}
	}

	part2 := 0
	for y := 0; y < isize; y++ {
		for i := 1; i < tsize-1; i++ {
			image = append(image, "")
			for x := 0; x < isize; x++ {
				image[y*(tsize-2)+i-1] += order[y][x][i][1 : tsize-1]
				part2 += strings.Count(order[y][x][i][1:tsize-1], "#")
			}
		}
	}

	orig := make([]string, len(image))
	copy(orig, image)
	for i := 1; i <= 8; i++ {
		monster := "..................#.\n#....##....##....###\n.#..#..#..#..#..#..."
		split := strings.Split(monster, "\n")

		for y := 0; y < len(image)-len(split); y++ {
		search:
			for x := 0; x < len(image[0])-len(split[0]); x++ {
				for i, s := range split {
					if match, _ := regexp.MatchString(s, image[y+i][x:x+len(s)]); !match {
						continue search
					}
				}
				part2 -= strings.Count(monster, "#")
			}
		}
		image = rotate(orig, i)
	}
	fmt.Println(part2)
}

func getEdges(tile []string) []string {
	return []string{tile[0], col(tile, len(tile[0])-1), tile[len(tile)-1], col(tile, 0)}
}

func hasEdge(tile []string, edge string) bool {
	for _, e := range getEdges(tile) {
		if canon(e) == canon(edge) {
			return true
		}
	}
	return false
}

func reverse(s string) (rs string) {
	for _, r := range s {
		rs = string(r) + rs
	}
	return
}

func canon(s string) string {
	if rs := reverse(s); rs < s {
		return rs
	}
	return s
}

func col(tile []string, i int) (col string) {
	for _, s := range tile {
		col += string(s[i])
	}
	return
}

func flip(tile []string) (flip []string) {
	for _, s := range tile {
		flip = append(flip, reverse(s))
	}
	return
}

func rotate(tile []string, n int) (rot []string) {
	if n%8 >= 4 {
		tile = flip(tile)
	}
	for ; n > 0; n-- {
		rot = []string{}
		for i := range tile {
			rot = append(rot, reverse(col(tile, i)))
		}
		tile = rot
	}
	return
}

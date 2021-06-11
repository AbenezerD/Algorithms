package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

/*
--- Day 23: Crab Cups ---
The small crab challenges you to a game! The crab is going to mix up some cups, and you have to predict where they'll end up.

The cups will be arranged in a circle and labeled clockwise (your puzzle input). For example, if your labeling were 32415, there would be five cups in the circle; going clockwise around the circle from the first cup, the cups would be labeled 3, 2, 4, 1, 5, and then back to 3 again.

Before the crab starts, it will designate the first cup in your list as the current cup. The crab is then going to do 100 moves.

Each move, the crab does the following actions:

The crab picks up the three cups that are immediately clockwise of the current cup. They are removed from the circle; cup spacing is adjusted as necessary to maintain the circle.
The crab selects a destination cup: the cup with a label equal to the current cup's label minus one. If this would select one of the cups that was just picked up, the crab will keep subtracting one until it finds a cup that wasn't just picked up. If at any point in this process the value goes below the lowest value on any cup's label, it wraps around to the highest value on any cup's label instead.
The crab places the cups it just picked up so that they are immediately clockwise of the destination cup. They keep the same order as when they were picked up.
The crab selects a new current cup: the cup which is immediately clockwise of the current cup.
For example, suppose your cup labeling were 389125467. If the crab were to do merely 10 moves, the following changes would occur:

-- move 1 --
cups: (3) 8  9  1  2  5  4  6  7
pick up: 8, 9, 1
destination: 2

-- move 2 --
cups:  3 (2) 8  9  1  5  4  6  7
pick up: 8, 9, 1
destination: 7

-- move 3 --
cups:  3  2 (5) 4  6  7  8  9  1
pick up: 4, 6, 7
destination: 3

-- move 4 --
cups:  7  2  5 (8) 9  1  3  4  6
pick up: 9, 1, 3
destination: 7

-- move 5 --
cups:  3  2  5  8 (4) 6  7  9  1
pick up: 6, 7, 9
destination: 3

-- move 6 --
cups:  9  2  5  8  4 (1) 3  6  7
pick up: 3, 6, 7
destination: 9

-- move 7 --
cups:  7  2  5  8  4  1 (9) 3  6
pick up: 3, 6, 7
destination: 8

-- move 8 --
cups:  8  3  6  7  4  1  9 (2) 5
pick up: 5, 8, 3
destination: 1

-- move 9 --
cups:  7  4  1  5  8  3  9  2 (6)
pick up: 7, 4, 1
destination: 5

-- move 10 --
cups: (5) 7  4  1  8  3  9  2  6
pick up: 7, 4, 1
destination: 3

-- final --
cups:  5 (8) 3  7  4  1  9  2  6
In the above example, the cups' values are the labels as they appear moving clockwise around the circle; the current cup is marked with ( ).

After the crab is done, what order will the cups be in? Starting after the cup labeled 1, collect the other cups' labels clockwise into a single string with no extra characters; each number except 1 should appear exactly once. In the above example, after 10 moves, the cups clockwise from 1 are labeled 9, 2, 6, 5, and so on, producing 92658374. If the crab were to complete all 100 moves, the order after cup 1 would be 67384529.

Using your labeling, simulate 100 moves. What are the labels on the cups after cup 1?

Your puzzle answer was 36472598.

The first half of this puzzle is complete! It provides one gold star: *

--- Part Two ---
Due to what you can only assume is a mistranslation (you're not exactly fluent in Crab), you are quite surprised when the crab starts arranging many cups in a circle on your raft - one million (1000000) in total.

Your labeling is still correct for the first few cups; after that, the remaining cups are just numbered in an increasing fashion starting from the number after the highest number in your list and proceeding one by one until one million is reached. (For example, if your labeling were 54321, the cups would be numbered 5, 4, 3, 2, 1, and then start counting up from 6 until one million is reached.) In this way, every number from one through one million is used exactly once.

After discovering where you made the mistake in translating Crab Numbers, you realize the small crab isn't going to do merely 100 moves; the crab is going to do ten million (10000000) moves!

The crab is going to hide your stars - one each - under the two cups that will end up immediately clockwise of cup 1. You can have them if you predict what the labels on those cups will be when the crab is finished.

In the above example (389125467), this would be 934001 and then 159792; multiplying these together produces 149245887792.

Determine which two cups will end up immediately clockwise of cup 1. What do you get if you multiply their labels together?
*/

func day23(c string, moves int) string {
	cups := make([]int, len(c))
	for i, char := range c {
		num, _ := strconv.Atoi(string(char))
		cups[i] = num
	}

	n := len(cups)
	currentCupIndex := 0
	for i := 0; i < moves; i++ {
		pickup := make([]int, 3)
		moduloCurrent := currentCupIndex % len(cups)
		pickup[0] = cups[(moduloCurrent+1)%n]
		pickup[1] = cups[(moduloCurrent+2)%n]
		pickup[2] = cups[(moduloCurrent+3)%n]
		destination := cups[moduloCurrent] - 1

		destinationFound := false
		destinationIndex := 0
		for !destinationFound && destination > 0 {
			for j := currentCupIndex + 4; j < currentCupIndex+len(cups); j++ {
				moduloJ := j % len(cups)
				if destination == cups[moduloJ] {
					destinationIndex = moduloJ
					destinationFound = true
					break
				}
			}
			if !destinationFound {
				destination -= 1
			}
		}

		max := 0
		maxIndex := 0
		if !destinationFound {
			for j := currentCupIndex + 4; j < currentCupIndex+n; j++ {
				moduloJ := j % len(cups)
				if cups[moduloJ] > max {
					max = cups[moduloJ]
					maxIndex = moduloJ
				}
			}
			destinationIndex = maxIndex
			destination = cups[destinationIndex]
		}

		// 3 (2) 8  9  1  5  4  6  7
		// 3  2 (5) 4  6  7  8  9  1
		// 7  2  5 (8) 9  1  3  4  6
		temp := []int{cups[moduloCurrent]}
		lim := (destinationIndex + n) % n
		if destinationIndex == 0 {
			lim = n
		}
		for j := (moduloCurrent + 4) % n; j < lim; j++ {
			moduloJ := j % len(cups)
			temp = append(temp, cups[moduloJ])
		}
		temp = append(temp, cups[destinationIndex])
		temp = append(temp, pickup...)
		for j := (destinationIndex + 1) % n; len(temp) < n; j++ {
			moduloJ := j % len(cups)
			temp = append(temp, cups[moduloJ])
		}

		tot := 0
		for _, cup := range cups {
			tot += cup
		}
		cups = temp
		//fmt.Printf("%v - current: %d, destination: %d, pickup: %v, tot: %d\n", cups, cups[currentCupIndex], destination, pickup, tot)
		currentCupIndex = 1
	}

	indexAt1 := 0
	for i, cup := range cups {
		if cup == 1 {
			indexAt1 = i
			break
		}
	}

	retString := ""
	for j := indexAt1 + 1; j < indexAt1+len(cups); j++ {
		moduloJ := j % n
		retString += strconv.Itoa(cups[moduloJ])
	}

	fmt.Printf("%s\n", retString)
	return retString
}

func day23c(c string, moves, max int) string {
	cups := make([]int, max)
	for i, char := range c {
		num, _ := strconv.Atoi(string(char))
		cups[i] = num
	}
	for i := 9; i < max; i++ {
		cups[i] = i + 1
	}

	circularList := initRing(cups)

	current := circularList
	ringMap := make(map[int]*ring.Ring, len(cups))
	for len(ringMap) < len(cups) {
		ringMap[current.Value.(int)] = current
		current = current.Next()
	}

	currentCup := ringMap[cups[0]]
	for i := 0; i < moves; i++ {
		pickup := currentCup.Unlink(3)

		destination := getDestination(currentCup.Value.(int), max, [3]int{pickup.Value.(int), pickup.Next().Value.(int), pickup.Next().Next().Value.(int)})

		// 3 (2) 8  9  1  5  4  6  7
		// 3  2 (5) 4  6  7  8  9  1
		// 7  2  5 (8) 9  1  3  4  6
		destinationCup := ringMap[destination]
		destinationCup.Link(pickup)

		//3-8-9-2-7-4-5-6-1 7-4-5
		//currentCup.Do(func(p interface{}) {
		//fmt.Printf("destination: %d", destination)
		//})
		//currentCup.Do(func(p interface{}) {
		//	fmt.Printf("%d -> ", p.(int))
		//})
		//fmt.Printf("pickup:")
		//for i := 0; i < 3; i++ {
		//	fmt.Printf("%d -> ", pickup.Value.(int))
		//	pickup = pickup.Next()
		//}
		//fmt.Printf("\n")
		//fmt.Printf("%v - current: %d, destination: %d, pickup: %v, tot: %d\n", cups, cups[currentCupIndex], destination, pickup, tot)
		currentCup = currentCup.Next()
	}

	indexAt1 := ringMap[1]

	i1 := indexAt1.Next().Value.(int)
	i2 := indexAt1.Next().Next().Value.(int)
	retString := strconv.Itoa(i1 * i2)
	fmt.Printf("%d * %d = %s", i1, i2, retString)
	return retString
}

func getDestination(current, max int, nextThree [3]int) int {
	destination := current - 1

	destinationFound := false
	for !destinationFound && destination > 0 {
		if destination != nextThree[0] && destination != nextThree[1] && destination != nextThree[2] {
			destinationFound = true
		}
		if !destinationFound {
			destination -= 1
		}
	}

	if destinationFound {
		return destination
	}

	maxIsDestination := false
	for !maxIsDestination {
		if max != nextThree[0] && max != nextThree[1] && max != nextThree[2] {
			maxIsDestination = true
		}
		if !maxIsDestination {
			max -= 1
		}
	}

	return max
}

func initRing(c []int) *ring.Ring {
	r := ring.New(len(c))
	next := r.Next()
	for i := 0; i < len(c); i++ {
		next.Value = c[i]
		next = next.Next()
	}

	//for i := 10; i <= 1000000; i++ {
	//	next.Value = i
	//	next = next.Next()
	//}

	return next
}

func day23b(c string, moves int) string {
	cups := make([]int, len(c)+1000000)
	for i, char := range c {
		num, _ := strconv.Atoi(string(char))
		cups[i] = num
	}

	for i := 10; i < 1000000+len(c); i++ {
		cups[i] = i
	}

	n := len(cups)
	currentCupIndex := 0
	for i := 0; i < moves; i++ {
		pickup := make([]int, 3)
		moduloCurrent := currentCupIndex % len(cups)
		pickup[0] = cups[(moduloCurrent+1)%n]
		pickup[1] = cups[(moduloCurrent+2)%n]
		pickup[2] = cups[(moduloCurrent+3)%n]
		destination := cups[moduloCurrent] - 1

		destinationFound := false
		destinationIndex := 0
		for !destinationFound && destination > 0 {
			for j := currentCupIndex + 4; j < currentCupIndex+len(cups); j++ {
				moduloJ := j % len(cups)
				if destination == cups[moduloJ] {
					destinationIndex = moduloJ
					destinationFound = true
					break
				}
			}
			if !destinationFound {
				destination -= 1
			}
		}

		max := 0
		maxIndex := 0
		if !destinationFound {
			for j := currentCupIndex + 4; j < currentCupIndex+n; j++ {
				moduloJ := j % len(cups)
				if cups[moduloJ] > max {
					max = cups[moduloJ]
					maxIndex = moduloJ
				}
			}
			destinationIndex = maxIndex
			destination = cups[destinationIndex]
		}

		// 3 (2) 8  9  1  5  4  6  7
		// 3  2 (5) 4  6  7  8  9  1
		// 7  2  5 (8) 9  1  3  4  6
		temp := []int{cups[moduloCurrent]}
		lim := (destinationIndex + n) % n
		if destinationIndex == 0 {
			lim = n
		}
		for j := (moduloCurrent + 4) % n; j < lim; j++ {
			moduloJ := j % len(cups)
			temp = append(temp, cups[moduloJ])
		}
		temp = append(temp, cups[destinationIndex])
		temp = append(temp, pickup...)
		for j := (destinationIndex + 1) % n; len(temp) < n; j++ {
			moduloJ := j % len(cups)
			temp = append(temp, cups[moduloJ])
		}

		tot := 0
		for _, cup := range cups {
			tot += cup
		}
		cups = temp
		//fmt.Printf("%v - current: %d, destination: %d, pickup: %v, tot: %d\n", cups, cups[currentCupIndex], destination, pickup, tot)
		currentCupIndex = 1
		fmt.Printf(".%d", i)
	}

	indexAt1 := 0
	for i, cup := range cups {
		if cup == 1 {
			indexAt1 = i
			break
		}
	}

	retString := strconv.Itoa(cups[indexAt1+1]) + " " + strconv.Itoa(cups[indexAt1+2])
	return retString
}

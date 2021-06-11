package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

/*
--- Day 1: Report Repair ---
After saving Christmas five years in a row, you've decided to take a vacation at a nice resort on a tropical island. Surely, Christmas will go on without you.

The tropical island has its own currency and is entirely cash-only. The gold coins used there have a little picture of a starfish; the locals just call them stars. None of the currency exchanges seem to have heard of them, but somehow, you'll need to find fifty of these coins by the time you arrive so you can pay the deposit on your room.

To save your vacation, you need to get all fifty stars by December 25th.

Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!

Before you leave, the Elves in accounting just need you to fix your expense report (your puzzle input); apparently, something isn't quite adding up.

Specifically, they need you to find the two entries that sum to 2020 and then multiply those two numbers together.

For example, suppose your expense report contained the following:

1721
979
366
299
675
1456
In this list, the two entries that sum to 2020 are 1721 and 299. Multiplying them together produces 1721 * 299 = 514579, so the correct answer is 514579.

Of course, your expense report is much larger. Find the two entries that sum to 2020; what do you get if you multiply them together?

Your puzzle answer was 55776.

--- Part Two ---
The Elves in accounting are thankful for your help; one of them even offers you a starfish coin they had left over from a past vacation. They offer you a second one if you can find three numbers in your expense report that meet the same criteria.

Using the above example again, the three entries that sum to 2020 are 979, 366, and 675. Multiplying them together produces the answer, 241861950.

In your expense report, what is the product of the three entries that sum to 2020?

Your puzzle answer was 223162626.
*/

func ReadInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}

func get2(intArray []int) int {
	diff := make(map[int]int, len(intArray))
	for _, num := range intArray {
		difference := 2020.0 - int(num)
		diff[difference] = int(num)
	}

	for _, num := range intArray {
		if difference, ok := diff[num]; ok {
			fmt.Printf("%d + %d = %d\n", difference, int(num), difference+int(num))
			return difference * num
		}
	}
	return 0
}

type tuple struct {
	x int
	y int
}

func get3(intArray []int) int {
	tuple := make(map[int][]int)
	for i := range intArray {
		for j := range intArray {
			if i != j {
				tuple[intArray[i]+intArray[j]] = []int{intArray[i], intArray[j]}
			}
		}
	}

	diff := make(map[int]int, 0)
	for k, v := range tuple {
		for i := range intArray {
			if i != v[0] && i != v[1] {
				difference := 2020 - k
				diff[difference] = k
			}
		}
	}

	for _, num := range intArray {
		if difference, ok := diff[num]; ok {
			arr := tuple[difference]
			fmt.Printf("%d + %d = %d\n", difference, num, difference+num)
			fmt.Printf("%d + %d + %d = %d\n", arr[0], arr[1], num, num+arr[0]+arr[1])
			return arr[0] * arr[1] * num
		}
	}
	return 0
}

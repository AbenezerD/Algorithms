package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
--- Day 16: Ticket Translation ---
As you're walking to yet another connecting flight, you realize that one of the legs of your re-routed trip coming up is on a high-speed train. However, the train ticket you were given is in a language you don't understand. You should probably figure out what it says before you get to the train station after the next flight.

Unfortunately, you can't actually read the words on the ticket. You can, however, read the numbers, and so you figure out the fields these tickets must have and the valid ranges for values in those fields.

You collect the rules for ticket fields, the numbers on your ticket, and the numbers on other nearby tickets for the same train service (via the airport security cameras) together into a single document you can reference (your puzzle input).

The rules for ticket fields specify a list of fields that exist somewhere on the ticket and the valid ranges of values for each field. For example, a rule like class: 1-3 or 5-7 means that one of the fields in every ticket is named class and can be any value in the ranges 1-3 or 5-7 (inclusive, such that 3 and 5 are both valid in this field, but 4 is not).

Each ticket is represented by a single line of comma-separated values. The values are the numbers on the ticket in the order they appear; every ticket has the same format. For example, consider this ticket:

.--------------------------------------------------------.
| ????: 101    ?????: 102   ??????????: 103     ???: 104 |
|                                                        |
| ??: 301  ??: 302             ???????: 303      ??????? |
| ??: 401  ??: 402           ???? ????: 403    ????????? |
'--------------------------------------------------------'
Here, ? represents text in a language you don't understand. This ticket might be represented as 101,102,103,104,301,302,303,401,402,403; of course, the actual train tickets you're looking at are much more complicated. In any case, you've extracted just the numbers in such a way that the first number is always the same specific field, the second number is always a different specific field, and so on - you just don't know what each position actually means!

Start by determining which tickets are completely invalid; these are tickets that contain values which aren't valid for any field. Ignore your ticket for now.

For example, suppose you have the following notes:

class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12
It doesn't matter which position corresponds to which field; you can identify invalid nearby tickets by considering only whether tickets contain values that are not valid for any field. In this example, the values on the first nearby ticket are all valid for at least one field. This is not true of the other three nearby tickets: the values 4, 55, and 12 are are not valid for any field. Adding together all of the invalid values produces your ticket scanning error rate: 4 + 55 + 12 = 71.

Consider the validity of the nearby tickets you scanned. What is your ticket scanning error rate?

Your puzzle answer was 23009.

The first half of this puzzle is complete! It provides one gold star: *

--- Part Two ---
Now that you've identified which tickets contain invalid values, discard those tickets entirely. Use the remaining valid tickets to determine which field is which.

Using the valid ranges for each field, determine what order the fields appear on the tickets. The order is consistent between all tickets: if seat is the third field, it is the third field on every ticket, including your ticket.

For example, suppose you have the following notes:

class: 0-1 or 4-19
row: 0-5 or 8-19
seat: 0-13 or 16-19

your ticket:
11,12,13

nearby tickets:
3,9,18
15,1,5
5,14,9
Based on the nearby tickets in the above example, the first position must be row, the second position must be class, and the third position must be seat; you can conclude that in your ticket, class is 12, row is 11, and seat is 13.

Once you work out which field is which, look for the six fields on your ticket that start with the word departure. What do you get if you multiply those six values together?
*/

func day16() {
	file, err := os.Open("./app/aoc/2020/day16_input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	nearbyTicketsNext := false
	ranges := make([][2]int, 0)
	numsNotInRange := make([]int, 0)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, "or") {
			strRange := strings.Split(text, ": ")
			strRanges := strings.Split(strRange[1], " or ")
			for _, r := range strRanges {
				nums := strings.Split(r, "-")
				min, _ := strconv.Atoi(nums[0])
				max, _ := strconv.Atoi(nums[1])

				ranges = append(ranges, [2]int{min, max})
			}
		}

		if nearbyTicketsNext {
			numStrings := strings.Split(text, ",")
			for _, numStr := range numStrings {
				num, _ := strconv.Atoi(numStr)
				if !numInAnyRange(num, ranges) {
					numsNotInRange = append(numsNotInRange, num)
				}
			}
		}
		if text == "nearby tickets:" {
			nearbyTicketsNext = true
		}
	}

	total := 0
	for _, num := range numsNotInRange {
		total += num
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%d", total)
	}
}

func day16b() {
	file, err := os.Open("./app/aoc/2020/day16_input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	nearbyTicketsNext := false
	myTicketNext := false
	ranges := make(map[string][][2]int, 0)
	myTickets := make([]int, 0)
	positions := 0
	positionsMap := make(map[string]map[int]int, 0)
	allRanges := make([][2]int, 0)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, "or") {
			strRange := strings.Split(text, ": ")
			strRanges := strings.Split(strRange[1], " or ")
			nums1 := strings.Split(strRanges[0], "-")
			min1, _ := strconv.Atoi(nums1[0])
			max1, _ := strconv.Atoi(nums1[1])
			allRanges = append(allRanges, [2]int{min1, max1})

			nums2 := strings.Split(strRanges[1], "-")
			min2, _ := strconv.Atoi(nums2[0])
			max2, _ := strconv.Atoi(nums2[1])
			allRanges = append(allRanges, [2]int{min2, max2})

			ranges[strRange[0]] = [][2]int{{min1, max1}, {min2, max2}}
			positions++
		}

		if text == "nearby tickets:" {
			nearbyTicketsNext = true
		}
		if myTicketNext && !nearbyTicketsNext {
			numStrings := strings.Split(text, ",")
			for _, numStr := range numStrings {
				num, _ := strconv.Atoi(numStr)
				myTickets = append(myTickets, num)
			}
			myTicketNext = false
		}
		if text == "your ticket:" {
			myTicketNext = true
		}

		if nearbyTicketsNext {
			skipTicket := false
			numStrings := strings.Split(text, ",")
			for _, numStr := range numStrings {
				num, _ := strconv.Atoi(numStr)
				if !numInAnyRange(num, allRanges) {
					skipTicket = true
				}
			}
			if !skipTicket {
				for i, numStr := range numStrings {
					num, _ := strconv.Atoi(numStr)
					for key, val := range ranges {
						if numInAnyRange(num, val) {
							if count, ok := positionsMap[key][i]; ok {
								positionsMap[key][i] = count + 1
							} else {
								if posMap, ok := positionsMap[key]; ok {
									posMap[i] = 1
								} else {
									positionsMap[key] = map[int]int{i: 1}
								}
							}
						}
					}
				}
			}
		}
	}

	//indexes := make([]int, 0)
	//for key, val := range positionsMap {
	//	fmt.Printf("%s ", key)
	//	for k, v := range val {
	//		if v == 190 {
	//			fmt.Printf("%d,", k)
	//		}
	//	}
	//	fmt.Printf("\n")
	//	if strings.HasPrefix(key, "departure") {
	//		maxCount := 0
	//		maxIndex := 0
	//		for k, v := range val {
	//			if v > maxCount {
	//				maxCount = v
	//				maxIndex = k
	//			}
	//		}
	//		indexes = append(indexes, maxIndex)
	//		//fmt.Printf("%s - index = %d\n", key, maxIndex)
	//	}
	//}

	multiple := 1
	for _, index := range []int{4, 7, 9, 10, 11, 14} {
		multiple *= myTickets[index]
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%d", multiple)
	}
}

func numInAnyRange(num int, ranges [][2]int) bool {
	// range[min, max] if min <= num <= max; num in range
	for _, numRange := range ranges {
		if num >= numRange[0] && num <= numRange[1] {
			return true
		}
	}

	return false
}

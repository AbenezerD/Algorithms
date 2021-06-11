package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
--- Day 7: Handy Haversacks ---
You land at the regional airport in time for your next flight. In fact, it looks like you'll even have time to grab some food: all flights are currently delayed due to issues in luggage processing.

Due to recent aviation regulations, many rules (your puzzle input) are being enforced about bags and their contents; bags must be color-coded and must contain specific quantities of other color-coded bags. Apparently, nobody responsible for these regulations considered how long they would take to enforce!

For example, consider the following rules:

light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.
These rules specify the required contents for 9 bag types. In this example, every faded blue bag is empty, every vibrant plum bag contains 11 bags (5 faded blue and 6 dotted black), and so on.

You have a shiny gold bag. If you wanted to carry it in at least one other bag, how many different bag colors would be valid for the outermost bag? (In other words: how many colors can, eventually, contain at least one shiny gold bag?)

In the above rules, the following options would be available to you:

A bright white bag, which can hold your shiny gold bag directly.
A muted yellow bag, which can hold your shiny gold bag directly, plus some other bags.
A dark orange bag, which can hold bright white and muted yellow bags, either of which could then hold your shiny gold bag.
A light red bag, which can hold bright white and muted yellow bags, either of which could then hold your shiny gold bag.
So, in this example, the number of bag colors that can eventually contain at least one shiny gold bag is 4.

How many bag colors can eventually contain at least one shiny gold bag? (The list of rules is quite long; make sure you get all of it.)

Your puzzle answer was 335.

The first half of this puzzle is complete! It provides one gold star: *

--- Part Two ---
It's getting pretty expensive to fly these days - not because of ticket prices, but because of the ridiculous number of bags you need to buy!

Consider again your shiny gold bag and the rules from the above example:

faded blue bags contain 0 other bags.
dotted black bags contain 0 other bags.
vibrant plum bags contain 11 other bags: 5 faded blue bags and 6 dotted black bags.
dark olive bags contain 7 other bags: 3 faded blue bags and 4 dotted black bags.
So, a single shiny gold bag must contain 1 dark olive bag (and the 7 bags within it) plus 2 vibrant plum bags (and the 11 bags within each of those): 1 + 1*7 + 2 + 2*11 = 32 bags!

Of course, the actual rules have a small chance of going several levels deeper than this example; be sure to count all of the bags, even if the nesting becomes topologically impractical!

Here's another example:

shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags.
In this example, a single shiny gold bag must contain 126 other bags.

How many individual bags are required inside your single shiny gold bag?
*/

type Node struct {
	Name  string
	Edges []*Node
}

type Node2 struct {
	Name  string
	Edges []*ChildNode
}

type ChildNode struct {
	*Node2
	Total int
	Count int
}

func getBagColors() {
	file, err := os.Open("./app/aoc/2020/day7_input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	bagsMap := make(map[string]*Node, 0)
	for scanner.Scan() {
		text := scanner.Text()
		arr := strings.Split(text, " bags contain ")

		node := createIfNotExist(arr[0], bagsMap)

		children := strings.Split(arr[1], ", ")
		for _, child := range children {
			if child == "no other bags." {
				continue
			}
			arr := strings.Split(child, " ")
			childText := arr[1] + " " + arr[2]
			childNode := createIfNotExist(childText, bagsMap)
			childNode.Edges = append(childNode.Edges, node)
		}
	}

	shinyGoldBag := bagsMap["shiny gold"]
	shinyBagsParents := make(map[string]bool, len(shinyGoldBag.Edges))
	countIfNotCounted(shinyGoldBag, shinyBagsParents)
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	} else if shinyGoldBag != nil {
		fmt.Printf("shiny gold bag has %d direct parents\n", len(shinyGoldBag.Edges))
		fmt.Printf("shiny gold bag has %d indirect parents\n", len(shinyBagsParents))
	} else {
		fmt.Printf("shiny gold bag is nil")
	}
}

func getBagColors2() {
	file, err := os.Open("./app/aoc/2020/day7_input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	bagsMap := make(map[string]*ChildNode, 0)
	for scanner.Scan() {
		text := scanner.Text()
		arr := strings.Split(text, " bags contain ")

		node := createIfNotExist2(arr[0], bagsMap)

		children := strings.Split(arr[1], ", ")
		for _, child := range children {
			if child == "no other bags." {
				continue
			}
			arr := strings.Split(child, " ")
			childText := arr[1] + " " + arr[2]
			childNode := createIfNotExist2(childText, bagsMap)
			count, _ := strconv.Atoi(arr[0])
			node.Edges = append(node.Edges, &ChildNode{Node2: childNode.Node2, Count: count})
		}
	}

	shinyGoldBag := bagsMap["shiny gold"]
	shinyGoldBag.Count = 1

	shinyBagsParents := make(map[string]bool, len(shinyGoldBag.Edges))
	count3 := countIfNotCounted3(shinyGoldBag)

	sum := 1
	var dfs func(node *ChildNode)

	visited := make(map[string]bool, 0)
	dfs = func(node *ChildNode) {
		for _, edge := range node.Edges {
			dfs(edge)
		}
		if !visited[node.Name] {
			sum += sum * node.Count
			visited[node.Name] = true
		}
		node.Total = sum
	}
	dfs(shinyGoldBag)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	} else if shinyGoldBag != nil {
		fmt.Printf("shiny gold bag has %d direct parents\n", len(shinyGoldBag.Edges))
		fmt.Printf("shiny gold bag has %d indirect parents\n", len(shinyBagsParents))
		fmt.Printf("shiny gold bag has %d indirect parents'\n", count3)
		fmt.Printf("shiny gold bag has %d indirect parents dfs'\n", shinyGoldBag.Total)
	} else {
		fmt.Printf("shiny gold bag is nil")
	}
}

func createIfNotExist(text string, bags map[string]*Node) *Node {
	if node, ok := bags[text]; ok {
		return node
	}

	node := &Node{
		Name: text,
	}
	bags[text] = node
	return node
}

func createIfNotExist2(text string, bags map[string]*ChildNode) *ChildNode {
	if node, ok := bags[text]; ok {
		return node
	}

	node := &ChildNode{
		Node2: &Node2{
			Name: text,
		},
	}
	bags[text] = node
	return node
}

func countIfNotCounted(node *Node, visited map[string]bool) {
	if node.Name != "shiny gold" {
		visited[node.Name] = true
	}
	for _, n := range node.Edges {
		if !visited[n.Name] {
			countIfNotCounted(n, visited)
		}
	}
}

func countIfNotCounted3(node *ChildNode) int {
	if len(node.Edges) == 0 {
		return node.Count
	}
	childCount := 0
	for _, n := range node.Edges {
		childCount += countIfNotCounted3(n)
	}

	return node.Count + node.Count*childCount
}

// shiny gold -> 1 dark olive bag, 2 vibrant plum bags.
// dark olive -> 3 faded blue bags, 4 dotted black bags.
// vibrant plum -> 5 faded blue bags, 6 dotted black bags.

// shiny gold -> bright white, muted yellow
// bright white -> light red, dark orange
// muted yellow -> light red, dark orange

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

/*
--- Day 21: Allergen Assessment ---
You reach the train's last stop and the closest you can get to your vacation island without getting wet. There aren't even any boats here, but nothing can stop you now: you build a raft. You just need a few days' worth of food for your journey.

You don't speak the local language, so you can't read any ingredients lists. However, sometimes, allergens are listed in a language you do understand. You should be able to use this information to determine which ingredient contains which allergen and work out which foods are safe to take with you on your trip.

You start by compiling a list of foods (your puzzle input), one food per line. Each line includes that food's ingredients list followed by some or all of the allergens the food contains.

Each allergen is found in exactly one ingredient. Each ingredient contains zero or one allergen. Allergens aren't always marked; when they're listed (as in (contains nuts, shellfish) after an ingredients list), the ingredient that contains each listed allergen will be somewhere in the corresponding ingredients list. However, even if an allergen isn't listed, the ingredient that contains that allergen could still be present: maybe they forgot to label it, or maybe it was labeled in a language you don't know.

For example, consider the following list of foods:

mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)
The first food in the list has four ingredients (written in a language you don't understand): mxmxvkd, kfcds, sqjhc, and nhms. While the food might contain other allergens, a few allergens the food definitely contains are listed afterward: dairy and fish.

The first step is to determine which ingredients can't possibly contain any of the allergens in any food in your list. In the above example, none of the ingredients kfcds, nhms, sbzzf, or trh can contain an allergen. Counting the number of times any of these ingredients appear in any ingredients list produces 5: they all appear once each except sbzzf, which appears twice.

Determine which ingredients cannot possibly contain any of the allergens in your list. How many times do any of those ingredients appear?

Your puzzle answer was 2798.

The first half of this puzzle is complete! It provides one gold star: *

--- Part Two ---
Now that you've isolated the inert ingredients, you should have enough information to figure out which ingredient contains which allergen.

In the above example:

mxmxvkd contains dairy.
sqjhc contains fish.
fvjkl contains soy.
Arrange the ingredients alphabetically by their allergen and separate them by commas to produce your canonical dangerous ingredient list. (There should not be any spaces in your canonical dangerous ingredient list.) In the above example, this would be mxmxvkd,sqjhc,fvjkl.

Time to stock your raft with supplies. What is your canonical dangerous ingredient list?

*/

func day21() {
	file, err := os.Open("./app/aoc/2020/day21_input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	ingredientsInAllergen := make(map[string][]string, 0)
	ingredientCount := make(map[string]int, 0)
	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Split(text, " (contains ")
		ingredients := strings.Split(split[0], " ")
		allergens := strings.Split(strings.Replace(split[1], ")", "", 1), ", ")

		for _, ingredient := range ingredients {
			ingredientCount[ingredient] += 1
		}
		for _, allergen := range allergens {
			if previous, ok := ingredientsInAllergen[allergen]; ok {
				current := make([]string, 0)
				for _, a := range previous {
					for _, ingredient := range ingredients {
						if a == ingredient {
							current = append(current, ingredient)
						}
					}
				}
				ingredientsInAllergen[allergen] = current
			} else {
				ingredientsInAllergen[allergen] = ingredients
			}
		}
	}

	total := 0
	for ingredient, count := range ingredientCount {
		broke := false
		for _, ingredients := range ingredientsInAllergen {
			for _, ing := range ingredients {
				if ingredient == ing {
					broke = true
					break
				}
			}
			if broke {
				break
			}
		}

		if broke {
			continue
		}
		fmt.Printf("%s\n", ingredient)
		total += count
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%d", total)
	}
}

func day21b() {
	file, err := os.Open("./app/aoc/2020/day21_input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	ingredientsInAllergen := make(map[string][]string, 0)
	ingredientCount := make(map[string]int, 0)
	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Split(text, " (contains ")
		ingredients := strings.Split(split[0], " ")
		allergens := strings.Split(strings.Replace(split[1], ")", "", 1), ", ")

		for _, ingredient := range ingredients {
			ingredientCount[ingredient] += 1
		}
		for _, allergen := range allergens {
			if previous, ok := ingredientsInAllergen[allergen]; ok {
				current := make([]string, 0)
				for _, a := range previous {
					for _, ingredient := range ingredients {
						if a == ingredient {
							current = append(current, ingredient)
						}
					}
				}
				ingredientsInAllergen[allergen] = current
			} else {
				ingredientsInAllergen[allergen] = ingredients
			}
		}
	}

	ingredientsMap := make(map[string]string, 0)
	allergenCount := len(ingredientsInAllergen)
	ingredientsMap = getValidAllergens(ingredientsMap, ingredientsInAllergen)
	for len(ingredientsMap) < allergenCount {
		for allergen, ingredients := range ingredientsInAllergen {
			newIngredients := make([]string, 0)
			for _, ingredient := range ingredients {
				if _, ok := ingredientsMap[ingredient]; !ok {
					newIngredients = append(newIngredients, ingredient)
				}
			}
			ingredientsInAllergen[allergen] = newIngredients
		}
		ingredientsMap = getValidAllergens(ingredientsMap, ingredientsInAllergen)
	}

	combo := make([]string, 0)
	for ingredient, allergen := range ingredientsMap {
		combo = append(combo, fmt.Sprintf("%s %s", allergen, ingredient))
	}
	sort.Strings(combo)

	ings := make([]string, len(combo))
	for i, ing := range combo {
		ings[i] = strings.Split(ing, " ")[1]
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s", strings.Join(ings, ","))
	}
}

func getValidAllergens(ingredients map[string]string, mapOf map[string][]string) map[string]string {
	for allergen, i := range mapOf {
		if len(i) == 1 {
			ingredients[i[0]] = allergen
		}
	}

	return ingredients
}

//kfcds, nhms, sbzzf, or trh

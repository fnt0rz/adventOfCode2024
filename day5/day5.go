package day5

import (
	"fmt"
	"main/utils"
	"slices"
	"sort"
	"strings"
)

func Run() {
	data := utils.GetData("5", false)
	defer utils.Timer("Day 5")()

	lines := strings.Split(data, "\n")
	settings, pages := splitParts(lines)

	part1Result, part2result := solve(createRules(settings), pages)
	fmt.Println(part1Result, part2result)
}

func solve(rules map[int][]int, pages []string) (int, int) {
	solution1 := 0
	solution2 := 0
	for _, page := range pages {
		numbers := utils.StringSlice(strings.Split(page, ",")).ToInt()

		// filter out even numbers
		if len(numbers)%2 == 0 {
			continue
		}

		valid := true
		for i, number := range numbers {
			position := checkCorrectPosition(rules, number, numbers)
			valid = numbers[i] == numbers[position]
			if !valid {
				break
			}
		}
		if valid {
			middle := findMiddleNumber(numbers)
			solution1 += middle
		} else {
			reOrderPage := reOrderPage(rules, numbers)
			middle := findMiddleNumber(reOrderPage)
			solution2 += middle
		}
	}

	return solution1, solution2
}

func checkCorrectPosition(rules map[int][]int, number int, numbers []int) int {
	r := rules[number]
	position := len(numbers) - 1
	for j := len(numbers) - 1; j >= 0; j-- {
		if slices.Contains(r, numbers[j]) {
			position--
		}
	}
	return position
}

func reOrderPage(rules map[int][]int, page []int) []int {
	orderedPage := []int{}
	queue := map[int]int{}

	for _, number := range page {
		position := checkCorrectPosition(rules, number, page)
		queue[position] = number
	}

	keys := make([]int, 0, len(queue))
	for key := range queue {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	for _, key := range keys {
		orderedPage = append(orderedPage, queue[key])
	}

	return orderedPage
}

func findMiddleNumber(numbers []int) int {
	middleIndex := len(numbers) / 2
	return numbers[middleIndex]
}

func splitParts(lines []string) ([]string, []string) {

	for index, line := range lines {
		if line == "" {
			return lines[:index], lines[index+1:]

		}
	}
	return nil, nil
}

func createRules(lines []string) map[int][]int {
	rules := map[int][]int{}

	for _, line := range lines {
		rule := utils.StringSlice(strings.Split(line, "|")).ToInt()

		if _, exists := rules[rule[0]]; !exists {
			rules[rule[0]] = rule[1:]
		} else {
			rules[rule[0]] = append(rules[rule[0]], rule[1:]...)
		}
	}
	return rules
}

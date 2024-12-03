package day3

import (
	"fmt"
	"main/utils"
	"regexp"
	"strings"
)

func Run() {
	data := utils.GetData("3", false)
	defer utils.Timer("Day 3")()

	totalValue := part1(data)
	totalValue2 := part2(data)

	fmt.Println(totalValue, totalValue2)
}

func part1(data string) int {
	matches := regexp.MustCompile(`mul\(\d+,\d+\)`).FindAllString(data, -1)
	totalValue := 0
	for _, match := range matches {
		totalValue += calculate(match)
	}
	return totalValue
}

func calculate(match string) int {
	numbers := utils.StringSlice(strings.Split(match[4:len(match)-1], ",")).ToInt()
	return numbers[0] * numbers[1]
}

func part2(data string) int {
	// This regex matches "mul(number,number)", "do()", or "don't()"
	matches := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`).FindAllString(data, -1)

	totalValue := 0
	isEnabled := true
	for _, match := range matches {
		switch {
		case match == "do()":
			isEnabled = true
		case match == "don't()":
			isEnabled = false
		default:
			if isEnabled {
				totalValue += calculate(match)
			}
		}
	}

	return totalValue
}

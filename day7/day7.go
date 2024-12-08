package day7

import (
	"fmt"
	"main/utils"
	"strconv"
	"strings"
)

func Run() {
	data := utils.GetData("7", false)
	defer utils.Timer("Day 7")()

	lines := strings.Split(data, "\n")
	result1, result2 := solve(lines)

	fmt.Println(result1, result2)
}

func solve(lines []string) (int, int) {
	result1 := 0
	result2 := 0
	for _, line := range lines {
		colunIndex := strings.Index(line, ":")
		result, _ := strconv.Atoi(line[:colunIndex])
		parts := utils.StringSlice(strings.Split(line[colunIndex+2:], " ")).ToInt()

		if trySolve(result, 0, parts, false) {
			result1 += result
		}
		if trySolve(result, 0, parts, true) {
			result2 += result
		}
	}
	return result1, result2
}

func trySolve(result, current int, parts []int, part2 bool) bool {
	if len(parts) == 0 {
		return result == current
	}

	if current > result {
		return false
	}

	if trySolve(result, calculate(current, parts[0], '+'), parts[1:], part2) {
		return true
	}

	if part2 && trySolve(result, calculate(current, parts[0], '|'), parts[1:], part2) {
		return true

	}

	return trySolve(result, calculate(current, parts[0], '*'), parts[1:], part2)
}

func calculate(a, b int, operator rune) int {
	calculation := 0
	switch operator {
	case '+':
		calculation = a + b
	case '*':
		calculation = a * b
	case '|':
		mul, q := 10, 10
		{
			for q != 0 {
				q = b / mul
				if q > 0 {
					mul *= 10
				}
			}
		}
		calculation = (a * mul) + b
	}

	return calculation
}

package day1

import (
	"fmt"
	"main/utils"
	"sort"
	"strings"
)

func Run() {
	data := utils.GetData("1", false)
	lines := strings.Split(string(data), "\n")

	defer utils.Timer("Day 1")()

	leftInts, rightInts := createIntSlices(lines)

	total1 := result1(leftInts, rightInts)
	total2 := result2(leftInts, rightInts)

	fmt.Println("Difference part1: ", total1)
	fmt.Println("Difference part2: ", total2)
}

func result2(leftInts, rightInts []int) int {

	counts := make(map[int]int)
	for _, rightInt := range rightInts {
		counts[rightInt]++
	}

	total := 0
	for _, leftInt := range leftInts {
		total += leftInt * counts[leftInt]
	}

	return total
}

func result1(leftInts, rightInts []int) int {
	var differences []int
	for i := 0; i < len(leftInts); i++ {
		difference := leftInts[i] - rightInts[i]
		differences = append(differences, utils.AbsInt(difference))
	}

	total := utils.Sum(differences)
	return total
}

func createIntSlices(lines []string) (leftInts, rightInts []int) {
	var left utils.StringSlice
	var right utils.StringSlice
	for _, line := range lines {
		values := strings.Split(line, "   ")
		left = append(left, values[0])
		right = append(right, values[1])
	}

	leftInts = left.ToInt()
	rightInts = right.ToInt()

	sort.Ints(leftInts)
	sort.Ints(rightInts)

	return leftInts, rightInts
}

package day4

import (
	"fmt"
	"main/utils"
	"strings"
)

func Run() {
	data := utils.GetData("4", false)
	defer utils.Timer("Day 4")()

	lines := strings.Split(data, "\n")

	xmas1Count := part1(lines)
	xmas2Count := part2(lines)

	fmt.Println(xmas1Count, xmas2Count)
}

func part1(lines []string) int {
	count := 0
	for y, line := range lines {
		for x, char := range line {
			if char == 'X' {
				count += checkXmas(x, y, lines)
			}
		}
	}
	return count
}

func part2(lines []string) int {
	count := 0
	for y, line := range lines {
		for x, char := range line {
			if char == 'A' {
				count += checkMas(x, y, lines)
			}
		}
	}
	return count
}

func checkXmas(x int, y int, lines []string) int {
	count := 0

	// check forward
	if (len(lines[0]) - x) >= 4 {
		if lines[y][x] == 'X' && lines[y][x+1] == 'M' && lines[y][x+2] == 'A' && lines[y][x+3] == 'S' {
			count++
		}
	}

	// check backwards
	if x >= 3 {
		if lines[y][x] == 'X' && lines[y][x-1] == 'M' && lines[y][x-2] == 'A' && lines[y][x-3] == 'S' {
			count++
		}
	}

	// check downwards
	if (len(lines) - 4) >= y {
		if lines[y][x] == 'X' && lines[y+1][x] == 'M' && lines[y+2][x] == 'A' && lines[y+3][x] == 'S' {
			count++
		}
	}

	// check upwards
	if y >= 3 {
		if lines[y][x] == 'X' && lines[y-1][x] == 'M' && lines[y-2][x] == 'A' && lines[y-3][x] == 'S' {
			count++
		}
	}

	// check forward-down
	if (len(lines[0])-4) >= x && (len(lines)-4) >= y {
		if lines[y][x] == 'X' && lines[y+1][x+1] == 'M' && lines[y+2][x+2] == 'A' && lines[y+3][x+3] == 'S' {
			count++
		}
	}

	// check forward-up
	if (len(lines[0])-4) >= x && y >= 3 {
		if lines[y][x] == 'X' && lines[y-1][x+1] == 'M' && lines[y-2][x+2] == 'A' && lines[y-3][x+3] == 'S' {
			count++
		}
	}

	// check backwards-up
	if y >= 3 && x >= 3 {
		if lines[y][x] == 'X' && lines[y-1][x-1] == 'M' && lines[y-2][x-2] == 'A' && lines[y-3][x-3] == 'S' {
			count++
		}
	}

	// check backwards-down
	if x >= 3 && (len(lines)-4) >= y {
		if lines[y][x] == 'X' && lines[y+1][x-1] == 'M' && lines[y+2][x-2] == 'A' && lines[y+3][x-3] == 'S' {
			count++
		}
	}

	return count
}

func checkMas(x int, y int, lines []string) int {
	count := 0

	if y == 0 || x == 0 || len(lines)-1 <= y || len(lines[0])-1 <= x {
		return 0
	}

	// normal
	if lines[y-1][x-1] == 'M' && lines[y-1][x+1] == 'M' && lines[y+1][x-1] == 'S' && lines[y+1][x+1] == 'S' {
		count++
	}

	// flipped vertical
	if lines[y-1][x-1] == 'S' && lines[y-1][x+1] == 'S' && lines[y+1][x-1] == 'M' && lines[y+1][x+1] == 'M' {
		count++
	}

	// flipped vertical and horizontal
	if lines[y-1][x-1] == 'S' && lines[y-1][x+1] == 'M' && lines[y+1][x-1] == 'S' && lines[y+1][x+1] == 'M' {
		count++
	}

	// flipped horizontal
	if lines[y-1][x-1] == 'M' && lines[y-1][x+1] == 'S' && lines[y+1][x-1] == 'M' && lines[y+1][x+1] == 'S' {
		count++
	}

	return count
}

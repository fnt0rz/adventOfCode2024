package day6

import (
	"fmt"
	"main/utils"
	"strings"
)

type Guard struct {
	x         int
	y         int
	positions map[postion]int
}

type postion struct {
	x int
	y int
}

func Run() {
	data := utils.GetData("6", true)
	defer utils.Timer("Day 6")()

	lines := strings.Split(data, "\n")
	part1Result := part1(lines)

	fmt.Println(part1Result)
}

func part1(lines []string) int {
	guard := Guard{0, 0, make(map[postion]int)}
	startingPos := postion{0, 0}
	for y, line := range lines {
		for x, char := range line {
			if char == '^' {
				guard.x = x
				guard.y = y
				startingPos = postion{x, y}
				break
			}
		}
	}

	finished := false
	nextDirection := "up"

stepLoop:
	for !finished {
		switch nextDirection {
		case "up":
			nextDirection, finished = moveUp(lines, &guard, startingPos)
		case "right":
			nextDirection, finished = moveRight(lines, &guard)
		case "down":
			nextDirection, finished = moveDown(lines, &guard)
		case "left":
			nextDirection, finished = moveLeft(lines, &guard)
		default:
			break stepLoop
		}
	}

	keys := make([]postion, 0, len(guard.positions))
	for p := range guard.positions {
		keys = append(keys, p)
	}

	return len(keys)
}

func moveUp(lines []string, guard *Guard, startingPosition postion) (nextDirection string, finished bool) {
	maxSteps := guard.y

	blocked := false
	for i := 0; i < maxSteps+1; i++ {
		if lines[guard.y-i][guard.x] == '#' {
			guard.y -= i - 1
			blocked = true
			break
		}
		pos := postion{guard.x, guard.y - i}
		if _, ok := guard.positions[pos]; ok {
			if pos != startingPosition {
				fmt.Println("loop")
			}
		}
		guard.positions[pos]++
	}

	return "right", !blocked
}

func moveRight(lines []string, guard *Guard) (nextDirection string, finished bool) {
	maxSteps := len(lines[guard.y]) - guard.x - 1

	blocked := false
	for i := 0; i < maxSteps+1; i++ {
		if lines[guard.y][guard.x+i] == '#' {
			guard.x += i - 1
			blocked = true
			break
		}
		pos := postion{guard.x + i, guard.y}
		guard.positions[pos]++
	}

	return "down", !blocked
}

func moveDown(lines []string, guard *Guard) (nextDirection string, finished bool) {
	maxSteps := len(lines) - guard.y - 1

	blocked := false
	for i := 0; i < maxSteps+1; i++ {
		if lines[guard.y+i][guard.x] == '#' {
			guard.y += i - 1
			blocked = true
			break
		}
		pos := postion{guard.x, guard.y + i}
		guard.positions[pos]++
	}

	return "left", !blocked
}

func moveLeft(lines []string, guard *Guard) (nextDirection string, finished bool) {
	maxSteps := guard.x

	blocked := false
	for i := 0; i < maxSteps+1; i++ {
		if lines[guard.y][guard.x-i] == '#' {
			guard.x -= i - 1
			blocked = true
			break
		}

		pos := postion{guard.x - i, guard.y}
		guard.positions[pos]++
	}

	return "up", !blocked
}

package utils

func AbsInt(x int) int {
	return absDiffInt(x, 0)
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func Sum(i []int) (total int) {
	for _, r := range i {
		total += r
	}
	return total
}

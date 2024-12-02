package day2

import (
	"fmt"
	"main/utils"
	"strings"
)

func Run() {

	data := utils.GetData("2", false)
	lines := strings.Split(string(data), "\n")

	defer utils.Timer("Day 2")()

	safeReports, safeReportsWithDampener := solve(lines)

	fmt.Println(safeReports, safeReports+safeReportsWithDampener)
}

func solve(lines []string) (int, int) {
	safeReports := 0
	safeReportsWithDampener := 0

	for _, line := range lines {
		reportNum := utils.StringSlice(strings.Split(line, " ")).ToInt()

		switch {
		case isReportSafe(reportNum):
			safeReports++
		case checkForSaftyWithDampener(reportNum):
			safeReportsWithDampener++
		}
	}

	return safeReports, safeReportsWithDampener
}

func checkForSaftyWithDampener(reportNum []int) bool {
	for i := 0; i < len(reportNum); i++ {
		isSafe := isSafeWithDampener(reportNum, i)
		if isSafe {
			return true
		}
	}
	return false
}

func isSafeWithDampener(reportNum []int, index int) bool {
	reportCopy := make([]int, len(reportNum))
	copy(reportCopy, reportNum)

	if index == len(reportCopy)-1 {
		reportCopy = reportCopy[:index]
	} else {
		reportCopy = append(reportCopy[:index], reportCopy[index+1:]...)
	}

	return isReportSafe(reportCopy)
}

func isReportSafe(reportNum []int) bool {
	isIncrease, isDecrease := false, false

	for i := 1; i < len(reportNum); i++ {
		diff := reportNum[i] - reportNum[i-1]
		if diff == 0 || (isIncrease && diff < 0) || (isDecrease && diff > 0) || diff > 3 || diff < -3 {
			return false
		}
		isIncrease = isIncrease || diff > 0
		isDecrease = isDecrease || diff < 0
	}

	return true
}

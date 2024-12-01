package utils

import (
	"strconv"
	"strings"
)

type StringSlice []string

func (s StringSlice) ToInt() []int {

	i := make([]int, len(s))
	for index, v := range s {
		r, err := strconv.Atoi(strings.TrimSpace(v))

		if err != nil {
			panic(err.Error())
		}

		i[index] = r
	}
	return i
}

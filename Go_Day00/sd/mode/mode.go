package mode

import (
	"math"
)

type Mode int

func ComputeMode(arr []int) Mode {
	hash := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		hash[arr[i]]++
	}
	maxCount := 0
	for _, v := range hash {
		if v >= maxCount {
			maxCount = v
		}
	}
	max := math.MaxInt
	for k, v := range hash {
		if v == maxCount && k < max {
			max = k
		}
	}
	return Mode(max)
}
